package activity

import (
	"database/sql"
	"errors"
	"log"
)

type Activity struct {
	ID int
	Title string
	Location string
}

type ActMenu struct{
	DB *sql.DB
}

func (am *ActMenu) AddActivity(newActivity Activity) (bool, error) {
	// Menyiapkan query untuk insert
	addQuery, err := am.DB.Prepare("INSERT INTO activities (id_user, title, location, create_date) VALUES (?, ?, ?, now())")
	if err != nil {
		log.Println("Prep insert a new activity : ", err.Error())
		return false, errors.New("prep statement insert a new activity error")
	}

	// Menjalankan query dengan parameter tertentu
	res, err := addQuery.Exec(newActivity.ID, newActivity.Title, newActivity.Location)
	if err != nil {
		log.Println("Insert a new activity error : ", err.Error())
		return false, errors.New("Insert a new activity error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert a new activity ", err.Error())
		return false, errors.New("error after insert a new activity")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}