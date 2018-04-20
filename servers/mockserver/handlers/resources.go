package handlers

import (
	"net/http"
)

// Resource ...
type Resource struct {
	Link  string `json:"link"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

var resources = []*Resource{
	{
		Link:  "www.google.com",
		Name:  "Google",
		Image: "http://www.stickpng.com/assets/images/580b57fcd9996e24bc43c51f.png",
	},
	{
		Link:  "www.facebook.com",
		Name:  "Facebook",
		Image: "http://pngimg.com/uploads/facebook_logos/facebook_logos_PNG19751.png",
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
