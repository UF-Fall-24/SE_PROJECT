// src/components/Payment.js

import React from 'react';

export default function Payment() {
  const pkg = {
    package_name: 'Mountain Adventure',
    package_description:
      'Experience a thrilling mountain adventure with trekking, camping, and breathtaking views.',
    location: 'Nepal',
    package_price: 799.5,
    days: 7,
    nights: 6,
  };

  return (
    <div className="payment-container max-w-xl mx-auto p-6 bg-white shadow rounded">
      <h2 className="text-2xl font-bold mb-2">{pkg.package_name}</h2>
      <p className="mb-4">{pkg.package_description}</p>

      <div className="flex justify-between text-gray-600 text-sm mb-6">
        <span>📍 {pkg.location}</span>
        <span>${pkg.package_price.toFixed(2)}</span>
        <span>{pkg.days} Days / {pkg.nights} Nights</span>
      </div>

      <div className="mb-4">
        <h3 className="text-xl font-semibold">Total</h3>
        <p className="text-3xl font-extrabold">${pkg.package_price.toFixed(2)}</p>
      </div>

      <button
        className="w-full py-2 bg-green-600 text-white rounded hover:bg-green-700"
        onClick={() => alert('Proceeding to payment gateway…')}
      >
        Proceed to Payment
      </button>
    </div>
  );
}
