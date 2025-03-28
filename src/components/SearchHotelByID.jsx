import React, { useState } from 'react';
import { getHotel } from '../services/hotelService';

const SearchHotelByID = () => {
  const [hotelID, setHotelID] = useState('');
  const [hotel, setHotel] = useState(null);
  const [error, setError] = useState('');

  const handleSearch = async () => {
    // Clear previous results/errors.
    setHotel(null);
    setError('');
    if (!hotelID) {
      setError('Please enter a valid hotel ID.');
      return;
    }
    try {
      const data = await getHotel(hotelID);
      setHotel(data);
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div>
      <h2>Search Hotel By ID</h2>
      <div>
        <input 
          type="number" 
          placeholder="Enter hotel ID" 
          value={hotelID}
          onChange={(e) => setHotelID(e.target.value)}
        />
        <button onClick={handleSearch}>Search</button>
      </div>
      {error && <p style={{ color: 'red' }}>Error: {error}</p>}
      {hotel && (
        <div style={{ marginTop: '1em', border: '1px solid #ccc', padding: '1em' }}>
          <h3>{hotel.hotel_name}</h3>
          <p><strong>Address:</strong> {hotel.address}</p>
          <p><strong>City:</strong> {hotel.city}</p>
          <p><strong>Description:</strong> {hotel.description}</p>
          <p><strong>Rating:</strong> {hotel.rating}</p>
          <p><strong>Room Type:</strong> {hotel.room_type}</p>
          <p><strong>Room Price:</strong> ${hotel.room_price}</p>
        </div>
      )}
    </div>
  );
};

export default SearchHotelByID;
