package messages

// [Client Messages]

// Resource ...
type Resource struct {
	ID        int    `json:"id"`
	Link      string `json:"link"`
	Name      string `json:"name"`
	ImageLink string `json:"imageLink"`
}
