package models

import "time"

type Message struct {
	ID        string    `bson:"_id,omitempty"`
	Author    string    `bson:"author"`
	Content   string    `bson:"content"`
	Timestamp time.Time `bson:"timestamp"`
}
