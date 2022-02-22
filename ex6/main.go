package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("/etc/passwd")
	defer f.Close()
	check(err)
	reader := bufio.NewReader(f)
	users := make(map[string]int)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line_splitted := strings.Split(line, ":")
		user := line_splitted[0]
		id, err := strconv.Atoi(line_splitted[2])
		users[user] = id
		fmt.Printf("User n°%d : %s\n", id, user)
	}
}
