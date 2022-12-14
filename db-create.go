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
	insertCategory(db, "Medieval & Fantasy", "RPG1", 0)         // 11
	insertCategory(db, "War, Present/Future", "RPG2", 0)        // 12
	insertCategory(db, "Realtime Multiplayer", "MULTI", 0)      // 13
	insertCategory(db, "Solitaire", "SOLIT", 0)                 // 14
	insertCategory(db, "Space", "SPACE", 0)                     // 15
	insertCategory(db, "Synchronet Games", "SYNC", 0)           // 16
	insertCategory(db, "Trivia", "TRIV", 0)                     // 17
	insertCategory(db, "Word Games", "WORD", 0)                 // 18
	insertCategory(db, "BBS, Hack and Crack", "BBS", 0)         // 19

	// INSERT AVAILABLE DOOR SERVERS
	insertServer(db, "Gold Mine", "Gold Mine is a community door server managed by aLPHA. You can find more information at https://goldminebbs.com")                                                        // 1                                                                                                                                       // 1
	insertServer(db, "BBSlink", "BBSlink was created in early 2014 to allow sysops to easily add door games to their BBS.")                                                                                 // 2
	insertServer(db, "Door Party", "Door Party offers either one-menu-for-all game access, or direct door access which makes it seem like you're hosting the games, but with DoorParty's large user base!") // 3                                                                                // 3

	// INSERT DOOR GAME NAMES
	/*id */
	/*1 */
	insertTitle(db, "Drug Lord", 1, 16, 0, 1, "DrugLord is a game similar to Dope Wars and its derivatives.", "1991")
	/*2 */ insertTitle(db, "Bordello", 1, 0, 0, 1, "Run your own whore house and fight against other players. Steal money, whores, destroy other houses, fight with opponents thugs, and much more!", "2004")
	/*3 */ insertTitle(db, "Pimp Wars", 1, 0, 0, 1, "Pimp Wars is a game where players assume the role of a Pimp and work the streets through 'extra-cirricular' activities.", "1990")
	/*4 */ insertTitle(db, "Rancho Nevada", 1, 0, 0, 1, "What once was the greatest Adult Door ever is now the greatest BBS Door of any kind! This final version has a real goal (beyond simply reaching level 100)!", "1998")
	/*5 */ insertTitle(db, "Piss-O-Rama", 1, 0, 0, 1, "This is the Second version of the Piss-O-Rama, a favorite among SysOps who aren't afraid to let people see a bad word on their board.", "1991")
	/*6 */ insertTitle(db, "Godfather of Crime", 1, 0, 0, 1, "A rip-roaring, knock-em-dead ANSI roleplaying door game of the 1930s New York organized Crime Scene!", "1995")
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
	/*71*/ insertTitle(db, "Baseball Dice", 6, 4, 0, 0, "", "")
	/*72*/ insertTitle(db, "Fat Fish", 6, 16, 0, 0, "", "")
	/*73*/ insertTitle(db, "GuTTeRBoWL", 6, 0, 0, 0, "", "")
	/*74*/ insertTitle(db, "Camptown Races", 6, 0, 0, 0, "", "")
	/*75*/ insertTitle(db, "Synkroban", 9, 16, 0, 0, "", "")
	/*76*/ insertTitle(db, "Thirstyville", 9, 16, 0, 0, "", "")
	/*77*/ insertTitle(db, "Port Victoria", 9, 0, 0, 0, "", "")
	/*78*/ insertTitle(db, "Oregon Trail", 9, 0, 0, 0, "", "")
	/*79*/ insertTitle(db, "Sim BBS", 19, 0, 0, 0, "", "")
	/*80*/ insertTitle(db, "Synchronet Minesweeper", 9, 16, 0, 0, "", "")
	/*81*/ insertTitle(db, "Uber Blox", 10, 16, 0, 0, "", "")
	/*82*/ insertTitle(db, "Synchrotetris", 16, 0, 0, 0, "", "")
	/*83*/ insertTitle(db, "Stacktion", 16, 0, 0, 0, "", "")
	/*84*/ insertTitle(db, "L.O.R.D.", 11, 16, 0, 0, "", "")
	/*85*/ insertTitle(db, "L.O.R.D. II", 11, 16, 0, 0, "", "")
	/*86*/ insertTitle(db, "The Beast's Domain", 11, 0, 0, 0, "", "")
	/*87*/ insertTitle(db, "For Honour", 11, 0, 0, 0, "", "")
	/*88*/ insertTitle(db, "Death Masters", 11, 0, 0, 0, "", "")
	/*89*/ insertTitle(db, "DoorMUD", 11, 0, 0, 0, "", "")
	/*90*/ insertTitle(db, "Dungeon Master", 11, 0, 0, 0, "", "")
	/*91*/ insertTitle(db, "Kannons & Katapults", 11, 0, 0, 0, "", "")
	/*92*/ insertTitle(db, "Arrowbridge I", 11, 0, 0, 0, "", "")
	/*93*/ insertTitle(db, "Arrowbridge II", 11, 0, 0, 0, "", "")
	/*94*/ insertTitle(db, "LEGION", 11, 0, 0, 0, "", "")
	/*95*/ insertTitle(db, "King's Crown", 11, 0, 0, 0, "", "")
	/*96*/ insertTitle(db, "Exitilus", 11, 0, 0, 0, "", "")
	/*97*/ insertTitle(db, "Usurper", 11, 0, 0, 0, "", "")
	/*98*/ insertTitle(db, "Adv. Maze II", 11, 0, 0, 0, "", "")
	/*99*/ insertTitle(db, "Rise to Power", 11, 0, 0, 0, "", "")
	/*100*/ insertTitle(db, "BladeMaster", 11, 7, 0, 0, "", "")
	/*101*/ insertTitle(db, "Crossroads", 11, 7, 0, 0, "", "")
	/*102*/ insertTitle(db, "Tele-Arena", 11, 7, 0, 0, "", "")
	/*103*/ insertTitle(db, "Colonies", 12, 0, 0, 0, "", "")
	/*104*/ insertTitle(db, "Atlantis", 12, 0, 0, 0, "", "")
	/*105*/ insertTitle(db, "Operation Overkill II", 12, 0, 0, 0, "", "")
	/*106*/ insertTitle(db, "Darkness 2.0", 12, 0, 0, 0, "", "")
	/*107*/ insertTitle(db, "Jedi Knight", 12, 0, 0, 0, "", "")
	/*108*/ insertTitle(db, "Barren Realms Elite", 12, 0, 0, 0, "", "")
	/*109*/ insertTitle(db, "Forces of Darkness", 12, 0, 0, 0, "", "")
	/*110*/ insertTitle(db, "Global War", 12, 0, 0, 0, "", "")
	/*111*/ insertTitle(db, "Iron War", 12, 0, 0, 0, "", "")
	/*112*/ insertTitle(db, "Grunt Fest", 12, 0, 0, 0, "", "")
	/*113*/ insertTitle(db, "Lunatix", 12, 7, 0, 0, "", "")
	/*114*/ insertTitle(db, "Mutants!", 12, 7, 0, 0, "", "")
	/*115*/ insertTitle(db, "Betsy Ross Solitaire", 14, 0, 0, 0, "", "")
	/*116*/ insertTitle(db, "Accordian Solitaire", 14, 0, 0, 0, "", "")
	/*117*/ insertTitle(db, "Galaxy 5", 15, 0, 0, 0, "", "")
	/*118*/ insertTitle(db, "Star Trek", 15, 16, 0, 0, "", "")
	/*119*/ insertTitle(db, "Trade Wars 2 - 500 Secors", 15, 16, 0, 0, "", "")
	/*120*/ insertTitle(db, "Galactic Dynasty 2", 15, 0, 0, 0, "", "")
	/*121*/ insertTitle(db, "Stellar Quest", 15, 0, 0, 0, "", "")
	/*122*/ insertTitle(db, "Yankee Trader", 15, 0, 0, 0, "", "")
	/*123*/ insertTitle(db, "Star Trek Quiz", 17, 0, 0, 0, "", "")
	/*124*/ insertTitle(db, "Showbiz Quiz", 17, 0, 0, 0, "", "")
	/*125*/ insertTitle(db, "Entertainment Awards Trivia", 17, 0, 0, 0, "", "")
	/*126*/ insertTitle(db, "Farwest Trivia", 17, 7, 0, 0, "", "")
	/*126*/ insertTitle(db, "Chain Reaction", 18, 0, 0, 0, "", "")
	/*127*/ insertTitle(db, "Jeopardized!", 18, 16, 0, 0, "", "")
	/*128*/ insertTitle(db, "Wordem", 18, 16, 0, 0, "", "")
	/*129*/ insertTitle(db, "Soduku", 18, 0, 0, 0, "", "")
	/*130*/ insertTitle(db, "BBS Wordle", 18, 0, 0, 0, "", "")

	/*131*/
	insertTitle(db, "MKZ LORD", 11, 0, 0, 0, "", "")
	/*132*/ insertTitle(db, "Planets: Exploration of Space", 15, 0, 0, 0, "", "")
	/*133*/ insertTitle(db, "Trade Wars 2002", 15, 0, 0, 0, "", "")
	/*134*/ insertTitle(db, "Global War", 12, 0, 0, 0, "", "")
	/*135*/ insertTitle(db, "Global Backgammon", 3, 0, 0, 0, "", "")
	/*136*/ insertTitle(db, "Falcon's Eye", 11, 0, 0, 0, "", "")
	/*137*/ insertTitle(db, "Falcon's Honor", 11, 0, 0, 0, "", "")
	/*138*/ insertTitle(db, "Arcadian Legends", 11, 0, 0, 0, "", "")
	/*139*/ insertTitle(db, "Netrunner", 12, 19, 0, 0, "", "")
	/*140*/ insertTitle(db, "Assassin", 11, 0, 0, 0, "", "")
	/*141*/ insertTitle(db, "BBS Crash", 19, 0, 0, 0, "", "")
	/*142*/ insertTitle(db, "Virtual Sysop", 19, 0, 0, 0, "", "")
	/*143*/ insertTitle(db, "The Pit", 11, 0, 0, 0, "", "")
	/*144*/ insertTitle(db, "Junkyard", 9, 0, 0, 0, "", "")
	/*145*/ insertTitle(db, "Murder Motel", 9, 0, 0, 0, "", "")
	/*146*/ insertTitle(db, "Dark Lands", 11, 0, 0, 0, "", "")
	/*147*/ insertTitle(db, "Hackin' Crackin'", 19, 0, 0, 0, "", "")
	/*148*/ insertTitle(db, "Mega Slots", 4, 0, 0, 0, "", "")
	/*149*/ insertTitle(db, "Fresh Water Fishing Sim", 6, 0, 0, 0, "", "")
	/*150*/ insertTitle(db, "Lemonade", 9, 0, 0, 0, "", "")
	/*151*/ insertTitle(db, "BCR Games Server", 9, 0, 0, 0, "", "")

	/*152*/
	insertTitle(db, "Solar Relams Elite", 15, 12, 0, 0, "", "")
	/*153*/ insertTitle(db, "The Clans", 11, 0, 0, 0, "", "")
	/*154*/ insertTitle(db, "LORE", 11, 0, 0, 0, "", "")
	/*155*/ insertTitle(db, "Drakeville", 11, 0, 0, 0, "", "")
	/*156*/ insertTitle(db, "LOD", 12, 0, 0, 0, "", "")
	/*157*/ insertTitle(db, "Starship Galactica", 15, 12, 0, 0, "", "")
	/*158*/ insertTitle(db, "Hackarama!", 19, 0, 0, 0, "", "")
	/*159*/ insertTitle(db, "Lost in Space", 15, 12, 0, 0, "", "")
	/*160*/ insertTitle(db, "Pentasim", 12, 0, 0, 0, "", "")
	/*161*/ insertTitle(db, "Time Travel", 12, 0, 0, 0, "", "")
	/*162*/ insertTitle(db, "Dog World 2", 12, 0, 0, 0, "", "")

	// INSERT DOOR RECORDS FOR EACH SERVER
	// code, title int, server int

	// Adult
	insertDoor(db, "DRUGLORD", 1, 1)
	insertDoor(db, "DRUGLORD", 1, 3)
	insertDoor(db, "BORDELLO", 2, 1)
	insertDoor(db, "BORD", 2, 2)
	insertDoor(db, "PIMPWARS", 3, 1)
	insertDoor(db, "PIMPWARS", 3, 3)
	insertDoor(db, "PIMP", 3, 2)
	insertDoor(db, "RN", 4, 1)
	insertDoor(db, "PISSORAM", 5, 1)
	insertDoor(db, "GODFATHE", 6, 1)
	// Arcade & Action
	insertDoor(db, "GOOBLE", 7, 1)
	insertDoor(db, "GOOBLE", 7, 3)
	insertDoor(db, "LEMONS", 8, 1)
	insertDoor(db, "CHICKEN", 9, 1)
	insertDoor(db, "CHXDELIV", 9, 3)
	insertDoor(db, "MYSTMAN", 10, 1)
	insertDoor(db, "BBSTRIS", 11, 1)
	insertDoor(db, "ESCAPE", 12, 1)
	// Board Games
	insertDoor(db, "4CORNERC", 13, 1)
	insertDoor(db, "BOGGLE", 14, 1)
	insertDoor(db, "BUBLBOGL", 14, 3)
	insertDoor(db, "CONCENTR", 15, 1)
	insertDoor(db, "RSEABATTL", 16, 1)
	insertDoor(db, "GGAM", 135, 1)
	// Cards & Casino
	insertDoor(db, "3PAIRPOK", 17, 1)
	insertDoor(db, "ACEYDEUC", 18, 1)
	insertDoor(db, "BBDICE", 19, 1)
	insertDoor(db, "BLACKJAC", 20, 1)
	insertDoor(db, "BOXPOKER", 21, 1)
	insertDoor(db, "DICEWARZ", 22, 1)
	insertDoor(db, "DICEWARZ", 22, 3)
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
	insertDoor(db, "MEGA", 148, 2)
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
	// Sports
	insertDoor(db, "BBDICE", 71, 1)
	insertDoor(db, "FATFISH", 72, 1)
	insertDoor(db, "FATFISH", 72, 3)
	insertDoor(db, "GUTBOWL", 73, 1)
	insertDoor(db, "CAMPTOWN", 74, 1)
	insertDoor(db, "FISH", 149, 2)
	// Other
	insertDoor(db, "SOKOBAN", 75, 1)
	insertDoor(db, "SOKOBAN", 75, 3)
	insertDoor(db, "THIRSTY", 76, 1)
	insertDoor(db, "THIRSTY", 76, 3)
	insertDoor(db, "PORTVIC", 77, 1)
	insertDoor(db, "OREGONTR", 78, 1)
	insertDoor(db, "JUNK", 144, 2)
	insertDoor(db, "MMOT", 145, 2)
	insertDoor(db, "LMON", 150, 2)
	insertDoor(db, "LEMONS", 150, 3)
	insertDoor(db, "BCR", 151, 2)
	insertDoor(db, "DW2", 162, 3)
	// Puzzle
	insertDoor(db, "MSWEEPER", 80, 1)
	insertDoor(db, "MSWEEPER", 80, 3)
	insertDoor(db, "UBERBLOX", 81, 1)
	insertDoor(db, "UBERBLOX", 81, 3)
	insertDoor(db, "TETRIS", 82, 1)
	insertDoor(db, "STACKTIO", 83, 1)
	// RPG1
	insertDoor(db, "LORD", 84, 1)
	insertDoor(db, "LORD", 84, 2)
	insertDoor(db, "LORD", 84, 3)
	insertDoor(db, "LORD2", 85, 1)
	insertDoor(db, "LORD2", 85, 2)
	insertDoor(db, "LORD2", 85, 3)
	insertDoor(db, "MKZLORD", 131, 2)
	insertDoor(db, "TBD", 86, 1)
	insertDoor(db, "FH", 87, 1)
	insertDoor(db, "DEATHMAS", 88, 1)
	insertDoor(db, "DEATHMAS", 88, 3)
	insertDoor(db, "DOORMUD", 89, 1)
	insertDoor(db, "DOORMUD", 89, 2)
	insertDoor(db, "DM", 90, 1)
	insertDoor(db, "DMAS", 90, 2)
	insertDoor(db, "KNK", 91, 1)
	insertDoor(db, "KNK", 91, 3)
	insertDoor(db, "AB1", 92, 1)
	insertDoor(db, "AB1", 92, 3)
	insertDoor(db, "AB2", 93, 1)
	insertDoor(db, "AB2", 93, 3)
	insertDoor(db, "LEGION", 94, 1)
	insertDoor(db, "LGN", 94, 2)
	insertDoor(db, "KC", 95, 1)
	insertDoor(db, "EXIT", 96, 1)
	insertDoor(db, "EXIT", 96, 3)
	insertDoor(db, "USURPER", 97, 1)
	insertDoor(db, "USRP", 97, 2)
	insertDoor(db, "USURPER", 97, 3)
	insertDoor(db, "AM2", 98, 1)
	insertDoor(db, "RTP", 99, 1)
	insertDoor(db, "BM", 100, 1)
	insertDoor(db, "XROADS", 101, 1)
	insertDoor(db, "TARENA", 102, 1)
	insertDoor(db, "FALC", 136, 2)
	insertDoor(db, "FE", 136, 3)
	insertDoor(db, "FHON", 137, 2)
	insertDoor(db, "FH", 137, 3)
	insertDoor(db, "ARCL", 138, 2)
	insertDoor(db, "TAL", 138, 3)
	insertDoor(db, "ASSN", 140, 2)
	insertDoor(db, "TPIT", 143, 2)
	insertDoor(db, "TPIT", 143, 3)
	insertDoor(db, "DARK", 146, 2)
	insertDoor(db, "CLANS", 153, 2)
	insertDoor(db, "LORE", 154, 2)
	// RPG2
	insertDoor(db, "COLONIES", 103, 1)
	insertDoor(db, "COLONIES", 103, 3)
	insertDoor(db, "ATLANTIS", 104, 1)
	insertDoor(db, "ATLANTIS", 104, 3)
	insertDoor(db, "OOII", 105, 1)
	insertDoor(db, "OOII", 132, 2)
	insertDoor(db, "OOK", 132, 3)
	insertDoor(db, "DARK2", 106, 1)
	insertDoor(db, "DARK2", 106, 3)
	insertDoor(db, "DKNS", 106, 2)
	insertDoor(db, "JEDI", 107, 1)
	insertDoor(db, "JEDI", 107, 3)
	insertDoor(db, "BRE", 108, 1)
	insertDoor(db, "BRE", 108, 2)
	insertDoor(db, "BRE", 108, 3)
	insertDoor(db, "FOD", 109, 1)
	insertDoor(db, "GWAR", 110, 1)
	insertDoor(db, "GWAR", 134, 2)
	insertDoor(db, "IRONWAR", 111, 1)
	insertDoor(db, "GRUNTF", 112, 1)
	insertDoor(db, "LUNATIX", 113, 1)
	insertDoor(db, "LUNA", 113, 2)
	insertDoor(db, "LUNATIX", 113, 3)
	insertDoor(db, "MUTANTS", 114, 1)
	insertDoor(db, "NETR", 139, 2)
	insertDoor(db, "NETR", 139, 3)
	insertDoor(db, "LOD", 156, 3)
	insertDoor(db, "PENTASIM", 160, 3)
	insertDoor(db, "TIMETRAV", 161, 3)
	// Solitaire
	insertDoor(db, "BETSYROS", 115, 1)
	insertDoor(db, "AS", 116, 1)
	// Space
	insertDoor(db, "GALAXY5", 117, 1)
	insertDoor(db, "STARTREK", 118, 1)
	insertDoor(db, "STARTREK", 118, 3)
	insertDoor(db, "TW2", 119, 1)
	insertDoor(db, "TW", 133, 2)
	insertDoor(db, "TW2002", 133, 3)
	insertDoor(db, "GD2", 120, 1)
	insertDoor(db, "SQUEST", 121, 1)
	insertDoor(db, "YT", 122, 1)
	insertDoor(db, "TEOS", 132, 2)
	insertDoor(db, "TEOS", 132, 3)
	insertDoor(db, "SRE", 152, 3)
	insertDoor(db, "GALACTIC", 157, 3)
	// Trivia
	insertDoor(db, "TREKQUIZ", 123, 1)
	insertDoor(db, "SBQUIZ", 124, 1)
	insertDoor(db, "AWARDS", 125, 1)
	insertDoor(db, "FWTRIVIA", 126, 1)
	//Word Games
	insertDoor(db, "CHAINREA", 127, 1)
	insertDoor(db, "JEOPARDY", 128, 1)
	insertDoor(db, "JEOPARDY", 128, 3)
	insertDoor(db, "WORDEM", 129, 1)
	insertDoor(db, "SODUKU", 130, 1)
	insertDoor(db, "WORDLE", 131, 1)
	// BBS Hack Crack
	insertDoor(db, "HACK", 147, 2)
	insertDoor(db, "SIMBBS", 79, 1)
	insertDoor(db, "BBSC", 141, 2)
	insertDoor(db, "VSYS", 142, 2)
	insertDoor(db, "HACKARAM", 158, 3)

}
