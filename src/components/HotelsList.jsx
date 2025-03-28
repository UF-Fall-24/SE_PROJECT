import React, { useEffect, useState } from 'react';
import { getHotels } from '../services/hotelService';
import { Link } from 'react-router-dom';

function HotelsList() {
  const [hotels, setHotels] = useState([]);

  useEffect(() => {
    async function fetchData() {
      try {
        const data = await getHotels();
        console.log('Hotels data:', data); // see the actual JSON
        setHotels(data);
      } catch (error) {
        console.error('Error fetching hotels:', error);
      }
    }
    fetchData();
  }, []);

  return (
    <div>
      <h2>Available Hotels</h2>
      <ul>
        {hotels.map((hotel) => (
          <li key={hotel.id}>
            <h3>{hotel.hotel_name}</h3>
            <p>Address: {hotel.address}</p>
            <p>City: {hotel.city}</p>
            <p>Description: {hotel.description}</p>
            <p>Rating: {hotel.rating}</p>
            <p>Room Type: {hotel.room_type}</p>
            <p>Room Price: ${hotel.room_price}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default HotelsList;
