package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	// go sql driver
	_ "github.com/go-sql-driver/mysql"

	"golang.org/x/net/context"

	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
	"github.com/leedsjb/capstone2k18/servers/elevate/parsers"

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

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("Google Cloud HTTPS L7 Load Balancer Health Check"))
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
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
		addr = ":80"
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
	sqlDbName := os.Getenv("SQLDBNAME")

	cfg := mysql.Cfg(sqlInstance, "proxyuser", "")
	cfg.DBName = sqlDbName
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		fmt.Printf("Error dialing MySQL: %v", err)
		os.Exit(1)
	}
	fmt.Println("Connected to SQL")
	defer db.Close()

	// [LOAD TRIES]
	var aircraftTrie = indexes.NewTrie()
	var groupsTrie = indexes.NewTrie()
	var peopleTrie = indexes.NewTrie()
	notifier := handlers.NewNotifier()

	handlerCtx := handlers.NewHandlerContext(aircraftTrie, groupsTrie, peopleTrie, db)
	parserCtx := parsers.NewParserContext(aircraftTrie, groupsTrie, peopleTrie, db, notifier)
	if err := handlerCtx.LoadAircraftTrie(aircraftTrie); err != nil {
		log.Fatalf("Error loading aircraft trie: %v", err)
	}
	if err := handlerCtx.LoadGroupsTrie(groupsTrie); err != nil {
		log.Fatalf("Error loading groups trie: %v", err)
	}
	if err := handlerCtx.LoadPeopleTrie(peopleTrie); err != nil {
		log.Fatalf("Error loading people trie: %v", err)
	}

	// [PUB/SUB]

	// TODO: temp workaround, maybe better soln?
	ctx := context.Background()

	// create pub/sub client
	psClient, err := pubsub.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}

	testTopicNames := [18]string{
		"test_mission_create",
		"test_mission_complete",
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

	testSubNames := [18]string{
		"test_mission_create_sub",
		"test_mission_complete",
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

		subName := subscription.ID()
		log.Printf("subscribed: [%v]", subName)

		go subscribe(subscription, notifier, parserCtx)
	}

	// [HTTPS]

	// Create a new mux for the web server.
	mux := http.NewServeMux()

	// Tell the mux to call your handlers
	wsh := handlers.NewWebSocketsHandler(notifier)
	mux.HandleFunc("/", HealthCheckHandler)
	mux.Handle("/v1/ws", wsh)
	mux.HandleFunc("/v1/aircraft", handlerCtx.AircraftHandler)
	mux.HandleFunc("/v1/aircraft/", handlerCtx.AircraftDetailHandler)
	mux.HandleFunc("/v1/people", handlerCtx.PeopleHandler)
	mux.HandleFunc("/v1/people/", handlerCtx.PersonDetailHandler)
	// TODO: write peopleMeHandler for auth
	// mux.HandleFunc("/v1/people/me", handlerCtx.PeopleMeHandler)
	mux.HandleFunc("/v1/groups", handlerCtx.GroupsHandler)
	mux.HandleFunc("/v1/groups/", handlerCtx.GroupDetailHandler)
	// TODO: write resourcesHandler after we set up cloud storage
	mux.HandleFunc("/v1/resources/", handlerCtx.ResourcesHandler)

	//Wrap this new mux with CORS middleware handler and add that
	//to the main server mux.
	wrappedMux := handlers.NewCORSHandler(mux)

	// Start a web server listening on the address you read from
	// the environment variable, using the mux you created as
	// the root handler. Use log.Fatal() to report any errors
	// that occur when trying to start the web server.
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
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

// listen for and process pubsub events
func subscribe(subscription *pubsub.Subscription, notifier *handlers.Notifier, parserCtx *parsers.ParserContext) {
	ctx := context.Background()
	err := subscription.Receive(ctx, func(ctx context.Context, pulledMsg *pubsub.Message) {
		// if subscription is topicName
		// var msg messages.TopicNameStruct
		subName := subscription.ID()
		switch subName {
		case "test_mission_create_sub":
			msg := &messages.Mission_Create{}
			msgType := "mission-create"
			// parses information into structs formatted for front-end
			// and delivers via websocket
			if err := parserCtx.ParseMissionCreate(msg, pulledMsg, msgType); err != nil {
				fmt.Printf("[MISSION CREATE] ERROR: %v", err)
			}
		case "test_mission_complete":
			msg := &messages.Mission_Complete{}
			msgType := "mission-complete"
			parserCtx.ParseMissionComplete(msg, pulledMsg, msgType)
		case "test_mission_waypoints_update_sub":
			msg := &messages.Mission_Waypoint_Update{}
			msgType := "mission-waypoints-update"
			parserCtx.ParseMissionWaypointsUpdate(msg, pulledMsg, msgType)
		case "test_mission_crew_update_sub":
			msg := &messages.Mission_Crew_Update{}
			msgType := "mission-crew-update"
			parserCtx.ParseMissionCrewUpdate(msg, pulledMsg, msgType)
		case "test_waypoint_create_sub":
			msg := &messages.Waypoint{}
			msgType := "waypoint-create"
			log.Printf("no current action: %v", subName)
			parserCtx.ParseWaypointCreate(msg, pulledMsg, msgType)
		case "test_waypoint_update_sub":
			msg := &messages.Waypoint{}
			msgType := "waypoint-update"
			log.Printf("no current action: %v", subName)
			parserCtx.ParseWaypointUpdate(msg, pulledMsg, msgType)
		case "test_waypoint_delete_sub":
			msg := &messages.Waypoint_Delete{}
			msgType := "waypoint-delete"
			log.Printf("no current action: %v", subName)
			parserCtx.ParseWaypointDelete(msg, pulledMsg, msgType)
		case "test_aircraft_create_sub":
			msg := &messages.Aircraft_Create{}
			msgType := "aircraft-create"
			parserCtx.ParseAircraftCreate(msg, pulledMsg, msgType)
		case "test_ac_properties_update_sub":
			msg := &messages.Aircraft_Props_Update{}
			msgType := "aircraft-props-update"
			parserCtx.ParseAircraftPropsUpdate(msg, pulledMsg, msgType)

		// [PENDING]: Wait for Brian to add ID to these actions
		case "test_ac_crew_update_sub":
			msg := &messages.Aircraft_Crew_Update{}
			msgType := "aircraft-crew-update"
			parserCtx.ParseAircraftCrewUpdate(msg, pulledMsg, msgType)
			// TODO: call sql sproc
		case "test_ac_service_schedule_sub":
			msg := &messages.Aircraft_Service_Schedule{}
			msgType := "aircraft-service-schedule"
			parserCtx.ParseAircraftServiceSchedule(msg, pulledMsg, msgType)
			// TODO: call sql sproc
		case "test_ac_position_update_sub":
			msg := &messages.Aircraft_Pos_Update{}
			msgType := "aircraft-position-update"
			parserCtx.ParseAircraftPositionUpdate(msg, pulledMsg, msgType)
			// TODO: call sql sproc
		//[END PENDING]

		case "test_user_create_sub":
			msg := &messages.User{}
			msgType := "user-create"
			log.Printf("no current action: %v", subName)
			parserCtx.ParseUserCreate(msg, pulledMsg, msgType)
		case "test_user_update_sub":
			msg := &messages.User{}
			msgType := "user-update"
			log.Printf("no current action: %v", subName)
			parserCtx.ParseUserUpdate(msg, pulledMsg, msgType)
		case "test_user_delete_sub":
			msg := &messages.User_Delete{}
			msgType := "user-delete"
			log.Printf("no current action: %v", subName)
			parserCtx.ParseUserDelete(msg, pulledMsg, msgType)
		case "test_group_create_sub":
			msg := &messages.Group{}
			msgType := "group-create"
			log.Printf("no current action: %v", subName)
			parserCtx.ParseGroupCreate(msg, pulledMsg, msgType)
		case "test_group_update_sub":
			msg := &messages.Group{}
			msgType := "group-update"
			log.Printf("no current action: %v", subName)
			parserCtx.ParseGroupUpdate(msg, pulledMsg, msgType)
		case "test_group_delete_sub":
			msg := &messages.Group_Delete{}
			msgType := "group-delete"
			log.Printf("no current action: %v", subName)
			parserCtx.ParseGroupDelete(msg, pulledMsg, msgType)
		default:
			log.Printf("not a valid subscription type")
		}

		// track number of messages processed
		// countMu.Lock()
		// count++
		// countMu.Unlock()

		pulledMsg.Ack()
		log.Printf("Message Acknowledged: [%v]", subName)
	})
	if err != nil {
		log.Fatalf("Could not receive subscription: %v", err)
	}
}
