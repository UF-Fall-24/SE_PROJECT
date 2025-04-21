// src/components/AccommodationHotels.js

import React, { useEffect, useState } from 'react';
import { getHotels } from '../services/hotelService';
import Payment from './PaymentCombined';

export default function AccommodationHotels() {
  const [hotels, setHotels]                   = useState([]);
  const [error, setError]                     = useState('');
  const [selectedHotel, setSelectedHotel]     = useState(null);
  const [showPaymentForm, setShowPaymentForm] = useState(false);

  // 1) Load hotel list
  useEffect(() => {
    (async () => {
      try {
        const data = await getHotels();
        setHotels(data);
      } catch (err) {
        setError(`Failed to fetch hotels: ${err.message}`);
      }
    })();
  }, []);

  // 2) If user has chosen a hotel and clicked "Complete Booking", show Payment form
  if (showPaymentForm) {
    return <Payment />;
  }

  // 3) If a hotel is selected but user hasn't clicked "Complete Booking" yet, show summary
  if (selectedHotel) {
    return (
      <div className="max-w-xl mx-auto mt-8 p-6 bg-white shadow-lg rounded-lg">
        <h2 className="text-2xl font-bold mb-4">Booking Summary</h2>

        {/* Hard‑coded package info */}
        <div className="mb-6 p-4 border rounded">
          <h3 className="text-xl font-semibold">Mountain Adventure</h3>
          <p className="text-gray-700 mb-2">
            Experience a thrilling mountain adventure with trekking, camping, and breathtaking views.
          </p>
          <div className="flex text-sm text-gray-600">
            <span className="mr-4">📍 Nepal</span>
            <span className="mr-4">$799.50</span>
            <span>7 Days / 6 Nights</span>
          </div>
        </div>

        {/* Selected hotel info */}
        <div className="mb-6 p-4 border rounded">
          <h3 className="text-xl font-semibold">{selectedHotel.hotel_name}</h3>
          <p className="text-gray-700">{selectedHotel.address}</p>
          <div className="flex text-sm text-gray-600 mt-2">
            <span className="mr-4">{selectedHotel.city}</span>
            <span className="mr-4">Rating: {selectedHotel.rating}</span>
            <span className="mr-4">Type: {selectedHotel.room_type}</span>
            <span>Price: ${selectedHotel.room_price}</span>
          </div>
        </div>

        <button
          className="w-full py-2 px-4 bg-green-600 text-white font-medium rounded hover:bg-green-700 transition"
          onClick={() => setShowPaymentForm(true)}
        >
          Complete Booking & Proceed to Payment
        </button>
      </div>
    );
  }

  // 4) Default: show list of hotels
  return (
    <div className="max-w-3xl mx-auto mt-8 p-6 bg-white shadow-md rounded-lg">
      <h2 className="text-2xl font-bold mb-4">Available Hotels</h2>
      {error && <p className="text-red-600">{error}</p>}

      {hotels.length > 0 ? (
        <table className="w-full table-auto border-collapse mb-6">
          <thead className="bg-gray-200">
            <tr>
              <th className="p-3 text-left">Hotel</th>
              <th className="p-3">City</th>
              <th className="p-3">Rating</th>
              <th className="p-3">Room Type</th>
              <th className="p-3">Price ($)</th>
              <th className="p-3">Action</th>
            </tr>
          </thead>
          <tbody>
            {hotels.map(hotel => (
              <tr key={hotel.id} className="border-t">
                <td className="p-3">
                  <strong>{hotel.hotel_name}</strong><br/>
                  <small className="text-gray-600">{hotel.address}</small>
                </td>
                <td className="p-3 text-center">{hotel.city}</td>
                <td className="p-3 text-center">{hotel.rating}</td>
                <td className="p-3 text-center">{hotel.room_type}</td>
                <td className="p-3 text-center">${hotel.room_price}</td>
                <td className="p-3 text-center">
                  <button
                    className="px-3 py-1 bg-blue-600 text-white rounded hover:bg-blue-700"
                    onClick={() => setSelectedHotel(hotel)}
                  >
                    Complete Booking
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        !error && <p>No hotels found.</p>
      )}
    </div>
  );
}
