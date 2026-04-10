package main

import "match-dos-estudos/src/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
