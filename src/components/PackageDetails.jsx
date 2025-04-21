// src/components/PackageDetails.js
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { getPackage } from '../services/packageService';
import Payment from './payment';

export default function PackageDetails() {
  const { id } = useParams();
  const [pkg, setPkg]           = useState(null);
  const [loading, setLoading]   = useState(true);
  const [error, setError]       = useState('');
  const [showPayment, setShowPayment] = useState(false);

  useEffect(() => {
    (async () => {
      try {
        const data = await getPackage(id);
        setPkg(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    })();
  }, [id]);

  if (loading) return <p>Loading package details…</p>;
  if (error)   return <p className="text-red-600">Error: {error}</p>;

  // Once user clicks "Payment", swap into the Payment card
  if (showPayment) {
    return <Payment />;
  }

  return (
    <div className="max-w-xl mx-auto p-6 bg-white shadow-md rounded-lg">
      <h2 className="text-2xl font-bold mb-4">{pkg.package_name}</h2>
      <p className="text-gray-700 mb-6">{pkg.package_description}</p>
      <table className="w-full text-sm text-gray-600 mb-6">
        <tbody>
          <tr>
            <td className="py-2 font-medium">Location</td>
            <td className="py-2">{pkg.location}</td>
          </tr>
          <tr>
            <td className="py-2 font-medium">Duration</td>
            <td className="py-2">{pkg.days} Days / {pkg.nights} Nights</td>
          </tr>
          <tr>
            <td className="py-2 font-medium">Base Price</td>
            <td className="py-2">${pkg.package_price.toFixed(2)}</td>
          </tr>
        </tbody>
      </table>
      <button
        className="w-full py-2 bg-blue-600 text-white font-semibold rounded hover:bg-blue-700 transition"
        onClick={() => setShowPayment(true)}
      >
        Payment
      </button>
    </div>
  );
}
