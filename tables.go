package main

import "database/sql"

func createTables(db *sql.DB) {

	// Create Database Tables
	createDoorsTable(db)
	createCategoriesTable(db)
	createServersTable(db)

	// INSERT TOP-LEVEL CATEGORY RECORDS
	insertCategory(db, "Adult: Sex, Drugs & Crime", 1)
	insertCategory(db, "Arcade & Action", 0)
	insertCategory(db, "Board Games", 0)
	insertCategory(db, "Cards, Dice & Casino", 0)
	insertCategory(db, "DOS Text-Based", 0)
	insertCategory(db, "Leisure & Sports", 0)
	insertCategory(db, "Other", 0)
	insertCategory(db, "Multiplayer", 0)
	insertCategory(db, "Puzzle", 0)
	insertCategory(db, "Realtime Multiplayer", 0)
	insertCategory(db, "RPGs: Medieval & Fantasy", 0)
	insertCategory(db, "Solitaire", 0)
	insertCategory(db, "Space", 0)
	insertCategory(db, "Synchronet Games", 0)
	insertCategory(db, "Trivia", 0)
	insertCategory(db, "Word Games", 0)

	// INSERT AVAILABLE DOOR SERVERS
	insertServer(db, "Gold Mine")
	insertServer(db, "BBS Link")
	insertServer(db, "Door Party")

	// INSERT RECORDS
	insertDoor(db, "LORD", "Legend of the Red Dragon", 11, 1, 0)
	insertDoor(db, "LORD", "Legend of the Red Dragon", 11, 2, 0)
	insertDoor(db, "LORD", "Legend of the Red Dragon", 11, 3, 0)

}
