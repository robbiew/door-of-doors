package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func initDb() {

	if _, err := os.Stat("./data.db"); errors.Is(err, fs.ErrNotExist) {
		fmt.Print(err.Error())

		file, err := os.Create("data.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		fmt.Println("data.db doesn't exist - created")

		sqliteDatabase, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File
		defer sqliteDatabase.Close()                          // Defer Closing the database
		createDoorTable(sqliteDatabase)                       // Create Database Tables
		createCategoryTable(sqliteDatabase)
		createServerTable(sqliteDatabase)

		insertCategory(sqliteDatabase, "RPGs: Medieval & Fantasy")

		insertServer(sqliteDatabase, "Gold Mine")
		insertServer(sqliteDatabase, "BBS Link")
		insertServer(sqliteDatabase, "Door Party")

		// INSERT RECORDS
		insertDoor(sqliteDatabase, "LORD", "Legend of the Red Dragon", 1, 1)
		insertDoor(sqliteDatabase, "LORD", "Legend of the Red Dragon", 1, 2)
		insertDoor(sqliteDatabase, "LORD", "Legend of the Red Dragon", 1, 3)

		// DISPLAY INSERTED RECORDS

	} else {
		fmt.Print("database file already exists - skipping")

	}

}

func createDoorTable(db *sql.DB) {
	createDoorTableSQL := `CREATE TABLE door (
		"idDoor" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT NOT NULL,
		"name" TEXT NOT NULL,
		"category" integer NOT NULL,
        "server" integer NOT NULL		
	  );` // SQL Statement for Create Table

	fmt.Println("Create door table...")
	statement, err := db.Prepare(createDoorTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("door table created")
}

func createCategoryTable(db *sql.DB) {
	createCategoryTableSQL := `CREATE TABLE category (
		"idCategory" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Println("Create category table...")
	statement, err := db.Prepare(createCategoryTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("category table created")
}

func createServerTable(db *sql.DB) {
	createServerTableSQL := `CREATE TABLE server (
		"idServer" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Println("Create server table...")
	statement, err := db.Prepare(createServerTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("server table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertDoor(db *sql.DB, code string, name string, category int, server int) {
	fmt.Println("Inserting door record ...")
	insertDoorSQL := `INSERT INTO door(code, name, category, server) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertDoorSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, category, server)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// We are passing db reference connection from main to our method with other parameters
func insertCategory(db *sql.DB, name string) {
	fmt.Println("Inserting category record ...")
	insertCategorySQL := `INSERT INTO category(name) VALUES (?)`
	statement, err := db.Prepare(insertCategorySQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertServer(db *sql.DB, name string) {
	fmt.Println("Inserting server record ...")
	insertServerSQL := `INSERT INTO server(name) VALUES (?)`
	statement, err := db.Prepare(insertServerSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayDoors(db *sql.DB) {
	row, err := db.Query("SELECT * FROM door ORDER BY idDoor")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var idDoor int
		var name string
		var server int
		var category int
		row.Scan(&idDoor, &name, &category, &server)
		fmt.Println(idDoor, name, category, server)
	}
}
