package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for golang")

	//comma ok syntax
	input , _ := reader.ReadString('\n')
	fmt.Println("thanks for the rating", input)
	fmt.Printf("type of input is %T \n", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err!= nil {
		fmt.Println(err)
	}else{
		fmt.Println("the rating is", numRating+1)
		fmt.Println("thank you for rataingg")
	}


	

}