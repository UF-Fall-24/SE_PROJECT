import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './Login.css';

const Login = ({ onLogin }) => {
  useEffect(() => {
    console.log("Debug: onLogin function received in Login.js?", onLogin);
  }, [onLogin]);

  if (!onLogin) {
    console.error("Error: `onLogin` function is not passed to Login.js");
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
    <div className="login-container">
      <div className="login-card">
        <h2>Login</h2>
        {error && <p className="error">{error}</p>}
        <form onSubmit={handleLoginClick}>
          <input 
            type="email" 
            name="email"
            placeholder="Email" 
            value={email} 
            onChange={(e) => setEmail(e.target.value)} 
            required 
          />
          <input 
            type="password" 
            name="password"
            placeholder="Password" 
            value={password} 
            onChange={(e) => setPassword(e.target.value)} 
            required 
          />
          <button type="submit">Login</button>
        </form>
      </div>
    </div>
  );
};

export default Login;
