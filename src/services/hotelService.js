// hotelService.js
import { apiCall } from './apiService';

// 1. Fetch all hotels
export async function getHotels() {
  return apiCall('http://localhost:8000/hotels', { method: 'GET' });
}

// 2. Fetch a single hotel by its ID
export async function getHotel(id) {
  return apiCall(`http://localhost:8000/hotels/${id}`, { method: 'GET' });
}

// 3. Create a new hotel
export async function createHotel(hotelData) {
  return apiCall('http://localhost:8000/hotels', {
    method: 'POST',
    body: JSON.stringify(hotelData),
  });
}

// 4. Update an existing hotel
export async function updateHotel(id, hotelData) {
  return apiCall(`http://localhost:8000/hotels/${id}`, {
    method: 'PUT',
    body: JSON.stringify(hotelData),
  });
}

// 5. Delete a hotel
export async function deleteHotel(id) {
  return apiCall(`http://localhost:8000/hotels/${id}`, { method: 'DELETE' });
}

// 6. Fetch hotels filtered by location (city)
export async function getHotelsByLocation(location) {
  return apiCall(`http://localhost:8000/hotels?location=${encodeURIComponent(location)}`, {
    method: 'GET'
  });
}

