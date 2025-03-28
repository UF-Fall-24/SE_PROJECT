// accommodationService.js
import { apiCall } from './apiService';

// Create accommodation based on selected hotel
export async function createAccommodation(accommodationData) {
  return apiCall('http://localhost:8000/accommodations', { 
    method: 'POST',
    body: JSON.stringify(accommodationData),
  });
}
