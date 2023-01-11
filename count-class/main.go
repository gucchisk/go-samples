package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/wreulicke/emc"
)

func main() {
	log.SetOutput(ioutil.Discard)
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	info, err := file.Stat()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	count, err := emc.CountClassFile(file, info, false)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("count: %d\n", count)
}
