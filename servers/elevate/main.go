package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	// go sql driver
	_ "github.com/go-sql-driver/mysql"

	"golang.org/x/net/context"

	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"

	"cloud.google.com/go/pubsub"
)

//NotificationsHandler handles requests for the /notifications resource
type NotificationsHandler struct {
	notifier *handlers.Notifier
}

//NewNotificationsHandler constructs a new NotificationsHandler
func NewNotificationsHandler(notifier *handlers.Notifier) *NotificationsHandler {
	return &NotificationsHandler{notifier}
}

//ServeHTTP handles HTTP requests for the NotificationsHandler
func (nh *NotificationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Notification pushed from the server at %s", time.Now().Format("15:04:05"))
	nh.notifier.Notify([]byte(msg))
}

//main is the main entry point for the server
func main() {

	// get the present working directory
	pwd := os.Getenv("PWD")

	// Read the ADDR environment variable to get the address
	// the server should listen on.
	addr := os.Getenv("ADDR")
	// If empty, default to ":443" for https
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
		fmt.Println(tlsKeyPath)
		fmt.Println(tlsCertPath)
	}

	// [CloudSQL]

	sqlInstance := os.Getenv("SQLINSTANCE")
	sqlUser := os.Getenv("SQLUSER")
	sqlPass := os.Getenv("SQLPASS")
	sqlDbName := os.Getenv("SQLDBNAME")
	// sqlTblName := os.Getenv("SQLTABLENAME")
	// first25Missions := os.Getenv("SQLMISSIONS")

	cfg := mysql.Cfg(sqlInstance, sqlUser, sqlPass)
	cfg.DBName = sqlDbName
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		fmt.Printf("Error dialing MySQL: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	// [PUB/SUB]

	// TODO: temp workaround, maybe better soln?
	ctx := context.Background()

	// create pub/sub client
	psClient, err := pubsub.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}

	testTopicNames := [17]string{
		"test_mission_create",
		"test_mission_waypoints_update",
		"test_mission_crew_update",
		"test_waypoint_create",
		"test_waypoint_update",
		"test_waypoint_delete",
		"test_aircraft_create",
		"test_ac_properties_update",
		"test_ac_crew_update",
		"test_ac_service_schedule",
		"test_ac_position_update",
		"test_user_create",
		"test_user_update",
		"test_user_delete",
		"test_group_create",
		"test_group_update",
		"test_group_delete",
	}

	testSubNames := [17]string{
		"test_mission_create_sub",
		"test_mission_waypoints_update_sub",
		"test_mission_crew_update_sub",
		"test_waypoint_create_sub",
		"test_waypoint_update_sub",
		"test_waypoint_delete_sub",
		"test_aircraft_create_sub",
		"test_ac_properties_update_sub",
		"test_ac_crew_update_sub",
		"test_ac_service_schedule_sub",
		"test_ac_position_update_sub",
		"test_user_create_sub",
		"test_user_update_sub",
		"test_user_delete_sub",
		"test_group_create_sub",
		"test_group_update_sub",
		"test_group_delete_sub",
	}

	notifier := handlers.NewNotifier()

	// create topics and subscriptions if don't yet exist
	for i, testTopicName := range testTopicNames {
		topic := psClient.Topic(testTopicName)

		// create topic if doesn't exist
		exists, err := topic.Exists(ctx)
		if err != nil {
			log.Fatalf("Error checking for topic: %v", err)
		}
		if !exists {
			if _, err := psClient.CreateTopic(ctx, testTopicName); err != nil {
				log.Fatalf("Failed to create topic: %v", err)
			}
		}

		// create subscription if doesn't exist
		subscription := psClient.Subscription(testSubNames[i])
		exists, err = subscription.Exists(ctx)
		if err != nil {
			log.Fatalf("Error checking for subscription: %v", err)
		}
		if !exists {
			if _, err := psClient.CreateSubscription(ctx, testSubNames[i],
				pubsub.SubscriptionConfig{Topic: topic}); err != nil {
				log.Fatalf("Failed to create subscription: %v", err)
			}
		}
		go subscribe(subscription, notifier)
	}

	// [HTTPS]

	// Create a new mux for the web server.
	mux := http.NewServeMux()

	//Wrap this new mux with CORS middleware handler and add that
	//to the main server mux.
	// wrappedMux := handlers.NewCORSHandler(mux)

	// Tell the mux to call your handlers
	wsh := handlers.NewWebSocketsHandler(notifier)
	mux.Handle("/v1/ws", wsh)

	// Start a web server listening on the address you read from
	// the environment variable, using the mux you created as
	// the root handler. Use log.Fatal() to report any errors
	// that occur when trying to start the web server.
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

type testStruct struct {
	Type    string   `json:"type"`
	Payload *payload `json:"payload"`
}

type payload struct {
	Status   string   `json:"status"`
	Callsign string   `json:"callsign"`
	Mission  *mission `json:"mission"`
}

type mission struct {
	MissionDetails string `json:"missionDetails"`
}

// listen for and process pubsub events
func subscribe(subscription *pubsub.Subscription, notifier *handlers.Notifier) {
	ctx := context.Background()
	err := subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var toClient testStruct
		log.Printf("BEFORE UNMARSHALING: %#v", msg.Data)
		if err := json.Unmarshal(msg.Data, &toClient); err != nil {
			log.Printf("PROBLEM contents of decoded json: %#v", toClient)
			log.Printf("Could not decode message data: %#v", msg.Data)
			msg.Ack()
			return
		}

		// TODO: process msg contents
		// TODO: send msg contents to websockets
		// TODO: save msg contents to CloudSQL using StoredProcedures

		// [sample processing message]
		// log.Printf("[ID %d] Processing. . .", id)
		// if err := update(id); err != nil {
		// 	log.Printf("[ID %d] could not update: %v", id, err)
		// 	msg.Nack()
		// 	return
		// }

		// track number of messages processed
		// countMu.Lock()
		// count++
		// countMu.Unlock()

		send, err := json.Marshal(toClient)
		if err != nil {
			log.Printf("PROBLEM marshaling json: %v", err)
			msg.Ack()
			return
		}
		notifier.Notify(send)

		msg.Ack()
		log.Printf("Message Acknowledged: (%v)\n", send)
	})
	if err != nil {
		log.Fatalf("Could not receive subscription: %v", err)
	}
}
