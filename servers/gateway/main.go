package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/leedsjb/capstone2k18/servers/gateway/indexes"

	"github.com/leedsjb/capstone2k18/servers/gateway/models/users"

	"gopkg.in/mgo.v2"

	"github.com/leedsjb/capstone2k18/servers/gateway/sessions"

	"github.com/go-redis/redis"

	"github.com/leedsjb/capstone2k18/servers/gateway/handlers"
)

//NewServiceProxy creates a reverse proxy for a microservice
func NewServiceProxy(addrs []string, signingKey string, store sessions.Store) *httputil.ReverseProxy {
	nextIndex := 0
	mx := sync.Mutex{}
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			//modify the request to indicate remote host
			sessionState := &handlers.SessionState{}
			_, err := sessions.GetState(r, signingKey, store, sessionState)
			if err == nil {
				userJSON, err := json.Marshal(sessionState.AuthenticatedUser)
				if err != nil {
					log.Printf("Error marshaling user: %v", err)
				}
				r.Header.Add("X-User", string(userJSON))
			} else {
				log.Printf("Error retrieving session state: %v", err)
				r.Header.Del("X-User")
			}

			mx.Lock()
			r.URL.Host = addrs[nextIndex%len(addrs)]
			nextIndex++
			mx.Unlock()
			r.URL.Scheme = "http"
		},
	}
}

func main() {
	//Read the following environment variables:
	//ADDR: the server port
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}
	//TLSKEY and TLSCERT: paths to TLS key and cert
	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")
	if tlsKeyPath == "" || tlsCertPath == "" {
		fmt.Println("Please provide both a TLS key and cert")
		os.Exit(1)
	}
	//SESSIONKEY: a string to use when signing and validating SessionIDs
	sessionKey := os.Getenv("SESSIONKEY")
	if sessionKey == "" {
		fmt.Println("Please provide a session key")
		os.Exit(1)
	}
	//REDISADDR: the address of your redis session store server
	redisAddr := os.Getenv("REDISADDR")
	if redisAddr == "" {
		fmt.Println("Please the address of your redis session store server")
		os.Exit(1)
	}
	// TODO: Instead of DBADDR, how to connect to MySQL
	//DBADDR: the address of your database server
		dbAddr := os.Getenv("DBADDR")
		if dbAddr == "" {
			fmt.Println("Please provide the address of your database server")
			os.Exit(1)
		}

	// TODO: Listen to Pub/Sub where we (subscriber) are listening to published msgs 
	//may contain a comma-delimited list of network addresses where
	//messaging microservice instances are listening
	messagesSvcAddrs := os.Getenv("MESSAGESSVCADDR")
	if len(messagesSvcAddrs) == 0 {
		fmt.Println("Please provide a list of messaging microservice addresses")
		os.Exit(1)
	}

	// TODO: load balance message addresses
	splitMessagesSvcAddrs := strings.Split(messagesSvcAddrs, ",")
	//may contain a comma-delimited list of network addresses where
	//page summary microservice instances are listening
	summarySvcAddrs := os.Getenv("SUMMARYSVCADDR")
	if len(summarySvcAddrs) == 0 {
		fmt.Println("Please provide a list of summary microservice addresses")
		os.Exit(1)
	}
	splitSummarySvcAddrs := strings.Split(summarySvcAddrs, ",")

	//Use the REDISADDR to create a new redis Client, which you can pass to your
	//sessions.NewRedisStore() function.
	sessionStore := sessions.NewRedisStore(redis.NewClient(&redis.Options{Addr: redisAddr}), time.Duration(300)*time.Second)

	// TODO: not MongoDB -> how to connect to MySQL
	//Use the DBADDR to dial MongoDB server
	sess, err := mgo.Dial(dbAddr)
	if err != nil {
		fmt.Printf("Error dialing Mongo: %v", err)
		os.Exit(1)
	}

	//Initialize user store.
	userStore := users.NewMongoStore(sess, "messages", "users")

	//Initialize trie
	trie := indexes.NewTrie()

	//Create and initialize a new instance of your handler context struct with
	//the signing key, the sesson store, and the user store.
	handlerCtx := handlers.NewHandlerContext(sessionKey, sessionStore, userStore, trie)
	if err := handlerCtx.UserStore.LoadTrie(trie); err != nil {
		fmt.Printf("Error loading trie: %v", err)
		os.Exit(1)
	}

	//Create a new mux
	mux := http.NewServeMux()

	//Add microservice handlers
	///v1/summary
	mux.Handle("/v1/summary", NewServiceProxy(splitSummarySvcAddrs, sessionKey, sessionStore))
	///v1/channels
	mux.Handle("/v1/channels", NewServiceProxy(splitMessagesSvcAddrs, sessionKey, sessionStore))
	mux.Handle("/v1/channels/", NewServiceProxy(splitMessagesSvcAddrs, sessionKey, sessionStore))
	///v1/messages
	mux.Handle("/v1/messages", NewServiceProxy(splitMessagesSvcAddrs, sessionKey, sessionStore))
	mux.Handle("/v1/messages/", NewServiceProxy(splitMessagesSvcAddrs, sessionKey, sessionStore))

	//Add regular handlers
	///v1/users: UsersHandler
	mux.HandleFunc("/v1/users", handlerCtx.UsersHandler)
	///v1/users/me: UsersMeHandler
	mux.HandleFunc("/v1/users/me", handlerCtx.UsersMeHandler)
	///v1/sessions: SessionsHandler
	mux.HandleFunc("/v1/sessions", handlerCtx.SessionsHandler)
	///v1/sessions/mine: SessionsMineHandler
	mux.HandleFunc("/v1/sessions/mine", handlerCtx.SessionsMineHandler)

	notifier := handlers.NewNotifier()

	//Create a new notifier
	mux.Handle("/v1/ws", handlerCtx.NewWebSocketsHandler(notifier))

	//Wrap this new mux with CORS middleware handler and add that
	//to the main server mux.
	wrappedMux := handlers.NewCORSHandler(mux)

	//use wrappedMux instead of mux as root handler
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))
}
