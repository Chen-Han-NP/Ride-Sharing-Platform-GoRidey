import AuthService from "../services/auth.service";
import RideServices from "../services/ride.service";
import React from "react";
import { useNavigate } from "react-router-dom";

const currentUser = AuthService.getCurrentUser();

function Table({tableData}){
  return(   
      <table className="table">
          <thead>
              <tr>
                  <th>ID</th>
                  <th>DateTime</th>
                  {currentUser.user_type === "passenger" ?(
                    <th>Rider Name</th>
                  ) : (
                    <th>Passenger Name</th>
                  )}
                  <th>Pickup Postal Code</th>
                  <th>Pickup Datetime</th>
                  <th>Dropoff Postal Code</th>
                  <th>Dropoff Datetime</th>
                  <th>Ride Status</th>
              </tr>
          </thead>
          <tbody>
          {
              tableData.map((data, index)=>{
                  return(
                      <tr key={index}>
                          <td>{data.ride_id}</td>
                          <td>{data.ride_dt}</td>
                          {
                            currentUser.user_type === "passenger" ? (
                              <td>{data.rider_name}</td>
                            ) : (
                              <td>{data.passenger_name}</td>
                            )
                          }
                          <td>{data.pickup_code}</td>
                          <td>{data.pickup_dt}</td>
                          <td>{data.dropoff_code}</td>
                          <td>{data.dropoff_dt}</td>
                          <td>{data.ride_status}</td>
                      </tr>
                  )
              })
          }
          </tbody>
      </table>
  )
}

const RideHistory = () => {
  var rides = RideServices.getAllRides();
  var message = "";
  const navigate = useNavigate();

  // Sort the rides in reverse chronological order



  function RefreshButton() {
    const handleRefreshButton = (e) => {
      e.preventDefault();

      RideServices.allrides("").then(
          (response) => {
            navigate("/rides");
            window.location.reload();
            rides = response;
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
    <button className="btn btn-success" onClick={(e) => handleRefreshButton(e)}>
      Refresh
    </button>
  );
}

  if (rides == null){
    return (
      <p><strong>You do not have any rides</strong></p>
    )
  } 
  return (
    <React.Fragment>
    <div className="rideHistory">
    <header>
      <h3>Ride History</h3>
      </header>
      <br></br>
      <RefreshButton />

      <div className="row">
            {message && (
              <div className="form-group">
                <div className="alert alert-danger" role="alert">
                  {message}
                </div>
              </div>
            )}
          <div className="col-sm-8">
          <Table tableData={rides}/>
          </div>
          <div className="col-sm-4">
          </div>
      </div>
  </div>
    </React.Fragment>
  )

  
};

export default RideHistory; 