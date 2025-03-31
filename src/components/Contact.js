import React from 'react';
import './ContentPage.css';
import heroImage from './accommodation.jpg'; // Replace with your hero image file
import hotelImg from './hotels.jpeg'; // Replace with your hotel image file
import packageImg from './packages.jpg'; // Replace with your package image file

const ContentPage = () => {
  return (
    <div className="content-page">
      <section className="hero-section">
        <img src={heroImage} alt="Travel Hero" className="hero-image" />
        <div className="hero-overlay">
          <h1>Welcome to Our Travel Booking Platform</h1>
          <p>
            Experience the best travel solutions with real-time updates, interactive visuals, and unbeatable offers.
          </p>
        </div>
      </section>

      <section className="info-section">
        <h2>About Our Platform</h2>
        <p>
          Our platform offers a variety of travel solutions including accommodation options such as hotels and customized packages for every traveler.
          We are committed to providing you with a seamless experience and exceptional service.
        </p>
      </section>

      <section className="services-section">
        <h2>Our Services</h2>
        <div className="services-grid">
          <div className="service-card">
            <img src={hotelImg} alt="Hotel Accommodation" />
            <h3>Hotel Accommodation</h3>
            <p>
              Find and book the best hotels with real-time availability, exclusive discounts, and personalized options.
            </p>
          </div>
          <div className="service-card">
            <img src={packageImg} alt="Travel Packages" />
            <h3>Travel Packages</h3>
            <p>
              Choose from a wide range of travel packages tailored to your preferences. Enjoy bundled deals and curated itineraries for a hassle-free journey.
            </p>
          </div>
        </div>
      </section>
    </div>
  );
};

export default ContentPage;
