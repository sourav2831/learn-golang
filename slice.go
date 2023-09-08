package main

import "fmt"

func slice(){
	cards:=[]string{"1","2"};
	cards=append(cards,"3")
	for i,card:=range cards{
		fmt.Println(card,i)
	}
}