// src/components/Home.js
import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './Home.css'; // Make sure to create this CSS file
import flightImage from './flight.jpg'; // Import the images
import carImage from './car.jpeg';
import busImage from './bus.jpg';

const Home = () => {
    const [currentImageIndex, setCurrentImageIndex] = useState(0);

    const images = [
        flightImage,
        carImage,
        busImage
    ];

    useEffect(() => {
        const interval = setInterval(() => {
            setCurrentImageIndex((prevIndex) => 
                prevIndex === images.length - 1 ? 0 : prevIndex + 1
            );
        }, 3000); // Change image every 3 seconds

        return () => clearInterval(interval);
    }, []);

    return (
        <div className="home-container">
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
            <div className="welcome-text">
                <h1>Welcome to Book Ease</h1>
                <p>Your one-stop solution for all your travel needs!</p>
            </div>
            <div className="auth-links">
                <Link to="/login">Login</Link>
                <Link to="/register">Register</Link>
            </div>
        </div>
    );
};

export default Home;
