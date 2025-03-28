// apiService.js

// Retrieves the authentication token from local storage (if used).
export function getAuthToken() {
  return localStorage.getItem('token');
}
// apiService.js
export async function apiCall(url, options = {}) {
  const token = localStorage.getItem('token');
  const headers = {
    'Content-Type': 'application/json',
    ...(token && { 'Authorization': `Bearer ${token}` }),
    ...options.headers,
  };

  const response = await fetch(url, { ...options, headers });
  let data;
  try {
    data = await response.json();
  } catch {
    data = null;
  }

  if (!response.ok) {
    throw new Error(data?.message || 'Invalid response');
  }
  return data;
}
