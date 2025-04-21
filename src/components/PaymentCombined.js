// src/components/PaymentCombined.js

import React, { useState } from 'react';

export default function PaymentCombined() {
  // Hard‑coded package and hotel details
  const pkg = {
    name: 'Mountain Adventure',
    description:
      'Experience a thrilling mountain adventure with trekking, camping, and breathtaking views.',
    location: 'Nepal',
    price: 799.5,
    days: 7,
    nights: 6,
  };

  const hotel = {
    name: 'Skyline Heights',
    address: '222 Highrise Ave',
    city: 'Seattle',
    rating: 4.4,
    roomType: 'Deluxe',
    price: 230,
  };

  const [step, setStep] = useState('review'); // 'review' | 'payment' | 'success'
  const [card, setCard] = useState({ number: '', expiry: '', cvv: '' });

  const handleProceed = () => setStep('payment');

  const handleSubmit = e => {
    e.preventDefault();
    // Here you'd integrate with your payment gateway...
    setStep('success');
  };

  // Review step: show both package & hotel
  if (step === 'review') {
    return (
      <div className="max-w-xl mx-auto p-6 bg-white shadow-lg rounded-lg">
        <h2 className="text-2xl font-bold mb-4">Review Your Booking</h2>

        {/* Package */}
        <div className="mb-6 p-4 border rounded">
          <h3 className="text-xl font-semibold mb-1">{pkg.name}</h3>
          <p className="text-gray-700 mb-2">{pkg.description}</p>
          <div className="flex text-sm text-gray-600">
            <span className="mr-4">📍 {pkg.location}</span>
            <span className="mr-4">${pkg.price.toFixed(2)}</span>
            <span>{pkg.days} Days / {pkg.nights} Nights</span>
          </div>
        </div>

        {/* Hotel */}
        <div className="mb-6 p-4 border rounded">
          <h3 className="text-xl font-semibold mb-1">{hotel.name}</h3>
          <p className="text-gray-700">{hotel.address}, {hotel.city}</p>
          <div className="flex text-sm text-gray-600 mt-2">
            <span className="mr-4">Rating: {hotel.rating}</span>
            <span className="mr-4">Type: {hotel.roomType}</span>
            <span>Price: ${hotel.price.toFixed(2)}</span>
          </div>
        </div>

        <button
          onClick={handleProceed}
          className="w-full py-2 bg-blue-600 text-white font-medium rounded hover:bg-blue-700 transition"
        >
          Complete Booking & Proceed to Payment
        </button>
      </div>
    );
  }

  // Payment form
  if (step === 'payment') {
    return (
      <div className="max-w-xl mx-auto p-6 bg-white shadow-lg rounded-lg">
        <h2 className="text-2xl font-semibold mb-4">Payment Details</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-sm font-medium mb-1">Card Number</label>
            <input
              type="text"
              value={card.number}
              onChange={e => setCard({ ...card, number: e.target.value })}
              required
              className="w-full border rounded px-3 py-2 focus:outline-none focus:ring"
              placeholder="1234 5678 9012 3456"
            />
          </div>
          <div className="flex gap-4 mb-4">
            <div className="flex-1">
              <label className="block text-sm font-medium mb-1">Expiry Date</label>
              <input
                type="text"
                value={card.expiry}
                onChange={e => setCard({ ...card, expiry: e.target.value })}
                required
                className="w-full border rounded px-3 py-2 focus:outline-none focus:ring"
                placeholder="MM/YY"
              />
            </div>
            <div className="flex-1">
              <label className="block text-sm font-medium mb-1">CVV</label>
              <input
                type="password"
                value={card.cvv}
                onChange={e => setCard({ ...card, cvv: e.target.value })}
                required
                className="w-full border rounded px-3 py-2 focus:outline-none focus:ring"
                placeholder="123"
              />
            </div>
          </div>
          <button
            type="submit"
            className="w-full py-2 bg-green-600 text-white font-medium rounded hover:bg-green-700 transition"
          >
            Pay ${(pkg.price + hotel.price).toFixed(2)}
          </button>
        </form>
      </div>
    );
  }

  // Success
  return (
    <div className="max-w-xl mx-auto p-6 bg-white shadow-lg rounded-lg text-center">
      <h2 className="text-2xl font-bold mb-4">Booking Successful!</h2>
      <p className="text-gray-700 mb-2">
        Your payment of ${(pkg.price + hotel.price).toFixed(2)} has been processed.
      </p>
      <p className="font-medium">Thank you for booking with Book Ease.</p>
    </div>
  );
}
