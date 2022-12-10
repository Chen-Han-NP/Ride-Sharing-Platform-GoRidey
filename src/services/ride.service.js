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



const newride = (pickup_code, dropoff_code, ride_status) => {

/*
    return axios.post(RIDE_URL + "newride", {
        "pickup_code": pickup_code,
        "dropoff_code": dropoff_code ,
        "ride_status": ride_status
    }, axiosConfig)
            .then((response) => {
                localStorage.setItem("ride", JSON.stringify(response.data));

        return response.data;
      });
*/



    
      
};

const allrides = () => {
    var FULL_URL = RIDE_URL + "allrides"
    var URL = "http://localhost:5050/api/auth/welcome"
    
    return axios.get(FULL_URL, axiosConfig)
            .then((response) => {
                localStorage.setItem("allrides", JSON.stringify(response.data));
        return response.data;
      });
      


};

const getCurrentRide = () => {
    return JSON.parse(localStorage.getItem("ride"));
};



const getAllRides = () => {
    
    return JSON.parse(localStorage.getItem("allrides"))
}


const RideServices = {
    newride,
    allrides,
    getCurrentRide,
    getAllRides
}

export default RideServices;