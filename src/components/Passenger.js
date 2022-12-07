import React from "react";
import AuthService from "../services/auth.service";

const Passenger = () => {
  const currentUser = AuthService.getCurrentUser();
  
  return (
    <div className="container">
      <header className="passengerPage">
        <h3>Passenger page</h3>
        <br></br>
      </header>
      <h4>Welcome passenger {currentUser.first_name} {currentUser.last_name}</h4>
    </div>
  );
};

export default Passenger; 