package activity

import (
	"database/sql"
	"errors"
	"log"
)

type Activity struct {
	ID int
	Title string
	CreateDate string
	Location string
	IDUser int
}

type ActMenu struct {
	DB *sql.DB
}

func (am *ActMenu) Insert(newActivity Activity) (int, error) {
	insertQry, err := am.DB.Prepare("INSERT INTO activities (title, location, create_date, id_user) VALUES (?,?,now(),?)")
	if err != nil {
		log.Println("Prepare insert newActivity : ", err.Error())
		return 0, errors.New("Prepare statement insert new activity error.")
	}

	res, err := insertQry.Exec(newActivity.Title, newActivity.Location, newActivity.IDUser)
	if err != nil {
		log.Println("Insert new activity : ", err.Error())
		return 0, errors.New("Insert new activity error.")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("Afer inser new activity : ", err.Error())
		return 0, errors.New("Error after insert new activity.")
	}

	if affRows <= 0 {
		log.Println("No rows affected.")
		return 0, errors.New("No record affected.")
	}

	id, _ := res.LastInsertId()

	return int(id), nil
}