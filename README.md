# How to run the app

## A Note On The Database
I used a MySQL database locally. I am not sure how the database will behave on your machine. That being said the Golang backend server should automatically allocate resources to run the MySQL server on an assigned port on your machine. However if you need credentials to the database for any reason, they can be found in the `server.go` file inside `func main ()`. I will also provide them here: 
```code
    //Connection creds to MySQL db
    dbHost := "localhost"
    dbPort := 3306 
    dbUser := "root"
    dbPassword := "showMessage();"
    dbName := "eightcig-test-db"
```

## Prerequisites

Before getting started, make sure you have:
<br>
[Node.js](https://nodejs.org/) and [Golang](https://go.dev/doc/install) installed on your machine.

## Clone the repository
Go to the GitHub repository of the project and clone it to your local machine using either Git CLI or GitHub Desktop. Or download the code as .zip file.

## Frontend Setup (Vite/React)
Navigate to the root directory of the project (you should see folders like: `src`, `public`, `node_modules`).
Run `npm install` to install the required frontend dependencies.

```bash
npm install

```

## Backend Setup (GoLang)
Navigate to the backend directory.

```bash
cd .\backend\
```

Make sure you have Golang installed correctly and the `go` command is available
<br>
<br>
Run `go mod download` to install the required backend dependencies.
```bash
go mod download
```

## Run The Backend Server

In the backend directory, run `go run .\server.go`. This should start the Golang backend server.
```bash
go run .\server.go
```

## Run The Frontend Server

Open a separate terminal and navigate back to the root directory. Run `npm run dev` to start the Vite development server for the frontend.
```bash
npm run dev
```
Vite should automatically start running the frontend server on an available port.
<br>
In your browser navigate to the IP Vite provides you `http://localhost:<port>/`
<br>
<br>
The fullstack application should <i>hopefully</i> now be running on your machine!
