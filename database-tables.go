package main

import "database/sql"

func createTables(db *sql.DB) {

	// Create Database Tables
	createTitlesTables(db)
	createCategoriesTable(db)
	createServersTable(db)
	createDoorsTable(db)

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

	// INSERT DOOR GAME NAMES
	// db, name, category, category1, category2, isMature
	insertTitle(db, "Drug Lord", 1, 16, 0, 1)
	insertTitle(db, "Bordello", 1, 0, 0, 1)
	insertTitle(db, "Pimp Wars", 1, 0, 0, 1)
	insertTitle(db, "Rancho Nevada", 1, 0, 0, 1)
	insertTitle(db, "Piss-O-Rama", 1, 0, 0, 1)
	insertTitle(db, "Godfather of Crime", 1, 0, 0, 1)

	insertTitle(db, "Gooble Gooble", 2, 0, 0, 0)
	insertTitle(db, "Lemons", 2, 0, 0, 0)
	insertTitle(db, "Chicken Delivery", 2, 0, 0, 0)
	insertTitle(db, "Myst Man", 2, 0, 0, 0)
	insertTitle(db, "BBS*Tris", 2, 0, 0, 0)
	insertTitle(db, "Escape Arcade", 2, 0, 0, 0)

	// INSERT DOOR RECORDS
	// code, title int, server int
	// Gold Mine - Adult
	insertDoor(db, "DRUGLORD", 1, 1)
	insertDoor(db, "BORDELLO", 2, 1)
	insertDoor(db, "PIMPWARS", 3, 1)
	insertDoor(db, "RN", 4, 1)
	insertDoor(db, "PISSORAM", 5, 1)
	insertDoor(db, "GODFATHE", 6, 1)

	// Gold Mine - Arcade & Action
	insertDoor(db, "GOOBLE", 7, 1)
	insertDoor(db, "LEMONS", 8, 1)
	insertDoor(db, "CHICKEN", 9, 1)
	insertDoor(db, "MYSTMAN", 10, 1)
	insertDoor(db, "BBSTRIS", 11, 1)
	insertDoor(db, "ESCAPE", 12, 1)

	//DoorParty
	insertDoor(db, "CHANGEME", 1, 2)

	// // Gold Mine - Board Games
	// insertDoor(db, "4CORNERC", "4 Corner checkers", 3, 0, 0, 1, 0)
	// insertDoor(db, "BOGGLE", "Bubble Boggle", 3, 16, 0, 1, 0)
	// insertDoor(db, "CONCENTR", "Concentration", 3, 0, 0, 1, 0)
	// insertDoor(db, "SEABATTL", "Sea Battle", 3, 0, 0, 1, 0)

	// insertDoor(db, "LORD", "L.O.R.D.", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "L.O.R.D. II", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "The Beast's Domain", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "For Honour", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Death Masters", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Dungeon Master", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Kannons & Katapults", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Arrowbridge I", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Arrowbridge II", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "LEGION RPG", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "King's Crown", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Exitilus", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Usurper", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Adv. Maze II", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Rise to Power", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "BladeMaster", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Crossroads", 11, 0, 0, 1, 0)
	// insertDoor(db, "LORD", "Tele-Arena", 11, 0, 0, 1, 0)

	// // BBS Link
	// insertDoor(db, "LORD", "L.O.R.D.", 11, 0, 0, 2, 0)
	// // Door Party
	// insertDoor(db, "LORD", "L.O.R.D.", 11, 0, 0, 3, 0)
}
