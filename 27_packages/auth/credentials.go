package auth

import (
	"github.com/fatih/color"
)

func LoginWithCred(username string, password string) {
	color.Green("Username: %s", username)
	color.Green("Password: %s", password)
}
