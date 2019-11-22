package main

import (
	"monke-server/storage"
	"fmt"
)

func main() {
	user := storage.NewUser("zero", "foobar2000")
	fmt.Println(user)
}
