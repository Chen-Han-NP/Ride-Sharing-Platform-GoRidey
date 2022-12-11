import AuthService from "../services/auth.service";
import RideServices from "../services/ride.service";
import React, { useState, useRef } from "react";
import { useNavigate } from "react-router-dom";
import Form from "react-validation/build/form";
import Input from "react-validation/build/input";
import CheckButton from "react-validation/build/button";

const required = (value) => {
  if (!value) {
    return (
      <div className="invalid-feedback d-block">
        This field is required!
      </div>
    );
  }
};

function containsOnlyNumbers(str) {
  return /^\d+$/.test(str);
}

const validPostalCode = (value) => {
  if (value.length !== 6 || !containsOnlyNumbers(value)) {
    return (
      <div className="invalid-feedback d-block">
        Please enter a valid 6 digit postal code
      </div>
    );
  }
};


const Passenger = () => {
  const form = useRef();
  const checkBtn = useRef();

  const currentUser = AuthService.getCurrentUser();
  const currentRide = RideServices.getCurrentRide();

  const [pickup_code, setPickupCode] = useState("");
  const [dropoff_code, setDropoffCode] = useState("");
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const navigate = useNavigate();

  const onChangePickUpCode = (e) => {
    const pickup_code = e.target.value;
    setPickupCode(pickup_code);
  };

  const onChangeDropoffCode = (e) => {
    const dropoff_code = e.target.value;
    setDropoffCode(dropoff_code);
  };

  
  const handleNewRide = (e) => {
    e.preventDefault();

    setMessage("");
    setLoading(true);

    form.current.validateAll();

    if (checkBtn.current.context._errors.length === 0) {
      var user_name = currentUser.first_name + " " + currentUser.last_name
      var user_phone = currentUser.mobile_number

      RideServices.newride(user_name, user_phone, pickup_code, dropoff_code, "Pending").then(
        () => {
          navigate("/");
          window.location.reload();
        },
        (error) => {
          const resMessage =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();

          setLoading(false);
          setMessage(resMessage);
        }
      );
    } else {
      setLoading(false);
    }
  };

  
  return (
    <div className="container">
      <header className="passengerPage">
      <h3>Welcome Passenger! {currentUser.first_name} {currentUser.last_name}</h3>
      </header>
      <br></br>

      {
        !currentRide ? (
          <div className="card card-container">
            <h3>Book a New Ride now!</h3>
       
            <Form onSubmit={handleNewRide} ref={form}>
              <div className="form-group">
                <label htmlFor="pickup_code">Pickup Postal Code</label>
                <Input
                  type="text"
                  className="form-control"
                  name="pickup_code"
                  value={pickup_code}
                  onChange={onChangePickUpCode}
                  validations={[required, validPostalCode]}
                />
              </div>
    
              <div className="form-group">
                <label htmlFor="dropoff_code">Dropoff Postal Code</label>
                <Input
                  type="text"
                  className="form-control"
                  name="dropoff_code"
                  value={dropoff_code}
                  onChange={onChangeDropoffCode}
                  validations={[required, validPostalCode]}
                />
              </div>
              
              <div className="form-group">
              <button className="btn btn-primary btn-block" disabled={loading}>
                {loading && (
                  <span className="spinner-border spinner-border-sm"></span>
                )}
                <span>Search For Rider</span>
              </button>
            </div>

              {message && (
                <div className="form-group">
                  <div className="alert alert-danger" role="alert">
                    {message}
                  </div>
                </div>
              )}
              <CheckButton style={{ display: "none" }} ref={checkBtn} />
            </Form>

          </div>
        ) : (

          <div className="showRidePending">
            <h4>Current Ride Information</h4>
            <p>
              <strong>Ride ID:</strong> {currentRide.ride_id}
            </p>

            {currentRide.ride_status === "Riding" && (
              <div>
                <p>
                <strong>Rider Name: </strong> {currentRide.rider_name}
                </p>
                <p>
                <strong>Rider Phone No.: </strong> {currentRide.rider_phone}
                </p>
              </div>
            )}

            <p>
              <strong>Pickup Postal Code: </strong> {currentRide.pickup_code}
            </p>
            <p>
              <strong>Dropoff Postal Code: </strong> {currentRide.dropoff_code}
            </p>
            <p>
              <strong>Ride Status: </strong> {currentRide.ride_status}
            </p>
            {currentRide.ride_status === "Pending" && (
            <p>
              Waiting for rider...
            </p>
            )}
            </div>
        )
      }

    </div>
  );
};

export default Passenger; 