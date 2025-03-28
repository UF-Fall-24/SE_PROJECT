// PackageDetails.jsx
import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getPackage } from '../services/packageService';

const PackageDetails = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [pkg, setPkg] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchPackage = async () => {
      try {
        const data = await getPackage(id);
        setPkg(data);
        // store location immediately when package data fetched
        localStorage.setItem('selectedLocation', data.location);
      } catch (err) {
        setError(err.message);
      }
    };
    fetchPackage();
  }, [id]);

  const handleAccommodationClick = () => {
    navigate('/accommodation-hotels');
  };

  const handlePaymentClick = () => {
    navigate('/payment');
  };

  if (error) return <p style={{ color: 'red' }}>Error: {error}</p>;
  if (!pkg) return <p>Loading package details...</p>;

  return (
    <div>
      <h2>Package Details</h2>
      <table border="1" cellPadding="10">
        <tbody>
          <tr><td>ID</td><td>{pkg.id}</td></tr>
          <tr><td>Name</td><td>{pkg.package_name}</td></tr>
          <tr><td>Description</td><td>{pkg.package_description}</td></tr>
          <tr><td>Price</td><td>{pkg.package_price}</td></tr>
          <tr><td>Days</td><td>{pkg.days}</td></tr>
          <tr><td>Nights</td><td>{pkg.nights}</td></tr>
          <tr><td>Location</td><td>{pkg.location}</td></tr>
        </tbody>
      </table>

      <div style={{ marginTop: '20px' }}>
        <button onClick={handlePaymentClick}>Payment</button>
        <button onClick={handleAccommodationClick} style={{ marginLeft: '10px' }}>
          Book Accommodation
        </button>
      </div>
    </div>
  );
};

export default PackageDetails;
