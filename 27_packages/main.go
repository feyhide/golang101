package main

import (
	"idkman/auth"
	"idkman/user"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCred("feyhide", "SecretSauce")

	session := auth.GetSession("feyhide")

	color.Cyan("Session for feyhide: %s", session)

	u := user.User{
		Email:    "feyhide@example.com",
		Password: "SecretSauce",
	}

	color.Red("User: %s | Password: %s", u.Email, u.Password)

}
