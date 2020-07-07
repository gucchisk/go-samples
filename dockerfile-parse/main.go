package main

import(
	"fmt"
	"os"
	"regexp"
	"github.com/asottile/dockerfile"
)


func main() {
	cmds, err := dockerfile.ParseFile("Dockerfile")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	file, err := os.Create("Dockerfile.new")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer file.Close()
	
	for _, cmd := range cmds {
		// fmt.Printf("%d: %v\n", i, v)
		org := cmd.Original
		m, _ := regexp.MatchString("yum ", org)
		if !m {
			_, err := file.WriteString(org + "\n")
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
		fmt.Printf("original: %s\n", org)
		fmt.Printf("  cmd:%s\n", cmd.Cmd)
		fmt.Println("  values:")
		fmt.Print("    ")
		for _, v := range cmd.Value {
			fmt.Printf("%s, ", v)
		}
		fmt.Println("")
	}
}
