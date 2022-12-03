package main

import (
	//"encoding/json"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// ==== STRUCTs ========
type User struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type Claims struct {
	EmailAddress string `json:"email_address"`
	jwt.RegisteredClaims
}

type Passenger struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MobileNumber string `json:"mobile_number"`
}

type Rider struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MobileNumber string `json:"mobile_number"`
	IcNumber     string `json:"ic_number"`
	CarLicNumber string `json:"car_lic_number"`
}

// ====== GLOBAL VARIABLES ========
var sqlConnectionString = "root:password@tcp(127.0.0.1:3306)/"
var database = "RideSharingPlatform"
var jwtKey = []byte("my_secret_key")

// RETURN 200 -> Registered
// RETURN 409 -> Duplicated account (email)
// RETURN 417 -> INSERT failed
func signup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_type := params["user_type"]
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)

	// Connect to the db
	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Convert passenger to Passenger
	//caser := cases.Title(language.English)
	//user_type_format := caser.String(user_type)
	//user_type_low := strings.ToLower(user_type)
	//user_type_format := cases.Title(language.English).String(user_type_low)

	// Step 1: Check if rider or passenger
	if user_type == "passenger" {
		var passenger Passenger
		json.Unmarshal(reqBody, &passenger)

		// check if email exists in the User table
		isExist := checkEmailIsExist(passenger.EmailAddress, db)
		if isExist {
			w.WriteHeader(http.StatusConflict) //409
			fmt.Println("Duplicated account: " + passenger.EmailAddress)
			return
		}

		// Insert into User table
		userQueryStatement := fmt.Sprintf(`
		INSERT INTO User(email_address, password)
		VALUES ('%s', '%s')`, passenger.EmailAddress, passenger.Password)
		result, err := db.Exec(userQueryStatement)
		if err != nil {
			panic(err.Error())
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err.Error())
		}

		// Upon getting the ID, now insert into Passenger table
		queryStatement := fmt.Sprintf(`
		INSERT INTO Passenger
		VALUES (%d, '%s', '%s', '%s')`,
			id,
			passenger.FirstName,
			passenger.LastName,
			passenger.MobileNumber)

		result2, err := db.Exec(queryStatement)
		if err != nil {
			panic(err.Error())
		}
		rows_affected, err := result2.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		if rows_affected == 1 {
			w.WriteHeader(http.StatusAccepted) //202
			json.NewEncoder(w).Encode("Insert Successfully")
			fmt.Println("insert successfully")
			return
		} else {
			w.WriteHeader(http.StatusExpectationFailed) //417
			fmt.Println("Error with inserting")
			return
		}

	} else if user_type == "rider" {
		var rider Rider
		json.Unmarshal(reqBody, &rider)
		// check if email exists
		isExist := checkEmailIsExist(rider.EmailAddress, db)
		if isExist {
			w.WriteHeader(http.StatusConflict) //409
			fmt.Println("Duplicated account: " + rider.EmailAddress)
			return
		}

		// Insert into User table
		userQueryStatement := fmt.Sprintf(`
		INSERT INTO User(email_address, password)
		VALUES ('%s', '%s')`, rider.EmailAddress, rider.Password)
		result, err := db.Exec(userQueryStatement)
		if err != nil {
			panic(err.Error())
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err.Error())
		}

		// Upon getting the ID, now insert into Passenger table
		queryStatement := fmt.Sprintf(`
		INSERT INTO Rider
		VALUES (%d, '%s', '%s', '%s', '%s', '%s')`,
			id,
			rider.FirstName,
			rider.LastName,
			rider.MobileNumber,
			rider.IcNumber,
			rider.CarLicNumber)

		result2, err := db.Exec(queryStatement)
		if err != nil {
			panic(err.Error())
		}
		rows_affected, err := result2.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		if rows_affected == 1 {
			w.WriteHeader(http.StatusAccepted) //202
			json.NewEncoder(w).Encode("Insert Successfully")
			fmt.Println("insert successfully")
			return
		} else {
			w.WriteHeader(http.StatusExpectationFailed) //417
			fmt.Println("Error with inserting")
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

}

func account(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println("Get account info and update account")

	w.Write([]byte(fmt.Sprintf("Change account details for: %s ", claims.EmailAddress)))
	//w.Write([]byte(fmt.Sprintf("Change account details for:%s", "your mum")))
}

// ----- Functional Tool-------//
func checkEmailIsExist(new_email string, db *sql.DB) bool {

	query := fmt.Sprintf(`SELECT email_address FROM User WHERE email_address = '%s'`, new_email)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var email string
	for results.Next() {
		err = results.Scan(&email)
		if err != nil {
			panic(err.Error())
		}
		if email == new_email {
			return true
		}
	}
	return false
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup/{user_type}", signup).Methods("POST", "OPTIONS")
	router.HandleFunc("/account", account)

	fmt.Println("Listening at port 5051")
	log.Fatal(http.ListenAndServe(":5051", router))

}
