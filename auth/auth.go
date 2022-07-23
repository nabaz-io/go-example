package auth

import "time"

var db map[string]string = map[string]string{
	"admin": "admin",
}

func Auth(username string, password string) bool {
	time.Sleep(1)
	return db[username] == password
}
