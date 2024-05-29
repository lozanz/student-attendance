package repository

import (
	"database/sql"
	"student-attendance/structs"
)

func GetAllClass(db *sql.DB) (err error, results []structs.Class) {
	sql := "SELECT * FROM class"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var c = structs.Class{}
		err = rows.Scan(&c.ID, &c.Name)
		if err != nil {
			panic(err)
		}
		results = append(results, c)
	}
	return
}

func InsertClass(db *sql.DB, c structs.Class) (err error) {
	sql := "INSERT INTO class(id, name) VALUES ($1, $2)"
	errs := db.QueryRow(sql, c.ID, c.Name)
	return errs.Err()
}

func UpdateClass(db *sql.DB, c structs.Class) (err error) {
	sql := "UPDATE class SET name=$1 WHERE id=$2"
	errs := db.QueryRow(sql, c.Name, c.ID)
	return errs.Err()
}

func DeleteClass(db *sql.DB, c structs.Class) (err error) {
	sql := "DELETE FROM class WHERE id=$1"
	errs := db.QueryRow(sql, c.ID)
	return errs.Err()
}
