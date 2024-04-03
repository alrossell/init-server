package main

import (
    "fmt"
    "log"
    "database/sql"

    _ "github.com/lib/pq"
)

type Song struct {
    Id int
    Title string
    Artist string
    Album string
    ReleaseYear int
    Genre string
    DurationSeconds int
}

func main() {
    connStr := "user=postgres password=mypasswpord dbname=new_database sslmode=disable"
   
    var err error
    client, err := sql.Open("postgres", connStr)

    if err != nil {
        log.Fatal("Failed to connect to server: ", err)
    }

    createTableSQL := `CREATE TABLE IF NOT EXISTS songs (
        song_id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        artist TEXT NOT NULL,
        album TEXT NOT NULL,
        release_year INT,
        genre TEXT,
        duration_seconds INT
    );`

    _, err = client.Exec(createTableSQL)

    if err != nil {
        log.Fatal("Failed to create songs table: ", err)
        return
    } else {
        log.Println("Songs table created successfully!")
    }

    getRowCount := `SELECT COUNT(*) FROM songs;`

    var numberOfRows int
    err = client.QueryRow(getRowCount).Scan(&numberOfRows)

    if err != nil {
        log.Fatal("Failed to get row count: ", err)
        return
    } else {
        log.Println("Row count: ", numberOfRows)
    }

    if numberOfRows == 0 {
        log.Println("Inserting songs into songs table.")
        var songs []Song = []Song{
            {Title: "A Title", Artist: "A Artist", Album: "A Album", ReleaseYear: 2020, Genre: "Pop", DurationSeconds: 180},
            {Title: "B Title", Artist: "B Artist", Album: "B Album", ReleaseYear: 2021, Genre: "Rock", DurationSeconds: 200},
            {Title: "C Title", Artist: "C Artist", Album: "C Album", ReleaseYear: 2022, Genre: "Jazz", DurationSeconds: 220},
            {Title: "D Title", Artist: "D Artist", Album: "D Album", ReleaseYear: 2020, Genre: "Pop", DurationSeconds: 180},
            {Title: "E Title", Artist: "E Artist", Album: "E Album", ReleaseYear: 2020, Genre: "Pop", DurationSeconds: 180},
            {Title: "F Title", Artist: "F Artist", Album: "F Album", ReleaseYear: 2021, Genre: "Rock", DurationSeconds: 200},
            {Title: "G Title", Artist: "G Artist", Album: "G Album", ReleaseYear: 2022, Genre: "Jazz", DurationSeconds: 220},
            {Title: "H Title", Artist: "H Artist", Album: "H Album", ReleaseYear: 2021, Genre: "Rock", DurationSeconds: 200},
            {Title: "I Title", Artist: "I Artist", Album: "I Album", ReleaseYear: 2022, Genre: "Jazz", DurationSeconds: 220},
        }

        for _, song := range songs {
            _, err = client.Exec(
                `INSERT INTO songs (title, artist, album, release_year, genre, duration_seconds)
                 VALUES ($1, $2, $3, $4, $5, $6)`,
                song.Title, song.Artist, song.Album, song.ReleaseYear, song.Genre, song.DurationSeconds)

            if err != nil {
                log.Fatal(err)
                fmt.Println("Error creating song")
            }
        }

        fmt.Println("Song created successfully!")
    }

    createReviewTableSQL := `CREATE TABLE IF NOT EXISTS reviews (
        review_id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        song_id INT NOT NULL,
        date DATE NOT NULL,
        review TEXT NOT NULL
    );`

    _, err = client.Exec(createReviewTableSQL)

    if err != nil {
        log.Fatal("Failed to create review table: ", err)
        return
    } else {
        log.Println("Reivew table created successfully!")
    }
    

    createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
        user_id SERIAL PRIMARY KEY,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        user_name TEXT NOT NULL,
        creation_date DATE NOT NULL,
        password TEXT NOT NULL
    );`

    _, err = client.Exec(createUserTableSQL)

    if err != nil {
        log.Fatal("Failed to create user table: ", err)
        return
    } else {
        log.Println("User table created successfully!")
    }
}

