// HotelDetails.jsx
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getHotel } from '../services/hotelService';

const HotelDetails = () => {
  const { id } = useParams();
  const [hotel, setHotel] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchHotel = async () => {
      try {
        const data = await getHotel(id);
        setHotel(data);
      } catch (err) {
        setError(`Failed to fetch hotel details: ${err.message}`);
      }
    };
    fetchHotel();
  }, [id]);

  if (error) return <p style={{ color: 'red' }}>{error}</p>;
  if (!hotel) return <p>Loading hotel details...</p>;

  return (
    <div>
      <h2>Hotel Details: {hotel.hotel_name}</h2>
      <table border="1" cellPadding="10">
        <tbody>
          <tr><td>ID</td><td>{hotel.id}</td></tr>
          <tr><td>Name</td><td>{hotel.hotel_name}</td></tr>
          <tr><td>Address</td><td>{hotel.address}</td></tr>
          <tr><td>City</td><td>{hotel.city}</td></tr>
          <tr><td>Description</td><td>{hotel.description}</td></tr>
          <tr><td>Rating</td><td>{hotel.rating}</td></tr>
          <tr><td>Room Type</td><td>{hotel.room_type}</td></tr>
          <tr><td>Price per Night</td><td>${hotel.room_price}</td></tr>
        </tbody>
      </table>
    </div>
  );
};

export default HotelDetails;
