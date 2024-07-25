package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}

	pranav := User{"pranav", "pranav@go.dev", true, 21}
	fmt.Println(pranav)
	fmt.Printf("pranav details are: %+v\n", pranav)
	fmt.Printf("Name is %v and email is %v.", pranav.Name, pranav.Email)

}


type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}