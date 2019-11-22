package storage

import (
    "crypto/sha256"
	"github.com/satori/go.uuid"
    "hash"
	"monke-server/models"
	"time"
)

// user id -> user object
var users map[string]models.User = map[string]models.User{}
// nickname -> user id
var nicks map[string]string = map[string]string{}
// user id -> password hash
var auths map[string][]byte = map[string][]byte{}

func NewUser(nick string, password string) models.User {
    var exists bool = true
    var id string
    for exists {
        id = uuid.NewV4().String()
        _, exists = users[id]
    }

    var user models.User = models.User{
		ID:                id,
		Nick:              nick,
		Bio:               "",
		SubscriberCount:   0,
		SubscriptionCount: 0,
		PostCount:         0,
		Created:           time.Now().Unix(),
	}

    var hashed hash.Hash = sha256.New()
    hashed.Write([]byte(password))

    users[id] = user
    nicks[user.Nick] = id
    auths[id] = hashed.Sum(nil)

	return user
}
