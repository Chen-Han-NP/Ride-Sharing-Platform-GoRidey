import axios from "axios";
const API_URL = "http://localhost:5051/api/";


const getUserBoard = () => {
  return axios.get(API_URL + "user");
};



const UserService = {
  getUserBoard
}

export default UserService;