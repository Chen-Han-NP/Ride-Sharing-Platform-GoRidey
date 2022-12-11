import AuthService from "../services/auth.service";
import RideServices from "../services/ride.service";
import React from "react";
import { useNavigate } from "react-router-dom";


const Rider = () => {

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
      e.preventDefault();
      RideServices.acceptRide(ride_id).then(
          (response) => {
            navigate("/rider");
            window.location.reload();
            localStorage.setItem("ride", JSON.stringify(response));
            localStorage.removeItem('pendingrides')
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
      <button className="btn btn-secondary" onClick={(e) => {

        const confirmBox = window.confirm(
          "Do you really want to accept this ride request?"
        )
        if (confirmBox === true) {
          handleAcceptRide(e, ride_id)
        }
      }}>
        Accept
      </button>
    );
  }

  function CancelRideButton({ride_id}) {

    const handleCancelRide = (e, ride_id) => {
      e.preventDefault();

      RideServices.cancelRide(ride_id).then(
          () => {
            navigate("/rider");
            window.location.reload();
            localStorage.removeItem("ride")
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
      <button className="btn btn-danger" onClick={(e) => 
      {
        const confirmBox = window.confirm(
          "Do you really want to cancel this ride?"
        )
        if (confirmBox === true) {
          handleCancelRide(e, ride_id)
        }
      }}>
        Cancel this ride
      </button>
    );
  }

  function CompleteRideButton({ride_id}) {

    const handleCompleteRide = (e, ride_id) => {
      e.preventDefault();

      RideServices.completeRide(ride_id).then(
          () => {
            navigate("/rider");
            window.location.reload();
            localStorage.removeItem("ride")
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
      <button className="btn btn-success" onClick={(e) => 
      {
        const confirmBox = window.confirm(
          "You must have reached the destination to complete your ride? \nDo you really want to complete your ride now?"
        )
        if (confirmBox === true) {
          handleCompleteRide(e, ride_id)
        }
      }}>
        Complete this ride
      </button>
    );
  }



  return (
    <div className="container">
      <header className="riderPage">
      <h3>Welcome Rider! {currentUser.first_name} {currentUser.last_name}</h3>
      </header>
      <br></br>

      { // If there is no rides happening at the moment
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
                    <th>Passenger Name</th>
                    <th>Passenger Phone No.</th>
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
                            <td>{data.passenger_name}</td>
                            <td>{data.passenger_phone}</td>
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


        ) : ( // if there are current ride and the ride is on-going
          <div className="showRide">
          <h4>Current Ride Information</h4>
          <p>
            <strong>Ride ID:</strong> {currentRide.ride_id}
          </p>
          <p>
            <strong>Passenger Name:</strong> {currentRide.passenger_name}
          </p>
          <p>
            <strong>Passenger Phone No.:</strong> {currentRide.passenger_phone}
          </p>
          <p>
            <strong>Pickup Postal Code: </strong> {currentRide.pickup_code}
          </p>
          <p>
            <strong>Dropoff Postal Code: </strong> {currentRide.dropoff_code}
          </p>
          <p>
            <strong>Ride Status: </strong> {currentRide.ride_status}
          </p>
          <CancelRideButton ride_id={currentRide.ride_id} />
          <br></br>
          <br></br>
          <CompleteRideButton ride_id={currentRide.ride_id} />

        </div>

        ) 
      }
    </div>

  );
  };
  
  export default Rider; 