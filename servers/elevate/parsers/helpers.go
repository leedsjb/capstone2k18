package parsers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// TODO: may not work with strongly typed messages
// parse data for delivery
func parse(msg interface{}, pulledMsg *pubsub.Message, msgType string) {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
	}

	log.Printf("Message contents: %#v", msg)
}

// send message to client
func (ctx *ParserContext) ClientNotify(msg interface{}, msgType string, pulledMsg *pubsub.Message) {
	fmt.Printf("[CLIENT NOTIFY] In method")

	// TODO: parse pubsub message into client struct
	toClient := &messages.ClientMsg{
		Type:    msgType,
		Payload: msg,
	}

	// send msg contents to websockets
	send, err := json.Marshal(toClient)
	if err != nil {
		log.Printf("PROBLEM marshaling json: %v", err)
		pulledMsg.Ack()
		return
	}
	fmt.Printf("[WEBSOCKET] Printing message: %v", toClient)
	ctx.Notifier.Notify(send)
}

func getAircraftCallsign(ID int, db *sql.DB) (string, error) {
	missionID := strconv.Itoa(ID)
	// get mission from db using missionID
	aircraftRow, err := db.Query("SELECT ac_callsign FROM tblAIRCRAFT JOIN tblMISSION ON tblMISSION.aircraft_id = tblAIRCRAFT.ac_id WHERE mission_id=" + missionID)
	if err != nil {
		fmt.Printf("Error querying MySQL for aircraftID: %v", err)
	}
	var aircraftCallsign string
	err = aircraftRow.Scan(&aircraftCallsign)
	if err != nil {
		return "", err
	}
	return aircraftCallsign, nil
}

func MissionNotify(callsign string, aircraftID int, mobile string) error {
	// Followed example https://github.com/alainakafkes/twilio-sms-tutorial/blob/master/main.go

	if len(mobile) > 10 {
		mobile = strings.Replace(mobile, " ", "", -1)
		mobile = strings.Replace(mobile, "-", "", -1)
		mobile = strings.Replace(mobile, "+", "", -1)
		mobile = strings.Replace(mobile, "(", "", -1)
		mobile = strings.Replace(mobile, ")", "", -1)
		mobile = strings.Replace(mobile, ".", "", -1)
	}
	if len(mobile) < 12 {
		mobile = "+1" + mobile
	}

	if mobile == "+14258941368" {
		mobile = "+14259986567"
	}

	fmt.Printf("[MISSION NOTIFICATION] Sending to: %v\n", mobile)

	alnwTwilio := os.Getenv("TWILIO_NUM")

	accountSid := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_TOKEN")
	fmt.Printf("[MISSION NOTIFICATION] SID: %v, AUTHTOKEN: %v\n", accountSid, authToken)
	twilioUrl := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	fmt.Printf("[MISSION NOTIFICATION] TwilioURL: %v\n", twilioUrl)

	notifyUrl := "https://test.elevate.airliftnw.org/" + strconv.Itoa(aircraftID)
	notification := "You've been assigned to a new mission on " + callsign + ": " + notifyUrl

	log.Printf("[MISSION NOTIFICATION] notification string: %v\n", notification)

	// Pack up the data for our message

	msgData := url.Values{}
	msgData.Set("To", mobile)
	msgData.Set("From", alnwTwilio)
	msgData.Set("Body", notification)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", twilioUrl, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Println("[MISSION NOTIFICATION] Success!!!")
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Printf("[MISSION NOTIFICATION] SID is: %v\n", data["sid"])
		}
	} else {
		return fmt.Errorf(resp.Status)
	}

	return nil
}
