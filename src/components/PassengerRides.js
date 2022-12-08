import AuthService from "../services/auth.service";
import RideServices from "../services/ride.service";
import React, { useState, useRef } from "react";


const PassengerRides = () => {
  const currentUser = AuthService.getCurrentUser();

  RideServices.allrides()
  var getAllRides = RideServices.getAllRides()
  console.log(getAllRides)

};

export default PassengerRides; 