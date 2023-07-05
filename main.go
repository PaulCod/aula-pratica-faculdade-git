package main

import "fmt"

type Person struct {
	name  string
	email string
}

func (p *Person) getdata() {
	fmt.Printf("nome: %s, email: %s\n", p.name, p.email)
}

func main() {
	var name, email string

	fmt.Println("Hello")

	fmt.Println("Escreva seu nome: ")
	fmt.Scanf("%s\n", &name)

	fmt.Println("escreva seu email: ")
	fmt.Scanf("%s\n", &email)

	p := Person{name: name, email: email}
	p.getdata()
}
