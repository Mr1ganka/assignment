package main

import (
	a "assignment/Config"
	b "assignment/Routes"
	"fmt"
)

func main() {
	a.InitialMigration()
	fmt.Println("Tables have been Created")
	b.InitializeRouter()

}
