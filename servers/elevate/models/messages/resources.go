package messages

// [Client Messages]

// Resource ...
type Resource struct {
	Key		  string `json:"key"`
	ID        int    `json:"id"`
	Link      string `json:"link"`
	Name      string `json:"name"`
	ImageLink string `json:"imageLink"`
}