package routes

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func FetchUsers(writer http.ResponseWriter, request *http.Request) {
	db, err := database.Connect()

	if err != nil {
		writer.Write([]byte("Error connecting to database!"))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		writer.Write([]byte("Error fetching users!"))
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user

		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			writer.Write([]byte("Error scanning user!"))
			return
		}
		users = append(users, user)
	}

	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(users); err != nil {
		writer.Write([]byte("Error encoding users to JSON"))
		return
	}
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		writer.Write([]byte("Error converting parameter to integer!"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		writer.Write([]byte("Error connecting to database!"))
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM user WHERE id = ?", ID)

	if err != nil {
		writer.Write([]byte("Error getting user!"))
		return
	}
	defer row.Close()

	var user user

	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			writer.Write([]byte("Error scanning user!"))
			return
		}
	}

	if user.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprintf("User with id %d not found!", ID)))
		return
	}

	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(user); err != nil {
		writer.Write([]byte("Error encoding user to JSON"))
		return
	}
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		writer.Write([]byte("Error parsing id!"))
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.Write([]byte("Error reading 	request body!"))
		return
	}

	var user user

	if err = json.Unmarshal(requestBody, &user); err != nil {
		writer.Write([]byte("Error parsing JSON to user struct!"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		writer.Write([]byte("Error connecting to database!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE user SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		writer.Write([]byte("Error creating statement!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		writer.Write([]byte("Error updating user!"))
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		writer.Write([]byte("Error parsing id!"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		writer.Write([]byte("Error connecting to database!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		writer.Write([]byte("Error creating statement!"))
		return
	}

	defer statement.Close()
	if _, err := statement.Exec(ID); err != nil {
		writer.Write([]byte("Error deleting user!"))
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
