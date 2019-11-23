package models

import (
	"time"
)

type User struct {
	ID                string
	Nick              string
	Bio               string
	SubscriberCount   int
	SubscriptionCount int
	PostCount         int
	Created           int64
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
