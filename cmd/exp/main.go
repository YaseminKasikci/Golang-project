package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main3() {
	password := "bobo"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("❌ Error hashing password:", err)
		return
	}

	fmt.Println("🔐 Hashed password:", string(hashedPassword))

	// hashedPassword = []byte("$2a$10$Tn0v/1JHDhLcGHTfmbRPnOExPYjIy2RAzustQWrajSST7MeM.HALG")

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("❌ Invalid password")
	} else {
		fmt.Println("✅ Password is valid")
	}
}
