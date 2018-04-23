package main

// TODO determine if order of Go imports is significant
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	// "strings"
	"sync"
	"time"
	
	// sql + go sql driver
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/leedsjb/capstone2k18/servers/gateway/indexes"

	"github.com/leedsjb/capstone2k18/servers/gateway/models/users"

	"gopkg.in/mgo.v2"

	"github.com/leedsjb/capstone2k18/servers/gateway/sessions"

	"github.com/go-redis/redis"

	"github.com/leedsjb/capstone2k18/servers/gateway/handlers"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

const (
	maxConnRetries = 5
	qName          = "messageQ"
	apiRoot        = "/v1/"
	apiUsers       = apiRoot + "users"
	apiSessions    = apiRoot + "sessions"
	webSocket      = apiRoot + "ws"
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

	// get the present working directory
	pwd := os.Getenv("PWD")

	//ADDR: the server port
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	//TLSKEY and TLSCERT: paths to TLS key and cert
	// TLS Key and Cert facilitate secure communication between apiserver and the Google Cloud
	// Platform HTTPS Load Balancer. Certificate can be self-signed. Keys are managed by GCP KMS.
	// See https://cloud.google.com/kms/
	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")
	if tlsKeyPath == "" || tlsCertPath == "" {
		fmt.Println("Attempting to use default TLSKEY and TLSCERT paths")
		tlsKeyPath = pwd + "/tls/privkey.pem"
		tlsCertPath = pwd + "/tls/fullchain.pem"
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

	// // TODO: Listen to Pub/Sub where we (subscriber) are listening to published msgs 
	// //may contain a comma-delimited list of network addresses where
	// //messaging microservice instances are listening
	// messagesSvcAddrs := os.Getenv("MESSAGESSVCADDR")
	// if len(messagesSvcAddrs) == 0 {
	// 	fmt.Println("Please provide a list of messaging microservice addresses")
	// 	os.Exit(1)
	// }

	// // TODO: load balance message addresses
	// splitMessagesSvcAddrs := strings.Split(messagesSvcAddrs, ",")
	// //may contain a comma-delimited list of network addresses where
	// //page summary microservice instances are listening
	// summarySvcAddrs := os.Getenv("SUMMARYSVCADDR")
	// if len(summarySvcAddrs) == 0 {
	// 	fmt.Println("Please provide a list of summary microservice addresses")
	// 	os.Exit(1)
	// }
	// splitSummarySvcAddrs := strings.Split(summarySvcAddrs, ",")

	//Use the REDISADDR to create a new redis Client, which you can pass to your
	//sessions.NewRedisStore() function.
	sessionStore := sessions.NewRedisStore(redis.NewClient(&redis.Options{Addr: redisAddr}), time.Duration(300)*time.Second)

	// TODO: not MongoDB -> how to connect to MySQL
	//Use the DBADDR to dial MongoDB server
	sess, err := mgo.Dial(dbAddr)
	if err != nil {
		fmt.Printf("Error dialing MongoDB: %v", err)
		os.Exit(1)
	}

	sqlInstance := os.Getenv("SQLINSTANCE")
	sqlUser := os.Getenv("SQLUSER")
	sqlPass := os.Getenv("SQLPASS")
	sqlDbName := os.Getenv("SQLDBNAME")
	sqlTblName := os.Getenv("SQLTABLENAME")

	cfg := mysql.Cfg(sqlInstance, sqlUser, sqlPass)
	cfg.DBName = sqlDbName
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		fmt.Printf("Error dialing MySQL: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM " + sqlTblName)
	if err != nil {
		fmt.Printf("Error querying MySQL: %v", err)
		os.Exit(1)
	}

	// TODO: DO SOMETHING WITH ROWS
	i := 0
	for rows.Next() {
		i = i + 1
		fmt.Println(i)
		// var aircraft_lat int
		// var aircraft_long string
		// var model_title string
		// var model_desc string
		// var aircraft_category string
		// var aircraft_callsign string
		// var mission_date string
		// var agency_name string
		// var agency_area_code string
		// var agency_phone string
		// var adderss_street string
		// var address_city string
		// var address_state string
		// var address_zip string
		// err := rows.Scan(
		// 	&aircraft_lat,
		// 	&aircraft_long,
		// 	&model_title,
		// 	&model_desc,
		// 	&aircraft_category,
		// 	&aircraft_callsign,
		// 	&mission_date,
		// 	&agency_name,
		// 	&agency_area_code,
		// 	&address_state,
		// 	&address_zip,
		// 	&address_city,
		// 	&address_state,
		// 	&address_zip
		// )
		// if err != nil {
		// 	fmt.Printf("Error parsing MySQL rows: %v", err)
		// 	os.Exit(1)
		// }
		// fmt.Println(uid)
		// fmt.Println(username)
		// fmt.Println(department)
		// fmt.Println(created)
	}

	//Initialize user store.
	userStore := users.NewMongoStore(sess, "messages", "users")

	//Initialize trie
	trie := indexes.NewTrie()

	//Create and initialize a new instance of your handler context struct with
	//the signing key, the sesson store, and the user store.
	handlerCtx := handlers.NewHandlerContext(sessionKey, sessionStore, userStore, trie)
	// handlerCtx := handlers.NewHandlerContext(sessionKey, sessionStore, trie)
	if err := handlerCtx.UserStore.LoadTrie(trie); err != nil {
		fmt.Printf("Error loading trie: %v", err)
		os.Exit(1)
	}

	//Create a new mux
	mux := http.NewServeMux()

	// //Add microservice handlers
	// // /v1/summary
	// mux.Handle(apiRoot+"summary", NewServiceProxy(splitSummarySvcAddrs, sessionKey, sessionStore))
	// // /v1/channels
	// mux.Handle(apiRoot+"channels", NewServiceProxy(splitMessagesSvcAddrs, sessionKey, sessionStore))
	// mux.Handle(apiRoot+"channels/", NewServiceProxy(splitMessagesSvcAddrs, sessionKey, sessionStore))
	// ///v1/messages
	// mux.Handle(apiRoot+"messages", NewServiceProxy(splitMessagesSvcAddrs, sessionKey, sessionStore))
	// mux.Handle(apiRoot+"messages/", NewServiceProxy(splitMessagesSvcAddrs, sessionKey, sessionStore))

	//Add regular handlers
	///v1/users: UsersHandler
	mux.HandleFunc("apiUsers", handlerCtx.UsersHandler)
	///v1/users/me: UsersMeHandler
	mux.HandleFunc(apiUsers+"me", handlerCtx.UsersMeHandler)
	///v1/sessions: SessionsHandler
	mux.HandleFunc(apiSessions, handlerCtx.SessionsHandler)
	///v1/sessions/mine: SessionsMineHandler
	mux.HandleFunc(apiSessions+"/mine", handlerCtx.SessionsMineHandler)

	notifier := handlers.NewNotifier()

	//Create a new notifier
	mux.Handle(webSocket, handlerCtx.NewWebSocketsHandler(notifier))

	//Wrap this new mux with CORS middleware handler and add that
	//to the main server mux.
	wrappedMux := handlers.NewCORSHandler(mux)

	//use wrappedMux instead of mux as root handler
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))
}
