package handlers

import (
	"net/http"
)

// Resource ...
type Resource struct {
	ID        int    `json:"id"`
	Link      string `json:"link"`
	Name      string `json:"name"`
	ImageLink string `json:"imageLink"`
}

var resources = []*Resource{
	{
		ID:        1,
		Link:      "www.google.com",
		Name:      "Google",
		ImageLink: "http://www.stickpng.com/assets/images/580b57fcd9996e24bc43c51f.png",
	},
	{
		ID:        2,
		Link:      "www.facebook.com",
		Name:      "Facebook",
		ImageLink: "http://pngimg.com/uploads/facebook_logos/facebook_logos_PNG19751.png",
	},
}

// ResourcesHandler ...
func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		respond(w, resources)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
