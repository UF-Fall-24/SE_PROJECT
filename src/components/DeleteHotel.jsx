import React, { useState } from 'react';
import { deleteHotel } from '../services/hotelService';

const DeleteHotel = () => {
  const [hotelId, setHotelId] = useState('');
  const [message, setMessage] = useState('');

  // Update the hotelId as the user types
  const handleChange = (e) => {
    setHotelId(e.target.value);
  };

  // Submit the deletion request
  const handleDelete = async (e) => {
    e.preventDefault();
    setMessage('');
    try {
      const response = await deleteHotel(hotelId);
      // Check the response message and update accordingly
      if (response.message && response.message.toLowerCase().includes("no hotel")) {
        setMessage("No hotel with given ID");
      } else {
        setMessage("Hotel deleted successfully");
      }
    } catch (error) {
      setMessage(`Error deleting hotel: ${error.message}`);
    }
  };

  return (
    <div>
      <h2>Delete Hotel</h2>
      {message && <p>{message}</p>}
      <form onSubmit={handleDelete}>
        <input
          type="number"
          placeholder="Enter Hotel ID"
          value={hotelId}
          onChange={handleChange}
          required
        />
        <button type="submit">Delete Hotel</button>
      </form>
    </div>
  );
};

export default DeleteHotel;
