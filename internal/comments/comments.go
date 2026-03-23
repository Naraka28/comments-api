package comments

import "time"

type Comment struct {
	Id int `json:"id"`
	Username string    `json:"username"`
    Message  string    `json:"message"`
    Date     time.Time `json:"date"`
}

type NewComment struct {
	Username string    `json:"username"`
    Message  string    `json:"message"`
}