package main

import "fmt"

func main(){
	//var variableName string="Hello world";
	variableName:=getVariableName()
	fmt.Println(variableName)
	//slice()
}

func getVariableName()string{
	return "Hello world"
}