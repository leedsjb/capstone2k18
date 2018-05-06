package main

import (
	"database/sql"
	"time"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	// go sql driver
	_ "github.com/go-sql-driver/mysql"

	"golang.org/x/net/context"

	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"

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
		var ac_type_id string
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
			&ac_type_id,
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
			"========================================================\nFLIGHT %d\nmission_id: %s\naircraft_id: %s\nagency_id: %s\nmission_date: %s\naircraft_id2: %s\naircraft_callsign: %s\nmodel_id: %s\naircraft_lat: %s\naircraft_long: %s\nac_type_id: %s\nagency_id2:%s\nagency_name: %s\nagency_area_code: %s\nagency_phone: %s\naddress_id: %s\n",
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
			ac_type_id,
			agency_id2,
			agency_name,
			agency_area_code,
			agency_phone,
			address_id,
		)
		i++
	}

	// sproc := os.Getenv("SPROC")
	arbnum := "10"
	sprocRows, err := db.Query("CALL uspGetRecentMissions(" + arbnum + ")")
	if err != nil {
		fmt.Printf("Error sproc-ing MySQL: %v", err)
		os.Exit(1)
	}
	j := 1
	for sprocRows.Next() { 
		var mission_id string
		var mission_date string
		var aircraft_callsign string
		var aircraft_lat string
		var aircraft_long string
		var agency_name string
		err := sprocRows.Scan(
			&mission_id,
			&mission_date,
			&aircraft_callsign,
			&aircraft_lat,
			&aircraft_long,
			&agency_name,
		)
		if err != nil {
			fmt.Printf("Error scanning sproc rows: %v", err)
			os.Exit(1)
		}
		fmt.Printf(
			"========================================================\nAIRCRAFT %d\nmission_id: %s\nmission_date: %s\naircraft_callsign: %s\naircraft_lat: %s\naircraft_long: %s\nagency_name: %s\n",
			j, 
			mission_id,
			mission_date,
			aircraft_callsign,
			aircraft_lat,
			aircraft_long,
			agency_name,
		)
		j++
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
		go subscribe(subscription, notifier, db)
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

type pubSubMessage struct {
	Message struct {
		Attributes map[string]string
		Data       []byte
		ID         string `json:"message_id"`
	}
	Subscription string
}

type testStruct struct {
	Key string `json:"key"`
	ID string `json:"ID"`
	Hello string `json:"hello"`
	World string `json:"world"`
}

// listen for and process pubsub events
func subscribe(subscription *pubsub.Subscription, notifier *handlers.Notifier, db *sql.DB) {
	ctx := context.Background()
	err := subscription.Receive(ctx, func(ctx context.Context, pulledMsg *pubsub.Message) {
		// if subscription is topicName
		// var msg messages.TopicNameStruct
		subName := subscription.ID()
		switch subName {
		case "test_mission_create_sub":
			msg := &messages.Mission_Create{}
			parseMissionCreate(msg, pulledMsg, subName, notifier)
			// parses information into structs formatted for front-end
			// and delivers via websocket
			// clientParse(msg, pulledMsg, subName, notifier)
			// TODO: call sql sproc here
		case "test_mission_waypoints_update_sub":
			msg := &messages.Mission_Waypoint_Update{}
			parseMissionWaypointsUpdate(msg, pulledMsg, subName, notifier)
		case "test_mission_crew_update_sub":
			msg := &messages.Mission_Crew_Update{}
			parseMissionCrewUpdate(msg, pulledMsg, subName, notifier)
		case "test_waypoint_create_sub":
			// msg := &messages.Waypoint{}
			// don't need to notify client
			// just call sql sproc
		case "test_waypoint_update_sub":
			msg := &messages.Waypoint{}
			// notify client, client must update if necessary
			parseWaypointUpdate(msg, pulledMsg, subName, notifier)
			// call sql sproc
		case "test_waypoint_delete_sub":
			msg := &messages.Waypoint_Delete{}
			parseWaypointDelete(msg, pulledMsg, subName, notifier)
		case "test_aircraft_create_sub":
			// msg := &messages.Aircraft_Create{}
			// don't notify client
			// call sproc
		case "test_ac_properties_update_sub":
			msg := &messages.Aircraft_Props_Update{}
			parseAircraftPropsUpdate(msg, pulledMsg, subName, notifier)
		case "test_ac_crew_update_sub":
			msg := &messages.Aircraft_Crew_Update{}
			parseAircraftCrewUpdate(msg, pulledMsg, subName, notifier)
		case "test_ac_service_schedule_sub":
			msg := &messages.Aircraft_Service_Schedule{}
			parseAircraftServiceSchedule(msg, pulledMsg, subName, notifier)
		case "test_ac_position_update_sub":
			msg := &messages.Aircraft_Pos_Update{}
			parseAircraftPositionUpdate(msg, pulledMsg, subName, notifier)
		case "test_user_create_sub":
			// don't need to notify client
			// msg := &messages.User{}
		case "test_user_update_sub":
			msg := &messages.User{}
			parseUserUpdate(msg, pulledMsg, subName, notifier)
		case "test_user_delete_sub":
			msg := &messages.User_Delete{}
			parseUserDelete(msg, pulledMsg, subName, notifier)
		case "test_group_create_sub":
			// don't need to notify client
			// msg := &messages.Group{}
		case "test_group_update_sub":
			msg := &messages.Group{}
			parseGroupUpdate(msg, pulledMsg, subName, notifier)
		case "test_group_delete_sub":
			msg := &messages.Group_Delete{}
			parseGroupDelete(msg, pulledMsg, subName, notifier)
		default:
			log.Printf("not a valid subscription type")
		}

		// TODO: process msg contents
		log.Printf("This subscription is: %v", subscription.ID())

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

		pulledMsg.Ack()
		// log.Printf("Message Acknowledged: (%v)\n", msg)
	})
	if err != nil {
		log.Fatalf("Could not receive subscription: %v", err)
	}
}

func parseMissionCreate(msg *messages.Mission_Create,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {
		// unmarshal json into correct struct
		log.Printf("before unmarshaling: %v", string(pulledMsg.Data))
		if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
			log.Printf("PROBLEM contents of decoded json: %#v", msg)
			log.Printf("Could not decode message data: %#v", pulledMsg)
			pulledMsg.Ack()
			return
		}

		// // TODO: parse pubsub message for client
		// type Mission_Create struct {
		// 	MissionID			string 				`json:"missionID"`
		// 	TCNum				string   			`json:"TCNum"`
		// 	Asset				string   			`json:"asset"`
		// 	RequestorID			string   			`json:"requestorID"`
		// 	ReceiverID			string 			 	`json:"receiverID"`
		// 	Priority			string 			 	`json:"priority"`
		// 	// CallType			string 	 			`json:"callType"`
		// 	Patient				*Patient 			`json:"patient"`
		// 	CrewMemberID		[]string 			`json:"crewMemberID"`
		// 	Waypoints			[]*MissionWaypoint  `json:"waypoints"`
		// }
		mission := &messages.Mission{
			Key:				"mission",
			Type:				msg.CallType,
			Status:				"",				// AircraftCreated, Aircraft Service Scheduled
			Vision:				"",				// Aircraft Service Scheduled -> should this come from there??
			NextWaypointETE: 	"",				// TODO: time.now() - msg.Waypoints[0].ETE
			Waypoints:			msg.Waypoints,
			FlightNum:			msg.TCNum,
		}
		missionDetail := &messages.MissionDetail{
			Key:				"mission-detail",
			Type:				msg.CallType,
			Status:				"",
			Vision:				"",
		}

		// to make go stop complaining
		log.Printf("mission: %v", mission)
		log.Printf("missionDetail: %v", missionDetail)
	}

func parseMissionWaypointsUpdate(msg *messages.Mission_Waypoint_Update,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseMissionCrewUpdate(msg *messages.Mission_Crew_Update, 
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseWaypointUpdate(msg *messages.Waypoint, 
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseWaypointDelete(msg *messages.Waypoint_Delete,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseAircraftPropsUpdate(msg *messages.Aircraft_Props_Update,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseAircraftCrewUpdate(msg *messages.Aircraft_Crew_Update,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseAircraftServiceSchedule(msg *messages.Aircraft_Service_Schedule,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseAircraftPositionUpdate(msg *messages.Aircraft_Pos_Update,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseUserUpdate(msg *messages.User,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseUserDelete(msg *messages.User_Delete, 
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}

func parseGroupUpdate(msg *messages.Group,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}	

func parseGroupDelete(msg *messages.Group_Delete,
	pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {

	}


// // parse data for delivery to client
// func clientParse(msg interface{}, pulledMsg *pubsub.Message, subName string, notifier *handlers.Notifier) {
// 	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))
// 	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
// 		log.Printf("PROBLEM contents of decoded json: %#v", msg)
// 		log.Printf("Could not decode message data: %#v", pulledMsg)
// 		pulledMsg.Ack()
// 		return
// 	}

// 	// TODO: parse pubsub message into client struct

// 	// TODO: send msg contents to websockets
// 	// add key to specify which type of client struct it is for easy reception on
// 	// client-side
// 	msg.Key = subName
// 	send, err := json.Marshal(msg)
// 	if err != nil {
// 		log.Printf("PROBLEM marshaling json: %v", err)
// 		pulledMsg.Ack()
// 		return
// 	}
// 	notifier.Notify(send)

// 	log.Printf("Message type is: %v", msg)
// }