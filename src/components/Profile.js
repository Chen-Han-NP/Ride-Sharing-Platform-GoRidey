import React, { useState, useRef } from "react";
import Form from "react-validation/build/form";
import Input from "react-validation/build/input";
import AuthService from "../services/auth.service";
import UserService from "../services/user.service";
import { useNavigate } from "react-router-dom";
import CheckButton from "react-validation/build/button";

function containsOnlyNumbers(str) {
  return /^\d+$/.test(str);
}

const validateNewPassword = (value) => {
  if (value.length < 8 || value.length > 40) {
    return (
      <div className="invalid-feedback d-block">
        The password must be between 8 and 40 characters.
      </div>
    );
  }
};




const validMobileNumber = (value) => {
  if (value.length !== 8 || !containsOnlyNumbers(value)) {
    return (
      <div className="invalid-feedback d-block">
        Please enter a valid 8 digit phone number
      </div>
    );
  }
};


const Profile = () => {
  const form = useRef();
  const PWform = useRef();

  const checkBtn = useRef();
  const PWform_checkBtn = useRef();

  const navigate = useNavigate();

  const currentUser = AuthService.getCurrentUser();
  const userType = currentUser.user_type  

  const [first_name, setFirstName] = useState(currentUser.first_name);
  const [last_name, setLastName] = useState(currentUser.last_name);
  const [mobile_number, setMobileNumber] = useState(currentUser.mobile_number);
  const [car_lic_number, setCarLicNumber] = useState(currentUser.car_lic_number);
  
  const [newPassword, setNewPassword] = useState("")
  const [newPassword2, setNewPassword2] = useState("")

  const [showForm, setShowForm] = useState(false)
  const [showPWform, setShowPWform] = useState(false)

  const [successful, setSuccessful] = useState(false);
  const [message, setMessage] = useState("");

  const onChangeFirstName = (e) => {
    const first_name = e.target.value;
    setFirstName(first_name);
  };

  const onChangeLastName = (e) => {
    const last_name = e.target.value;
    setLastName(last_name);
  };

  const onChangeMobileNumber = (e) => {
    const mobile_number = e.target.value;
    setMobileNumber(mobile_number);
  };

  const onChangeCarLicNumber = (e) => {
    const car_lic_number = e.target.value;
    setCarLicNumber(car_lic_number);
  };

  const onChangeShowForm = (e) => {
    if (!showForm) {
      setShowForm(true)
      setShowPWform(false)
    } else {
      setShowForm(false)
      setShowPWform(false)
    }
    setFirstName(currentUser.first_name)
    setLastName(currentUser.last_name)
    setMobileNumber(currentUser.mobile_number)
    setCarLicNumber(currentUser.car_lic_number)
      
  }

  const onChangeNewPassword = (e) => {
    const newpass = e.target.value
    setNewPassword(newpass)
  };

  const onChangeNewPassword2 = (e) => {
    const newpass2 = e.target.value
    setNewPassword2(newpass2)
  };

  const onChangeShowPWform = (e) => {
    if (!showPWform) {
      setShowPWform(true)
      setShowForm(false)
    } else {
      setShowPWform(false)
      setShowForm(false)
    }
    setNewPassword("")
    setNewPassword2("")
  }



  const handleEditForm = (e) => {
    e.preventDefault();

    setMessage("");
    setSuccessful(false);

    form.current.validateAll();

    // If passed validation, call auth service to send the API request
    if (checkBtn.current.context._errors.length === 0) {
      if (currentUser.user_type === "passenger"){
        UserService.updatePassengerProfile(first_name, last_name, mobile_number).then(
          () => {
            setMessage("Updated successful!");
            setSuccessful(true);
            navigate("/profile");
            window.location.reload();
          },
          (error) => {
            var resMessage = ""
            if (error.response.status === 400) {
              resMessage = "You do not need to update the same information."
            } else {
              resMessage =
              (error.response &&
                error.response.data &&
                error.response.data.message) ||
              error.message ||
              error.toString();
            }
            setMessage(resMessage);
            setSuccessful(false);
          }
        );

      } else {
        UserService.updateRiderProfile(first_name, last_name, mobile_number, car_lic_number).then(
          () => {
            setMessage("Updated successful!");
            setSuccessful(true);
            setTimeout(function () {
                  navigate("/profile");
                  window.location.reload();
          }, 2000);

          },
          (error) => {
            var resMessage = ""
            if (error.response.status === 400) {
              resMessage = "You do not need to update the same information."
            } else {
              resMessage =
              (error.response &&
                error.response.data &&
                error.response.data.message) ||
              error.message ||
              error.toString();
            }
            setMessage(resMessage);
            setSuccessful(false);
          }
        );

      }
    }
  };


  const handleEditPwForm = (e) => {
    e.preventDefault();

    setMessage("");
    setSuccessful(false);
    PWform.current.validateAll();
    // If passed validation, call user service to send the API request
    if (PWform_checkBtn.current.context._errors.length === 0) {
      if (newPassword !== newPassword2) {
        alert("Passwords not match!")

      } else {

        UserService.updatePassword(newPassword).then(
          () => {
            setMessage("Updated successful!");
            setSuccessful(true);
            setTimeout(function () {
                  navigate("/profile");
                  window.location.reload();
          }, 2000);

          },
          (error) => {
            var resMessage = ""
            if (error.response.status === 400) {
              resMessage = error.response.data
            } else {
              resMessage =
              (error.response &&
                error.response.data &&
                error.response.data.message) ||
              error.message ||
              error.toString();
            }
            setMessage(resMessage);
            setSuccessful(false);
          }
        );

      }

    }
  };


  return (
    <div className="container">

      { !showForm && !showPWform ? (
        <div>
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
      
      <button className="btn btn-secondary" onClick={(e) => onChangeShowForm(e)}>Edit Profile</button>
      <br></br>
      <br></br>
      <button className="btn btn-danger" onClick={(e) => onChangeShowPWform(e)}>Edit Password</button>
      
      </div>
      ) : ( 
        showForm ? (
                    
      <div>
        <header className="jumbotron">
          <h3>
          <button className="btn btn-primary" onClick={(e) => onChangeShowForm(e)}>Back</button>   Update Profile 
          </h3>
        </header>
        <br></br>
      <Form onSubmit={handleEditForm} ref={form}>
        {!successful && (
          
          <div>
            <div className="form-group">
              <label htmlFor="first_name">First name</label>
              <Input
                type="text"
                className="form-control"
                name="first_name"
                value={first_name}
                onChange={onChangeFirstName}
                placeholder = {first_name}
              />
            </div>

            <div className="form-group">
              <label htmlFor="last_name">Last name</label>
              <Input
                type="text"
                className="form-control"
                name="last_name"
                value={last_name}
                onChange={onChangeLastName}
                placeholder={mobile_number}
              />
            </div>

            <div className="form-group">
              <label htmlFor="mobile_number">Mobile number</label>
              <Input
                type="text"
                className="form-control"
                name="mobile_number"
                value={mobile_number}
                onChange={onChangeMobileNumber}
                validations={[validMobileNumber]}
                placeholder={mobile_number}
              />
            </div>
          {
            currentUser.user_type === "rider" && (
              <div className="form-group">
              <label htmlFor="car_lic_number">Car License Number</label>
              <Input
                type="text"
                className="form-control"
                name="car_lic_number"
                value={car_lic_number}
                onChange={onChangeCarLicNumber}
                placeholder={car_lic_number}
              />
            </div>
            )
          }

            <br></br>
            <div className="form-group">
              <button className="btn btn-success btn-block">Update</button>
            </div>
          </div>
        )}

        {message && (
          <div className="form-group">
            <div
              className={
                successful ? "alert alert-success" : "alert alert-danger"
              }
              role="alert"
            >
              {message}
            </div>
          </div>
        )}
        <CheckButton style={{ display: "none" }} ref={checkBtn} />
      </Form>
    </div>
        ) : (

                
    <div>
    <header className="jumbotron">
      <h3>
      <button className="btn btn-primary" onClick={(e) => onChangeShowPWform(e)}>Back</button>   Update Password
      </h3>
    </header>
    <br></br>
  <Form onSubmit={handleEditPwForm} ref={PWform}>
    {!successful && (
      
      <div>
        <div className="form-group">
          <label htmlFor="new_password">New Password</label>
          <Input
            type="password"
            className="form-control"
            name="new_password"
            value={newPassword}
            onChange={onChangeNewPassword}
            validations={[validateNewPassword]}
          />
        </div>

        <div className="form-group">
          <label htmlFor="new_password2">Repeat New Password</label>
          <Input
            type="password"
            className="form-control"
            name="new_password2"
            value={newPassword2}
            onChange={onChangeNewPassword2}
            validations={[]}
          />
        </div>

        <br></br>
        <div className="form-group">
          <button className="btn btn-success btn-block">Update</button>
        </div>
      </div>
    )}

    {message && (
      <div className="form-group">
        <div
          className={
            successful ? "alert alert-success" : "alert alert-danger"
          }
          role="alert"
        >
          {message}
        </div>
      </div>
    )}
    <CheckButton style={{ display: "none" }} ref={PWform_checkBtn} />
  </Form>
</div>

        )
        


    )}
    </div>
  );
};

export default Profile;