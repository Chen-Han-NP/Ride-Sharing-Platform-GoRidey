import React, {useState, useEffect } from 'react'


const API_LOGIN = "http://localhost:5050/login"


class User {
    constructor (email_address, password){
      this.email_address = email_address;
      this.password = password;
    }
  }
  
const LoginPage = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    
    useEffect(() => {
        
    });

    async function login(){
        var user = new User(email, password);
        var userStringify = JSON.stringify(user)
        
        let result = await fetch(API_LOGIN, {
            method: "POST",
            headers: {
                "Content-type": "application/json",
                "Accept": "application/json"
            },
            body: userStringify 
        });
        result = await result.json();
        localStorage.setItem()

            
    };

    return (
        <>
        <h1>Welcome to Ridey</h1>
        <p>Your awesome journey starts here</p>
        <br></br>
        <h3>Login</h3>
        <form>
        <div className="form-group">
            <label htmlFor="exampleInputEmail1">Email address</label>
            <input type="email" className="form-control" onChange={(e) =>
                 setEmail(e.target.value)} id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email"/>
            <small id="emailHelp" className="form-text text-muted">We'll never share your email with anyone else.</small>
        </div>
        <div className="form-group">
            <label htmlFor="exampleInputPassword1">Password</label>
            <input type="password" className="form-control" onChange={(e) => 
                setPassword(e.target.value)} id="exampleInputPassword1" placeholder="Password"/>
        </div>
        <div className="form-check">
            <input type="checkbox" className="form-check-input" id="exampleCheck1"/>
            <label className="form-check-label" htmlFor="exampleCheck1">Check me out</label>
        </div>
        <button type="submit" onClick={login} className="btn btn-primary">Login</button>
        </form>
        </>
    );
}
export default LoginPage;
  