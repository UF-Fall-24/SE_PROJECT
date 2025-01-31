// src/components/Dashboard.js
import React from 'react';

const Dashboard = ({ bookedFlights }) => {
    return (
        <div className="dashboard">
            <h2>Your Dashboard</h2>
            <p>Welcome back! Here you can manage your bookings and search for travel options.</p>
            <h3>My Booked Flights:</h3>
            {bookedFlights.length > 0 ? (
                <ul>
                    {bookedFlights.map((flight, index) => (
                        <li key={index}>{flight}</li>
                    ))}
                </ul>
            ) : (
                <p>No flights booked yet.</p>
            )}
            <div className="search-options">
                <h3>Search for Flights, Cars, and Buses</h3>
                <button onClick={() => alert('Searching for flights...')}>Search Flights</button>
                <button onClick={() => alert('Searching for cars...')}>Search Cars</button>
                <button onClick={() => alert('Searching for buses...')}>Search Buses</button>
            </div>
        </div>
    );
};

export default Dashboard;