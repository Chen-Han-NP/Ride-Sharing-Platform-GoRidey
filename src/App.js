//import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import { Routes, Route, Link, Router } from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import Profile from "./components/Profile";
import Passenger from "./components/Passenger";
import Rider from "./components/Rider";

import EventBus from "./common/EventBus";
import AuthService from './services/auth.service';


const App = () => {
  const [currentUser, setCurrentUser] = useState(undefined);
  const [showPassenger, setShowPassenger] = useState(false);
  const [showRider, setShowRider] = useState(false);

  useEffect(() => {
    const user = AuthService.getCurrentUser();
    
    if (user) {
      setCurrentUser(user);

    }

    EventBus.on("logout", () => {
      logOut();
    });
    return () => {
      EventBus.remove("logout");
    };
  }, []);

  const logOut = () => {
    AuthService.logout();
    setCurrentUser(undefined);
    localStorage.removeItem('user');

  };

  return (
    <div>
      {
        !currentUser ? (
          <nav className="navbar navbar-expand navbar-dark bg-dark">
          <Link to={"/"} className="navbar-brand">
            GoRidey
          </Link>
          <div className="navbar-nav ml-auto">
          <li className="nav-item">
            <Link to={"/login"} className="nav-link">
              Login
            </Link>
          </li>
          <li className="nav-item">
            <Link to={"/register"} className="nav-link">
              Register
            </Link>
          </li>
            </div>
            </nav>
        ) : (
          <nav className="navbar navbar-expand navbar-dark bg-dark">
          <Link to={"/"} className="navbar-brand">
            GoRidey
          </Link>
          <div className="navbar-nav ml-auto">
          <li className="nav-item">
            <Link to={"/passenger"} className="nav-link">
              Home
            </Link>
          </li>
          <li className="nav-item">
            <Link to={"/profile"} className="nav-link">
              Profile
            </Link>
          </li>
          <li className="nav-item">
            <Link to={"/"} className="nav-link" onClick={logOut}>
              Logout
            </Link>
          </li>
            </div>
            </nav> 
        )
        
      }
      {
        !currentUser ? (
          <div className="container mt-3">
          <Routes>
            <Route exact path={"/"} element={<Login />} />
            <Route exact path="/login" element={<Login />} />
            <Route exact path="/register" element={<Register />} />
          </Routes>
        </div>
        ) : (
          <div className="container mt-3">
          <Routes>
            <Route exact path={"/"} element={<Passenger />} />
            <Route path="/passenger" element={<Passenger />} />
            <Route exact path="/profile" element={<Profile />} />
          </Routes>
        </div>
        )
      }


          
    </div>
  );
};

export default App;
