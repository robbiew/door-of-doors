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
		log.Print(err.Error())

		file, err := os.Create("data.db") // Create SQLite file
		if err != nil {
			log.Println(err.Error())
		}
		file.Close()
		log.Println("data.db doesn't exist - created")
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
		"category2" integer NOT NULL,
		"category3" integer NOT NULL,
        "serverId" integer NOT NULL,
        "isMature" integer NOT NULL		
	  );` // SQL Statement for Create Table

	log.Println("Create doors table...")
	statement, err := db.Prepare(createDoorsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("doors table created")
}

func createCategoriesTable(db *sql.DB) {
	createCategoriesTableSQL := `CREATE TABLE categories (
		"idCategory" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"categoryName" TEXT NOT NULL,
		"categoryCode" TEXT NOT NULL,
        "isMature" integer NOT NULL
	  );` // SQL Statement for Create Table

	log.Println("Create categories table...")
	statement, err := db.Prepare(createCategoriesTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("categories table created")
}

func createServersTable(db *sql.DB) {
	createServersTableSQL := `CREATE TABLE servers (
		"idServer" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"serverName" TEXT NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Println("Create servers table...")
	statement, err := db.Prepare(createServersTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("servers table created")
}

func insertDoor(db *sql.DB, code string, title string, category int, category2 int, category3 int, server int, isMature int) {
	log.Println("Inserting door record ...")
	insertDoorSQL := `INSERT INTO doors(code, title, categoryId, category2, category3, serverId, isMature) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertDoorSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(code, title, category, category2, category3, server, isMature)
	if err != nil {
		log.Println(err.Error())
	}
}

func insertCategory(db *sql.DB, categoryName string, categoryCode string, isMature int) {
	log.Println("Inserting category record ...")
	insertCategorySQL := `INSERT INTO categories(categoryName, categoryCode, isMature) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertCategorySQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(categoryName, categoryCode, isMature)
	if err != nil {
		log.Println(err.Error())
	}
}

func insertServer(db *sql.DB, serverName string) {
	log.Println("Inserting server record ...")
	insertServerSQL := `INSERT INTO servers(serverName) VALUES (?)`
	statement, err := db.Prepare(insertServerSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(serverName)
	if err != nil {
		log.Println(err.Error())
	}
}

func categoryList(db *sql.DB) []CategoryList {
	rows, err := db.Query(`
    SELECT
        idCategory,
        categoryName,
		categoryCode
    FROM 
        categories
	ORDER BY
		categoryName
  `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var categoryList []CategoryList
	for rows.Next() {

		var idCategory int
		var categoryName string
		var categoryCode string

		err := rows.Scan(&idCategory, &categoryName, &categoryCode)
		if err != nil {
			log.Fatal(err)
		}

		categoryList = append(categoryList, CategoryList{CategoryId: idCategory, CategoryName: categoryName, CategoryCode: categoryCode})
	}
	return categoryList

}

func doorsByCategory(db *sql.DB, realCat int) []DoorsList {
	rows, err := db.Query(`
    SELECT DISTINCT
        title as DoorTitle
    FROM 
        doors 
    WHERE
        categoryId = ? OR category2 = ? OR category3 = ? 
	ORDER BY
		title
  `, realCat, realCat, realCat)

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

func doorByServer(db *sql.DB) []ServerList {
	rows, err := db.Query(`
    SELECT
        title,
		serverName
    FROM
        doors
	INNER JOIN servers ON servers.IdServer = serverId 
    WHERE
        title = ?
	ORDER BY
		title
  `, currTitle)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var serverList []ServerList
	for rows.Next() {

		var title string
		var serverName string

		err := rows.Scan(&title, &serverName)
		if err != nil {
			log.Fatal(err)
		}

		serverList = append(serverList, ServerList{DoorTitle: title, ServerName: serverName})
	}
	return serverList

}
