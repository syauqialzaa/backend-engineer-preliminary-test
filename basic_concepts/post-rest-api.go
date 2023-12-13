package basic_concepts

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
}

// Dos
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Do: Parse JSON request body
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Do: Process data and send a response
	// (In a real-world scenario, you might want to store the user in a database)
	responseUser := User{Name: newUser.Name, Email: newUser.Email}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUser)
}

// Dont's
// func createUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// Don't: Access request data without checking Content-Type
// 	var newUser User
// 	json.NewDecoder(r.Body).Decode(&newUser)

// 	// Don't: Directly manipulate response writer or header
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte("User created successfully"))

// 	// Don't: Mix responsibilities, keep handler clean
// 	responseUser := processUserData(r.Body)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(responseUser)
// }


func PostRestAPI() {
	http.HandleFunc("/api/users", createUserHandler)

	port := ":3000"
	http.ListenAndServe(port, nil)
}
