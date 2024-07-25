package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time in go")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2002, time.November, 2, 4, 30, 40, 0, time.UTC )

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
