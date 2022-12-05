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
	UserID       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type Claims struct {
	EmailAddress string `json:"email_address"`
	UserType     string `json:"user_type"`
	UserID       string `json:"user_id"`
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

type PassengerMap map[string]Passenger
type RiderMap map[string]Rider

// ====== GLOBAL VARIABLES ========
var sqlConnectionString = "root:password@tcp(127.0.0.1:3306)/"
var database = "RideSharingPlatform"
var jwtKey = []byte("my_secret_key")

var passengerMap PassengerMap
var riderMap RiderMap

// ====== FUNCTONS =========
func verifyJWT(w http.ResponseWriter, r *http.Request) (Claims, error) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return Claims{}, err
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return Claims{}, err
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
			return *claims, err
		}
		w.WriteHeader(http.StatusBadRequest)
		return *claims, err
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return *claims, err
	}
	// Token is valid
	return *claims, nil
}

// ======= HANDLER FUNCTIONS ========
// GET passenger info
// UPDATE Passenger in the db
func GetUser(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	claims, err := verifyJWT(w, r)

	if err != nil {
		panic(err.Error())
		return
	}

	email_address := claims.EmailAddress
	user_type := claims.UserType
	user_id := claims.UserID

	w.Header().Set("Access-Control-Allow-Origin", "*")

	var passenger_found Passenger
	var rider_found Rider
	var select_query string

	if r.Method == "GET" {
		if user_type == "passenger" {
			select_query = fmt.Sprintf(`
			SELECT u.email_address, u.password, p.first_name, p.last_name, p.mobile_number FROM User u
			INNER JOIN Passenger p
			ON u.user_id = p.passenger_id
			WHERE email_address = '%s'`, email_address)

			results, err := db.Query(select_query)
			if err != nil {
				w.WriteHeader(http.StatusNotFound) //404
				panic(err.Error())
				return
			}
			for results.Next() {
				err = results.Scan(&passenger_found.EmailAddress, &passenger_found.Password, &passenger_found.FirstName, &passenger_found.LastName, &passenger_found.MobileNumber)
				if err != nil {
					w.WriteHeader(http.StatusNotFound) //404
					panic(err.Error())
					return
				}
			}
			w.WriteHeader(http.StatusAccepted) //202
			passengerMap[user_id] = passenger_found
			json.NewEncoder(w).Encode(passengerMap)
			return

		} else if user_type == "rider" {
			select_query = fmt.Sprintf(`
			SELECT u.email_address, u.password, r.first_name, r.last_name, r.mobile_number, r.ic_number, r.car_lic_number FROM User u
			INNER JOIN Rider r
			ON u.user_id = r.rider_id
			WHERE email_address = '%s'`, email_address)

			results, err := db.Query(select_query)
			if err != nil {
				w.WriteHeader(http.StatusNotFound) //404
				panic(err.Error())
				return
			}
			for results.Next() {
				err = results.Scan(&rider_found.EmailAddress, &rider_found.Password, &rider_found.FirstName, &rider_found.LastName, &rider_found.MobileNumber, &rider_found.IcNumber, &rider_found.CarLicNumber)
				if err != nil {
					w.WriteHeader(http.StatusNotFound) //404
					panic(err.Error())
					return
				}
			}
			w.WriteHeader(http.StatusAccepted) //202
			riderMap[user_id] = rider_found
			json.NewEncoder(w).Encode(riderMap)
			return
		}

	} else if r.Method == "PUT" {
		var update_query string

		if user_type == "passenger" {

			// get the body of our POST request
			reqBody, _ := ioutil.ReadAll(r.Body)
			// unmarshal this into a new Diploma struct
			var passenger Passenger

			json.Unmarshal(reqBody, &passenger)

			update_query = fmt.Sprintf(`
UPDATE Passenger p
INNER JOIN User u
ON p.passenger_id = u.user_id
SET u.password = '%s', p.first_name = '%s', p.last_name = '%s', p.mobile_number = '%s'
WHERE u.user_id = %s;`, passenger.Password, passenger.FirstName, passenger.LastName, passenger.MobileNumber, user_id)

			// Update to db
			result, err := db.Exec(update_query)

			if err != nil {
				panic(err.Error())
			}
			rows_affected, err := result.RowsAffected()
			if err != nil {
				panic(err.Error())
			}
			if rows_affected == 1 {
				w.WriteHeader(http.StatusAccepted) //202
				passengerMap[user_id] = passenger
				json.NewEncoder(w).Encode(passengerMap)
				return
			} else {
				w.WriteHeader(http.StatusNotFound) // 404
				return
			}

		} else if user_type == "rider" {

			// get the body of our POST request
			reqBody, _ := ioutil.ReadAll(r.Body)
			// unmarshal this into a new Diploma struct
			var rider Rider

			json.Unmarshal(reqBody, &rider)

			update_query = fmt.Sprintf(`
UPDATE Rider r
INNER JOIN User u
ON r.rider_id = u.user_id
SET u.password = '%s', r.first_name = '%s', r.last_name = '%s', r.mobile_number = '%s', r.car_lic_number = '%s'
WHERE u.user_id = %s;`, rider.Password, rider.FirstName, rider.LastName, rider.MobileNumber, rider.CarLicNumber, user_id)

			// Update to db
			result, err := db.Exec(update_query)

			if err != nil {
				panic(err.Error())
			}

			rows_affected, err := result.RowsAffected()
			if err != nil {
				panic(err.Error())
			}
			if rows_affected == 1 {
				w.WriteHeader(http.StatusAccepted) //202
				riderMap[user_id] = rider
				json.NewEncoder(w).Encode(riderMap)
				return
			} else {
				w.WriteHeader(http.StatusNotFound) // 404
				return
			}
		}

	} else if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK) //200
		return
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/user", GetUser).Methods("GET", "PUT", "OPTIONS")

	passengerMap = PassengerMap{}
	riderMap = RiderMap{}

	fmt.Println("Listening at port 5051")
	log.Fatal(http.ListenAndServe(":5051", router))

}
