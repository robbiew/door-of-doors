package main

import "database/sql"

func createTables(db *sql.DB) {

	// Create Database Tables
	createDoorsTable(db)
	createCategoriesTable(db)
	createServersTable(db)

	// INSERT TOP-LEVEL CATEGORY RECORDS
	insertCategory(db, "Adult: Sex, Drugs & Crime", "ADULT", 1)
	insertCategory(db, "Arcade & Action", "ARCDE", 0)
	insertCategory(db, "Board Games", "BOARD", 0)
	insertCategory(db, "Cards, Dice & Casino", "CARDS", 0)
	insertCategory(db, "DOS Text-Based", "DOS", 0)
	insertCategory(db, "Leisure & Sports", "SPORT", 0)
	insertCategory(db, "Other", "OTHR", 0)
	insertCategory(db, "Multiplayer", "MULTI", 0)
	insertCategory(db, "Puzzle", "PUZZL", 0)
	insertCategory(db, "Realtime Multiplayer", "MULTI", 0)
	insertCategory(db, "RPG: Medieval & Fantasy", "RPG1", 0)
	insertCategory(db, "RPG: War, Present/Future", "RPG2", 0)
	insertCategory(db, "Solitaire", "SOLIT", 0)
	insertCategory(db, "Space", "SPACE", 0)
	insertCategory(db, "Synchronet Games", "SYNC", 0)
	insertCategory(db, "MBBS Games", "MBBS", 0)
	insertCategory(db, "Trivia", "TRIV", 0)
	insertCategory(db, "Word Games", "WORD", 0)

	// INSERT AVAILABLE DOOR SERVERS
	insertServer(db, "Gold Mine")
	insertServer(db, "BBS Link")
	insertServer(db, "Door Party")

	// INSERT RECORDS
	// Gold Mine
	insertDoor(db, "DRUGLORD", "Drug Lord", 1, 1, 0)
	insertDoor(db, "BORDELLO", "Bordello", 1, 1, 0)
	insertDoor(db, "PIMPWARS", "Pimp Wars", 1, 1, 0)
	insertDoor(db, "RN", "Rancho Nevada", 1, 1, 0)
	insertDoor(db, "PISSORAM", "Piss-o-Rama", 1, 1, 0)
	insertDoor(db, "GODFATHE", "Godfather of Crime", 1, 1, 0)

	insertDoor(db, "LORD", "L.O.R.D.", 11, 1, 0)
	insertDoor(db, "LORD", "L.O.R.D. II", 11, 1, 0)
	insertDoor(db, "LORD", "The Beast's Domain", 11, 1, 0)
	insertDoor(db, "LORD", "For Honour", 11, 1, 0)
	insertDoor(db, "LORD", "Death Masters", 11, 1, 0)
	insertDoor(db, "LORD", "Dungeon Master", 11, 1, 0)
	insertDoor(db, "LORD", "Kannons & Katapults", 11, 1, 0)
	insertDoor(db, "LORD", "Arrowbridge I", 11, 1, 0)
	insertDoor(db, "LORD", "Arrowbridge II", 11, 1, 0)
	insertDoor(db, "LORD", "LEGION RPG", 11, 1, 0)
	insertDoor(db, "LORD", "King's Crown", 11, 1, 0)
	insertDoor(db, "LORD", "Exitilus", 11, 1, 0)
	insertDoor(db, "LORD", "Usurper", 11, 1, 0)
	insertDoor(db, "LORD", "Adv. Maze II", 11, 1, 0)
	insertDoor(db, "LORD", "Rise to Power", 11, 1, 0)
	insertDoor(db, "LORD", "BladeMaster", 11, 1, 0)
	insertDoor(db, "LORD", "Crossroads", 11, 1, 0)
	insertDoor(db, "LORD", "Tele-Arena", 11, 1, 0)

	// BBS Link
	insertDoor(db, "LORD", "L.O.R.D.", 11, 2, 0)
	// Door Party
	insertDoor(db, "LORD", "L.O.R.D.", 11, 3, 0)
}
