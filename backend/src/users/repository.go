package users

import (
	"commons"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func InsertUser(user User) []byte {
	fmt.Printf("Insert %s user in dabatabase...\n", user.Name)

	db, err = sql.Open(PostgresDriver, DataSourceName)

	if HasErrors(err) {
		return commons.BuildResponse(commons.RESPONSE_STATUS["INTERNAL_SERVER_ERROR"], "Unable to connect with database.")
	}

	sqlStatement := "INSERT INTO users(name, birthdate, email) VALUES ($1, $2, $3)"
	_, err = db.Exec(sqlStatement, user.Name, user.BirthDate, user.Email)

	if HasErrors(err) {
		return commons.BuildResponse(commons.RESPONSE_STATUS["INTERNAL_SERVER_ERROR"], "Unable to save user data.")
	}

	defer db.Close()

	return commons.BuildResponse(commons.RESPONSE_STATUS["SUCCESS"], "Success.")
}

func ListUsers() []byte {
	fmt.Println("Listing users...")

	db, err = sql.Open(PostgresDriver, DataSourceName)

	if HasErrors(err) {
		response, _ := json.Marshal(User{})
		return response
	}

	sqlStatement := "SELECT id, name, birthdate, email FROM Users"
	result, err := db.Query(sqlStatement)

	if HasErrors(err) {
		response, _ := json.Marshal(User{})
		return response
	}

	users := []User{}
	for result.Next() {
		user := User{}
		result.Scan(&user.ID, &user.Name, &user.BirthDate, &user.Email)
		users = append(users, user)
	}

	response, _ := json.Marshal(users)
	return response
}
