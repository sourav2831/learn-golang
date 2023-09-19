package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (sb spanishBot) getGreeting() string {
	//custom logic
	return "hola"
}

func (eb englishBot) getGreeting() string {
	//custom logic
	return "Hi"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}