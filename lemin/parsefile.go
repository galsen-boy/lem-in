package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ParseFile extracts number of ants, rooms, rooms connections from given file
func ParseFile(filename string) (int, Graph) {
	var g Graph

	data, err := os.ReadFile(filename)
	check(err)
	Task = string(data)

	lines := strings.Split(string(data), "\n")

	n, err1 := strconv.Atoi(lines[0])
	check(err1)
	if n == 0 {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}

	flag := "room"
	lines = lines[1:]
	countStart := 0
	countEnd := 0

	for _, line := range lines {
		if line == "##start" {
			flag = "start"
			countStart++
			continue
		} else if line == "##end" {
			flag = "end"
			countEnd++
			continue
		} else if len(line) == 0 || line[:1] == "#" || line[:1] == "L" { // if it is comment
			if len(strings.Split(line, " ")) == 3 {
				fmt.Println("ERROR: invalid data format")
				os.Exit(1)
			}
			continue
		} else if countEnd > 1 || countStart > 1 {
			fmt.Println("ERROR: invalid data format")
			os.Exit(1)
		}

		splitted := strings.Split(line, " ")

		if len(splitted) == 1 { // if it is a link
			rooms := strings.Split(splitted[0], "-")
			if rooms[0] == "L" || rooms[1] == "L" || rooms[0] == "#" || rooms[1] == "#" {
				fmt.Println("ERROR: invalid data format")
				os.Exit(0)
			} else {
				g.Connect(rooms[0], rooms[1])
				continue
			}

		}

		if flag == "start" {
			g.Start = splitted[0]
			// fmt.Println(g.Start)
		} else if flag == "end" {
			g.End = splitted[0]
		}
		if g.Start == "L" || g.End == "L" {
			fmt.Println("ERROR: invalid data format")
			os.Exit(1)
		}
		flag = "room"
	}

	return n, g
}
