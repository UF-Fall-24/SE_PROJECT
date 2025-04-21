// src/components/Payment.js
import React, { useState, useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

const Payment = () => {
  const navigate = useNavigate();
  const { search } = useLocation();
  const params = new URLSearchParams(search);
  const packageId = params.get('package_id');

  const [total, setTotal] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!packageId) {
      setError('No package_id provided');
      setLoading(false);
      return;
    }

    // Adjust baseURL as needed, or use an environment var
    fetch(`http://localhost:8080/api/payment?package_id=${packageId}`)
      .then((res) => {
        if (!res.ok) throw new Error(`Server responded ${res.status}`);
        return res.json();
      })
      .then((data) => {
        setTotal(data.total_price);
      })
      .catch((err) => {
        console.error('Error fetching total price:', err);
        setError('Failed to fetch price');
      })
      .finally(() => {
        setLoading(false);
      });
  }, [packageId]);

  const handleBack = () => navigate(-1);

  if (loading) return <div>Loading total price…</div>;
  if (error)   return <div style={{ color: 'red' }}>{error}</div>;

  return (
    <div className="payment-page">
      <h2>Payment</h2>
      <p>
        <strong>Package ID:</strong> {packageId}
      </p>
      <p>
        <strong>Total Price:</strong> ${total.toFixed(2)}
      </p>
      <button onClick={handleBack}>Go Back</button>
      {/* Future: integrate actual payment gateway here */}
    </div>
  );
};

export default Payment;
