// src/App.js
import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, Link, Navigate } from 'react-router-dom';
import Home from './components/Home';
import About from './components/About';
import Contact from './components/Contact';
import Search from './components/Search';
import Dashboard from './components/Dashboard'; // Import Dashboard component
import ResetPassword from './components/ResetPassword'; // Import Reset Password component
import ForgotPassword from './components/ForgotPassword'; // Import Forgot Password component
import { FaBars } from 'react-icons/fa'; // Import hamburger icon

const App = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [users, setUsers] = useState([]);
    const [showDashboard, setShowDashboard] = useState(false);
    const [bookedFlights, setBookedFlights] = useState([]); // Store booked flights

    useEffect(() => {
        const storedUsers = JSON.parse(localStorage.getItem('users')) || [];
        setUsers(storedUsers);
    }, []);

    const handleRegister = (username, email) => {
        const newUser = { username, email };
        const updatedUsers = [...users, newUser];
        setUsers(updatedUsers);
        localStorage.setItem('users', JSON.stringify(updatedUsers));
        setIsLoggedIn(true);
    };

    const handleLogin = (username, password) => {
        const user = users.find(user => user.username === username && user.password === password);
        if (user) {
            setIsLoggedIn(true);
            // Simulate fetching booked flights for the logged-in user
            setBookedFlights(["Flight A", "Flight B"]); // Example booked flights
        } else {
            alert('Invalid credentials');
        }
    };

    const handleLogout = () => {
        setIsLoggedIn(false);
        setBookedFlights([]); // Clear booked flights on logout
    };

    const toggleDashboard = () => {
        setShowDashboard(!showDashboard);
    };

    return (
        <Router>
            <nav>
                <ul>
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/about">About Us</Link></li>
                    <li><Link to="/contact">Contact Information</Link></li>
                    {isLoggedIn && (
                        <li><button onClick={handleLogout} style={{ background: 'transparent', border: 'none' }}>Logout</button></li>
                    )}
                </ul>
                {isLoggedIn && (
                    <button onClick={toggleDashboard} style={{ float: 'right', fontSize: '24px', background: 'transparent', border: 'none' }}>
                        <FaBars /> {/* Hamburger icon */}
                    </button>
                )}
            </nav>
            <div className="container">
                <Routes>
                    <Route path="/" element={isLoggedIn ? <Navigate to="/dashboard" /> : <Home onRegister={handleRegister} onLogin={handleLogin} />} />
                    <Route path="/about" element={<About />} />
                    <Route path="/contact" element={<Contact />} />
                    <Route path="/dashboard" element={isLoggedIn ? <Dashboard bookedFlights={bookedFlights} /> : <Navigate to="/" />} />
                    <Route path="/reset-password" element={<ResetPassword />} /> {/* Route for Reset Password */}
                    <Route path="/forgot-password" element={<ForgotPassword />} /> {/* Route for Forgot Password */}
                </Routes>
                {showDashboard && isLoggedIn && (
                    <div className="dashboard-overlay">
                        <Dashboard bookedFlights={bookedFlights} />
                        <button onClick={toggleDashboard} style={{ position: 'absolute', top: '10px', right: '10px' }}>Close</button>
                    </div>
                )}
            </div>
            {isLoggedIn && <Search />} {/* Show Search component only if logged in */}
        </Router>
    );
};

export default App;