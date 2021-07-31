package srvfoodlog

import "fmt"

type FoodLogger struct {}

func(serivce *FoodLogger) CreateLog() {
	fmt.Println("you called this service...")
}