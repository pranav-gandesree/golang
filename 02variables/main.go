package main

import "fmt"

func main() {
	var name string = "pranavv"
	fmt.Println(name)
	fmt.Printf("the type of variable is: %T \n", name)

	var age int = 21
	fmt.Println(age)
	fmt.Printf("the type of variable is: %T \n", age)

	var isCollege bool = false
	fmt.Println("isCollege")
	fmt.Printf("the type of variable is: %T \n", isCollege)

	//default values of string anf int

	var college string
	fmt.Println(college)
	fmt.Printf("the type of variable is: %T \n", college)

	var salary int
	fmt.Println(salary)
	fmt.Printf("the type of variable is: %T \n", salary)

	//implicit type 
	var website = "pranav.cowjf"
	fmt.Printf("the type of variable is: %T \n", website)


	noOfUser := 300
	fmt.Println(noOfUser)
	fmt.Printf("the type of variable is: %T \n", noOfUser)

}