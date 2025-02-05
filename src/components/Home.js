// src/components/Home.js
import React, { useState } from 'react';
import Login from './Login';
import Register from './Register';

const Home = ({ onRegister, onLogin }) => {
    const [isRegistering, setIsRegistering] = useState(false);

    const handleSwitchToRegister = () => setIsRegistering(true);
    const handleSwitchToLogin = () => setIsRegistering(false);

    return (
        <div>
            {isRegistering ? (
                <Register onRegister={onRegister} onSwitchToLogin={handleSwitchToLogin} />
            ) : (
                <Login onLogin={onLogin} onSwitchToRegister={handleSwitchToRegister} />
            )}
        </div>
    );
};

export default Home;