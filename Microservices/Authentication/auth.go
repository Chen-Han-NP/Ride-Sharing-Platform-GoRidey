package main

import (
	//"encoding/json"

	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// ==== STRUCTs ========

type Credentials struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

// Create a struct that can be encoded into a JWT
type Claims struct {
	EmailAddress string `json:"email_address"`
	jwt.RegisteredClaims
}

type User struct {
	UserID       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type Users []User

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

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

// ====== GLOBAL VARIABLES ========
var users Users
var jwtKey = []byte("my_secret_key")
var sqlConnectionString = "root:password@tcp(127.0.0.1:3306)/"
var database = "RideSharingPlatform"

// ======= DB Functions ==========
func checkEmailIsExist(db *sql.DB, new_email string) bool {
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
func getAllUserFromDB(db *sql.DB) (Users, error) {
	users := Users{}

	results, err := db.Query("SELECT * FROM User")
	if err != nil {
		return users, err

	}
	for results.Next() {
		var user User

		err = results.Scan(&user.UserID, &user.UserType, &user.EmailAddress, &user.Password)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
func getPassenger(db *sql.DB, email_address string) (Passenger, error) {

	var passenger_found Passenger
	select_query := fmt.Sprintf(`
	SELECT u.email_address, u.password, p.first_name, p.last_name, p.mobile_number FROM User u
	INNER JOIN Passenger p
	ON u.user_id = p.passenger_id
	WHERE email_address = '%s'`, email_address)

	results, err := db.Query(select_query)
	if err != nil {
		return passenger_found, err
	}
	for results.Next() {
		err = results.Scan(&passenger_found.EmailAddress, &passenger_found.Password, &passenger_found.FirstName, &passenger_found.LastName, &passenger_found.MobileNumber)
		if err != nil {
			return passenger_found, err
		}
	}
	return passenger_found, nil
}
func getRider(db *sql.DB, email_address string) (Rider, error) {

	var rider_found Rider
	select_query := fmt.Sprintf(`
	SELECT u.email_address, u.password, r.first_name, r.last_name, r.mobile_number, r.ic_number, r.car_lic_number FROM User u
	INNER JOIN Rider r
	ON u.user_id = r.rider_id
	WHERE email_address = '%s'`, email_address)

	results, err := db.Query(select_query)
	if err != nil {
		return rider_found, err
	}
	for results.Next() {
		err = results.Scan(&rider_found.EmailAddress, &rider_found.Password, &rider_found.FirstName, &rider_found.LastName, &rider_found.MobileNumber, &rider_found.IcNumber, &rider_found.CarLicNumber)
		if err != nil {
			return rider_found, err
		}
	}
	return rider_found, nil
}

// ========= HANDLER FUNCTIONS ==========

// RETURN 200 -> Registered
// RETURN 409 -> Duplicated account (email)
// RETURN 417 -> INSERT failed
func SignUp(w http.ResponseWriter, r *http.Request) {
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
		isExist := checkEmailIsExist(db, passenger.EmailAddress)
		if isExist {
			w.WriteHeader(http.StatusConflict) //409
			fmt.Println("Duplicated account: " + passenger.EmailAddress)
			return
		}

		// Insert into User table
		userQueryStatement := fmt.Sprintf(`
		INSERT INTO User(user_type, email_address, password)
		VALUES ('%s', '%s', '%s')`, user_type, passenger.EmailAddress, passenger.Password)
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
		isExist := checkEmailIsExist(db, rider.EmailAddress)
		if isExist {
			w.WriteHeader(http.StatusConflict) //409
			fmt.Println("Duplicated account: " + rider.EmailAddress)
			return
		}

		// Insert into User table
		userQueryStatement := fmt.Sprintf(`
		INSERT INTO User(user_type, email_address, password)
		VALUES ('%s', '%s', '%s')`, user_type, rider.EmailAddress, rider.Password)
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

// RETURN Passenger OR Rider object string in JSON
func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	var user_type string
	// Receive user login information in JSON
	// and decode into User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Get the expected password from the list
	expectedPassword := ""

	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	users, err = getAllUserFromDB(db)
	if err != nil {
		panic(err.Error())
	}

	for _, user := range users {
		if creds.EmailAddress == user.EmailAddress {
			expectedPassword = user.Password
			user_type = user.UserType
		}
	}
	if expectedPassword == "" {
		w.WriteHeader(http.StatusUnauthorized) //401
		fmt.Fprintf(w, "Password not found!")
		return
	}

	//Declare the expiration time of the token to 30min
	expirationTime := time.Now().Add(30 * time.Minute)

	//Create JWT claims, which includes the email and expiry time
	claims := &Claims{
		EmailAddress: creds.EmailAddress,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	// If the user logs in with the correct credentials, this handler will then set a cookie on the client
	// side with the JWT value. Once the cookie is set on a client, it is sent along with every request henceforth

	//return the user credential
	if user_type == "passenger" {
		passenger, err := getPassenger(db, creds.EmailAddress)
		if err != nil {
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(passenger)

	} else if user_type == "rider" {
		rider, err := getRider(db, creds.EmailAddress)
		if err != nil {
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(rider)
	} else {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}
}

// CHECK COOKIE JWT SIGNATURE AND DO SOMETHING
func Welcome(w http.ResponseWriter, r *http.Request) {
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

	// Finally, return the welcome message to the user, along with their
	// email given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.EmailAddress)))

}

// Reset the token expiration time
func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code until this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
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
	// (END) The code until this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(30 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// CLEAR THE TOKEN COOKIE AND JWT
func Logout(w http.ResponseWriter, r *http.Request) {
	// immediately clear the token cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}

func main() {
	router := mux.NewRouter()

	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	users, err = getAllUserFromDB(db)
	if err != nil {
		panic(err.Error())
	}

	router.HandleFunc("/api/auth/signup/{user_type}", SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/login", Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/welcome", Welcome).Methods("GET")
	router.HandleFunc("/api/auth/refresh", Refresh)
	router.HandleFunc("/api/auth/logout", Logout)

	fmt.Println("Listening at port 5050")
	log.Fatal(http.ListenAndServe(":5050", router))

}
