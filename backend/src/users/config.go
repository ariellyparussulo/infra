package users

import "fmt"

type User struct {
	ID        int
	Name      string
	BirthDate string
	Email     string
}

const PostgresDriver = "postgres"
const Username = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "postgres"
const DbName = "users"
const TableName = "users"

var DataSourceName = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", Username, Password, Host, DbName)

func HasErrors(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
