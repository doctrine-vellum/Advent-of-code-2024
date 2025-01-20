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
	var fileContent_list1 = []int{};
	var fileContent_list2 = []int{};
	file, err := os.Open("input.txt");
	if err != nil {
		log.Fatal(err);
	}
	defer file.Close();

	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		var temp []string = strings.Split(scanner.Text(), " ");
		num1, err1 := strconv.Atoi(temp[0]);
		num2, err2 := strconv.Atoi(temp[3]);
    		if err1 != nil || err2 != nil {
			fmt.Println("Fuck");
			panic("fuck");
    		}
		fileContent_list1 = append(fileContent_list1, num1);
		fileContent_list2 = append(fileContent_list2, num2);
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err);
	}
	fileContent := [][]int{fileContent_list1, fileContent_list2};
	return fileContent;
}
