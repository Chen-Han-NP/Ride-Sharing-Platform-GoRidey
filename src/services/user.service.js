import axios from "axios";
import AuthService from "./auth.service";

const API_URL = "http://localhost:5051/api/user/";
const currentUser = AuthService.getCurrentUser();

axios.defaults.withCredentials = true

let axiosConfig = {
    headers: {
        'Content-Type': 'text/plain',
    },
    withCredentials : true,
}


const updatePassengerProfile = (first_name, last_name, mobile_number) => {

    const putBody = {
        "first_name" : first_name,
        "last_name": last_name,
        "mobile_number": mobile_number
    }
      return axios.post(API_URL + "getuser", putBody , axiosConfig)
        .then((response) => {
            var new_user = {
              "user_id" : currentUser.user_id,
              "user_type": currentUser.user_type,
              "email_address": currentUser.email_address,
              "first_name": response.data.first_name ,
              "last_name": response.data.last_name,
              "mobile_number": response.data.mobile_number,
              "car_lic_number": currentUser.car_lic_number,
              "ic_number": currentUser.ic_number
            }
            localStorage.setItem("user", JSON.stringify(new_user));
            return response.data;
        });
};

const updateRiderProfile = (first_name, last_name, mobile_number, car_lic_number) => {
  const currentUser = AuthService.getCurrentUser();

  const putBody = {
    "first_name" : first_name,
    "last_name": last_name,
    "mobile_number": mobile_number,
    "car_lic_number": car_lic_number
}

      return axios.post(API_URL + "getuser", putBody , axiosConfig)
      .then((response) => {
          var new_user = {
            "user_id" : currentUser.user_id,
            "user_type": currentUser.user_type,
            "email_address": currentUser.email_address,
            "first_name": response.data.first_name ,
            "last_name": response.data.last_name,
            "mobile_number": response.data.mobile_number,
            "car_lic_number": response.data.car_lic_number,
            "ic_number": currentUser.ic_number
          }
          localStorage.setItem("user", JSON.stringify(new_user));
          return response.data;
      });
};


const updatePassword = (password) => {

  const putBody = {
    "password" : password
}

  return axios.post(API_URL + "password" , putBody, axiosConfig)
              .then((response) => {
                return response.status;
            });
};


const UserService = {
  updatePassengerProfile,
  updateRiderProfile,
  updatePassword
}

export default UserService;