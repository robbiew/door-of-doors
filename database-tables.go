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
	insertCategory(db, "RPG: Medieval & Fantasy", 0)
	insertCategory(db, "RPG: War, Present/Future", 0)
	insertCategory(db, "Solitaire", 0)
	insertCategory(db, "Space", 0)
	insertCategory(db, "Synchronet Games", 0)
	insertCategory(db, "MBBS Games", 0)
	insertCategory(db, "Trivia", 0)
	insertCategory(db, "Word Games", 0)

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

	insertDoor(db, "LORD", "L.O.R.D", 11, 1, 0)
	insertDoor(db, "LORD", "L.O.R.D II: New World", 11, 1, 0)
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
	insertDoor(db, "LORD", "BladeNaster", 11, 1, 0)
	insertDoor(db, "LORD", "Crossroads", 11, 1, 0)
	insertDoor(db, "LORD", "Tele-Arena", 11, 1, 0)

	// BBS Link
	insertDoor(db, "LORD", "L.O.R.D", 11, 2, 0)
	// Door Party
	insertDoor(db, "LORD", "L.O.R.D", 11, 3, 0)

}
