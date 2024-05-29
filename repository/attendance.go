package repository

import (
	"database/sql"
	"student-attendance/structs"
)

func GetAllAttendance(db *sql.DB) (results []structs.Attendance, err error) {
	sql := "SELECT * FROM attendance"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var a = structs.Attendance{}

		err := rows.Scan(&a.ID, &a.StudentID, &a.Date, &a.Status)
		if err != nil {
			panic(err)
		}

		results = append(results, a)
	}

	return
}

func InsertAttendance(db *sql.DB, a structs.Attendance) (err error) {
	sql := "INSERT INTO attendance (student_id, date, status) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, a.StudentID, a.Date, a.Status)
	return errs.Err()
}

func UpdateAttendance(db *sql.DB, a structs.Attendance) (err error) {
	sql := "UPDATE attendance SET date= $1, status=$2 WHERE id=$3"
	errs := db.QueryRow(sql, a.Date, a.Status, a.ID)
	return errs.Err()
}

func DeleteAttendance(db *sql.DB, a structs.Attendance) (err error) {
	sql := "DELETE FROM attendance WHERE id=$1"
	errs := db.QueryRow(sql, a.ID)
	return errs.Err()
}
