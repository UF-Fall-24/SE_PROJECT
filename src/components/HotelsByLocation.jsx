import React, { useState } from 'react';
import { getHotelsByLocation } from '../services/hotelService';

const HotelsByLocation = () => {
  const [location, setLocation] = useState('');
  const [hotels, setHotels] = useState([]);
  const [message, setMessage] = useState('');

  // Update the location as the user types
  const handleLocationChange = (e) => {
    setLocation(e.target.value);
  };

  // Fetch hotels and filter for an exact match (case-insensitive)
  const handleSearch = async (e) => {
    e.preventDefault();
    setMessage('');
    setHotels([]);
    try {
      const data = await getHotelsByLocation(location);
      // Filter hotels: trim spaces and compare case-insensitively
      const filteredHotels = data.filter(hotel => 
        hotel.city.trim().toLowerCase() === location.trim().toLowerCase()
      );
      
      if (filteredHotels.length === 0) {
        setMessage('No hotels found for this location.');
      } else {
        setHotels(filteredHotels);
      }
    } catch (error) {
      setMessage(`Error retrieving hotels: ${error.message}`);
    }
  };

  return (
    <div>
      <h2>Search Hotels By Location</h2>
      <form onSubmit={handleSearch}>
        <input
          type="text"
          placeholder="Enter city"
          value={location}
          onChange={handleLocationChange}
          required
        />
        <button type="submit">Search</button>
      </form>
      {message && <p>{message}</p>}
      {hotels.length > 0 && (
        <div>
          <h3>Hotels in {location.trim()}</h3>
          <ul>
            {hotels.map((hotel) => (
              <li key={hotel.id}>
                <h4>{hotel.hotel_name}</h4>
                <p>{hotel.address}, {hotel.city}</p>
                <p>{hotel.description}</p>
                <p>Rating: {hotel.rating}</p>
                <p>Room Type: {hotel.room_type}</p>
                <p>Room Price: ${hotel.room_price}</p>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default HotelsByLocation;
