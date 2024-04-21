package main

import (
    "bufio"
    // "fmt"
    "log"
    "os"
    "strings"
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func openDB() *sql.DB {
    db, err := sql.Open("sqlite3", "database.db")
    if err != nil {
        log.Fatal(err)
    }
    return db
}

func getApplicationFromCommand(input string) string {
    words := strings.Fields(input) // Split the input string into words
    if len(words) == 0 {
        return "" // Return empty if there are no words
    }
    if words[0] == "sudo" && len(words) > 1 {
        return words[1] // Return the second word if the first is 'sudo'
    }

	stripped := strings.TrimPrefix(words[0], "./")

    return stripped // Return the first word otherwise
}

func checkIfCommandIsInDB(db *sql.DB, command string) bool {
    var exists bool
    query := `SELECT EXISTS(SELECT 1 FROM commands WHERE full_command = ? LIMIT 1)`
    err := db.QueryRow(query, command).Scan(&exists)
    if err != nil {
        log.Fatal("Failed to check if command exists:", err)
    }
    return exists
}

func checkIfAppIsInDB(db *sql.DB, application string) bool {
    var exists bool
    query := `SELECT EXISTS(SELECT 1 FROM applications WHERE application = ? LIMIT 1)`
    err := db.QueryRow(query, application).Scan(&exists)
    if err != nil {
        log.Fatal("Failed to check if application exists:", err)
    }
    return exists
}

func writeNewAppToDB(db *sql.DB, application string) {
	query := `INSERT INTO applications (application, occurrences, display, last_used, favorite) VALUES (?, 1, 'True', strftime('%s','now'), 'False');`
	_, err := db.Exec(query, application)
	if err != nil {
		log.Fatalf("Failed to insert new application %s: %v", application, err)
	}
}

func writeNewCommandToDB(db *sql.DB, application string, command string) {
	query := `INSERT INTO commands (application, full_command, deleted, last_used, favorite) VALUES (?, ?, 'False', strftime('%s','now'), 'False');`
	_, err := db.Exec(query, application, command)
	if err != nil {
		log.Fatalf("Failed to insert new command for application %s: %v", application, err)
	}
}

func updateLastUsedDateOfCommand() {

}

func updateLastUsedDateOfApp() {

}

func main() {
    file, err := os.Open("bash_history")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
		
		command := scanner.Text()

		db := openDB()
		defer db.Close()

		if checkIfCommandIsInDB(db, command) == false {
            application := getApplicationFromCommand(command)
			writeNewCommandToDB(db, application, command)

			if checkIfAppIsInDB(db, application) == false {
				writeNewAppToDB(db, application)
		// 	} else {
				// updateLastUsedDateOfApp()
			}
		// } else {
            // updateLastUsedDateOfCommand()
        }

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}



// generate side bar
// get all entries in applications table sorted by occurences

// generate main window
// get all entries of a choosen application from commands table sorted by last_used

