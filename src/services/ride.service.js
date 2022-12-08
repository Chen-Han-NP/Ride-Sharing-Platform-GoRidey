import axios from "axios";

const RIDE_URL = "http://localhost:5052/api/ride/"

axios.defaults.withCredentials = true

let axiosConfig = {
    headers: {
        'Content-Type': 'text/plain'
    },
    withCredentials : true
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
    

    return axios.get(RIDE_URL + "allrides", axiosConfig)
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