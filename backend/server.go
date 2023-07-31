package main

import (
    "fmt"
    "net/http"
    "database/sql"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
    _ "github.com/go-sql-driver/mysql"
)

// Function to perform a test query to the database
func connectToMySQL(dbUser, dbPassword, dbHost, dbName string, dbPort int) {
    // Database connection string
    connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    // Open a connection to the database
    db, err := sql.Open("mysql", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Perform a test query
    rows, err := db.Query("SELECT 1")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    // Process the query results
    for rows.Next() {
        var result int
        err := rows.Scan(&result)
        if err != nil {
            panic(err)
        }
        fmt.Println("Query Result:", result)
    }

    // Print a success message indicating a successful database connection
	fmt.Println("Connection to the MySQL database was successful and running on port", dbPort)
}

// Function to fetch and display data from the table
func fetchEmployeeData(dbUser, dbPassword, dbHost, dbName string, dbPort int) []map[string]interface{} {
    // Database connection string
    connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    // Open a connection to the database
    db, err := sql.Open("mysql", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Perform the SELECT query to fetch data from the table
    rows, err := db.Query("SELECT id, name, performance, date FROM employees ORDER BY performance DESC")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    // Prepare a slice to store the employee performance data
    var employeeData []map[string]interface{}

    // Process the query results and store the data in the slice
    for rows.Next() {
        var id int
        var name string
        var performance float64
        var date string
        err := rows.Scan(&id, &name, &performance, &date)
        if err != nil {
            panic(err)
        }
        // Create a map for each row with the appropriate keys and values
        employee := map[string]interface{}{
            "ID":         id,
            "Name":       name,
            "Performance": performance,
            "Date":       date,
        }
        // Append the map to the slice
        employeeData = append(employeeData, employee)
    }

    // Return the employee performance data slice
    return employeeData
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is GET
    if r.Method == http.MethodGet {
        http.ServeFile(w, r, "partTwo.html")

        // Write a response back to the client
        fmt.Fprintf(w, "This is the home page for GET request!")
    } else {
        // If the request method is not GET, send an error response
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
        // Check if the request method is GET
        if r.Method == http.MethodGet {
            // Your logic to handle the home route ("/") for GET requests goes here
            dbUser := "root"
            dbPassword := "showMessage();"
            dbHost := "localhost"
            dbPort := 3306
            dbName := "eightcig-test-db"
            employeeData := fetchEmployeeData(dbUser, dbPassword, dbHost, dbName, dbPort)
    
            // Convert the employeeData array to JSON
            jsonResponse, err := json.Marshal(employeeData)
            if err != nil {
                http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
                return
            }

            // Set the Content-Type header to application/json
            w.Header().Set("Content-Type", "application/json")

            // Write the JSON response back to the client
            w.Write(jsonResponse)
        } else {
            // If the request method is not GET, send an error response
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
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

    //Connection creds to MySQL db
    dbHost := "localhost"
    dbPort := 3306 
    dbUser := "root"
    dbPassword := "showMessage();"
    dbName := "eightcig-test-db"
    
     // Test the database connection and query
    connectToMySQL(dbUser, dbPassword, dbHost, dbName, dbPort)

    // Fetch and display data from the table
    fetchEmployeeData(dbUser, dbPassword, dbHost, dbName, dbPort)

    // Create a new router
    router := mux.NewRouter()

    // Register the handler functions for different paths
    router.HandleFunc("/", homeHandler)
    router.HandleFunc("/api/dashboard", dashboardHandler)
    router.HandleFunc("/contact", contactHandler)

    // Add CORS middleware to the router
    corsMiddleware := cors.Default().Handler(router)

    // Start the server on port 8080 using the CORS-enabled router
    go startServer(corsMiddleware, PORT)

    // Print a message indicating the port
    fmt.Println("Web server is running on port", PORT)

    // Keep the main goroutine running
    select {}
}
