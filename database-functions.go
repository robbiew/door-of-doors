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

		clearScreen()

		fmt.Print("data.db doesn't exist - creating...\r\n")
		db, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File
		defer db.Close()                          // Defer Closing the database

		// Create all tables, insert records
		createTables(db)
	}
	// database exists, so do nothing
}

func createTitlesTables(db *sql.DB) {
	createTitlesTableSQL := `CREATE TABLE titles (
		"idTitle" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"titleName" TEXT NOT NULL,
		"categoryId" integer NOT NULL,
		"category2" integer NOT NULL,
		"category3" integer NOT NULL,
		"isMature" integer NOT NULL,
		"desc" TEXT NOT NULL,
		"year" TEXT NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Print("Creating TITLES table...\r\n")
	statement, err := db.Prepare(createTitlesTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	// fmt.Println("TITLES table created.")
}

func createDoorsTable(db *sql.DB) {
	createDoorsTableSQL := `CREATE TABLE doors (
		"idDoor" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT NOT NULL,
		"titleId" integer NOT NULL,
        "serverId" integer NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Print("Creating DOORS table...\r\n")
	statement, err := db.Prepare(createDoorsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	// fmt.Println("DOORS table created.")
}

func createCategoriesTable(db *sql.DB) {
	createCategoriesTableSQL := `CREATE TABLE categories (
		"idCategory" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"categoryName" TEXT NOT NULL,
		"categoryCode" TEXT NOT NULL,
        "isMature" integer NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Print("Creating CATEGORIES table...\r\n")
	statement, err := db.Prepare(createCategoriesTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	// fmt.Println("CATEGORIES table created.")
}

func createServersTable(db *sql.DB) {
	createServersTableSQL := `CREATE TABLE servers (
		"idServer" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"serverName" TEXT NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Print("Creating SERVERS table...\r\n")
	statement, err := db.Prepare(createServersTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	// fmt.Println("SERVERS table created.")
}

func insertTitle(db *sql.DB, titleName string, categoryId int, category2 int, category3 int, isMature int, desc string, year string) {
	// fmt.Println("Inserting TITLES records...")
	insertTitleSQL := `INSERT INTO titles(titleName, categoryId, category2, category3, isMature, desc, year) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertTitleSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(titleName, categoryId, category2, category3, isMature, desc, year)
	if err != nil {
		log.Println(err.Error())
	}
}

func insertDoor(db *sql.DB, code string, title int, server int) {
	// fmt.Println("Inserting DOORS records...")
	insertDoorSQL := `INSERT INTO doors(code, titleId, serverId) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertDoorSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(code, title, server)
	if err != nil {
		log.Println(err.Error())
	}
}

func insertCategory(db *sql.DB, categoryName string, categoryCode string, isMature int) {
	// fmt.Println("Inserting CATEGORIES records...")
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
	// fmt.Println("Inserting SERVERS records...")
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
        titleName AS DoorTitle 
    FROM 
        titles
	INNER JOIN doors ON doors.titleId = idTitle 
    WHERE
        categoryId = ? OR category2 = ? OR category3 = ? 
	ORDER BY
		DoorTitle
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

func doorByServer(db *sql.DB) []ServersList {
	rows, err := db.Query(`
    SELECT
        titles.titleName AS title,
		servers.serverName as serverName,
		titles.desc,
		titles.year,
		doors.code
    FROM
        titles
	LEFT JOIN doors ON doors.titleId = titles.idTitle
	LEFT JOIN servers ON servers.idServer = doors.serverId
    WHERE
        titles.titleName = ?
	ORDER BY
		title
  `, currTitle)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var serverList []ServersList
	for rows.Next() {

		var title string
		var serverName string
		var desc string
		var year string
		var code string

		err := rows.Scan(&title, &serverName, &desc, &year, &code)
		if err != nil {
			log.Fatal(err)
		}

		serverList = append(serverList, ServersList{DoorTitle: title, ServerName: serverName, Desc: desc, Year: year, DoorCode: code})
	}
	return serverList

}
