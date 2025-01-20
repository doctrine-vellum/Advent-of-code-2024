package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)


func ReadFile() [][]int {
	fileContent := [][]int{[]int{}};
	file, err := os.Open("input.txt");
	if err != nil {
		log.Fatal(err);
	}
	defer file.Close();

	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		temp := []int{};
		for _, word := range strings.Split(scanner.Text(), " ") {
			num, err := strconv.Atoi(word);
			if err != nil {
				fmt.Println("Fuck");
				panic("fuck");
				fmt.Println(err);
			}
			temp = append(temp, num);
		}
		fileContent = append(fileContent, temp);
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err);
	}
	return fileContent;
}
