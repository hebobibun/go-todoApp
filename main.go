package main

import (
	"bufio"
	"fmt"
	"os"
	"todo-list/activity"
	"todo-list/config"
	"todo-list/user"
)

func main() {
	var inputMenu int = 1
	var cfg = config.ReadConfig
	var conn = config.ConnectSQL(*cfg())
	var authMenu = user.AuthMenu{DB: conn}
	var actMenu = activity.ActMenu{DB: conn}

	var isLogged bool

	for inputMenu != 0 {
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Println("Input Menu : ")
		fmt.Scanln(&inputMenu)
	
		switch inputMenu {
		case 1:
			var newUser user.User
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Input name : ")
			text, _ := reader.ReadString('\n')
			newUser.Nama = text
			fmt.Print("Input password : ")
			fmt.Scanln(&newUser.Password)
			res, err := authMenu.Register(newUser)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res {
				fmt.Println("succesfully registered")
			} else {
				fmt.Println("Failed to register")
			}

		case 2:
			var newLogin user.User
			fmt.Print("Input your name : ")
			fmt.Scanln(&newLogin.Nama)
			fmt.Print("Input your password : ")
			fmt.Scanln(&newLogin.Password)
			res, idLogin, err := authMenu.Login(newLogin)

			if err != nil {
				fmt.Println("Error login : ", err.Error())
			}

			if res {
				fmt.Println("Logged in succesfully")
				isLogged = true
				for isLogged {
					fmt.Println("1. Add a new activity")
					fmt.Println("9. Logout")
					fmt.Print("Choose a menu : ")
					fmt.Scanln(&inputMenu)

					switch inputMenu {
					case 1:
						var newActivity activity.Activity
						newActivity.ID = idLogin
						fmt.Print("Input activity : ")
						fmt.Scanln(&newActivity.Title)
						fmt.Print("Input activity location : ")
						fmt.Scanln(&newActivity.Location)
						res, err := actMenu.AddActivity(newActivity)
						fmt.Println(newActivity)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Succesfully added a new activity")
						} else {
							fmt.Println("Failed to add a new activity")
							inputMenu = 1
						}
					case 9:
						isLogged = false
					}
				}
			}
		}
	}
}