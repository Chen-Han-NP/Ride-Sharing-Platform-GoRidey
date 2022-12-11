package main

import (
	//"encoding/json"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

// This struct is used as a response to the client
type CommonUser struct {
	UserID       string `json:"user_id"`
	UserType     string `json:"user_type"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MobileNumber string `json:"mobile_number"`
	IcNumber     string `json:"ic_number"`
	CarLicNumber string `json:"car_lic_number"`
}

type Ride struct {
	RideID         string `json:"ride_id" `
	PassengerID    string `json:"passenger_id"`
	PassengerName  string `json:"passenger_name"`
	PassengerPhone string `json:"passenger_phone"`
	RiderID        string `json:"rider_id"`
	RiderName      string `json:"rider_name"`
	RiderPhone     string `json:"rider_phone"`
	CarLicNumber   string `json:"car_lic_number"`
	PickupCode     string `json:"pickup_code"`
	DropoffCode    string `json:"dropoff_code"`
	RideStatus     string `json:"ride_status"`
}

// ====== GLOBAL VARIABLES ========
var sqlConnectionString = "root:password@tcp(127.0.0.1:3306)/"
var database = "RideSharingPlatform"
var jwtKey = []byte("lhdrDMjhveyEVcvYFCgh1dBR2t7GM0YJ")

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

// ======= Functions ==========
// Check if an on-going ride exist for passenger
func checkExistingRide(db *sql.DB, passenger_id string) bool {
	query := fmt.Sprintf(`SELECT COUNT(*) FROM Ride WHERE passenger_id = %s && (ride_status != 'Completed' && ride_status != 'Cancelled')`, passenger_id)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var count int
	for results.Next() {
		if err := results.Scan(&count); err != nil {
			panic(err.Error())
			return true
		}
	}
	return count > 0
}

// Check if the ride_id provided by the API exists
func checkRideIdExists(db *sql.DB, ride_id string) bool {
	query := fmt.Sprintf(`SELECT COUNT(*) FROM Ride WHERE ride_id = %s`, ride_id)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var count int
	for results.Next() {
		if err := results.Scan(&count); err != nil {
			panic(err.Error())
			return false
		}
	}
	return count > 0
}

// Get the ride info by providing ride_id
func getRide(db *sql.DB, ride_id string) (Ride, error) {
	var ride Ride
	var rideId int
	var passengerId int
	var riderId sql.NullInt64
	var riderName sql.NullString
	var riderPhone sql.NullString
	var carLicNumber sql.NullString

	select_query := fmt.Sprintf(`SELECT * FROM Ride WHERE ride_id = %s;`, ride_id)

	results, err := db.Query(select_query)
	if err != nil {
		return ride, err
	}

	for results.Next() {
		err = results.Scan(&rideId, &passengerId, &ride.PassengerName, &ride.PassengerPhone, &riderId, &riderName, &riderPhone, &carLicNumber, &ride.PickupCode, &ride.DropoffCode, &ride.RideStatus)
		if err != nil {
			return ride, err
		}
	}

	ride.RideID = strconv.Itoa(int(rideId))
	ride.PassengerID = strconv.Itoa(int(passengerId))
	if riderId.Valid {
		ride.RiderID = strconv.Itoa(int(riderId.Int64))
	}
	if riderName.Valid {
		ride.RiderName = riderName.String
	}
	if riderPhone.Valid {
		ride.RiderPhone = riderPhone.String
	}
	if carLicNumber.Valid {
		ride.CarLicNumber = carLicNumber.String
	}
	return ride, nil
}

// Get the current existing ride's ride_id and later can call the getRide func to save code
func getCurrentRideId(db *sql.DB, user_id string, user_type string) (string, error) {
	var rideIdInt int
	var rideId string

	var select_query string

	if user_type == "passenger" {
		select_query = fmt.Sprintf(`SELECT ride_id FROM Ride WHERE passenger_id = %s && ride_status = "Riding";`, user_id)
	} else if user_type == "rider" {
		select_query = fmt.Sprintf(`SELECT ride_id FROM Ride WHERE rider_id = %s && ride_status = "Riding";`, user_id)
	}

	results, err := db.Query(select_query)
	if err != nil {
		return rideId, err
	}

	for results.Next() {
		err = results.Scan(&rideIdInt)
		if err != nil {
			return rideId, err
		}
	}
	rideId = strconv.Itoa(int(rideIdInt))
	return rideId, nil
}

func getAllRides(db *sql.DB, user_type string, user_id string, ride_status string) ([]Ride, error) {
	var select_query string
	var rides []Ride

	if ride_status == "" {
		select_query = fmt.Sprintf(`SELECT * FROM Ride WHERE %s_id = %s;`, user_type, user_id)
	} else if ride_status == "Pending" { // If a rider want to see all the pending rides
		select_query = `SELECT * FROM Ride WHERE ride_status = 'Pending'`
	} else {
		select_query = fmt.Sprintf(`SELECT * FROM Ride WHERE %s_id = %s && ride_status = '%s';`, user_type, user_id, ride_status)
	}

	results, err := db.Query(select_query)
	if err != nil {
		return rides, err
	}

	for results.Next() {
		var ride Ride
		var rideId int
		var passengerId int
		var riderId sql.NullInt64
		var riderName sql.NullString
		var riderPhone sql.NullString
		var carLicNumber sql.NullString

		err = results.Scan(&rideId, &passengerId, &ride.PassengerName, &ride.PassengerPhone, &riderId, &riderName, &riderPhone, &carLicNumber, &ride.PickupCode, &ride.DropoffCode, &ride.RideStatus)
		if err != nil {
			return rides, err
		}

		ride.RideID = strconv.Itoa(int(rideId))
		ride.PassengerID = strconv.Itoa(int(passengerId))
		if riderId.Valid {
			ride.RiderID = strconv.Itoa(int(riderId.Int64))
		}
		if riderName.Valid {
			ride.RiderName = riderName.String
		}
		if riderPhone.Valid {
			ride.RiderPhone = riderPhone.String
		}
		if carLicNumber.Valid {
			ride.CarLicNumber = carLicNumber.String
		}
		rides = append(rides, ride)
	}
	return rides, nil
}

// ======= HANDLER FUNCTIONS ========
// Allow Passenger to initate a new ride
// And return the ride-info back to the user
func NewRide(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	// Verify JWT Token
	claims, err := verifyJWT(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		panic(err.Error())
		return
	}
	// Define the db
	// Connect to the db
	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Variables
	var ride Ride
	user_id := claims.UserID
	//email := claims.EmailAddress
	user_type := claims.UserType

	// Check that only passsenger can initiate the ride
	if user_type != "passenger" {
		w.WriteHeader(http.StatusNotAcceptable) // 406
		json.NewEncoder(w).Encode("Error: Only passenger is able to initiate the ride")
		return
	}

	// Check that passenger has current rides or not
	hasExistingRide := checkExistingRide(db, user_id)
	if hasExistingRide {
		w.WriteHeader(http.StatusNotAcceptable) //406
		json.NewEncoder(w).Encode("Error: Passenger has an existing ride")
		return
	}

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK) //200
		return
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusAccepted) // 202
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &ride)
		ride.PassengerID = user_id

		// Insert into Ride table
		userQueryStatement := fmt.Sprintf(`
		INSERT INTO Ride(passenger_id, passenger_name, passenger_phone, pick_up_code, drop_off_code, ride_status)
		VALUES ( %s, '%s', '%s', '%s', '%s', '%s');`, ride.PassengerID, ride.PassengerName, ride.PassengerPhone, ride.PickupCode, ride.DropoffCode, ride.RideStatus)
		result, err := db.Exec(userQueryStatement)
		if err != nil {
			panic(err.Error())
			return
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err.Error())
			return
		}

		// Set the ID in the Ride class
		ride.RideID = strconv.Itoa(int(id))

		json.NewEncoder(w).Encode(ride)
		return

	} else {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}
}

// Allow passenger to get a ride information by entering a ride id
// and return the ride-info back to the user
func GetRide(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get ride_id from params
	params := mux.Vars(r)
	ride_id := params["ride_id"]

	// Verify JWT Token
	_, err := verifyJWT(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		panic(err.Error())
		return
	}
	// Define the db
	// Connect to the db
	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method == "GET" {
		// check if ride id exists
		rideIdExists := checkRideIdExists(db, ride_id)
		if !rideIdExists {
			w.WriteHeader(http.StatusNotAcceptable) // 406
			json.NewEncoder(w).Encode("Ride ID does not exist")
			return
		}

		ride, err := getRide(db, ride_id)
		if err != nil {
			panic(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(ride)
		return

	} else {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}
}

// Allow a user to get all the rides or can choose a filter on ride status
// All ride status => [Pending, Riding, Completed, Cancelled]
// Note: Only pending request will return ALL ride records, other status type only return one belong to the user.
func AllRides(w http.ResponseWriter, r *http.Request) {

	// Check whether got query string first
	querystringmap := r.URL.Query()

	// Verify JWT Token
	claims, err := verifyJWT(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		panic(err.Error())
		return
	}

	// Variables
	var rides []Ride
	user_id := claims.UserID
	user_type := claims.UserType
	//user_email := claims.EmailAddress

	// Connect to db
	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method == "GET" {
		if len(querystringmap) == 0 {
			rides, err = getAllRides(db, user_type, user_id, "")
			if err != nil {
				w.WriteHeader(http.StatusNotAcceptable) // 406
				panic(err.Error())
				return
			}
			json.NewEncoder(w).Encode(rides)
		} else {
			for k, v := range querystringmap {
				if k == "status" {
					rides, err = getAllRides(db, user_type, user_id, v[0])
					if err != nil {
						w.WriteHeader(http.StatusNotAcceptable) // 406
						panic(err.Error())
						return
					}
					json.NewEncoder(w).Encode(rides)
				} else {
					// if the query string provided is not found
					fmt.Fprintf(w, "Un-identifiable query string!")
				}
			}
		}
	}
}

// Allow a rider to get his current on-going ride
func CurrentRide(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")

	// Verify JWT Token
	cookie, err := verifyJWT(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		panic(err.Error())
		return
	}

	user_id := cookie.UserID
	user_type := cookie.UserType

	// Define the db
	// Connect to the db
	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method == "GET" {
		// First get the ride id of the current ride -> On-going
		rideId, err := getCurrentRideId(db, user_id, user_type)
		if err != nil {
			panic(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Now call the get ride function
		ride, err := getRide(db, rideId)
		if err != nil {
			panic(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(ride)
		return

	} else {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}
}

// Allow a Rider to accept any ride been posted by passenger E.g. Rides with Pending status
func AcceptRide(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get ride_id from params
	params := mux.Vars(r)
	ride_id := params["ride_id"]

	// Verify JWT Token
	claims, err := verifyJWT(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		panic(err.Error())
		return
	}
	// Define the db
	// Connect to the db
	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// check if the user_type is rider or not
	if claims.UserType != "rider" {
		w.WriteHeader(http.StatusNotAcceptable) //406
		json.NewEncoder(w).Encode("Passenger does not have the permission to accept a ride")
		return
	}

	// check if the rider has any on-going rides or not
	all_rides, err := getAllRides(db, claims.UserType, claims.UserID, "Riding")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable) // 406
		panic(err.Error())
		return
	}
	// If the rider has on-going rides
	if len(all_rides) > 0 {
		w.WriteHeader(http.StatusNotAcceptable) //406
		json.NewEncoder(w).Encode("You have an on-going ride, cannot accept anymore rides")
		return
	}

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method == "GET" {
		ride, err := getRide(db, ride_id)

		// Check if the ride is Pending
		if ride.RideStatus != "Pending" {
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode("This ride is not pending, thus cannot accept")
			return
		}

		if err != nil {
			panic(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Get the rider info from the db
		var rider Rider
		var select_query string
		select_query = fmt.Sprintf(`
			SELECT u.email_address, u.password, r.first_name, r.last_name, r.mobile_number, r.ic_number, r.car_lic_number FROM User u
			INNER JOIN Rider r
			ON u.user_id = r.rider_id
			WHERE email_address = '%s'`, claims.EmailAddress)

		results, err := db.Query(select_query)
		if err != nil {
			w.WriteHeader(http.StatusNotFound) //404
			panic(err.Error())
			return
		}
		for results.Next() {
			err = results.Scan(&rider.EmailAddress, &rider.Password, &rider.FirstName, &rider.LastName, &rider.MobileNumber, &rider.IcNumber, &rider.CarLicNumber)
			if err != nil {
				w.WriteHeader(http.StatusNotFound) //404
				panic(err.Error())
				return
			}
		}

		ride.RiderID = claims.UserID
		ride.RideStatus = "Riding"
		ride.RiderName = rider.FirstName + " " + rider.LastName
		ride.RiderPhone = rider.MobileNumber
		ride.CarLicNumber = rider.CarLicNumber

		update_statement := fmt.Sprintf(`UPDATE Ride
				SET rider_id = %s, rider_name = '%s', rider_phone = '%s', car_lic_number = '%s', ride_status = '%s'
				WHERE ride_id = %s;`, ride.RiderID, ride.RiderName, ride.RiderPhone, ride.CarLicNumber, ride.RideStatus, ride.RideID)

		result, err := db.Exec(update_statement)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			panic(err.Error())
			return
		}
		rows_affected, err := result.RowsAffected()
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			panic(err.Error())
			return
		}

		if rows_affected == 1 {
			w.WriteHeader(http.StatusAccepted) //202
			// If successful, return the new ride info to the user
			json.NewEncoder(w).Encode(ride)
			return
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// Allow Rider the complete the ride after they have reached the destination,
// The ride status changes to "Completed"
func CompleteRide(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get ride_id from params
	params := mux.Vars(r)
	ride_id := params["ride_id"]

	// Verify JWT Token
	claims, err := verifyJWT(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		panic(err.Error())
		return
	}
	// check if the user_type is rider or not
	if claims.UserType != "rider" {
		w.WriteHeader(http.StatusNotAcceptable) //406
		json.NewEncoder(w).Encode("Passenger does not have the permission to complete a ride")
		return
	}

	// Define the db
	// Connect to the db
	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method == "GET" {
		ride, err := getRide(db, ride_id)
		if err != nil {
			panic(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Check if the ride is Riding
		if ride.RideStatus != "Riding" {
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode("This ride is not riding, thus cannot complete")
			return
		}

		ride.RideStatus = "Completed"

		update_statement := fmt.Sprintf(`UPDATE Ride
		SET ride_status = '%s'
		WHERE ride_id = %s;`, ride.RideStatus, ride.RideID)

		result, err := db.Exec(update_statement)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			panic(err.Error())
			return
		}
		rows_affected, err := result.RowsAffected()
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			panic(err.Error())
			return
		}

		if rows_affected == 1 {
			w.WriteHeader(http.StatusAccepted) //202
			json.NewEncoder(w).Encode("Ride is completed!")
			return
		}
	} else { // Other request method
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// Only passenger is able to cancel the ride only when the Ride status is "Pending"
// When a rider has Accpeted the ride, the passenger is unable to cancel the ride
// The passenger will have to communicate with the Rider to cancel the ride for him when the ride status is "Riding"
func CancelRide(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get ride_id from params
	params := mux.Vars(r)
	ride_id := params["ride_id"]

	// Verify JWT Token
	claims, err := verifyJWT(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		panic(err.Error())
		return
	}

	// Define the db
	// Connect to the db
	db, err := sql.Open("mysql", sqlConnectionString+database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method == "GET" {
		ride, err := getRide(db, ride_id)
		if err != nil {
			panic(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// 1. Check if ride cancellation policy works for Passenger
		if claims.UserType == "passenger" {
			if ride.RideStatus != "Pending" {
				w.WriteHeader(http.StatusNotAcceptable)
				json.NewEncoder(w).Encode("Error: You are not allowed to cancel the ride")
				return
			}
		}

		// 2. Check if the ride cancellation policy works for Rider
		if claims.UserType == "rider" {
			if ride.RideStatus != "Riding" {
				w.WriteHeader(http.StatusNotAcceptable)
				json.NewEncoder(w).Encode("Error: You are not allowed to cancel the ride")
				return
			}
		}

		// If can, change the ride status to Cancelled
		ride.RideStatus = "Cancelled"
		update_statement := fmt.Sprintf(`UPDATE Ride
		SET ride_status = '%s'
		WHERE ride_id = %s;`, ride.RideStatus, ride.RideID)

		result, err := db.Exec(update_statement)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			panic(err.Error())
			return
		}
		rows_affected, err := result.RowsAffected()
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			panic(err.Error())
			return
		}

		if rows_affected == 1 {
			w.WriteHeader(http.StatusAccepted) //202
			json.NewEncoder(w).Encode("Ride is cancelled!")
			return
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func main() {
	router := mux.NewRouter()
	/*
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost", "http://localhost:3000"},
			AllowCredentials: true,
			Debug:            true,
		})
	*/

	router.HandleFunc("/api/ride/newride", NewRide).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ride/getride/{ride_id}", GetRide).Methods("GET", "PUT", "OPTIONS")
	router.HandleFunc("/api/ride/current", CurrentRide).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ride/allrides", AllRides).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ride/accept/{ride_id}", AcceptRide).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ride/complete/{ride_id}", CompleteRide).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ride/cancel/{ride_id}", CancelRide).Methods("GET", "OPTIONS")

	fmt.Println("Listening at port 5052")
	//handler := c.Handler(router)

	//log.Fatal(http.ListenAndServe(":5052", handler))
	log.Fatal(http.ListenAndServe(":5052", router))
}
