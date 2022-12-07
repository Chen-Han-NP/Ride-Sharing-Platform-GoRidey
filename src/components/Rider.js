import React from "react";
import AuthService from "../services/auth.service";

const Rider = () => {
  const currentUser = AuthService.getCurrentUser();

  return (
    <div className="container">
      <header className="riderPage">
        <h3>Rider page</h3>
        <br></br>
      </header>


      <h4>Welcome rider {currentUser.first_name} {currentUser.last_name}</h4>
    </div>
  );
  };
  
  export default Rider; 