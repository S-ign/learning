package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// connect to database
	conn, err := sql.Open("pgx", "host=localhost port= 5432 dbname=test_connection user=sign password=password")
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to connect: %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to database!")

	// test connection
	err = conn.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("cannot ping database: %v\n", err))
	}

	log.Println("Pinged database")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// insert a row
	query := `insert into users (first_name, last_name)
	values ($1, $2)`

	_, err = conn.Exec(query, "Jack", "Brown")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("inserted a row")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// update a row
	stmt := `update users set first_name = $1 where id = $2`
	_, err = conn.Exec(stmt, "Jackie", 5)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("updated one or more rows")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// get one row by id
	query = `select id, first_name, last_name from users where id = $1`
	var firstName, lastName string
	var id int

	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("QueryRow returns", id, firstName, lastName)

	// delete a row
	query = `delete from users where id = $1`
	rowID := 6
	_, err = conn.Exec(query, rowID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("deleted row: ", rowID)

	// get rows from table again
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

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("record is", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("error scanning rows", err)
	}

	fmt.Println("--------------------------------")

	return nil
}
