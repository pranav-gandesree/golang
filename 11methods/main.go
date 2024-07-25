package main

import "fmt"

func main() {
	fmt.Println("methods in go ")

	pranav := User{"pranav", "pranav@go.dev", true, 21}
	fmt.Println(pranav)
	fmt.Printf("pranav details are: %+v\n", pranav)
	fmt.Printf("Name is %v and email is %v. \n", pranav.Name, pranav.Email)
	pranav.getStatus()
	pranav.NewMail()
	fmt.Printf("Name is %v and email is %v. \n", pranav.Name, pranav.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus(){

	fmt.Println("user status is ", u.Status)
}

func (u User) NewMail(){
	u.Email = "newmail@gmail.com"
	fmt.Println("user status is ", u.Email)
}