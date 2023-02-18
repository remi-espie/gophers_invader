//go:build linux

package screen

import (
	"database/sql"
	"fmt"
	tl "github.com/JoelOtter/termloop"
	_ "github.com/mattn/go-sqlite3"
	"gophers_invader/entities"
	"log"
	"os"
	"strconv"
)

func Scoreboard(game *tl.Game, mainMenuLevel *tl.BaseLevel) {

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	file, err := os.OpenFile("sqlite-database.db", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = file.Close()
	if err != nil {
		return
	}

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer func(sqliteDatabase *sql.DB) {
		err := sqliteDatabase.Close()
		if err != nil {
			return
		}
	}(sqliteDatabase)

	createTable(sqliteDatabase)

	// DISPLAY INSERTED RECORDS
	scoreboard := getScoreboard(sqliteDatabase)

	buttonMainMenu := entities.Button{
		Rectangle: tl.NewRectangle(-8, 4, 16, 3, tl.ColorGreen),
		Level:     level,
		Text:      "Main menu",
		Action: func() {
			game.Screen().SetLevel(mainMenuLevel)
		},
	}

	level.AddEntity(&buttonMainMenu)

	level.AddEntity(tl.NewRectangle(-20, -12, 40, 1, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(-20, -12, 1, 25, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(20, -12, 1, 25, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(-20, 13, 40, 1, tl.ColorWhite))

	level.AddEntity(tl.NewText(-15, -11, "Name:", tl.ColorBlue, tl.ColorBlack))
	level.AddEntity(tl.NewText(-5, -11, "Score:", tl.ColorBlue, tl.ColorBlack))
	level.AddEntity(tl.NewText(5, -11, "Duration:", tl.ColorBlue, tl.ColorBlack))

	for index, score := range scoreboard {
		if index > 10 {
			break
		}
		level.AddEntity(tl.NewText(-15, -10+index, score.name, tl.ColorBlue, tl.ColorBlack))
		level.AddEntity(tl.NewText(-5, -10+index, strconv.Itoa(score.score), tl.ColorBlue, tl.ColorBlack))
		level.AddEntity(tl.NewText(5, -10+index, fmt.Sprintf("%.2f", score.duration), tl.ColorBlue, tl.ColorBlack))
	}

	player := entities.Player{
		Entity: tl.NewEntityFromCanvas(0, 10, entities.CreateCanvas(playerBytes)),
		Level:  level,
		Game:   game,
	}
	level.AddEntity(&player)

	game.Screen().SetLevel(level)
	//game.Start()
}

func AddScore(name string, score int, duration float32) {
	file, err := os.OpenFile("sqlite-database.db", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = file.Close()
	if err != nil {
		return
	}

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer func(sqliteDatabase *sql.DB) {
		err := sqliteDatabase.Close()
		if err != nil {
			return
		}
	}(sqliteDatabase)

	insertScore(sqliteDatabase, score, duration, name)
}

func createTable(db *sql.DB) {
	createScoreTable := `CREATE TABLE IF NOT EXISTS scoreboard (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"score" int NOT NULL,
		"duration" float NOT NULL,
		"name" varchar(32) NOT NULL		
	  );`

	statement, err := db.Prepare(createScoreTable) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		return
	} // Execute SQL Statements
}

// We are passing db reference connection from main to our method with other parameters
func insertScore(db *sql.DB, score int, duration float32, name string) {
	insertStudentSQL := `INSERT INTO scoreboard(score, duration, name) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(score, duration, name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getScoreboard(db *sql.DB) []scoreBoard {
	row, err := db.Query("SELECT * FROM scoreboard ORDER BY score DESC")
	if err != nil {
		log.Fatal(err)
	}
	scoreboard := make([]scoreBoard, 0)
	defer row.Close()
	for row.Next() {
		var id int
		var score int
		var duration float32
		var name string
		row.Scan(&id, &score, &duration, &name)
		scoreboard = append(scoreboard, scoreBoard{id, score, duration, name})
	}
	return scoreboard
}

type scoreBoard struct {
	id       int
	score    int
	duration float32
	name     string
}
