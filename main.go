package main

import (
	"fmt"
	"main/lemin"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		n, g := lemin.ParseFile(filename)
		group := g.GetBestPath(n)
		g.SendAnts(group, n)
	} else if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguments.")
	}
}
