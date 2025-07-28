package main

import (
	"errors"
	"fmt"
	"log"
)

// type UserMeta struct {
// 	Visit string
// }

// type User struct {
// 	Name    string
// 	Bio     string
// 	Age     int
// 	Higth   float32
// 	Country []UserMeta //slices, arr
// 	Like    []int
// 	Notes   map[string]int //maps
// }

// func main() {

// 	t, err := template.ParseFiles("hello.gohtml")
// 	if err != nil {
// 		panic(err)
// 	}

// 	user := User{
// 		Name:  "Suzane Smith",
// 		Bio:   "<script>alert(\"Haha, you have been h4x0r3d!\");</script>",
// 		Age:   123,
// 		Higth: 1.89,
// 		Country: []UserMeta{
// 			{Visit: "France"},
// 			{Visit: "Turkey"},
// 			{Visit: "Belgium"},
// 		},
// 		Like: []int{1, 2, 4},
// 		Notes: map[string]int{
// 			"Math":     13,
// 			"Francais": 12,
// 			"Art":      15,
// 		},
// 	}

// 	if user.Name == "Suzane Smith" {
// 		fmt.Printf("---------- OUI---------------")
// 	} else {
// 		fmt.Printf("----------- YOU NEED TO BE MEMBERS----------------")
// 	}

// 	err = t.Execute(os.Stdout, user)
// 	if err != nil {
// 		panic(err)

// 	}
// }

func main() {
	err := CreateUser()
	if err != nil {
		log.Println(err)
	}
	err = CreateOrg()
	if err != nil {
		log.Println(err)
	}
}

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		fmt.Printf("interger: %d string: %s anywalue:%v\n", 2123, "a-string", "another string")
		return fmt.Errorf("create user : %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org: %w", err)
	}
	return nil
}
