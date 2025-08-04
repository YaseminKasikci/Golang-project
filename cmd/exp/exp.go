package main

import (
	"fmt"
	"github/yaseminkasikci/lenslocked/models"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected!")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("bob4@bob.com", "bob123")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	//CREATE A TABLE
	// _, err = db.Exec(`
	// CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	name TEXT,
	// 	email TEXT UNIQUE NOT NULL
	// 	);

	// 	CREATE TABLE IF NOT EXISTS orders (
	// 	id SERIAL PRIMARY KEY,
	// 	user_id INT NOT NULL,
	// 	amount INT,
	// 	description TEXT
	// 	);
	// `)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Tables created")

	// // INSERT some data
	// name := "new user"
	// email := "new@calhoun.io"

	// // query := fmt.Sprintf(`
	// // INSERT INTO users (name, email)
	// // VALUES ('%s', '%s');
	// // `, name, email)
	// // fmt.Printf("Executing : %s\n", query)
	// // _ , err = db.Exec(query)

	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2) RETURNING id;`, name, email)

	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("User created. id = ", id)
	// id := 1
	// row := db.QueryRow(`
	// 	SELECT name, email
	// 	FROM users
	// 	WHERE id=$1;`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("Error, no rows!")
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User information: name=%s, email=%s\n", name, email)

	// userID := 1
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("FAKE orders #%d", i)
	// 	_, err := db.Exec(`
	// 	INSERT INTO orders(user_id, amount, description)
	// 	VALUES($1,$2,$3)`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created fake orders.")

	// type Order struct {
	// 	ID          int
	// 	UserId      int
	// 	Amount      int
	// 	Description string
	// }

	// var orders []Order
	// userID := 1
	// rows, err := db.Query(`
	// 	SELECT id, amount, description
	// 	FROM orders
	// 	WHERE user_id=$1`, userID)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var order Order
	// 	order.UserId = userID
	// 	err := rows.Scan(&order.ID, &order.Amount, &order.Description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	orders = append(orders, order)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Order:", orders)
	// check for an err
}
