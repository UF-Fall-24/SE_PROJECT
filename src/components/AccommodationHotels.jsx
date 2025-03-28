// AccommodationHotels.jsx
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { getHotels } from '../services/hotelService';

const AccommodationHotels = () => {
  const navigate = useNavigate();
  const [hotels, setHotels] = useState([]);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchHotels = async () => {
      try {
        const data = await getHotels();
        setHotels(data);
      } catch (err) {
        setError(`Failed to fetch hotels: ${err.message}`);
      }
    };
    fetchHotels();
  }, []);

  const handleSelectHotel = (id) => {
    navigate(`/hotel-details/${id}`);
  };

  return (
    <div>
      <h2>Available Hotels</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {hotels.length > 0 ? (
        <table border="1" cellPadding="12" cellSpacing="0" style={{ width: '100%', borderCollapse: 'collapse' }}>
          <thead style={{ backgroundColor: 'grey' }}>
            <tr>
              <th>Hotel</th>
              <th>City</th>
              <th>Rating</th>
              <th>Room Type</th>
              <th>Price per Night ($)</th>
              <th>Select</th>
            </tr>
          </thead>
          <tbody>
            {hotels.map((hotel) => (
              <tr key={hotel.id}>
                <td>
                  <strong>{hotel.hotel_name}</strong><br/>
                  <small>{hotel.address}</small>
                </td>
                <td style={{ textAlign: 'center' }}>{hotel.city}</td>
                <td style={{ textAlign: 'center' }}>{hotel.rating}</td>
                <td style={{ textAlign: 'center' }}>{hotel.room_type}</td>
                <td style={{ textAlign: 'center' }}>${hotel.room_price}</td>
                <td style={{ textAlign: 'center' }}>
                  <button onClick={() => handleSelectHotel(hotel.id)}>
                    Select
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
};

export default AccommodationHotels;
