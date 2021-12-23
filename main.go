package main 

import (
	"fmt"
	"github.com/pluralsight/webservice/models"
)

func main(){
	fmt.Println("hello world")

	u := models.User {
		ID: 2,
		FirstName: "Trish",
		LastName: "Yable",
	}

	fmt.Println(u)
}

