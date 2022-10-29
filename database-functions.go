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
	CategoryId   string
	CategoryName string
}

type DoorsList struct {
	DoorTitle string
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

		// Create all tables, insert records
		createTables(db)
	}
	// database exists, so do nothing
}

func createDoorsTable(db *sql.DB) {
	createDoorsTableSQL := `CREATE TABLE doors (
		"idDoor" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT NOT NULL,
		"title" TEXT NOT NULL,
		"categoryId" integer NOT NULL,
        "serverId" integer NOT NULL,
        "isMature" integer NOT NULL		
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
		"categoryName" TEXT NOT NULL,
        "isMature" integer NOT NULL
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

func insertDoor(db *sql.DB, code string, title string, category int, server int, isMature int) {
	fmt.Println("Inserting door record ...")
	insertDoorSQL := `INSERT INTO doors(code, title, categoryId, serverId, isMature) VALUES (?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertDoorSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, title, category, server, isMature)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertCategory(db *sql.DB, categoryName string, isMature int) {
	fmt.Println("Inserting category record ...")
	insertCategorySQL := `INSERT INTO categories(categoryName, isMature) VALUES (?, ?)`
	statement, err := db.Prepare(insertCategorySQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(categoryName, isMature)
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

func categoryList(db *sql.DB) []CategoryList {
	rows, err := db.Query(`
    SELECT
        idCategory,
        categoryName
    FROM 
        categories
  `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var categoryList []CategoryList
	for rows.Next() {

		var idCategory int
		var categoryName string

		err := rows.Scan(&idCategory, &categoryName)
		if err != nil {
			log.Fatal(err)
		}

		categoryList = append(categoryList, CategoryList{CategoryId: fmt.Sprint(idCategory), CategoryName: categoryName})
	}
	return categoryList

}

func doorsByCategory(db *sql.DB, cat int) []DoorsList {
	rows, err := db.Query(`
    SELECT DISTINCT
        title as DoorTitle
    FROM 
        doors 
    WHERE
        categoryId = ?
  `, cat)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var doorsList []DoorsList
	for rows.Next() {

		var DoorTitle string

		err := rows.Scan(&DoorTitle)
		if err != nil {
			log.Fatal(err)
		}

		doorsList = append(doorsList, DoorsList{DoorTitle: DoorTitle})
	}
	return doorsList
}
