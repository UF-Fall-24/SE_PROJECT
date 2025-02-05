// src/components/ResetPassword.js
import React, { useState } from 'react';

const ResetPassword = () => {
    const [email, setEmail] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [message, setMessage] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        // Reset password logic here
        setMessage('Password reset successful!');
    };

    return (
        <div>
            <h3>Reset Password</h3>
            {message && <p>{message}</p>}
            <form onSubmit={handleSubmit}>
                <input type="email" placeholder="Enter your email" value={email} onChange={(e) => setEmail(e.target.value)} required />
                <input type="password" placeholder="New Password" value={newPassword} onChange={(e) => setNewPassword(e.target.value)} required />
                <button type="submit">Reset Password</button>
            </form>
        </div>
    );
};

export default ResetPassword;