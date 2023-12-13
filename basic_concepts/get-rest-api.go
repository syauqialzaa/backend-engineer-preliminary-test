package basic_concepts

import (
	"encoding/json"
	"net/http"
)

// type User struct {
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// Dos
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Do: Access query parameters
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	// Do: Fetch data (In a real-world scenario, retrieve users from a database based on page and limit)
	users := fetchUsersFromDatabase(page, limit)

	// Do: Send a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func fetchUsersFromDatabase(page, limit string) []User {
	// In a real-world scenario, implement logic to retrieve users based on page and limit
	// For simplicity, we return a static list of users here.
	return []User{
		{Name: "Anies Baswedan", Email: "anies@example.com"},
		{Name: "Prabowo Subianto", Email: "prabowo@example.com"},
		{Name: "Ganjar Pranowo", Email: "ganjar@example.com"},
	}
}

// Dont's
// func getUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	// Don't: Access query parameters directly from req.URL
// 	page := r.URL.Path

// 	// Don't: Perform heavy computation in the handler
// 	users := computeAllUsers()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)

// 	// Don't: Overuse nested functions
// 	getUserData(func(userData UserData) {
// 	  getUserDetails(userData.ID, func(userDetails UserDetails) {
// 	    // Process user details
// 	    w.Header().Set("Content-Type", "application/json")
// 	    json.NewEncoder(w).Encode(userDetails)
// 	  })
// 	})
// }

func GetRestAPI() {
	http.HandleFunc("/api/users", getUsersHandler)
	port := ":3000"

	// Start the HTTP server
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
