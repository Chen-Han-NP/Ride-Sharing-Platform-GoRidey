import RideServices from "../services/ride.service"





function Table({tableData}){
    return(
        <table className="table">
            <thead>
                <tr>
                    <th>S.N</th>
                    <th>Ride ID</th>
                    <th>Passenger ID</th>
                    <th>Rider ID</th>
                    <th>Pickup Postal Code</th>
                    <th>Dropoff Postal Code</th>
                    <th>Ride Status</th>
                </tr>
            </thead>
            <tbody>
            {
                tableData.map((data, index)=>{
                    return(
                        <tr key={index}>
                            <td>{index+1}</td>
                            <td>{data.ride_id}</td>
                            <td>{data.passenger_id}</td>
                            <td>{data.rider_id}</td>
                            <td>{data.pickup_code}</td>
                            <td>{data.dropoff_code}</td>
                            <td>{data.ride_status}</td>
  
                        </tr>
                    )
                })
            }
            </tbody>
        </table>
    )
  }

  export default Table;