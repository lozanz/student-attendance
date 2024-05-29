package repository

import (
	"database/sql"
	"student-attendance/structs"
)

func GetAllUsers(db *sql.DB) (err error, results []structs.Users) {
	sql := "SELECT id,username,role FROM users"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user = structs.Users{}
		err = rows.Scan(&user.ID, &user.Username, &user.Role)
		if err != nil {
			panic(err)
		}
		results = append(results, user)
	}
	return
}

func GetUserByID(db *sql.DB, userID int64) (structs.Users, error) {
	var user structs.Users
	sql := "SELECT id, username, password, role FROM users WHERE id = $1"
	err := db.QueryRow(sql, userID).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return structs.Users{}, err
	}
	return user, nil
}

func RegisterUser(db *sql.DB, user structs.Users) (err error) {
	sql := "INSERT INTO users (id,username, password, role) VALUES($1, $2, $3, $4)"
	errs := db.QueryRow(sql, user.ID, user.Username, user.Password, user.Role)

	return errs.Err()
}

func LoginUser(db *sql.DB, user structs.Users) (result structs.Users, err error) {
	sql := "SELECT * FROM users WHERE username = $1"
	rows, err := db.Query(sql, user.Username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = structs.Users{}

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
		if err != nil {
			panic(err)
		}

		result = user
	}

	return
}

func AuthenticateUser(db *sql.DB, user structs.Users) (results []structs.Users, err error) {
	sql := "SELECT * FROM users WHERE username = $1 AND password = $2"

	rows, err := db.Query(sql, user.Username, user.Password)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var user = structs.Users{}
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
		if err != nil {
			break
		}
		results = append(results, user)
	}

	return
}
