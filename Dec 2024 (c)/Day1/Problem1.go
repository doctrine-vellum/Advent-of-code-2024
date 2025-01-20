package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	diff := 0
	colums := readFile()
	column1 := colums[0]
	column2 := colums[1]
	sort.Ints(column1)
	sort.Ints(column2)
	for index, num := range column1 {
		diff += abs(num - column2[index])
	}
	fmt.Println(diff)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func readFile() [][]int {
	var fileContent_list1 = []int{}
	var fileContent_list2 = []int{}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var temp []string = strings.Split(scanner.Text(), " ")
		num1, err1 := strconv.Atoi(temp[0])
		num2, err2 := strconv.Atoi(temp[3])
		if err1 != nil || err2 != nil {
			fmt.Println("Fuck")
		}
		fileContent_list1 = append(fileContent_list1, num1)
		fileContent_list2 = append(fileContent_list2, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fileContent := [][]int{fileContent_list1, fileContent_list2}
	return fileContent
}
