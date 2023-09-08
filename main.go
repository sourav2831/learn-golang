package main

import "fmt"

func main(){
	//var variableName string="Hello world";
	variableName:=getVariableName()
	fmt.Println(variableName)
}

func getVariableName()string{
	return "Hello world"
}