// src/components/Home.js
import React, { useState } from 'react';
import Login from './Login';
import Register from './Register';

const Home = ({ onLogin, onRegister }) => {
    const [isRegistering, setIsRegistering] = useState(false);

    const handleSwitchToRegister = () => setIsRegistering(true);
    const handleSwitchToLogin = () => setIsRegistering(false);

    return (
        <div>
            <h2>Welcome to Travel Booking</h2>
            <p>Your adventure starts here!</p>
            {isRegistering ? (
                <Register onRegister={onRegister} onSwitchToLogin={handleSwitchToLogin} />
            ) : (
                <Login onLogin={onLogin} onSwitchToRegister={handleSwitchToRegister} />
            )}
        </div>
    );
};

export default Home;