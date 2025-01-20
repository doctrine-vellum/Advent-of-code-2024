package main
import "fmt"

func main() {
	fmt.Println("This is problem 1: ", Problem1());
	fmt.Println("This is problem 2: ", Problem2());
	fmt.Println("The has", len(ReadFile()), "Reports");
}
