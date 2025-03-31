import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Register.css'; // Make sure this file is in the same folder
                         
const Register = ({ onRegister }) => {
  if (!onRegister) {
    console.error("Error: `onRegister` function is not passed to Register.js");
  }
  
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    
    if (!onRegister) {
      setError("Register function is missing.");
      return;
    }
    
    await onRegister(username, email, password);
    navigate('/login');
  };
  
  return (
    <div className="register-container">
      <div className="register-card">
        <h2>Register</h2>
        {error && <p className="error">{error}</p>}
        <form onSubmit={handleSubmit}>
          <input 
            type="text" 
            name="username"
            placeholder="Username" 
            value={username} 
            onChange={(e) => setUsername(e.target.value)} 
            required 
          />
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
          <button type="submit">Register</button>
        </form>
      </div>
    </div>
  );
};

export default Register;
