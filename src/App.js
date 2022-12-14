//import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import { Routes, Route, Link } from "react-router-dom";
import Login from "./components/Login";
import RegisterPassenger from "./components/RegisterPassenger";
import RegisterRider from "./components/RegisterRider";
import Profile from "./components/Profile";
import Passenger from "./components/Passenger";
import RideHistory from './components/RideHistory';
import Rider from "./components/Rider";
import '../node_modules/bootstrap/dist/css/bootstrap.min.css'


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

      if (user.user_type === "passenger") {
        setShowPassenger(true)
        setShowRider(false)
      } else if (user.user_type === "rider") {
        setShowRider(true)
        setShowPassenger(false)  
      }
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
    setShowPassenger(false)
    setShowRider(false)
    localStorage.removeItem('user');

  };

  return (
    <div>
      {
        !currentUser ? (
          <nav className="navbar navbar-expand-lg navbar-light">
            <div className='container'>
              <Link to={"/"} className="navbar-brand">
                GoRidey
              </Link>
              <div className='collapse navbar-collapse'>
                <ul className="navbar-nav ml-auto" >
                  <li className='nav-item'>
                    <Link to={"/login"} className="nav-link">
                      Login
                    </Link>
                  </li>
                  <li className='nav-item'>
                    <Link to={"/registerpassenger"} className="nav-link">
                      Register Passenger
                    </Link>
                  </li>
                  <li className='nav-item'>
                    <Link to={"/registerrider"} className="nav-link">
                      Register Rider
                    </Link>
                  </li>        
                </ul>
              </div>
            </div>
          </nav>

        ) : (
            showPassenger === true && showRider === false ? (
              <nav className="navbar navbar-expand-lg navbar-light">
              <div className='container'>
                <Link to={"/"} className="navbar-brand">
                  GoRidey
                </Link>
                <div className='collapse navbar-collapse'>
                  <ul className="navbar-nav ml-auto" >
                    <li className='nav-item'>
                      <Link to={"/passenger"} className="nav-link">
                        Home
                      </Link>
                    </li>
                    <li className='nav-item'>
                      <Link to={"/profile"} className="nav-link">
                        Profile
                      </Link>
                    </li>
                    <li className='nav-item'>
                      <Link to={"/rides"} className="nav-link">
                        Ride History
                      </Link>
                    </li>
                    <li className="nav-item">
                      <Link to={"/"} className="nav-link" onClick={logOut}>
                        Logout
                      </Link>
                    </li>        
                  </ul>
                </div>
              </div>
            </nav> 
            ) : (
              showPassenger === false && showRider === true && (              
              <nav className="navbar navbar-expand-lg navbar-light">
                <div className='container'>
                  <Link to={"/"} className="navbar-brand">
                    GoRidey
                  </Link>
                  <div className='collapse navbar-collapse'>
                    <ul className="navbar-nav ml-auto" >
                      <li className='nav-item'>
                        <Link to={"/rider"} className="nav-link">
                          Home
                        </Link>
                      </li>
                      <li className='nav-item'>
                        <Link to={"/profile"} className="nav-link">
                          Profile
                        </Link>
                      </li>
                      <li className='nav-item'>
                        <Link to={"/rides"} className="nav-link">
                          Ride History
                        </Link>
                      </li>
                      <li className="nav-item">
                        <Link to={"/"} className="nav-link" onClick={logOut}>
                          Logout
                        </Link>
                      </li>        
                    </ul>
                  </div>
                </div>
              </nav> 
              )
            )
        ) 
      }

      
      { !currentUser ? (
          <div className="auth-wrapper">
            <div className='auth-inner'>
              <Routes>
                <Route exact path={"/"} element={<Login />} />
                <Route exact path="/login" element={<Login />} />
                <Route exact path="/registerpassenger" element={<RegisterPassenger />} />
                <Route exact path="/registerrider" element={<RegisterRider/>} />
              </Routes>
            </div>
        </div>
        ) : ( 
          showPassenger === true && showRider === false ? (
            <div className="auth-wrapper">
              <div>
              <Routes>
                <Route exact path={"/"} element={<Passenger />} />
                <Route path="/passenger" element={<Passenger />} />
                <Route exact path="/profile" element={<Profile />} />
                <Route exact path="/rides" element={<RideHistory />} />
              </Routes>
              </div>
           </div>
          ) : (
            showPassenger === false && showRider === true && (
              <div className="auth-wrapper">
                <div>
                <Routes>
                  <Route exact path={"/"} element={<Rider />} />
                  <Route path="/rider" element={<Rider  />} />
                  <Route exact path="/profile" element={<Profile />} />
                  <Route exact path="/rides" element={<RideHistory />} />
                </Routes>
                </div>
              </div>
            )      
          )
        )
      }

    </div>
  );
};

export default App;
