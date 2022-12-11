import AuthService from "../services/auth.service";
import RideServices from "../services/ride.service";
import React, { useState, useRef } from "react";
import Table from "../objects/Table.js"


const RideHistory = () => {
  const currentUser = AuthService.getCurrentUser();
  RideServices.allrides()

  const getAllRides = RideServices.getAllRides()

 


  return (
    <React.Fragment>
    <div className="container">
      <div className="row">
          <div className="col-sm-8">
          <Table tableData={getAllRides}/>
          </div>
          <div className="col-sm-4">
          </div>
      </div>
  </div>
      
    </React.Fragment>


  )


};

export default RideHistory; 