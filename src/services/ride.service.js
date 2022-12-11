import axios from "axios";
import React from "react";
import ReactDOM from "react-dom";


const RIDE_URL = "http://localhost:5052/api/ride/"

axios.defaults.withCredentials = true

let axiosConfig = {
    headers: {
        'Content-Type': 'text/plain',
       // "Access-Control-Allow-Origin": "http://localhost:3000"
    },
    withCredentials : true,
}

// For Passenger
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

const currentRide = () => {
    return axios.get(RIDE_URL + "current" , axiosConfig)
    .then((response) => {
        if (response.data.ride_id == 0) {
            return
        } else {
            localStorage.setItem("ride", JSON.stringify(response.data));
            return response.data;
        }

    });
}

const allrides = (status) => {
    var FULL_URL;
    var status_list = ["Pending", "Riding", "Completed", "Cancelled"]

    if (status === ""){
        FULL_URL = RIDE_URL + "allrides"
        return axios.get(FULL_URL, axiosConfig)
            .then((response) => {
                localStorage.setItem("allrides", JSON.stringify(response.data));
                return response.data;
            });
        
    } else if (status === "Pending") {
        FULL_URL = RIDE_URL + "allrides?status=Pending" 
        return axios.get(FULL_URL, axiosConfig)
        .then((response) => {
            localStorage.setItem("pendingrides", JSON.stringify(response.data));
            return response.data;
        });
        
    } else if (status_list.includes(status)) {



    } else {
        return 
    }
      
};

const getCurrentRide = () => {
    const ride = localStorage.getItem("ride")
    if (ride == null) {
        return null
    } 
    return JSON.parse(localStorage.getItem("ride"));
};

const getAllRides = () => {
    const rides = localStorage.getItem("allrides")
    if (rides == null) {
        return null
    } 
    return JSON.parse(localStorage.getItem("allrides"))
    

}

// For Rider

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


const acceptRide = (ride_id) => {

    var FULL_URL = RIDE_URL + "accept/" + ride_id
    
    return axios.get(FULL_URL, axiosConfig)
            .then((response) => {
                return response.data;
      });
      
}

const completeRide = (ride_id) => {
    var FULL_URL = RIDE_URL + "complete/" + ride_id
    
    return axios.get(FULL_URL, axiosConfig)
            .then((response) => {
                return response.data;
      });
      
}

// For both Rider and Passenger
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