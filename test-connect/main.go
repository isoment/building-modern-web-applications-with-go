package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Connect to database and defer closing
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=user password=secret")
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}
	defer conn.Close()
	log.Println("Connected to database")

	// Test connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database")
	}
	log.Println("Pinged database")

	// Get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a row, use placeholders in the query
	query := `insert into users (first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(query, "Jack", "Brown")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a row")

	// Get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// update a row
	stmt := `update users set first_name = $1 where first_name = $2`
	_, err = conn.Exec(stmt, "Jackson", "Jack")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Updated a row")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// Get one row by id using QueryRow
	query = `select id, first_name, last_name from users where id = $1`

	var firstName, lastName string
	var id int

	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("QueryRow returns", id, firstName, lastName)

	// Delete a row
	query = `delete from users where id = $1`
	_, err = conn.Exec(query, 6)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted a row!")

	// get rows again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name from users")

	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	// Iterate over all the rows and assign
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("-----------------------------")

	return nil
}
