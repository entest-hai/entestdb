// hai tran 01/01/2025
// test concurrent connections to database
// configure max_connection: /var/lib/pgsql/{version_number}/data/postgresql.conf
// test: go test -run TestConcurrentConnection
// SELECT sum(numbackends) FROM pg_stat_database;
// SELECT count(*) as num_connections FROM pg_stat_activity;
// SELECT sum(numbackends) FROM pg_stat_database WHERE datname is not null;
// SELECT count(*) FROM pg_stat_activity WHERE datname is not null;
// check max connection configuration 
// SELECT current_setting('max_connections');
// SELECT * FROM pg_settings WHERE name = 'max_connections'; 
package test

import (
	"database/sql"
	"entestdb/config"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestConcurrentConnection(t *testing.T) {

	//
	var num_user int = 50

	// create a channel
	channel := make(chan int, num_user)

	// loop to run 10 connections using goroutine
	for i := 0; i < num_user; i++ {
		go func(user int) {
			//
			// fmt.Println("test sending connections to database", i)

			// connection string to database postgresql
			connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.HOST, config.PORT, config.USER, config.DB_NAME, config.PASSWORD)

			// ping database
			db, error := sql.Open("postgres", connStr)

			if error != nil {
				panic(error)
			}

			// infinite loop
			var count int = 0
			for {

				//

				// query a table
				rows, error := db.Query("SELECT count(*) FROM student;")

				if error != nil {
					panic(error)
				}

        // close db connection 
				defer rows.Close()
        defer db.Close()

				var n_student int

				for rows.Next() {
					rows.Scan(&n_student)
				}

				fmt.Printf("loop: %d, user: %d num student: %d\n", count, user, n_student)

				// parse rows response
				// for rows.Next() {
				// 	var id int
				// 	var name string
				// 	var login string
				// 	var age int
				// 	var gpa float32

				// 	error = rows.Scan(&id, &name, &login, &age, &gpa)

				// 	if error != nil {
				// 		panic(error)
				// 	}

				// 	fmt.Println(id, name, login, age, gpa)
				// }

				// sleep few seconds
				time.Sleep(3 * time.Second)

				count += 1

			}

			// send a message to channel
			channel <- 1
		}(i)
	}

	// retrieve values from the channel to wait for go routine completed
	for i := 0; i < num_user; i++ {
		fmt.Println(<-channel)
	}
}

func TestQuerySelectStudent(t *testing.T) {
	fmt.Println("test sending connections to database")

	// connection string to database postgresql
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.HOST, config.PORT, config.USER, config.DB_NAME, config.PASSWORD)

	// ping database
	db, error := sql.Open("postgres", connStr)

	if error != nil {
		panic(error)
	}

	// query a table
	rows, error := db.Query("SELECT * FROM student LIMIT 100;")

	if error != nil {
		panic(error)
	}

	// parse rows response
	for rows.Next() {
		var id int
		var name string
		var login string
		var age int
		var gpa float32

		error = rows.Scan(&id, &name, &login, &age, &gpa)

		if error != nil {
			panic(error)
		}

		fmt.Println(id, name, login, age, gpa)
	}
}

func TestPingDatabase(t *testing.T) {
	fmt.Println("test sending connections to database")

	// connection string to database postgresql
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.HOST, config.PORT, config.USER, config.DB_NAME, config.PASSWORD)

	// ping database
	db, error := sql.Open("postgres", connStr)

	if error != nil {
		panic(error)
	}

	error = db.Ping()

	if error != nil {
		panic(error)
	}

	fmt.Println("Successfully connected to database")
}
