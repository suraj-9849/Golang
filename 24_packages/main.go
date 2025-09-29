package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/suraj-9849/Golang/auth"
	"github.com/suraj-9849/Golang/user"
)

func main() {
	auth.LoginWithCredentials("Suraj","secretPassword")
	data:= auth.GetSession()
	fmt.Println(data)
	userObj  := user.User{Email: "suraj@example.com",Name: "Baburao"}
	userData := user.UserData(userObj.Email, userObj.Name)
	fmt.Println(userData)
	color.Green("Hello world!")
	color.Red("Hello world!")
}