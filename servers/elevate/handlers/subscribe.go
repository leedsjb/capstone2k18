package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type pushRequest struct {
	Message struct {
		Attributes map[string]string
		Data       []byte
		ID         string `json:"message_id"`
	}
	Subscription string
}

func PushHandler(w http.ResponseWriter, r *http.Request) {
	msg := &pushRequest{}
	if err := json.NewDecoder(r.Body).Decode(msg); err != nil {
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	// TODO: process msg.Message.Data by msg.Subscription

	// notifier.Notify([]byte(msg.Message.Data))

	// if msg.Subscription = topicID1
		// SQL stored procedure for topicID1
	// if msg.Subscription = topicID2
		// SQL stored procedure for topic ID2
	// etc...
}