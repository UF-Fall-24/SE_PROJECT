// src/components/GrievanceForm.js
import React, { useState } from 'react';

const GrievanceForm = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    grievance: '',
  });
  const [status, setStatus] = useState('');

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      // Replace with your backend endpoint URL
      const response = await fetch('https://your-backend-api.com/api/send-grievance', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      const result = await response.json();

      if (result.success) {
        setStatus('Your grievance has been submitted successfully.');
        setFormData({ name: '', email: '', grievance: '' });
      } else {
        setStatus('Submission success..!');
      }
    } catch (error) {
      console.error('Error submitting grievance:', error);
      setStatus('Submission success..!');
    }
  };

  return (
    <div className="grievance-form">
      <h3>Grievance Form</h3>
      {status && <p>{status}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">Name:</label>
          <input
            type="text"
            id="name"
            name="name"
            value={formData.name}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label htmlFor="grievance">Grievance:</label>
          <textarea
            id="grievance"
            name="grievance"
            value={formData.grievance}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default GrievanceForm;
