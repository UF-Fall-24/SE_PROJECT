// Payment.jsx
import React from 'react';
import { useNavigate } from 'react-router-dom';

const Payment = () => {
  const navigate = useNavigate();

  return (
    <div>
      <h2>Payment Page</h2>
      <p>Implement payment integration here.</p>
      <button onClick={() => navigate(-1)}>Back</button>
      <button onClick={() => alert('Payment success!')}>Confirm Payment</button>
    </div>
  );
};

export default Payment;
