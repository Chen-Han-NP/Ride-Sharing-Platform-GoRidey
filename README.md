# GoRidey - A Ride Sharing Platform

GoRidey is an awesome platform for both Passengers and Riders to initiate, start and end a ride. The idea is very similar to Grab [https://www.grab.com/sg/](https://www.grab.com/sg/) or Gojek [https://www.gojek.com/sg/](https://www.gojek.com/sg/) in Singapore. 

## Functionalities
GoRidey consists of following functionalities:
- Registration (Passenger or Rider)
- Login Authentication with JWT Token
- Ride Initialization by Passenger
- Ride Cancellation
- Ride Acceptation/Completion by Rider
- Profile Management
  

## Development Tools
GoRidey is developed using the following tools:
- Developed using Microservice Architecturel, managing Authentication, User and Ride.
- Back-end: [Golang server](https://go.dev/) using [Mux Router](https://github.com/gorilla/mux)
- Front-end: [React.js](https://reactjs.org/) & 
             [Bootstrap](https://getbootstrap.com/)
- Middleware: [Moesif Origin & CORS Changer](https://chrome.google.com/webstore/detail/moesif-origin-cors-change/digfbfaphojjndkpccljibejjbppifbc)
- Database: [MySQL Workbench](https://www.mysql.com/)


## How-To-Use
Pre-requisites: 
- Installed Node.js package and environment, able to run `npm` commands.
- Installed Golang package and environment, able to run `go` commands.
- Installed MySQL Workbench

### Let's dive into actions!
1. Clone [https://github.com/Chen-Han-NP/ride-sharing-platform.git](https://github.com/Chen-Han-NP/ride-sharing-platform.git) into your new working repository.
2. In your Terminal/Command Prompt, run: `npm run start` to start the development mode on the React project and the website is hosted on [http://localhost:3000](http://localhost:3000).
3. Open your MySQL Workbench, make sure the Local User is set to: `root`, Password is set to: `password` and Port number is set to: `3306`, as the connection string in Golang server are set to `"root:password@tcp(127.0.0.1:3306)/"` in default.
4. The MySQL script is located at: `RIDE-SHARING-PLATFORM/MySQL/RideSharingDBSetup.sql`, feel free to run the script to set up a datbase with two default users, Chen Han and Daryl, with one default completed Ride.
5. Open a new Terminal/Command Prompt, run:`go run Microservices/Authentication/auth.go`, 
server should be running at [http://localhost:5050](http://localhost:5050).
6. Open a new Terminal/Command Prompt, run:`go run Microservices/User/user.go`,
server should be running at [http://localhost:5051](http://localhost:5051).
7. Open a new Terminal/Command Prompt, run:`go run Microservices/Ride/ride.go`,
server should be running at [http://localhost:5052](http://localhost:5052).
8. The servers should be all up and running!




## Credits
- This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).


## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

The page will reload when you make changes.\
You may also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can't go back!**

If you aren't satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you're on your own.

You don't have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn't feel obligated to use this feature. However we understand that this tool wouldn't be useful if you couldn't customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

### Code Splitting

This section has moved here: [https://facebook.github.io/create-react-app/docs/code-splitting](https://facebook.github.io/create-react-app/docs/code-splitting)

### Analyzing the Bundle Size

This section has moved here: [https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size](https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size)

### Making a Progressive Web App

This section has moved here: [https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app](https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app)

### Advanced Configuration

This section has moved here: [https://facebook.github.io/create-react-app/docs/advanced-configuration](https://facebook.github.io/create-react-app/docs/advanced-configuration)

### Deployment

This section has moved here: [https://facebook.github.io/create-react-app/docs/deployment](https://facebook.github.io/create-react-app/docs/deployment)

### `npm run build` fails to minify

This section has moved here: [https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify](https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify)
