package user

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	ID int
	Name string
	Password string
}

type AuthMenu struct {
	DB *sql.DB
}

func (am *AuthMenu) Register(newUser User) (bool, error) {
	// Preparing the query to INSERT
	registerQry, err := am.DB.Prepare("INSERT INTO users (nama, password) VALUES (?,?)")
	if err != nil {
		log.Println("Prepare insert newUser : ", err.Error())
		return false, errors.New("Prepare statement insert user error")
	}

	// Executing INSERT query with certain parameters (name, password)
	res, err := registerQry.Exec(newUser.Name, newUser.Password)
	if err != nil {
		log.Println("Inser newUser : ", err.Error())
		return false, errors.New("Insert user error")
	}

	// Check how many rows affected after INSERT
	affRows, err := res.RowsAffected()
	if err != nil {
		log.Println("After insert newUser", err.Error())
		return false, errors.New("Error after insert")
	}

	if affRows <= 0 {
		log.Println("No rows affeced")
		return false, errors.New("No record affected")
	}

	return true, nil
}

func (am *AuthMenu) Login(name, password string) (User, error) {
	loginQry, err := am.DB.Prepare("SELECT id_user FROM users WHERE nama = ? AND password = ?")
	if err != nil {
		log.Println("Prepare select user ", err.Error())
		return User{}, errors.New("Prepare login error")
	}

	row := loginQry.QueryRow(name, password)

	if row.Err() != nil {
		log.Println("Login query : ", row.Err().Error())
		return User{}, errors.New("User data not found.")
	}

	res := User{}
	err = row.Scan(&res.ID)

	if err != nil {
		log.Println("After login query : ", err.Error())
		return User{}, errors.New("Error, can't login.")
	}

	res.Name = name

	return res, nil
}

func (am *AuthMenu) UpdatePassword(newPassword string, id int) (bool, error) {
	updatePwQry, err := am.DB.Prepare("UPDATE users SET password = ? WHERE id_user = ?")
	if err != nil {
		log.Println("Prepare update password : ", err.Error())
		return false, errors.New("Prepare statement update password error.")
	}

	res, err := updatePwQry.Exec(newPassword, id)
	if err != nil {
		log.Println("Update password : ", err.Error())
		return false, errors.New("Update password error.")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("After update password : ", err.Error())
		return false, errors.New("after update password error.")
	}

	if affRows <= 0 {
		log.Println("No rows affected")
		return false, errors.New("No record affected.")
	}

	return true, nil
}