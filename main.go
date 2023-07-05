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

	teste := []string{"Joao", "Maria", "Ana", "Joao"}

	for _, name := range teste {
		fmt.Println(name)
	}

	p := Person{name: name, email: email}
	p.getdata()
}
