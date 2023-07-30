package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

// Function to perform a test query to the database
// func testDBQuery() {
// 	// Database connection string
// 	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

// 	// Open a connection to the database
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	// Perform a test query
// 	rows, err := db.Query("SELECT 1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	// Process the query results
// 	for rows.Next() {
// 		var result int
// 		err := rows.Scan(&result)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Query Result:", result)
// 	}
// }

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is GET
    if r.Method == http.MethodGet {
        // Your logic to handle the home route ("/") for GET requests goes here

        // Write a response back to the client
        fmt.Fprintf(w, "This is the home page for GET request!")
    } else {
        // If the request method is not GET, send an error response
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    // Your logic to handle the "/about" route goes here

    // Write a response back to the client
    fmt.Fprintf(w, "This is the about page!")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    // Your logic to handle the "/contact" route goes here

    // Write a response back to the client
    fmt.Fprintf(w, "This is the contact page!")
}

func startServer(router http.Handler, port int) {
    err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
    if err != nil {
        panic(err)
    }
}

func main() {
    const PORT int = 8080
	// Test the database connection and query
	// testDBQuery()

    // Create a new router
    router := mux.NewRouter()

    // Register the handler functions for different paths
    router.HandleFunc("/", homeHandler)
    router.HandleFunc("/about", aboutHandler)
    router.HandleFunc("/contact", contactHandler)

    // Add CORS middleware to the router
    corsMiddleware := cors.Default().Handler(router)

    // Start the server on port 8080 using the CORS-enabled router
    go startServer(corsMiddleware, PORT)

    // Print a message indicating the port
    fmt.Println("Server is running on port", PORT)

    // Keep the main goroutine running
    select {}
}
