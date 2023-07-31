//? REACT
import React, { useEffect } from 'react';

//? AXIOS
import axios from 'axios';

const Dashboard = () => {

    useEffect(() => {
        // Function to make the Axios GET request
        const fetchData = async () => {
          try {
            const response = await axios.get('http://localhost:8080/api/dashboard');
            // setData(response.data);
            console.log(response.data)
          } catch (error) {
            console.error('Error fetching data:', error);
          }
        };
    
        fetchData(); // Call the function to trigger the GET request when the component is loaded
      }, []);

  return (
    <div>dashboard</div>
  )
}

export default Dashboard