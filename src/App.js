import React, { lazy, Suspense, useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom';
import { FaHome, FaInfoCircle, FaPhone } from 'react-icons/fa';

// Lazy load components
const Home = lazy(() => import('./components/Home'));
const Register = lazy(() => import('./components/Register'));
const Login = lazy(() => import('./components/Login'));
const About = lazy(() => import('./components/About'));
const Contact = lazy(() => import('./components/Contact'));
const Search = lazy(() => import('./components/Search'));
const Dashboard = lazy(() => import('./components/Dashboard'));
const ResetPassword = lazy(() => import('./components/ResetPassword'));
const ForgotPassword = lazy(() => import('./components/ForgotPassword'));

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

    const handleRegister = async (username, email, password) => {
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

    const handleLogin = async (email, password) => {
        try {
            const response = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password }),
            });
            if (response.ok) {
                const data = await response.json();
                localStorage.setItem("token", data.token);
                setIsLoggedIn(true);
                setToken(data.token);
            } else {
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
                <Suspense fallback={<div>Loading...</div>}>
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
                </Suspense>
            </div>

            {isLoggedIn && <Search />}
        </Router>
    );
};

export default App;