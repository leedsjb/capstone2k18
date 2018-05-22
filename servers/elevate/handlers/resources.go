package handlers

import (
	"fmt"
	"net/http"

	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// Resource ...
type resourceRow struct {
	ID        int    `json:"id"`
	Link      string `json:"link"`
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
	ImageLink string `json:"imageLink"`
}

// var resources = []*Resource{
// 	{
// 		ID:        1,
// 		Link:      "www.google.com",
// 		Name:      "Google",
// 		ImageLink: "http://www.stickpng.com/assets/images/580b57fcd9996e24bc43c51f.png",
// 	},
// 	{
// 		ID:        2,
// 		Link:      "www.facebook.com",
// 		Name:      "Facebook",
// 		ImageLink: "http://pngimg.com/uploads/facebook_logos/facebook_logos_PNG19751.png",
// 	},
// }

// ResourcesHandler ...
func (ctx *HandlerContext) ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		resourceList := []*messages.Resource{}
		resourceRows, err := ctx.GetAllResources()
		if err != nil {
			fmt.Printf("Couldn't get resources from the database: %v", err)
			return
		}
		resourceRow := &resourceRow{}
		for resourceRows.Next() {
			err = resourceRows.Scan(
				&resourceRow.ID,
				&resourceRow.ShortName,
				&resourceRow.LongName,
				&resourceRow.Link,
				&resourceRow.ImageLink,
			)
			if err != nil {
				fmt.Printf("Error scanning resource row: %v", err)
			}
			resource := &messages.Resource{
				ID:        resourceRow.ID,
				Link:      resourceRow.Link,
				Name:      resourceRow.LongName,
				ImageLink: resourceRow.ImageLink,
			}
			resourceList = append(resourceList, resource)
		}
		respond(w, resourceList)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
