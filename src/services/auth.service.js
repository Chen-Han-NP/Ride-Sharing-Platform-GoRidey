import axios from "axios";

const AUTH_URL = "http://localhost:5050/api/auth/"

axios.defaults.withCredentials = true
let axiosConfig = {
    headers: {
        'Content-Type': 'text/plain'
    },
    withCredentials : true
}


const register_passenger = (email_address, password, first_name, last_name, mobile_number) => {
    return axios.post(AUTH_URL + "signup/passenger", {
      "email_address" : email_address,
      "password": password,
      "first_name" : first_name,
      "last_name" : last_name,
      "mobile_number": mobile_number
    }, axiosConfig)
};

const register_rider = (email_address, password, first_name, last_name, mobile_number, ic_number, car_lic_number) => {

    return axios.post(AUTH_URL + "signup/rider", {
        "email_address" : email_address,
        "password": password,
        "first_name" : first_name,
        "last_name" : last_name,
        "mobile_number": mobile_number,
        "ic_number": ic_number,
        "car_lic_number": car_lic_number
    }, axiosConfig);
};


const login = (email_address, password) => {

    return axios.post(AUTH_URL + "login", {
        "email_address": email_address,
        "password": password }, 
        axiosConfig)
            .then((response) => {
                
            if (response.data.email_address) {
                delete response.data.password
                localStorage.setItem("user", JSON.stringify(response.data));
            }
        return response.data;
      });
      
};
  
const logout = () => {
    localStorage.removeItem("user");
    return axios.post(AUTH_URL + "logout").then((response) => {
      return response.data;
    });
};
  
const getCurrentUser = () => {
    return JSON.parse(localStorage.getItem("user"));
};
  
const AuthService = {
    register_passenger,
    register_rider,
    login,
    logout,
    getCurrentUser,
}

  
export default AuthService;



