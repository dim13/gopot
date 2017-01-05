package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("teapot.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	b := Parse(fd)
	for _, v := range b {
		for _, p := range v.X() {
			fmt.Println(p)
		}
		fmt.Println("")
	}
}
