package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	// go sql driver
	_ "github.com/go-sql-driver/mysql"

	"golang.org/x/net/context"

	// "github.com/leedsjb/capstone2k18/servers/gateway/handlers"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"

	"cloud.google.com/go/pubsub"
)

// var subscription *pubsub.Subscription

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
	first25Missions := os.Getenv("SQLMISSIONS")

	cfg := mysql.Cfg(sqlInstance, sqlUser, sqlPass)
	cfg.DBName = sqlDbName
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		fmt.Printf("Error dialing MySQL: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query(first25Missions)
	// rows, err := db.Query("SELECT * FROM " + sqlTblName)
	if err != nil {
		fmt.Printf("Error querying MySQL: %v", err)
		os.Exit(1)
	}

	// TODO: DO SOMETHING WITH ROWS
	i := 0
	for rows.Next() { 
		var mission_id string
		var aircraft_id string
		var agency_id string
		var mission_date string
		var aircraft_id2 string
		var aircraft_callsign string
		var model_id string
		var aircraft_lat string
		var aircraft_long string
		var agency_id2 string
		var agency_name string
		var agency_area_code string
		var agency_phone string
		var address_id string
		err := rows.Scan(
			&mission_id,
			&aircraft_id,
			&agency_id,
			&mission_date,
			&aircraft_id2,
			&aircraft_callsign,
			&model_id,
			&aircraft_lat,
			&aircraft_long,
			&agency_id2,
			&agency_name,
			&agency_area_code,
			&agency_phone,
			&address_id,
		)
		if err != nil {
			fmt.Printf("Error parsing MySQL rows: %v", err)
			os.Exit(1)
		}
		fmt.Printf(
			"========================================================\nFLIGHT %d\nmission_id: %s\naircraft_id: %s\nagency_id: %s\nmission_date: %s\naircraft_id2: %s\naircraft_callsign: %s\nmodel_id: %s\naircraft_lat: %s\naircraft_long: %s\nagency_id2:%s\nagency_name: %s\nagency_area_code: %s\nagency_phone: %s\naddress_id: %s\n",
			i, 
			mission_id,
			aircraft_id,
			agency_id,
			mission_date,
			aircraft_id2,
			aircraft_callsign,
			model_id,
			aircraft_lat,
			aircraft_long,
			agency_id2,
			agency_name,
			agency_area_code,
			agency_phone,
			address_id,
		)
		i++
	}

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
		go subscribe(subscription)
	}

	// [HTTPS]

	// Create a new mux for the web server.
	mux := http.NewServeMux()

	//Wrap this new mux with CORS middleware handler and add that
	//to the main server mux.
	// wrappedMux := handlers.NewCORSHandler(mux)

	// Tell the mux to call your handlers

	// Start a web server listening on the address you read from
	// the environment variable, using the mux you created as
	// the root handler. Use log.Fatal() to report any errors
	// that occur when trying to start the web server.
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

type testStruct struct {
	Message string `json:"msg"`
}

// listen for and process pubsub events
func subscribe(subscription *pubsub.Subscription) {
	ctx := context.Background()
	err := subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var data testStruct
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("PROBLEM contents of decoded json: %#v", data)
			log.Printf("Could not decode message data: %#v", msg)
			msg.Ack()
			return
		}
		log.Printf("contents of decoded json: %#v\r\n", data)

		// TODO: process msg contents
		// TODO: send msg contents to websockets
		// TODO: save msg contents to CloudSQL using StoredProcedures

		log.Printf("Message data: %v\n", data)

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

		msg.Ack()
		log.Printf("Msg Acknowledged: (%v)\n", data)
	})
	if err != nil {
		log.Fatalf("Could not receive subscription: %v", err)
	}
}