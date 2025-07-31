package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
	SSLMode string
}

func (cgf PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cgf.Host, cgf.Port, cgf.User, cgf.Password, cgf.Database, cgf.SSLMode )
}

func main() {

	cfg := PostgresConfig{
		Host: "localhost",
		Port: "5432",
		User: "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode: "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected!")

	//CREATE A TABLE 
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount INT,
		description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created")

	// INSERT some data
	name := "new user"
	email := "new@calhoun.io"

	// query := fmt.Sprintf(`
	// INSERT INTO users (name, email)
	// VALUES ('%s', '%s');
	// `, name, email)
	// fmt.Printf("Executing : %s\n", query)
	// _ , err = db.Exec(query)

	row := db.QueryRow(`
		INSERT INTO users (name, email)
		VALUES ($1, $2) RETURNING id;`, name, email)
		
	var id int 
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("User created. id = ", id)
}