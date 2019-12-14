package main

import (
	. "awesomeProject/src/controllers"
)

func main() {
	router := Health()
	router.Run()
}