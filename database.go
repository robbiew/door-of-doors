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

type CategoryList struct {
	DoorId       string
	DoorCode     string
	DoorTitle    string
	CategoryName string
	ServerName   string
}

func initDb() {

	if _, err := os.Stat("./data.db"); errors.Is(err, fs.ErrNotExist) {
		fmt.Print(err.Error())

		file, err := os.Create("data.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		fmt.Println("data.db doesn't exist - created")

		db, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File
		defer db.Close()                          // Defer Closing the database
		createDoorsTable(db)                      // Create Database Tables
		createCategoriesTable(db)
		createServersTable(db)

		insertCategory(db, "RPGs: Medieval & Fantasy")

		insertServer(db, "Gold Mine")
		insertServer(db, "BBS Link")
		insertServer(db, "Door Party")

		// INSERT RECORDS
		insertDoor(db, "LORD", "Legend of the Red Dragon", 1, 1)
		insertDoor(db, "LORD", "Legend of the Red Dragon", 1, 2)
		insertDoor(db, "LORD", "Legend of the Red Dragon", 1, 3)

		// DISPLAY INSERTED RECORDS

	} else {
		fmt.Print("database file already exists - skipping")

	}

}

func createDoorsTable(db *sql.DB) {
	createDoorsTableSQL := `CREATE TABLE doors (
		"idDoor" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT NOT NULL,
		"title" TEXT NOT NULL,
		"categoryId" integer NOT NULL,
        "serverId" integer NOT NULL		
	  );` // SQL Statement for Create Table

	fmt.Println("Create doors table...")
	statement, err := db.Prepare(createDoorsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("doors table created")
}

func createCategoriesTable(db *sql.DB) {
	createCategoriesTableSQL := `CREATE TABLE categories (
		"idCategory" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"categoryName" TEXT NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Println("Create categories table...")
	statement, err := db.Prepare(createCategoriesTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("categories table created")
}

func createServersTable(db *sql.DB) {
	createServersTableSQL := `CREATE TABLE servers (
		"idServer" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"serverName" TEXT NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Println("Create servers table...")
	statement, err := db.Prepare(createServersTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("servers table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertDoor(db *sql.DB, code string, title string, category int, server int) {
	fmt.Println("Inserting door record ...")
	insertDoorSQL := `INSERT INTO doors(code, title, categoryId, serverId) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertDoorSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, title, category, server)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// We are passing db reference connection from main to our method with other parameters
func insertCategory(db *sql.DB, categoryName string) {
	fmt.Println("Inserting category record ...")
	insertCategorySQL := `INSERT INTO categories(categoryName) VALUES (?)`
	statement, err := db.Prepare(insertCategorySQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(categoryName)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertServer(db *sql.DB, serverName string) {
	fmt.Println("Inserting server record ...")
	insertServerSQL := `INSERT INTO servers(serverName) VALUES (?)`
	statement, err := db.Prepare(insertServerSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(serverName)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func doorsByCategory(db *sql.DB) []CategoryList {

	rows, err := db.Query(`
    
    SELECT
        idDoor,
        doors.title AS title, 
        doors.code AS code, 
        categories.categoryName AS categoryName,
        servers.serverName AS serverName
    FROM 
        doors
    INNER JOIN categories ON categories.idCategory = doors.categoryId  
    INNER JOIN servers ON servers.idServer = doors.serverId
    WHERE
        doors.categoryId = 1;

  `)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var categoryList []CategoryList
	for rows.Next() {

		var title string
		var idDoor int
		var code string
		var categoryName string
		var serverName string

		err := rows.Scan(&idDoor, &title, &code, &categoryName, &serverName)
		if err != nil {
			log.Fatal(err)
		}

		categoryList = append(categoryList, CategoryList{DoorId: fmt.Sprint(idDoor), DoorTitle: title, CategoryName: categoryName, DoorCode: code, ServerName: serverName})
	}
	return categoryList
}
