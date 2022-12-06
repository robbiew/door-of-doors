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
		"adult" integer NOT NULL,
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
        "adult" integer NOT NULL
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
		"serverName" TEXT NOT NULL,
		"serverDesc" TEXT NOT NULL
	  );` // SQL Statement for Create Table

	fmt.Print("Creating SERVERS table...\r\n")
	statement, err := db.Prepare(createServersTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	// fmt.Println("SERVERS table created.")
}

func insertTitle(db *sql.DB, titleName string, categoryId int, category2 int, category3 int, adult int, desc string, year string) {
	// fmt.Println("Inserting TITLES records...")
	insertTitleSQL := `INSERT INTO titles(titleName, categoryId, category2, category3, adult, desc, year) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertTitleSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(titleName, categoryId, category2, category3, adult, desc, year)
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

func insertCategory(db *sql.DB, categoryName string, categoryCode string, adult int) {
	// fmt.Println("Inserting CATEGORIES records...")
	insertCategorySQL := `INSERT INTO categories(categoryName, categoryCode, adult) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertCategorySQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(categoryName, categoryCode, adult)
	if err != nil {
		log.Println(err.Error())
	}
}

func insertServer(db *sql.DB, serverName string, serverDesc string) {
	// fmt.Println("Inserting SERVERS records...")
	insertServerSQL := `INSERT INTO servers(serverName, serverDesc) VALUES (?, ?)`
	statement, err := db.Prepare(insertServerSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(serverName, serverDesc)
	if err != nil {
		log.Println(err.Error())
	}
}

func categoryList(db *sql.DB) []CategoryList {
	var rows *sql.Rows
	var err error

	if C.Adult == "0" {
		rows, err = db.Query(`
    SELECT
        idCategory,
        categoryName,
		categoryCode,
		adult
    FROM 
        categories
	WHERE
		adult = 0
	ORDER BY
		categoryName
  `)
		if err != nil {
			log.Fatal(err)
		}
	}

	if C.Adult == "1" {

		rows, err = db.Query(`
    SELECT
        idCategory,
        categoryName,
		categoryCode,
		adult
    FROM 
        categories
	ORDER BY
		categoryName
  `)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()
	var categories []CategoryList

	for rows.Next() {

		var idCategory int
		var categoryName string
		var categoryCode string
		var adult string

		err := rows.Scan(&idCategory, &categoryName, &categoryCode, &adult)
		if err != nil {
			log.Fatal(err)
		}
		if excludeEmptyCat(db, idCategory) > 0 {
			categories = append(categories, CategoryList{CategoryId: idCategory, CategoryName: categoryName, CategoryCode: categoryCode, CategoryAdult: adult})
		}
	}
	return categories

}

func doorsByCategory(db *sql.DB, realCat int) []DoorsList {

	var rows *sql.Rows
	var err error

	if C.Adult == "0" {
		rows, err = db.Query(`
    SELECT DISTINCT
        titleName AS DoorTitle, 
		desc AS DoorDesc,
		year AS DoorYear,
		adult as DoorAdult
    FROM 
        titles
	INNER JOIN doors ON doors.titleId = idTitle 
    WHERE
        categoryId = ? OR category2 = ? OR category3 = ? 
	AND 
		adult = 0
	ORDER BY
		DoorTitle

  `, realCat, realCat, realCat)

	}
	if C.Adult == "1" {
		rows, err = db.Query(`
    SELECT DISTINCT
        titleName AS DoorTitle, 
		desc AS DoorDesc,
		year AS DoorYear,
		adult AS DoorAdult 
    FROM 
        titles
	INNER JOIN doors ON doors.titleId = idTitle 
    WHERE
        categoryId = ? OR category2 = ? OR category3 = ?
	ORDER BY
		DoorTitle
  `, realCat, realCat, realCat)

	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var doors []DoorsList

	for rows.Next() {

		var DoorTitle string
		var DoorDesc string
		var DoorYear string
		var DoorAdult string
		err := rows.Scan(&DoorTitle, &DoorDesc, &DoorYear, &DoorAdult)
		if err != nil {
			log.Fatal(err)
		}

		doors = append(doors, DoorsList{DoorTitle: DoorTitle, DoorDesc: DoorDesc, DoorYear: DoorYear, DoorAdult: DoorAdult})
	}
	return doors
}

func doorByServer(db *sql.DB) []ServersList {
	rows, err := db.Query(`
    SELECT
        titles.titleName AS title,
		servers.serverName as serverName,
		servers.serverDesc as serverDesc,
		titles.desc,
		titles.year,
		doors.code,
		doors.serverId
		
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

	var serversList []ServersList
	for rows.Next() {

		var title string
		var serverName string
		var serverDesc string
		var desc string
		var year string
		var code string
		var serverId string

		err := rows.Scan(&title, &serverName, &serverDesc, &desc, &year, &code, &serverId)
		if err != nil {
			log.Fatal(err)
		}

		serversList = append(serversList, ServersList{DoorTitle: title, ServerName: serverName, ServerDesc: serverDesc, Desc: desc, Year: year, DoorCode: code, ServerId: serverId})
	}
	return serversList

}

func doorCount(db *sql.DB, server int) int {

	rows, err := db.Query(`
    SELECT 
		COUNT(*)
    FROM 
        doors
    WHERE
        serverId = ? 

  `, server)

	if err != nil {
		log.Fatal(err)
	}

	var count int
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

	}
	return count

}

func excludeEmptyCat(db *sql.DB, category int) int {
	rows, err := db.Query(`
    SELECT 
		COUNT(*)
    FROM 
        titles
    WHERE
        categoryId = ? 
  `, category)

	if err != nil {
		log.Fatal(err)
	}

	var count int
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

	}
	return count

}
