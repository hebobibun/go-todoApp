package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-app/activity"
	"todo-app/config"
	"todo-app/user"
)

func main() {
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = user.AuthMenu{DB: conn}
	var actMenu = activity.ActMenu{DB: conn}

	var inputMenu = 1

	for inputMenu != 0 {
		fmt.Println("TODO LIST APP")
		fmt.Println("------------------")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("-")
		fmt.Println("0. Exit")
		fmt.Println("------------------")
		fmt.Print("choose a menu [1, 2, 0] : ")
		fmt.Scanln(&inputMenu)

		if inputMenu == 1 {
			
			// REGISTER
			fmt.Println("=======================")
			var newUser user.User
			fmt.Println("REGISTER")
			fmt.Print("Insert your name : ")
			consoleReader := bufio.NewReader(os.Stdin)
			newName, _ := consoleReader.ReadString('\n')
			newName = strings.TrimSuffix(newName, "\n")
			newUser.Name = newName
			fmt.Print("Insert new password : ")
			fmt.Scanln(&newUser.Password)

			res, err := authMenu.Register(newUser)
			if err != nil {
				fmt.Println("------------------")
				fmt.Println(err.Error())
			}
			if res {
				fmt.Println("------------------")
				fmt.Println("Registered a new user successfully!")
			} else {
				fmt.Println("------------------")
				fmt.Println("Failed to register a new user.")
			}
			fmt.Println("=======================")

		} else if inputMenu == 2 {

			// LOGIN
			fmt.Println("=======================")
			var inputName, inputPassword string
			fmt.Println("LOGIN PAGE")
			fmt.Println("------------------")
			fmt.Print("Input your name : ")
			consoleReader := bufio.NewReader(os.Stdin)
			inputName, _ = consoleReader.ReadString('\n')
			inputName = strings.TrimSuffix(inputName, "\n")
			fmt.Print("Input your password : ")
			fmt.Scanln(&inputPassword)

			res, err := authMenu.Login(inputName, inputPassword)
			if err != nil {
				fmt.Println("------------------")
				fmt.Println(err.Error())
			} else {
				fmt.Println("...")
				fmt.Println("[Logged in successfully]")
			}

			if res.ID > 0 {

				// HOME - LOGGED IN USER
				fmt.Println("=======================")
				isLogin := true
				loginMenu := 0

				for isLogin {

					// MENU - LOGGED IN USER
					fmt.Printf("Welcome to todoApp, %v!\n", res.Name)
					fmt.Println("------------------")
					fmt.Println("1. Insert a new activity")
					fmt.Println("2. My activity")
					fmt.Println("3. My profile")
					fmt.Println("4. Update password")
					fmt.Println("-")
					fmt.Println("9. Logout")
					fmt.Println("------------------")
					fmt.Print("Choose a menu [1, 2, 3, 4, 9] : ")
					fmt.Scanln(&loginMenu)
					fmt.Println("=======================")

					if loginMenu == 1 {

						// INSERT A NEW ACTIVITY
						inputActivity := activity.Activity{}
						inputActivity.IDUser = res.ID
						fmt.Println("INSERT A NEW ACTIVITY")
						fmt.Println("------------------")
						fmt.Print("Insert activity title : ")
						consoleReader := bufio.NewReader(os.Stdin)
						newAct, _ := consoleReader.ReadString('\n')
						newAct = strings.TrimSuffix(newAct, "\n")
						inputActivity.Title = newAct
						fmt.Print("Insert activity location : ")
						inputLoc, _ := consoleReader.ReadString('\n')
						inputLoc = strings.TrimSuffix(inputLoc, "\n")
						inputActivity.Location = inputLoc

						actRes, err := actMenu.Insert(inputActivity)
						if err != nil {
							fmt.Println("------------------")
							fmt.Println(err.Error())
						} else {
							fmt.Println("------------------")
							fmt.Println("Inserted a new activity successfully!")
							fmt.Println("=======================")
						}
						inputActivity.ID = actRes
						
					} else if loginMenu == 2 {

						// MY ACTIVITY
						fmt.Println("MY ACTIVITY")
						fmt.Println("------------------")
						activities, err := actMenu.Show(res.ID)
						if err != nil {
							fmt.Println("------------------")
							fmt.Println(err.Error())
						} else if len(activities) == 0 {
							fmt.Println("You have 0 activity right now.")
							fmt.Println("=======================")
						} else {
							for i := 0; i < len(activities); i++ {
								fmt.Println("Activity", i+1, "   : ", activities[i].Title)
								fmt.Println("Location      : ", activities[i].Location)
								fmt.Println("Date          : ", activities[i].CreateDate)
								fmt.Println("------------------")
							}
							fmt.Println("=======================")
						}
						

					} else if loginMenu == 3 {

						// MY PROFILE
						fmt.Println("MY PROFILE")
						fmt.Println("------------------")
						fmt.Println("Nama : ", res.Name)
						fmt.Println("=======================")

					} else if loginMenu == 4 {

						// UPDATE PASSWORD
						fmt.Println("UPDATE PASSWORD")
						fmt.Println("------------------")
						var newPass string
						fmt.Print("Input new password : ")
						fmt.Scanln(&newPass)

						isChanged, err := authMenu.UpdatePassword(newPass, res.ID)
						if err != nil {
							fmt.Println("------------------")
							fmt.Println(err.Error())
						}
						if isChanged {
							fmt.Println("------------------")
							fmt.Println("Password updated successfully!")
							isLogin = false
						}

					} else if loginMenu == 9 {

						// EXIT
						isLogin = false

					}
				}
			}
			fmt.Println("=======================")
		}
	}
}
