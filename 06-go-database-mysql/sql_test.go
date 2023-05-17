package go_database_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('3', 'Pamungkas')"
	_, error := db.ExecContext(ctx, script)

	if error != nil {
		panic(error)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM customer"
	rows, error := db.QueryContext(ctx, script)

	if error != nil {
		panic(error)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id", id)
		fmt.Println("Name", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, error := db.QueryContext(ctx, script)

	if error != nil {
		panic(error)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)

		if err != nil {
			panic(err)
		}
		fmt.Println("=======================")
		fmt.Println("Id", id)
		fmt.Println("Name", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance", balance)
		fmt.Println("Rating", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date", birthDate.Time)
		}
		fmt.Println("Married", married)
		fmt.Println("Created At", createdAt)

	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, error := db.QueryContext(ctx, script)

	fmt.Println(script)

	if error != nil {
		panic(error)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}
