package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID                string `json:"id"`
	Nick              string `json:"nick"`
	Bio               string `json:"bio"`
	SubscriberCount   int    `json:"subscriber_count"`
	SubscriptionCount int    `json:"subscription_count"`
	PostCount         int    `json:"post_count"`
	Created           int64  `json:"created"`
}

func (user User) New(nick string, bio string, id string) User {
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

func (user User) JSON() ([]byte, error) {
	return json.Marshal(user)
}
