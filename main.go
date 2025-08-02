package main

import (
	"fmt"
	"goSre/database"
)

func main() {
	db := database.UpdatePassword("carlos@hotmail.com", "102030")
	fmt.Println(db)
	/*
		newUser := database.CreateUser("carlos@gmail.com", "123456")
		fmt.Println(newUser)*/

}
