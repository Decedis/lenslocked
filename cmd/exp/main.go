package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string //there is a template.HTML type we can use if we want to keep the script tags as they are.
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Doe",
		Bio:  `<script>alert("Hello scrublord")</script>`,
		Age:  42,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
