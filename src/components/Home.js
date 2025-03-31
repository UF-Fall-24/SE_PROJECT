import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './Home.css';

const Home = () => {
  // Use localStorage to check if a user is logged in
  const token = localStorage.getItem('token');

  // Live clock state
  const [currentTime, setCurrentTime] = useState(new Date());

  useEffect(() => {
    // Update the clock every second
    const timer = setInterval(() => setCurrentTime(new Date()), 1000);
    return () => clearInterval(timer);
  }, []);

  return (
    <div className="home-container">
      {/* Full-Screen Hero Section (unchanged) */}
      <div className="hero-section">
        <div className="hero-overlay">
          <h1>Welcome to Book Ease</h1>
          <p>Your one-stop solution for all your travel needs!</p>
          <p className="live-clock">Current Time: {currentTime.toLocaleTimeString()}</p>
        </div>
      </div>

      {/* Lower Content for Non-Logged-in Users */}
      {!token && (
        <div className="auth-landing-section">
          <div className="auth-box new-user">
            <h2>New to Book Ease?</h2>
            <p>
              Create an account to unlock exclusive deals, personalized recommendations, and easy itinerary management.
            </p>
            <Link to="/register" className="btn">Create Account</Link>
          </div>
          <div className="auth-box returning-user">
            <h2>Already a Member?</h2>
            <p>
              Sign in to manage your bookings, track travel history, and explore more travel options.
            </p>
            <Link to="/login" className="btn">Sign In</Link>
          </div>
        </div>
      )}

      {/* (Optional) Content for Logged-in Users */}
      {token && (
        <div className="logged-in-content">
          {/* For example, a hotel dropdown or user dashboard can be rendered here */}
          <p>Welcome back! Access your dashboard below.</p>
          {/* <HotelDropdown /> */}
        </div>
      )}
    </div>
  );
};

export default Home;
