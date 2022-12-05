//import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import { Routes, Route, Link, Router } from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import Home from "./components/Home";
import Profile from "./components/Profile";
import Passenger from "./components/Passenger";
import Rider from "./components/Rider";

import EventBus from "./common/EventBus";
import AuthService from './services/auth.service';

const App = () => {
  const [currentUser, setCurrentUser] = useState(undefined);
  const [currentUserType, setCurrentUserType] = useState(null);

  useEffect(() => {
    const user = AuthService.getCurrentUser();


    if (user) {
      setCurrentUser(user);
      setCurrentUserType(user.user_type);
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
    setCurrentUser(undefined)
  };

  return (
    <div>
      <nav className="navbar navbar-expand navbar-dark bg-dark">
        <Link to={"/"} className="navbar-brand">
          GoRidey
        </Link>
        <div className="navbar-nav mr-auto">
          <li className="nav-item">
            <Link to={"/home"} className="nav-link">
              Home
            </Link>
          </li>

          {currentUserType === "passenger" && currentUserType !== null ? (
            <li className="nav-item">
              <Link to={"/passenger"} className="nav-link">
                Passenger
              </Link>
            </li>
          ) : (
            <li className="nav-item">
            <Link to={"/rider"} className="nav-link">
              Rider
            </Link>
          </li>
          )
        }
       
        </div>

        {currentUser ? (
          <div className="navbar-nav ml-auto">
            <li className="nav-item">
              <Link to={"/profile"} className="nav-link">
                {currentUser.username}
              </Link>
            </li>
            <li className="nav-item">
              <a href="/logout" className="nav-link" onClick={logOut}>
                LogOut
              </a>
            </li>
          </div>
        ) : (
          <div className="navbar-nav ml-auto">
            <li className="nav-item">
              <Link to={"/login"} className="nav-link">
                Login
              </Link>
            </li>

            <li className="nav-item">
              <Link to={"/register"} className="nav-link">
                Sign Up
              </Link>
            </li>
          </div>
        )}
      </nav>

      <div className="container mt-3">
        <Routes>
          <Route exact path={"/"} element={<Home />} />
          <Route exact path={"/home"} element={<Home />} />
          <Route exact path="/login" element={<Login />} />
          <Route exact path="/register" element={<Register />} />
          <Route exact path="/profile" element={<Profile />} />
          <Route exact path="/passenger" element={<Passenger />} />
          <Route exact path="/rider" element={<Rider />} />

        </Routes>
      </div>
    </div>
  );
};

export default App;
