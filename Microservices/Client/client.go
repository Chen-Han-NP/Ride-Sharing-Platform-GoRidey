package main

/*
This GO server helps to provide an endpoint to:
1. Create a new account for the user upon signing up
2. Log in valiadtion for the user
3. Account modification for both driver and passengers
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const URL string = "http://localhost:8888/api/v1/account"

type User struct {
	userType     string `json:"user_type"`
	firstName    string `json:"first_name"`
	lastName     string `json:"last_name"`
	mobileNumber string `json:"mobile_number"`
	emailAddress string `json:"email_ddress"`
	idNumber     string `json:"id_number"`
	carLicNumber string `json:"car_lic_number"`
}

func login_page() {
	fmt.Printf(`
========
Log In Page
========`)
}

func signup_page() {
	fmt.Printf(`
========
Sign Up Page
========`)
	var accountType int
	var userType, firstName, lastName, mobileNumber, emailAddress string
	var idNumber, carLicNumber string

	fmt.Printf(`
Choose your account type
1. Passenger
2. Driver
Enter your option: `)
	fmt.Scanln(&accountType)
	if accountType == 1 {
		userType = "Passenger"
	} else {
		userType = "Driver"
	}

	fmt.Print("Enter your first name: ")
	fmt.Scanf("%s", &firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scanf("%s", &lastName)

	fmt.Print("Enter your mobile number: ")
	fmt.Scanf("%s", &mobileNumber)

	fmt.Print("Enter your email address: ")
	fmt.Scanf("%s", &emailAddress)

	if userType == "Driver" {
		fmt.Print("Enter your IC number")
		fmt.Scanf("%s", &idNumber)

		fmt.Print("Enter your Car License Number")
		fmt.Scanf("%s", &carLicNumber)
	} else {
		idNumber = "NULL"
		carLicNumber = "NULL"
	}

	var confirmSignUp string
	fmt.Printf(`
=========
Confirm your details
%s
First Name: %s
Last Name: %s
Mobile no: %s
Email Address: %s
IC Number: %s
Car License Number: %s
=========
Confirm? (Y/N): `, idNumber, firstName, lastName, mobileNumber, emailAddress, idNumber, carLicNumber)

	fmt.Scanln(&confirmSignUp)
	if confirmSignUp == "Y" || confirmSignUp == "y" {

		new_user := User{
			userType:     userType,
			firstName:    firstName,
			lastName:     lastName,
			mobileNumber: mobileNumber,
			emailAddress: emailAddress,
			idNumber:     idNumber,
			carLicNumber: carLicNumber,
		}

		client := &http.Client{}
		postBody, _ := json.Marshal(new_user)
		resBody := bytes.NewBuffer(postBody) // convert from strings to bytes
		if req, err := http.NewRequest("POST", URL+"/"+"signup", resBody); err == nil {
			if res, err := client.Do(req); err == nil {
				if res.StatusCode == 202 {
					fmt.Printf("New user created successfully!")
				} else if res.StatusCode == 409 { //conflict status
					fmt.Printf("Error - the user with email %s exists", emailAddress)
				}
				defer res.Body.Close()

			} else if err != nil {
				fmt.Println(err)
			}
		}

		return

	} else {
		return
	}

}

func main() {

	fmt.Printf(`
========
Shared-Riding Platform Console
 1. Login
 2. Signup
 9. Quit
Enter an option: `)

	var option int
	fmt.Scanln(&option)

	for option != 9 {
		if option == 1 {
			login_page()
		} else if option == 2 {
			signup_page()
		} else {
			fmt.Println("Please try again!")
		}

		fmt.Printf(`
========
Shared-Riding Platform Console
 1. Login
 2. Signup
 9. Quit
Enter an option: `)
		fmt.Scanln(&option)
	}

}
