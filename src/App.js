import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom';
import Home from './components/Home';
import Register from './components/Register';
import Login from './components/Login';
import About from './components/About';
import Contact from './components/Contact';
import Search from './components/Search';
import Dashboard from './components/Dashboard';
import ResetPassword from './components/ResetPassword';
import ForgotPassword from './components/ForgotPassword';
import { FaHome, FaInfoCircle, FaPhone } from 'react-icons/fa';

const App = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [token, setToken] = useState(null);
    const [bookedFlights, setBookedFlights] = useState([]);

    useEffect(() => {
        const savedToken = localStorage.getItem('token');
        if (savedToken) {
            setToken(savedToken);
            setIsLoggedIn(true);
        }
    }, []);

    // âœ… Fix: Define handleRegister function and pass it to Register.js
    const handleRegister = async (username, email, password) => {
        console.log("ðŸ› ï¸ Debug: Register function is called!");
        try {
            const response = await fetch("http://localhost:8080/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ username, email, password }),
            });

            if (response.ok) {
                alert("Registration successful! Please log in.");
            } else {
                alert("Registration failed. Try again.");
            }
        } catch (error) {
            console.error("Error registering:", error);
            alert("Server error. Please try again.");
        }
    };

    // âœ… Fix: Ensure handleLogin function exists
    const handleLogin = async (email, password) => {
        console.log("🛠️ Debug: Login function is called!");
        console.log("🛠️ Debug: Email:", email);
        console.log("🛠️ Debug: Password:", password);
        try {
            const response = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password }),
            });
            if (response.ok) {
                const data = await response.json();
                console.log("✅ Login successful:", data); // Log the entire response
                localStorage.setItem("token", data.token);
                setIsLoggedIn(true);
                setToken(data.token);
            } else {
                console.log("❌ Login failed:", response.status); // Log the status code
                alert("Invalid credentials. Please try again.");
            }
        } catch (error) {
            console.error("Error logging in:", error);
            alert("Server error. Please try again.");
        }
    };
    const handleLogout = () => {
        setIsLoggedIn(false);
        setToken(null);
        localStorage.removeItem('token');
        setBookedFlights([]);
    };

    
    return (
        <Router>
            <header className="header">
            <h1>Book Ease</h1>
            <nav className='nav-container'>
                <Link to="/" className="nav-item"><FaHome className="icon" /> Home</Link>
                    <Link to="/about" className="nav-item"><FaInfoCircle className="icon" /> About Us</Link>
                    <Link to="/contact" className="nav-item"><FaPhone className="icon" /> Contact Us</Link>
                    {isLoggedIn ? (
                        <button onClick={handleLogout} className="nav-item">Logout</button>
                    ) : (
                        <>
                            <Link to="/login" className="nav-item">Login</Link>
                            <Link to="/register" className="nav-item">Register</Link>
                        </>
                    )}
            </nav>
            </header>

        <div className="container">
        <Routes>
        <Route path="/" element={<Home />} />
                    <Route path="/about" element={<About />} />
                    <Route path="/contact" element={<Contact />} />
                    <Route path="/register" element={<Register onRegister={handleRegister} />} />
                    <Route path="/login" element={<Login onLogin={handleLogin} />} />
                    <Route path="/dashboard" element={isLoggedIn ? <Dashboard bookedFlights={bookedFlights} /> : <Navigate to="/login" />} />
                    <Route path="/reset-password" element={<ResetPassword />} />
                    <Route path="/forgot-password" element={<ForgotPassword />} />
            </Routes>
        </div>
        {isLoggedIn && <Search />}
            
        </Router>
    );
};

export default App;