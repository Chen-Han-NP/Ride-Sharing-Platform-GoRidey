import React from "react";
import AuthService from "../services/auth.service";

const Profile = () => {
  const currentUser = AuthService.getCurrentUser();

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
    </div>
  );
};

export default Profile;