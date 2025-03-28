import React, { useState } from 'react';
import { createHotel } from '../services/hotelService';

const CreateHotelForm = () => {
  const [hotelData, setHotelData] = useState({
    hotel_name: '',
    address: '',
    city: '',
    description: '',
    rating: 0,
    room_type: '',
    room_price: 0,
  });
  
  const [message, setMessage] = useState('');

  const handleChange = (e) => {
    const { name, value, type } = e.target;
    // Convert numeric inputs to numbers
    setHotelData({ 
      ...hotelData, 
      [name]: type === 'number' ? parseFloat(value) : value 
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const newHotel = await createHotel(hotelData);
      setMessage(`Hotel "${newHotel.hotel_name}" created successfully!`);
      // Optionally clear the form:
      setHotelData({
        hotel_name: '',
        address: '',
        city: '',
        description: '',
        rating: 0,
        room_type: '',
        room_price: 0,
      });
    } catch (err) {
      setMessage(`Error: ${err.message}`);
    }
  };

  return (
    <div>
      <h2>Create New Hotel</h2>
      {message && <p>{message}</p>}
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="hotel_name"
          placeholder="Hotel Name"
          value={hotelData.hotel_name}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="address"
          placeholder="Address"
          value={hotelData.address}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="city"
          placeholder="City"
          value={hotelData.city}
          onChange={handleChange}
          required
        />
        <textarea
          name="description"
          placeholder="Description"
          value={hotelData.description}
          onChange={handleChange}
          required
        />
        <input
          type="number"
          name="rating"
          placeholder="Rating"
          value={hotelData.rating}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="room_type"
          placeholder="Room Type"
          value={hotelData.room_type}
          onChange={handleChange}
          required
        />
        <input
          type="number"
          name="room_price"
          placeholder="Room Price"
          value={hotelData.room_price}
          onChange={handleChange}
          required
        />
        <button type="submit">Create Hotel</button>
      </form>
    </div>
  );
};

export default CreateHotelForm;
