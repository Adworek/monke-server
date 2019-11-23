package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID                string `json:"id,omitempty"`
	Nick              string `json:"nick,omitempty"`
	Bio               string `json:"bio,omitempty"`
	SubscriberCount   int    `json:"subscriber_count,omitempty"`
	SubscriptionCount int    `json:"subscription_count,omitempty"`
	PostCount         int    `json:"post_count,omitempty"`
	Created           int64  `json:"created,omitempty"`
}

func (user User) New(nick string, about string, id string) User {
	var created User = User{
		ID:                id,
		Nick:              nick,
		Bio:               bio,
		SubscriberCount:   0,
		SubscriptionCount: 0,
		Created:           time.Now().Unix(),
	}

	return created
}
