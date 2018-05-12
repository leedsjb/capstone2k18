package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// GroupsHandler ...
func (ctx *HandlerContext) GroupsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		groups := []*messages.Group{}
		groupsRow, err := ctx.DB.Query("SELECT * FROM Groups")
		if err != nil {
			fmt.Printf("Error querying MySQL for requestor: %v", err)
		}
		// TODO create variables and fill contents from retrieved rows
		for groupsRow.Next() {
			err = groupsRow.Scan(&requestor)
			if err != nil {
				fmt.Printf("Error scanning requestor row: %v", err)
				os.Exit(1)
			}
		}

		for _, v := range groupDetails {
			g := &Group{
				ID:            v.ID,
				Name:          v.Name,
				PeoplePreview: v.PeoplePreview,
			}
			groups = append(groups, g)
		}
		respond(w, groups)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// GroupDetailHandler ...
func GroupDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding ID: %v", err), http.StatusBadRequest)
		return
	}
	var gd *messages.GroupDetail
	for _, v := range groupDetails {
		if v.ID == id {
			gd = v
			break
		}
	}
	if gd == nil {
		http.Error(w, "No group with that ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		respond(w, gd)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
