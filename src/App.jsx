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
import Dashboard from './pages/dashboard';

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/checkout" element={<Checkout />} />
        <Route path="/" element={<Store />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </Router>
  )
}

export default App
