package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq" // Importing pq package to enable PostgreSQL driver
)

func main() {
	// Set up the connection string
	connStr := "meow"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening connection to the database: ", err)
	}
	defer db.Close()

	// Check if the connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Call directory function to print file paths
	files := directory(`C:\Users\12813\Downloads\musi`)
	file, err := os.Open(files)
	if err != nil {
		print(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	query := "SELECT file_data FROM audio_files LIMIT 1"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error execution query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var file_data byte
		err := rows.Scan(&file_data)
		if err != nil {
			log.Fatal("Error Scanning row: ")
		}
	}
}

func directory(startingDirectory string) string {
	var filePaths []string
	err := filepath.Walk(startingDirectory, func(path string, info os.FileInfo, err error) error {
		// Handle any errors encountered during the walk
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Print the current file or directory path
		fmt.Println(path)
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		// Return nil to continue walking the directory
		return nil
	})

	// Check for errors from filepath.Walk
	if err != nil {
		fmt.Println("Error walking the directory:", err)
	}

	return filePaths[0]
}

// func upload(file_name string, file_data []byte) {
// 	return
// }
