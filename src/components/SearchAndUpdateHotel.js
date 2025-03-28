import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { getHotel, updateHotel } from '../services/hotelService';

const SearchAndUpdateHotel = () => {
  const navigate = useNavigate();
  // State to hold the entered hotel ID (as string) and the fetched hotel data.
  const [searchId, setSearchId] = useState('');
  const [hotelData, setHotelData] = useState(null);
  const [message, setMessage] = useState('');

  // Handle changes in the hotel ID search field.
  const handleSearchIdChange = (e) => {
    setSearchId(e.target.value);
  };

  // Fetch hotel details based on the entered ID.
  const handleSearch = async () => {
    setMessage('');
    try {
      const data = await getHotel(searchId);
      setHotelData(data);
    } catch (error) {
      setMessage(`Error fetching hotel: ${error.message}`);
      setHotelData(null);
    }
  };

  // Handle input changes for the update form; converts number fields.
  const handleChange = (e) => {
    const { name, value, type } = e.target;
    let newValue = value;
    if (type === 'number') {
      newValue = value === '' ? 0 : parseFloat(value);
      if (isNaN(newValue)) newValue = 0;
    }
    setHotelData((prevData) => ({
      ...prevData,
      [name]: newValue,
    }));
  };

  // Submit the updated hotel data to the backend.
  const handleSubmit = async (e) => {
    e.preventDefault();
    setMessage('');
    try {
      const updatedHotel = await updateHotel(searchId, hotelData);
      setMessage('Hotel updated successfully!');
      setHotelData(updatedHotel);
      // Optionally navigate to the hotel details page:
      // navigate(`/hotels/${searchId}`);
    } catch (error) {
      setMessage(`Error updating hotel: ${error.message}`);
    }
  };

  return (
    <div>
      <h2>Search and Update Hotel</h2>
      
      {/* Search Section */}
      {!hotelData && (
        <div>
          <input
            type="number"
            placeholder="Enter Hotel ID"
            value={searchId}
            onChange={handleSearchIdChange}
          />
          <button onClick={handleSearch}>Search</button>
        </div>
      )}

      {/* Display message if present */}
      {message && <p>{message}</p>}

      {/* Update Form Section */}
      {hotelData && (
        <form onSubmit={handleSubmit}>
          <div>
            <label>Hotel Name:</label>
            <input
              type="text"
              name="hotel_name"
              placeholder="Hotel Name"
              value={hotelData.hotel_name}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Address:</label>
            <input
              type="text"
              name="address"
              placeholder="Address"
              value={hotelData.address}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>City:</label>
            <input
              type="text"
              name="city"
              placeholder="City"
              value={hotelData.city}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Description:</label>
            <textarea
              name="description"
              placeholder="Description"
              value={hotelData.description}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Rating:</label>
            <input
              type="number"
              name="rating"
              placeholder="Rating"
              value={hotelData.rating}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Room Type:</label>
            <input
              type="text"
              name="room_type"
              placeholder="Room Type"
              value={hotelData.room_type}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label>Room Price:</label>
            <input
              type="number"
              name="room_price"
              placeholder="Room Price"
              value={hotelData.room_price}
              onChange={handleChange}
              required
            />
          </div>
          <button type="submit">Update Hotel</button>
        </form>
      )}
    </div>
  );
};

export default SearchAndUpdateHotel;
