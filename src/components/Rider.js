import AuthService from "../services/auth.service";
import RideServices from "../services/ride.service";
import React, { useState, useRef } from "react";
import { useNavigate } from "react-router-dom";
import Form from "react-validation/build/form";
import Input from "react-validation/build/input";
import CheckButton from "react-validation/build/button";


const Rider = () => {
  // Calling API
  RideServices.currentRide()
  RideServices.allPendingRides()

  var message = "";
  const navigate = useNavigate();
  
  const currentUser = AuthService.getCurrentUser();
  const currentRide = RideServices.getCurrentRide();
  var currentPendingRides = RideServices.getAllPendingRides();


  const handleAllPendingRides = () => {
    //e.preventDefault();

    RideServices.allPendingRides().then(
        () => {
          navigate("/rider");
          window.location.reload();
        },
        (error) => {
          const resMessage =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();

          message = resMessage;
        }
      );
  };

  function AcceptRideButton({ride_id}) {

    const handleAcceptRide = (e, ride_id) => {
      console.log("HAHA" + ride_id)
      e.preventDefault();


      RideServices.acceptRide(ride_id).then(
          () => {
            navigate("/rider");
            window.location.reload();

            localStorage.setItem("ride", currentPendingRides.find(o => o.ride_id === ride_id));

          },
          (error) => {
            const resMessage =
              (error.response &&
                error.response.data &&
                error.response.data.message) ||
              error.message ||
              error.toString();

              message = resMessage;
          }
        );
        
    };

  
    return (
      <button className="btn btn-secondary" onClick={(e) => handleAcceptRide(e, ride_id)}>
        Click me
      </button>
    );
    
  }



  return (
    <div className="container">
      <header className="riderPage">
      <h3>Welcome Rider! {currentUser.first_name} {currentUser.last_name}</h3>
      </header>
      <br></br>

      { // If there is no rides happening at the moment and 
        !currentRide ? (
          
          <div>
            <p><strong>You do not have any on-going rides</strong></p>
            <p><strong>Please press the button to check for any in-coming ride requests</strong></p>
            <button className="btn btn-primary" onClick={handleAllPendingRides}>
              Check
            </button>
            <br></br>
            <br></br>

            {currentPendingRides && (
            <table className="table">
            <thead>
                <tr>
                    <th>S.N</th>
                    <th>Ride ID</th>
                    <th>Passenger ID</th>
                    <th>Pickup Postal Code</th>
                    <th>Dropoff Postal Code</th>
                    <th>Ride Status</th>
                    <th>Accept?</th>
                </tr>
            </thead>
            <tbody>
            {
                currentPendingRides.map((data, index)=>{
                    return(
                        <tr key={index}>
                            <td>{index+1}</td>
                            <td>{data.ride_id}</td>
                            <td>{data.passenger_id}</td>
                            <td>{data.pickup_code}</td>
                            <td>{data.dropoff_code}</td>
                            <td>{data.ride_status}</td>
                            <td><AcceptRideButton ride_id={data.ride_id} /> </td>

                        </tr>
                    )
                })
            }
            </tbody>
        </table>

            )}
            

            {message && (
              <div className="form-group">
                <div className="alert alert-danger" role="alert">
                  {message}
                </div>
              </div>
            )}
            </div>


        ) : (
          <div>
            You have ride
            </div>

        ) 
      }



    </div>

  );
  };
  
  export default Rider; 