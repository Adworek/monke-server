package models

type User struct {
	ID                string
	Nick              string
	Bio               string
	SubscriberCount   int
	SubscriptionCount int
	PostCount         int
	Created           int64
}
