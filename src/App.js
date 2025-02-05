// src/App.js
import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, Link, Navigate } from 'react-router-dom';
import { FaHome, FaInfoCircle, FaPhone } from 'react-icons/fa';
import Home from './components/Home';
import About from './components/About';
import Contact from './components/Contact';
import Search from './components/Search';
import Dashboard from './components/Dashboard';
import ResetPassword from './components/ResetPassword';
import ForgotPassword from './components/ForgotPassword';

const App = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [users, setUsers] = useState([]);
    const [bookedFlights, setBookedFlights] = useState([]);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    useEffect(() => {
        const storedUsers = JSON.parse(localStorage.getItem('users')) || [];
        setUsers(storedUsers);
    }, []);

    const handleRegister = (username, email) => {
        const newUser = { username, email, password }; // Ensure that password is included
        const updatedUsers = [...users, newUser];
        setUsers(updatedUsers);
        localStorage.setItem('users', JSON.stringify(updatedUsers));
        setIsLoggedIn(true);
    };

    const handleLogin = () => {
        const user = users.find(user => user.username === username && user.password === password);
        if (user) {
            setIsLoggedIn(true);
        } else {
            alert('Invalid credentials');
        }
    };

    const handleLogout = () => {
        setIsLoggedIn(false);
        setBookedFlights([]);
    };

    return (
        <Router>
            <header className="header">
                <h1>Book Ease</h1>
                <nav className="nav-container">
                    <Link to="/" className="nav-item"><FaHome className="icon" /> Home</Link>
                    <Link to="/about" className="nav-item"><FaInfoCircle className="icon" /> About Us</Link>
                    <Link to="/contact" className="nav-item"><FaPhone className="icon" /> Contact Us</Link>
                </nav>
            </header>

            <div className="container">
                <Routes>
                    <Route path="/" element={isLoggedIn ? <Navigate to="/dashboard" /> : <Home onRegister={handleRegister} onLogin={handleLogin} />} />
                    <Route path="/about" element={<About />} />
                    <Route path="/contact" element={<Contact />} />
                    <Route path="/dashboard" element={isLoggedIn ? <Dashboard bookedFlights={bookedFlights} /> : <Navigate to="/" />} />
                    <Route path="/reset-password" element={<ResetPassword />} />
                    <Route path="/forgot-password" element={<ForgotPassword />} />
                </Routes>
            </div>

            {isLoggedIn && <Search />}
        </Router>
    );
};

export default App;