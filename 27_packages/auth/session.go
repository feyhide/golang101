package auth

import (
	"math/rand"
)

// Private function (accessible only within the same package).
func createSession(username string) string {
	sessionToken := ""
	for _, char := range username {
		sessionToken += string(char)
		sessionToken += string('0' + rand.Intn(10))
	}
	return sessionToken
}

// public func (exported and accessible from other packages).
func GetSession(username string) string {
	sessionToken := createSession(username)
	return sessionToken
}
