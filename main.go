package main

import (
	"fmt"
	"diegofonseca/sing-me-a-music/sing"
)


func main() {

	data, err := sing.Sing()

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}