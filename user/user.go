package user

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	ID int
	Nama string
	Password string
}

type AuthMenu struct{
	DB *sql.DB
}

func (am *AuthMenu) Duplicate(name string) bool {
	res := am.DB.QueryRow("SELECT id_user FROM users where nama = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
        if err.Error() != "sql: no rows in result set" {
            log.Println("Result scan error", err.Error())
        }
        return false
    }
	if idExist > 0 {
		return true
	}
	return true
}

func (am *AuthMenu) Register(newUser User) (bool, error) {
	// Menyiapkan query untuk insert
	rgsQurey, err := am.DB.Prepare("INSERT INTO users (nama, password) VALUES (?, ?)")
	if err != nil {
		log.Println("Prep insert user : ", err.Error())
		return false, errors.New("prep statement insert user error")
	}

	if am.Duplicate(newUser.Nama) {
		log.Println("Duplicated")
		return false, errors.New("Nama sudah digunakan")
	}

	// Menjalankan query dengan parameter tertentu
	res, err := rgsQurey.Exec(newUser.Nama, newUser.Password)
	if err != nil {
		log.Println("Insert user : ", err.Error())
		return false, errors.New("insert user error")
	}

	// Cek berapa baris yang affected oleh query di atas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert user ", err.Error())
		return false, errors.New("error setelah insert")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (am *AuthMenu) Login(user User) (bool, int, error) {
	row := am.DB.QueryRow("SELECT id_user FROM users WHERE nama = ? AND password = ?", user.Nama, user.Password)

	var idLogin int
	err := row.Scan(&idLogin)
	if err != nil {
		log.Println(err.Error())
		return false, 0, errors.New("Error scan")
	}
	if idLogin > 0 {
		return true, idLogin, nil
	}

	return false, idLogin, errors.New("username or password is invalid")
}
