// src/components/Payment.js

import React, { useState } from 'react';

export default function Payment() {
  // Hard‑coded package details
  const pkg = {
    package_name: 'Mountain Adventure',
    package_description:
      'Experience a thrilling mountain adventure with trekking, camping, and breathtaking views.',
    location: 'Nepal',
    package_price: 799.5,
    days: 7,
    nights: 6,
  };

  const [step, setStep] = useState('details'); // 'details' | 'form' | 'success'
  const [cardNumber, setCardNumber] = useState('');
  const [expiry, setExpiry] = useState('');
  const [cvv, setCvv] = useState('');

  const handleProceed = () => {
    setStep('form');
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Here you’d normally validate & send to your payment gateway
    setStep('success');
  };

  if (step === 'details') {
    return (
      <div className="max-w-md mx-auto mt-8 p-6 bg-white shadow-lg rounded-lg">
        <h1 className="text-3xl font-bold mb-2">{pkg.package_name}</h1>
        <p className="text-gray-700 mb-4">{pkg.package_description}</p>
        <div className="flex justify-between text-sm text-gray-500 mb-6">
          <span>📍 {pkg.location}</span>
          <span>${pkg.package_price.toFixed(2)}</span>
          <span>{pkg.days} Days / {pkg.nights} Nights</span>
        </div>
        <button
          className="w-full py-2 px-4 bg-blue-600 text-white font-medium rounded hover:bg-blue-700 transition"
          onClick={handleProceed}
        >
          Proceed to Payment
        </button>
      </div>
    );
  }

  if (step === 'form') {
    return (
      <div className="max-w-md mx-auto mt-8 p-6 bg-white shadow-lg rounded-lg">
        <h2 className="text-2xl font-semibold mb-4">Enter Payment Details</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-sm font-medium mb-1">Card Number</label>
            <input
              type="text"
              value={cardNumber}
              onChange={(e) => setCardNumber(e.target.value)}
              placeholder="1234 5678 9012 3456"
              className="w-full border rounded px-3 py-2 focus:outline-none focus:ring"
              required
            />
          </div>
          <div className="mb-4 grid grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium mb-1">Expiry Date</label>
              <input
                type="text"
                value={expiry}
                onChange={(e) => setExpiry(e.target.value)}
                placeholder="MM/YY"
                className="w-full border rounded px-3 py-2 focus:outline-none focus:ring"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">CVV</label>
              <input
                type="password"
                value={cvv}
                onChange={(e) => setCvv(e.target.value)}
                placeholder="123"
                className="w-full border rounded px-3 py-2 focus:outline-none focus:ring"
                required
              />
            </div>
          </div>
          <button
            type="submit"
            className="w-full py-2 bg-green-600 text-white font-medium rounded hover:bg-green-700 transition"
          >
            Pay ${pkg.package_price.toFixed(2)}
          </button>
        </form>
      </div>
    );
  }

  // step === 'success'
  return (
    <div className="max-w-md mx-auto mt-8 p-6 bg-white shadow-lg rounded-lg text-center">
      <h2 className="text-2xl font-bold mb-4">Payment Successful!</h2>
      <p className="text-gray-700 mb-6">
        Thank you for booking the <strong>{pkg.package_name}</strong>.
      </p>
      <p className="text-green-600 font-semibold">Your transaction has been completed.</p>
    </div>
  );
}
