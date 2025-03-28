import React from 'react';
import { useNavigate } from 'react-router-dom';
import packagesImage from './packages.jpg'; // Make sure this image exists in your project
import accommodationImage from './accommodation.jpg';
import hotelsImage from './hotels.jpeg';
import './Dashboard.css'; // Create and customize this CSS file

const Dashboard = () => {
  const navigate = useNavigate();

  // Function to navigate to the corresponding route when an image is clicked
  const handleClick = (route) => {
    navigate(route);
  };

  return (
    <div className="dashboard-container">
      <h2>User Dashboard</h2>
      <div className="dashboard-grid">
        <div className="dashboard-item" onClick={() => handleClick('/packages')}>
          <img src={packagesImage} alt="Packages" />
          <h3>Packages</h3>
        </div>
        <div className="dashboard-item" onClick={() => handleClick('/accommodations')}>
          <img src={accommodationImage} alt="Accommodations" />
          <h3>Accommodations</h3>
        </div>
        <div className="dashboard-item" onClick={() => handleClick('/hotels')}>
          <img src={hotelsImage} alt="Hotels" />
          <h3>Hotels</h3>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
