# GoRidey - A Ride Sharing Platform

GoRidey is an awesome platform for both Passengers and Riders to initiate, start and end a ride. The idea is very similar to Grab [https://www.grab.com/sg/](https://www.grab.com/sg/) or Gojek [https://www.gojek.com/sg/](https://www.gojek.com/sg/) in Singapore. 

<img width="1266" alt="Screenshot 2022-12-14 at 3 39 49 PM" src="https://user-images.githubusercontent.com/73086331/207540959-93b48d96-7e01-420c-bb8c-4519d56b3ad5.png">

## Design Architecture
![ETI Asg 1 Diagram_2022-12-14_03-34-29](https://user-images.githubusercontent.com/73086331/207539716-2f2f3214-51be-4867-aef9-d2232c981ade.png)
Created using [Mural](https://www.mural.co/)

## Microservice Design Considerations
### 1. Authentication Management
- User Registration
- User Login with credentials, assign JWT Token embeded in Cookie
- User Logout
### 2. User Management
- User Get Profile Data
- User Update Profile Data
- User Update Password
<img width="1176" alt="Screenshot 2022-12-14 at 3 41 17 PM" src="https://user-images.githubusercontent.com/73086331/207542314-dac98ae7-d226-4caf-8c5a-8724f0f50f5c.png">

### 3. Ride Management
- Ride Initialization by Passenger
- Get All Pending Ride Requests for Rider
- Ride Acception by Rider
- Ride Cancellation by Passenger for Pending Rides and Rider for On-going Rides
- Ride Completion by Rider upon reaching the destination
- Get All Ride Histories and Display in reverse choronological order


## What are the Development Tools & Methods Used?
GoRidey is developed using the following tools:
- Developed using **Microservice architecture**, specifically **Authentication**, **User** and **Ride Managements**
- **Back-end**: [Golang server](https://go.dev/) using [Mux Router](https://github.com/gorilla/mux)
- **Front-end**: [React.js](https://reactjs.org/) & 
             [Bootstrap](https://getbootstrap.com/)
- **CORS Middleware**: [Moesif Origin & CORS Changer](https://chrome.google.com/webstore/detail/moesif-origin-cors-change/digfbfaphojjndkpccljibejjbppifbc)
- **Database**: [MySQL](https://www.mysql.com/)


## How to Use?
**Pre-requisites**: 
- Installed [Node.js](https://nodejs.org/en/) package and environment, able to run `npm` commands.
- Installed [Golang](https://go.dev/) package and environment, able to run `go` commands.
- Installed [MySQL Workbench](https://www.mysql.com/)

### Let's Dive into Actions!
1. Clone [https://github.com/Chen-Han-NP/ride-sharing-platform.git](https://github.com/Chen-Han-NP/ride-sharing-platform.git) into your new working repository.
2. In your Terminal/Command Prompt, run: `npm run start` to start the development mode on the React project and the website is hosted on **[http://localhost:3000](http://localhost:3000)**.
3. Open your MySQL Workbench, make sure your user is **`root`**, password is **`password`** and database server port number is **`3306`**, as the connection string in Golang server are set to **`"root:password@tcp(127.0.0.1:3306)/"`** in default.
4. *MySQL script* is located at: **`RIDE-SHARING-PLATFORM/MySQL/RideSharingDBSetup.sql`**, please find and run the script to set up a **database** with two default users - *Chen Han and Daryl*, and *one default completed Ride*.
5. Open a new Terminal/Command Prompt, run: **`go run Microservices/Authentication/auth.go`**, 
server should be running at [http://localhost:5050](http://localhost:5050).
6. Open a new Terminal/Command Prompt, run:`go run Microservices/User/user.go`,
server should be running at [http://localhost:5051](http://localhost:5051).
7. Open a new Terminal/Command Prompt, run:`go run Microservices/Ride/ride.go`,
server should be running at [http://localhost:5052](http://localhost:5052).
8. **All the servers should be up and running now! Feel free to explore the features :)**


## Deployment
- Branch `gh-pages` is set for future deployment at [Github Pages](https://pages.github.com/), 
using command `npm run deploy`, page is currently active at [https://chen-han-np.github.io/ride-sharing-platform/](https://chen-han-np.github.io/ride-sharing-platform/), but not up-to-date with developing codes.


## Credits
- This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).
- Some part of the codes are referred from: [https://www.bezkoder.com/react-login-example-jwt-hooks/](https://www.bezkoder.com/react-login-example-jwt-hooks/) for Register and Login functions.
- Special credit to: 
[https://www.youtube.com/watch?v=b9eMGE7QtTk&t=767s&ab_channel=JavaScriptMastery](https://www.youtube.com/watch?v=b9eMGE7QtTk&t=767s&ab_channel=JavaScriptMastery) & 
[https://www.youtube.com/watch?v=hQAHSlTtcmY&ab_channel=WebDevSimplified](https://www.youtube.com/watch?v=hQAHSlTtcmY&ab_channel=WebDevSimplified) for teaching the basics of React.js
- And many more awesome Stack Overflowers for solving existing issues regarding Cors, Cookies and JWT tokens! :)



