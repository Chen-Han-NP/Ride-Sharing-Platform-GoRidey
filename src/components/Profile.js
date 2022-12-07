import React from "react";
import AuthService from "../services/auth.service";

const Profile = () => {
  const currentUser = AuthService.getCurrentUser();
  const userType = currentUser.user_type  

  return (
    <div className="container">
      <header className="jumbotron">
        <h3>
          Your Profile
        </h3>
        <br></br>
      </header>
      <p>
        <strong>First Name:</strong> {currentUser.first_name}
      </p>
      <p>
        <strong>Last Name:</strong> {currentUser.last_name}
      </p>
      <p>
        <strong>Email:</strong> {currentUser.email_address}
      </p>
      <p>
        <strong>Mobile Number:</strong> {currentUser.mobile_number}
      </p>
      {
        userType === "rider" && (
          <div>
          <p>
          <strong>Ic Number:</strong> {currentUser.ic_number}
        </p>
        <p>
          <strong>Car Lic Number:</strong> {currentUser.car_lic_number}
        </p>
        </div>
        )
      }
    </div>
  );
};

export default Profile;