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
		ImageLink: "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/272px-Google_2015_logo.svg.png",
	},
	{
		ID:        2,
		Link:      "www.facebook.com",
		Name:      "Facebook",
		ImageLink: "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c2/F_icon.svg/240px-F_icon.svg.png",
	},
	{
		ID:        3,
		Link:      "www.twitter.com",
		Name:      "Twitter",
		ImageLink: "https://upload.wikimedia.org/wikipedia/en/thumb/9/9f/Twitter_bird_logo_2012.svg/295px-Twitter_bird_logo_2012.svg.png",
	},
	{
		ID:        4,
		Link:      "www.amazon.com",
		Name:      "Amazon",
		ImageLink: "https://upload.wikimedia.org/wikipedia/commons/thumb/7/70/Amazon_logo_plain.svg/602px-Amazon_logo_plain.svg.png",
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
