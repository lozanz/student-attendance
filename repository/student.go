package repository

import (
	"database/sql"
	"student-attendance/structs"
)

func GetAllStudent(db *sql.DB) (results []structs.Student, err error) {
	sql := "SELECT * FROM student"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var s = structs.Student{}

		err := rows.Scan(&s.ID, &s.UserID, &s.Name, &s.JenisKelamin, &s.Alamat, &s.ClassID)
		if err != nil {
			panic(err)
		}

		results = append(results, s)
	}

	return
}

func InsertStudent(db *sql.DB, s structs.Student) (err error) {
	sql := "INSERT INTO student ( user_id, name,jenis_kelamin,alamat,class_id) VALUES ($1, $2, $3, $4,$5)"
	errs := db.QueryRow(sql, s.UserID, s.Name, s.JenisKelamin, s.Alamat, s.ClassID)

	return errs.Err()
}
func UpdateStudent(db *sql.DB, s structs.Student) (err error) {
	sql := "UPDATE student SET user_id=$1 ,jenis_kelamin = $2, alamat = $3, class_id = $4 WHERE id = $5"
	errs := db.QueryRow(sql, s.UserID, s.JenisKelamin, s.Alamat, s.ClassID, s.ID)

	return errs.Err()
}

func DeleteStudent(db *sql.DB, s structs.Student) (err error) {
	sql := "DELETE FROM student WHERE id = $1"
	errs := db.QueryRow(sql, s.ID)

	return errs.Err()
}
