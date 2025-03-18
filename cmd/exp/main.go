package main

import (
	"html/template"
	"os"
)

type User struct {
	Name      string
	Bio       string //there is a template.HTML type we can use if we want to keep the script tags as they are.
	Age       int
	Inventory map[int]string
	Friends   []string
	Address   string
}

func main() {
	t, err := template.ParseFiles("hello.go.tpl")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Doe",
		Bio:  `<script>alert("Hello scrublord")</script>`,
		Age:  42,
		Inventory: map[int]string{
			1: "Baseball",
			2: "Bat",
			3: "Glove",
			4: "Sword",
		},
		Friends: []string{"Jason", "John", "Kyle", "Samantha", "Jennifer"},
		Address: "123 Cyprus Ave",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
