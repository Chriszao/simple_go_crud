package routes

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InsertUser(writer http.ResponseWriter, request *http.Request) {
	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.Write([]byte("Failed to read request body!"))
		return
	}

	var user user

	if err = json.Unmarshal(requestBody, &user); err != nil {
		writer.Write([]byte("Failed to convert user to struct!"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		writer.Write([]byte("Error connecting to database!"))
		return
	}

	defer db.Close()

	// Prepare statement
	statement, err := db.Prepare("INSERT INTO user (name, email) VALUES (?, ?)")

	if err != nil {
		writer.Write([]byte("Error statement creation!"))
		return
	}

	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)

	if err != nil {
		writer.Write([]byte("Error statement execution!"))
		return
	}

	insertedId, err := insert.LastInsertId()

	if err != nil {
		writer.Write([]byte("Error getting id entered!"))
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("User successfully inserted! Id: %d", insertedId)))
}
