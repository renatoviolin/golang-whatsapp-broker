package dto

import "time"

type MessageRepository struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Id        string    `json:"id" bson:"id"`
	WaID      string    `json:"waID" bson:"waID"`
	Type      string    `json:"type" bson:"type"`
	Body      string    `json:"body" bson:"body"`
	IsRead    bool      `json:"isRead" bson:"isRead"`
}
