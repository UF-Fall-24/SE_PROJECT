import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = ({ onLogin }) => {
    useEffect(() => {
        console.log("ðŸ› ï¸ Debug: onLogin function received in Login.js?", onLogin);
    }, []);

    if (!onLogin) {
        console.error("âŒ Error: `onLogin` function is not passed to Login.js");
    }

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleLoginClick = async (e) => {
        e.preventDefault();
        setError('');

        if (!onLogin) {
            setError("Login function is missing.");
            return;
        }

        await onLogin(email, password);
        navigate('/dashboard');
    };

    return (
        <div>
            <h2>Login</h2>
            {error && <p style={{ color: "red" }}>{error}</p>}
            <form onSubmit={handleLoginClick}>
                <input 
                    type="email" 
                    placeholder="Email" 
                    value={email} 
                    onChange={(e) => setEmail(e.target.value)} 
                    required 
                />
                <input 
                    type="password" 
                    placeholder="Password" 
                    value={password} 
                    onChange={(e) => setPassword(e.target.value)} 
                    required 
                />
                <button type="submit">Login</button>
            </form>
        </div>
    );
};

export default Login;