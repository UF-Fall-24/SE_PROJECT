// src/components/PackagesList.jsx
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getPackages } from '../services/packageService';

const PackagesList = () => {
  const [packages, setPackages] = useState([]);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchPackages = async () => {
      try {
        const data = await getPackages();
        setPackages(data);
      } catch (err) {
        setError(`Error fetching packages: ${err.message}`);
      }
    };
    fetchPackages();
  }, []);

  const handleBookNow = (id) => {
    navigate(`/package-details/${id}`);
  };

  return (
    <div>
      <h2>Available Packages</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {packages.length > 0 ? (
        <table border="1" cellPadding="12" cellSpacing="0" style={{ width: '100%', borderCollapse: 'collapse' }}>
          <thead style={{ backgroundColor: 'grey' }}>
            <tr>
              <th>Package</th>
              <th>Price ($)</th>
              <th>Duration</th>
              <th>Action</th>
            </tr>
          </thead>
          <tbody>
            {packages.map((pkg) => (
              <tr key={pkg.id}>
                <td>
                  <strong>{pkg.package_name}</strong>
                  <p>{pkg.package_description}</p>
                  <small><em>📍 {pkg.location}</em></small>
                </td>
                <td style={{ textAlign: 'center' }}>{pkg.package_price}</td>
                <td style={{ textAlign: 'center' }}>
                  {pkg.days} Days / {pkg.nights} Nights
                </td>
                <td style={{ textAlign: 'center' }}>
                  <button onClick={() => handleBookNow(pkg.id)}>Book Now</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        !error && <p>No packages available.</p>
      )}
    </div>
  );
};

export default PackagesList;
