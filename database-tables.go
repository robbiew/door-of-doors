package main

import "database/sql"

func createTables(db *sql.DB) {

	// Create Database Tables
	createTitlesTables(db)
	createCategoriesTable(db)
	createServersTable(db)
	createDoorsTable(db)

	// INSERT TOP-LEVEL CATEGORY RECORDS
	insertCategory(db, "Adult: Sex, Drugs & Crime", "ADULT", 1) // 1
	insertCategory(db, "Arcade & Action", "ARCDE", 0)           // 2
	insertCategory(db, "Board Games", "BOARD", 0)               // 3
	insertCategory(db, "Cards, Dice & Casino", "CARDS", 0)      // 4
	insertCategory(db, "DOS Text-Based", "DOS", 0)              // 5
	insertCategory(db, "Leisure & Sports", "SPORT", 0)          // 6
	insertCategory(db, "MBBS Games", "MBBS", 0)                 // 7
	insertCategory(db, "Multiplayer", "MULTI", 0)               // 8
	insertCategory(db, "Other", "OTHR", 0)                      // 9
	insertCategory(db, "Puzzle", "PUZZL", 0)                    // 10
	insertCategory(db, "RPG: Medieval & Fantasy", "RPG1", 0)    // 11
	insertCategory(db, "RPG: War, Present/Future", "RPG2", 0)   // 12
	insertCategory(db, "Realtime Multiplayer", "MULTI", 0)      // 13
	insertCategory(db, "Solitaire", "SOLIT", 0)                 // 14
	insertCategory(db, "Space", "SPACE", 0)                     // 15
	insertCategory(db, "Synchronet Games", "SYNC", 0)           // 16
	insertCategory(db, "Trivia", "TRIV", 0)                     // 17
	insertCategory(db, "Word Games", "WORD", 0)                 // 18

	// INSERT AVAILABLE DOOR SERVERS
	insertServer(db, "Gold Mine")  // 1
	insertServer(db, "BBS Link")   // 2
	insertServer(db, "Door Party") // 3

	// INSERT DOOR GAME NAMES
	/*id */
	/*1  */
	insertTitle(db, "Drug Lord", 1, 16, 0, 1, "DrugLord is a game similar to Dope Wars and its derivatives.", "1991")
	/*2 */ insertTitle(db, "Bordello", 1, 0, 0, 1, "Run your own whore house and fight against other players. Steal money, whores, destroy other houses, fight with opponents thugs, and much more!", "2004")
	/*3 */ insertTitle(db, "Pimp Wars", 1, 0, 0, 1, "Pimp Wars is a game where players assume the role of a Pimp and work the streets through 'extra-cirricular' activities.", "1990")
	/*4 */ insertTitle(db, "Rancho Nevada", 1, 0, 0, 1, "What once was the greatest Adult Door ever is now the greatest BBS Door of any kind! This final version has a real goal (beyond simply reaching level 100)!", "1998")
	/*5 */ insertTitle(db, "Piss-O-Rama", 1, 0, 0, 1, "", "")
	/*6 */ insertTitle(db, "Godfather of Crime", 1, 0, 0, 1, "", "")
	/*7 */ insertTitle(db, "Gooble Gooble", 2, 16, 0, 0, "", "")
	/*8 */ insertTitle(db, "Lemons", 2, 16, 0, 0, "", "")
	/*9 */ insertTitle(db, "Chicken Delivery", 2, 16, 0, 0, "", "")
	/*10*/ insertTitle(db, "Myst Man", 2, 0, 0, 0, "", "")
	/*11*/ insertTitle(db, "BBS*Tris", 2, 0, 0, 0, "", "")
	/*12*/ insertTitle(db, "Escape Arcade", 2, 0, 0, 0, "", "")
	/*13*/ insertTitle(db, "4 Corner checkers", 3, 0, 0, 0, "", "")
	/*14*/ insertTitle(db, "Bubble Boggle", 3, 16, 0, 0, "", "")
	/*15*/ insertTitle(db, "Concentration", 3, 0, 0, 0, "", "")
	/*16*/ insertTitle(db, "Sea Battle", 3, 0, 0, 0, "", "")
	/*17*/ insertTitle(db, "3 Pair Poker", 4, 0, 0, 0, "", "")
	/*18*/ insertTitle(db, "Acey Deucy", 4, 0, 0, 0, "", "")
	/*19*/ insertTitle(db, "Baseball Dice", 4, 0, 0, 0, "", "")
	/*20*/ insertTitle(db, "Blackjack", 4, 0, 0, 0, "", "")
	/*21*/ insertTitle(db, "Box Poker", 4, 0, 0, 0, "", "")
	/*22*/ insertTitle(db, "Dice Warz", 4, 16, 0, 0, "", "")
	/*23*/ insertTitle(db, "Dice Warz II", 4, 16, 0, 0, "", "")
	/*24*/ insertTitle(db, "Domain Poker", 4, 16, 0, 0, "", "")
	/*25*/ insertTitle(db, "Milliway's Casino", 4, 16, 0, 0, "", "")
	/*26*/ insertTitle(db, "Synchronet Blackjack", 4, 16, 0, 0, "", "")
	/*27*/ insertTitle(db, "Star Stocks", 4, 16, 0, 0, "", "")
	/*28*/ insertTitle(db, "Titan Slots", 4, 0, 0, 0, "", "")
	/*29*/ insertTitle(db, "Las Vegas Keno", 4, 0, 0, 0, "", "")
	/*30*/ insertTitle(db, "Las Vegas Blackjack", 4, 0, 0, 0, "", "")
	/*31*/ insertTitle(db, "Las Vegas Craps", 4, 0, 0, 0, "", "")
	/*32*/ insertTitle(db, "Las Vegas Video Poker", 4, 0, 0, 0, "", "")
	/*33*/ insertTitle(db, "Las Vegas Slots", 4, 0, 0, 0, "", "")
	/*34*/ insertTitle(db, "Las Vegas Casino Manager", 4, 0, 0, 0, "", "")
	/*35*/ insertTitle(db, "5 Card Stud", 4, 0, 0, 0, "", "")
	/*36*/ insertTitle(db, "Zork I", 5, 0, 0, 0, "", "")
	/*37*/ insertTitle(db, "Zork II", 5, 0, 0, 0, "", "")
	/*38*/ insertTitle(db, "Zork III", 5, 0, 0, 0, "", "")
	/*39*/ insertTitle(db, "Beyond Zork", 5, 0, 0, 0, "", "")
	/*40*/ insertTitle(db, "Zork: Undiscovered Underground", 5, 0, 0, 0, "", "")
	/*41*/ insertTitle(db, "Arthur: Quest for Excalibur", 5, 0, 0, 0, "", "")
	/*42*/ insertTitle(db, "A Mind Forever Voyaging", 5, 0, 0, 0, "", "")
	/*43*/ insertTitle(db, "Ballyhoo", 5, 0, 0, 0, "", "")
	/*44*/ insertTitle(db, "Border Zone", 5, 0, 0, 0, "", "")
	/*45*/ insertTitle(db, "Bureaucracy", 5, 0, 0, 0, "", "")
	/*46*/ insertTitle(db, "Cutthroats", 5, 0, 0, 0, "", "")
	/*47*/ insertTitle(db, "Enchanter", 5, 0, 0, 0, "", "")
	/*48*/ insertTitle(db, "Hitchhiker's Guide to the Galaxy", 5, 0, 0, 0, "", "")
	/*49*/ insertTitle(db, "Hollywood Hijynx", 5, 0, 0, 0, "", "")
	/*50*/ insertTitle(db, "Infidel", 5, 0, 0, 0, "", "")
	/*51*/ insertTitle(db, "Journey: The Quest Begins", 5, 0, 0, 0, "", "")
	/*52*/ insertTitle(db, "The Lurking Horror", 5, 0, 0, 0, "", "")
	/*53*/ insertTitle(db, "Leather Goddesses of Phobos", 5, 0, 0, 0, "", "")
	/*54*/ insertTitle(db, "Sherlock: Riddle of the Crown", 5, 0, 0, 0, "", "")
	/*55*/ insertTitle(db, "Starcross", 5, 0, 0, 0, "", "")
	/*56*/ insertTitle(db, "Stationfall", 5, 0, 0, 0, "", "")
	/*57*/ insertTitle(db, "Shogun", 5, 0, 0, 0, "", "")
	/*58*/ insertTitle(db, "Moonmist", 5, 0, 0, 0, "", "")
	/*59*/ insertTitle(db, "Nord and Bert", 5, 0, 0, 0, "", "")
	/*60*/ insertTitle(db, "Planetfall", 5, 0, 0, 0, "", "")
	/*61*/ insertTitle(db, "Plundered", 5, 0, 0, 0, "", "")
	/*62*/ insertTitle(db, "Seastalker", 5, 0, 0, 0, "", "")
	/*63*/ insertTitle(db, "Spellbreaker", 5, 0, 0, 0, "", "")
	/*64*/ insertTitle(db, "Suspect", 5, 0, 0, 0, "", "")
	/*65*/ insertTitle(db, "Suspended", 5, 0, 0, 0, "", "")
	/*66*/ insertTitle(db, "Trinity", 5, 0, 0, 0, "", "")
	/*67*/ insertTitle(db, "Wishbringer", 5, 0, 0, 0, "", "")
	/*68*/ insertTitle(db, "Witness", 5, 0, 0, 0, "", "")
	/*69*/ insertTitle(db, "Forbidden Planet", 5, 0, 0, 0, "", "")
	/*70*/ insertTitle(db, "Kyrandia", 5, 7, 0, 0, "", "")

	// INSERT DOOR RECORDS FOR EACH SERVER
	// code, title int, server int

	// Gold Mine
	// Adult
	insertDoor(db, "DRUGLORD", 1, 1)
	insertDoor(db, "BORDELLO", 2, 1)
	insertDoor(db, "PIMPWARS", 3, 1)
	insertDoor(db, "RN", 4, 1)
	insertDoor(db, "PISSORAM", 5, 1)
	insertDoor(db, "GODFATHE", 6, 1)
	// Arcade & Action
	insertDoor(db, "GOOBLE", 7, 1)
	insertDoor(db, "LEMONS", 8, 1)
	insertDoor(db, "CHICKEN", 9, 1)
	insertDoor(db, "MYSTMAN", 10, 1)
	insertDoor(db, "BBSTRIS", 11, 1)
	insertDoor(db, "ESCAPE", 12, 1)
	// Board Games
	insertDoor(db, "4CORNERC", 13, 1)
	insertDoor(db, "BOGGLE", 14, 1)
	insertDoor(db, "CONCENTR", 15, 1)
	insertDoor(db, "RSEABATTL", 16, 1)
	// Cards & Casino
	insertDoor(db, "3PAIRPOK", 17, 1)
	insertDoor(db, "ACEYDEUC", 18, 1)
	insertDoor(db, "BBDICE", 19, 1)
	insertDoor(db, "BLACKJAC", 20, 1)
	insertDoor(db, "BOXPOKER", 21, 1)
	insertDoor(db, "DICEWARZ", 22, 1)
	insertDoor(db, "DICEWAR2", 23, 1)
	insertDoor(db, "DPOKER", 24, 1)
	insertDoor(db, "MILLIWAY", 25, 1)
	insertDoor(db, "SBJ", 26, 1)
	insertDoor(db, "STARSTOX", 27, 1)
	insertDoor(db, "TITANSLO", 28, 1)
	insertDoor(db, "LVKENO", 29, 1)
	insertDoor(db, "LVBJACK", 30, 1)
	insertDoor(db, "LVCRAPS", 31, 1)
	insertDoor(db, "LVPOKER", 32, 1)
	insertDoor(db, "LVSLOTS", 33, 1)
	insertDoor(db, "LVCASINO", 34, 1)
	insertDoor(db, "FIVECARD", 35, 1)
	// DOS TEXT BASED
	insertDoor(db, "ZORK1", 36, 1)
	insertDoor(db, "ZORK2", 37, 1)
	insertDoor(db, "ZORK3", 38, 1)
	insertDoor(db, "BEYONDZ", 39, 1)
	insertDoor(db, "ZTUU", 40, 1)
	insertDoor(db, "ARTHUR", 41, 1)
	insertDoor(db, "AMFV", 42, 1)
	insertDoor(db, "BALLYHOO", 43, 1)
	insertDoor(db, "BORDERZO", 44, 1)
	insertDoor(db, "BUREAUCR", 45, 1)
	insertDoor(db, "CUTTHROA", 46, 1)
	insertDoor(db, "ENCHANTE", 47, 1)
	insertDoor(db, "HHGTG", 48, 1)
	insertDoor(db, "HOLLYWOO", 49, 1)
	insertDoor(db, "INFIDEL", 50, 1)
	insertDoor(db, "JOURNEY", 51, 1)
	insertDoor(db, "LURKING", 52, 1)
	insertDoor(db, "LEATHER", 53, 1)
	insertDoor(db, "SHERLOCK", 54, 1)
	insertDoor(db, "SORCERER", 55, 1)
	insertDoor(db, "STARCROS", 56, 1)
	insertDoor(db, "STATION", 57, 1)
	insertDoor(db, "SHOGUN", 58, 1)
	insertDoor(db, "MOONMIST", 59, 1)
	insertDoor(db, "NORDANDB", 60, 1)
	insertDoor(db, "PLANETFA", 61, 1)
	insertDoor(db, "PLUNDERE", 62, 1)
	insertDoor(db, "SEASTALK", 63, 1)
	insertDoor(db, "SPELLBRE", 64, 1)
	insertDoor(db, "SUSPECT", 65, 1)
	insertDoor(db, "SUSPENDE", 66, 1)
	insertDoor(db, "TRINITY", 67, 1)
	insertDoor(db, "WISHBRIN", 68, 1)
	insertDoor(db, "WITNESS", 69, 1)
	insertDoor(db, "PLANET", 70, 1)

	//DoorParty
	insertDoor(db, "CHANGEME", 1, 2)

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
