import axios from "axios";

const AUTH_URL = "http://localhost:5050/login"
const REGISTER = "http://localhost:5051"


const register_passenger = (email_address, password, first_name, last_name, mobile_number) => {
    return axios.post(REGISTER + "signup/passenger", {
      "email_address" : email_address,
      "password": password,
      "first_name" : first_name,
      "last_name" : last_name,
      "mobile_number": mobile_number
    });
};

const register_rider = (email_address, password, first_name, last_name, mobile_number, ic_number, car_lic_number) => {
    return axios.post(REGISTER + "signup/rider", {
        "email_address" : email_address,
        "password": password,
        "first_name" : first_name,
        "last_name" : last_name,
        "mobile_number": mobile_number,
        "ic_number": ic_number,
        "car_lic_number": car_lic_number
    });
};


const login = (email_address, password) => {
    return axios.post(AUTH_URL + "login", {
        "email_address": email_address,
        "password": password,
      })
      .then((response) => {
        if (response.data.email_address) {
          localStorage.setItem("user", JSON.stringify(response.data));
        }
        return response.data;
      });
};
  
const logout = () => {
    localStorage.removeItem("user");
    return axios.post(AUTH_URL + "signout").then((response) => {
      return response.data;
    });
};
  
  const getCurrentUser = () => {
    return JSON.parse(localStorage.getItem("user"));
  };
  
  const AuthService = {
    register,
    login,
    logout,
    getCurrentUser,
  }
  
  export default AuthService;



