package main
import (
	"fmt"
)

func main() {
	similarity := 0;
	columns := ReadFile();
	column1 := columns[0];
	column2 := columns[1];
	rightListFreq := make(map[int]int);
	for _, num := range column2 {
		rightListFreq[num] += 1;
	}
	for _, num := range column1 {
		similarity += num * rightListFreq[num];
	}
	fmt.Println(similarity);
}
