//import logo from './logo.svg';
import './App.css';
import { useState } from 'react';
import LoginPage from './components/Login.js';

class User {
  constructor (emailAddress, password){
    this.emailAddress = emailAddress;
    this.password = password;
  }
}

class Passenger {
  constructor(passengerId, firstName, lastName, emailAddress, password, mobileNumber) {
    this.passengerId = passengerId;
    this.firstName = firstName;
    this.lastName = lastName;
    this.emailAddress = emailAddress;
    this.password = password;
    this.mobileNumber = mobileNumber
  }
}

const PassengerPage = (props) => {
  var passenger = props.passenger;

  return  (
    <>
    <h1>Welcome to Passenger Home page</h1>
    <h2>--Detail--</h2>
    <h3>First Name: {passenger.firstName}</h3>
    <h3>Last Name: {passenger.lastName}</h3>
    <h3>Email: {passenger.emailAddress}</h3>
    <h3>Mobile Number: {passenger.mobileNumber}</h3>

    </>
  )
}
const App = () => {
  const isUserLogin = false;


  var passenger = new Passenger("U123", "Chen", "Han", "chenhan@gmail.com", "12345678", "14114121");

  return (
    <div className='App'>
      {isUserLogin === true ? (
        <PassengerPage passenger={passenger}>
        </PassengerPage>
      ) : (
        <>
        <LoginPage></LoginPage>
        </>
      )}
    </div>
  );
}



const StateExample = () => {
  // react state can only can changed using the function attached
  // so cant use counter = 100 as it is not mutable
  const [counter, setCounter] = useState(100);

  return (
    <div className='App'>
      <button onClick={() => setCounter(prevCount => prevCount - 1)}>-</button>
      <h1>{counter}</h1>
      <button>+</button>
    </div>
  );
}

export default App;
