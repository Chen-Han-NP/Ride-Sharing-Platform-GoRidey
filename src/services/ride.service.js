// This ride.service.js containing all the API request to the microservices - Ride.go required in the project.
import axios from "axios";

const RIDE_URL = "http://localhost:5052/api/ride/"

// This allows the react to send cookie over to the golang server for authentication
axios.defaults.withCredentials = true
let axiosConfig = {
    headers: {
        'Content-Type': 'text/plain',
    },
    withCredentials : true,
}

// For passenger to initiate a new ride
// Set the current 'ride' in the localstorage
const newride = (user_name, user_phone, pickup_code, dropoff_code, ride_status) => {

    const postBody = {
        "passenger_name" : user_name,
        "passenger_phone": user_phone,
        "pickup_code": pickup_code,
        "dropoff_code": dropoff_code ,
        "ride_status": ride_status
    }

    return axios.post(RIDE_URL + "newride", postBody , axiosConfig)
            .then((response) => {
                
                localStorage.setItem("ride", JSON.stringify(response.data));
        return response.data;
      });
};

// This allows the passenger/rider to get their current ride in the status
// For them to check and get updated to their ride status if changed
const currentRide = () => {
    return axios.get(RIDE_URL + "current" , axiosConfig)
    .then((response) => {
        if (response.data.ride_id === 0) {
            return
        } else {
            localStorage.setItem("ride", JSON.stringify(response.data));
            return response.data;
        }

    });
}


function padTo2Digits(num) {
    return num.toString().padStart(2, '0');
  }
  
function formatDate(date) {
return (
    [
    date.getFullYear(),
    padTo2Digits(date.getMonth() + 1),
    padTo2Digits(date.getDate()),
    ].join('-') +
    ' ' +
    [
    padTo2Digits(date.getHours()),
    padTo2Digits(date.getMinutes()),
    padTo2Digits(date.getSeconds()),
    ].join(':')
);
}
  

// This allows the passenger/rider to retrieve all their past ride from the server
// IN REVERSE CHRONOLOGICAL ORDER
// They can also choose to filter all the rides according to the ride status
const allrides = (status) => {
    var FULL_URL;
   // var status_list = ["Pending", "Riding", "Completed", "Cancelled"]

    if (status === ""){
        FULL_URL = RIDE_URL + "allrides"
        return axios.get(FULL_URL, axiosConfig)
            .then((response) => {
                var rides = response.data.map(obj => {
                    return {...obj, ride_dt: new Date(obj.ride_dt)};
                });
                // Sort in reverse chronological order
                var sortedDesc = rides.sort(
                    (objA, objB) => Number(objB.ride_dt) - Number(objA.ride_dt),
                  )
                sortedDesc = sortedDesc.map(obj => {
                    return {...obj, ride_dt: formatDate(obj.ride_dt)}
                });

                localStorage.setItem("allrides", JSON.stringify(sortedDesc));
                return sortedDesc;
            });
        
    } else if (status === "Pending") {
        FULL_URL = RIDE_URL + "allrides?status=Pending" 
        return axios.get(FULL_URL, axiosConfig)
        .then((response) => {
            localStorage.setItem("pendingrides", JSON.stringify(response.data));
            return response.data;
        });
        
    } else {
        return
    }
      
};

const getCurrentRide = () => {
    const ride = localStorage.getItem("ride")
    if (ride === null) {
        return null
    } 
    return JSON.parse(localStorage.getItem("ride"));
};

const getAllRides = () => {
    const rides = localStorage.getItem("allrides")
    if (rides === null) {
        return null
    } 
    return JSON.parse(localStorage.getItem("allrides"))
    

}

// For Rider to retrieve a list pending rides from passengers
const allPendingRides = () => {
    var FULL_URL = RIDE_URL + "allrides?status=Pending" 

    return axios.get(FULL_URL, axiosConfig)
    .then((response) => {
        localStorage.setItem("pendingrides", JSON.stringify(response.data));
        return response.data;
    });

};

const getAllPendingRides = () => {
    return JSON.parse(localStorage.getItem("pendingrides"));
}

// For rider to accept a pending ride req to start a trip
const acceptRide = (ride_id) => {

    var FULL_URL = RIDE_URL + "accept/" + ride_id
    
    return axios.get(FULL_URL, axiosConfig)
            .then((response) => {
                return response.data;
      });
      
}

// For rider to complete the rider upon reach destination to end a trip
const completeRide = (ride_id) => {
    var FULL_URL = RIDE_URL + "complete/" + ride_id
    
    return axios.get(FULL_URL, axiosConfig)
            .then((response) => {
                return response.data;
      });
      
}

// For both Rider and Passenger
// Rider: Only when the ride status is "Riding"
// Passenger: Only when the ride status is "Pending"
const cancelRide = (ride_id) => {
    var FULL_URL = RIDE_URL + "cancel/" + ride_id
    return axios.get(FULL_URL, axiosConfig)
            .then((response) => {
                return response.data;
      });
      
}


const RideServices = {
    newride,
    currentRide,
    allrides,
    acceptRide,
    cancelRide,
    completeRide,
    getCurrentRide,
    getAllRides,
    allPendingRides,
    getAllPendingRides

}

export default RideServices;