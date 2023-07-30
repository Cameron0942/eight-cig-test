//? REACT
import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

//? CSS
import './App.css'

//? AXIOS
import axios from 'axios';

//? COMPONENTS
// import About from './About';
import Store from './pages/store';
import Checkout from './pages/checkout';

function App() {

  useEffect(() => {
    // Function to make the Axios GET request
    const fetchData = async () => {
      try {
        const response = await axios.get('http://localhost:8080/'); // Replace 'api/data' with your actual endpoint
        // setData(response.data);
        console.log(response.data)
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData(); // Call the function to trigger the GET request when the component is loaded
  }, []);

  return (
    <Router>
      <Routes>
        <Route path="/checkout" element={<Checkout />} />
        <Route path="/" element={<Store />} />
        {/* <Route path="/contact" element={Contact} /> */}
      </Routes>
    </Router>
  )
}

export default App
