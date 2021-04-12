package blog

import "time"

// Post represents the json response when
// the server asks for a blogpost or all blogposts
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Markdown  []byte    `json:"markdown"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}
