// src/services/packageService.js
import { apiCall } from './apiService';

// Fetch all packages
export async function getPackages() {
  return apiCall('http://localhost:8000/packages', { method: 'GET' });
}

// Fetch a single package by ID
export async function getPackage(id) {
  return apiCall(`http://localhost:8000/packages/${id}`, { method: 'GET' });
}
