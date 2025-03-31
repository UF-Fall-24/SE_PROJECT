// src/components/Home.js
import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './Home.css';
import flightImage from './flight.jpg';
import carImage from './car.jpeg';
import busImage from './bus.jpg';
import HotelDropdown from './HotelDropdown';

const Home = () => {
  // Check if a token exists (indicating the user is logged in)
  const token = localStorage.getItem('token');

  // Only use the carousel if the user is not logged in.
  const [currentImageIndex, setCurrentImageIndex] = useState(0);
  const images = [flightImage, carImage, busImage];

  useEffect(() => {
    // Only set up the carousel interval if no token (user not logged in)
    if (!token) {
      const interval = setInterval(() => {
        setCurrentImageIndex((prevIndex) =>
          prevIndex === images.length - 1 ? 0 : prevIndex + 1
        );
      }, 3000); // Change image every 3 seconds

      return () => clearInterval(interval);
    }
  }, [token, images.length]);

  return (
    <div className="home-container">
      {/* Show carousel and auth links only when user is not logged in */}
      {!token && (
        <>
          <div className="image-carousel">
            {images.map((img, index) => (
              <img
                key={index}
                src={img}
                alt={`Travel ${index + 1}`}
                className={index === currentImageIndex ? 'active' : ''}
              />
            ))}
          </div>
          <div className="auth-links">
            <Link to="/login">Login</Link>
            <Link to="/register">Register</Link>
          </div>
        </>
      )}
      {/* Always show welcome text */}
      <div className="welcome-text">
        <h1>Welcome to Book Ease</h1>
        <p>Your one-stop solution for all your travel needs!</p>
      </div>
      {/* Show HotelDropdown only when user is logged in */}
      {token && (
        <div className="hotel-dropdown">
          <HotelDropdown />
        </div>
      )}
    </div>
  );
};

export default Home;
  