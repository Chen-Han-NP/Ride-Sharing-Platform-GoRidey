package main

import (
	//"encoding/json"

	"database/sql"
	"encoding/json"
	"fmt"
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

func get_all_users_from_db(db *sql.DB) (Users, error) {
	users := Users{}

	results, err := db.Query("SELECT * FROM User")
	if err != nil {
		return users, err

	}
	for results.Next() {
		var user User

		err = results.Scan(&user.UserID, &user.EmailAddress, &user.Password)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
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

	users, err = get_all_users_from_db(db)
	if err != nil {
		panic(err.Error())
	}

	for _, user := range users {
		if creds.EmailAddress == user.EmailAddress {
			expectedPassword = user.Password
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

}

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

	users, err = get_all_users_from_db(db)
	if err != nil {
		panic(err.Error())
	}

	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/welcome", Welcome).Methods("GET")
	router.HandleFunc("/refresh", Refresh)
	router.HandleFunc("/logout", Logout)

	fmt.Println("Listening at port 5050")
	log.Fatal(http.ListenAndServe(":5050", router))

}
