import React, { useState, useRef, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './HotelDropdown.css'; // Create this file to style the dropdown

const HotelDropdown = () => {
  const navigate = useNavigate();
  const [isOpen, setIsOpen] = useState(false);
  const dropdownRef = useRef(null);

  // Toggle dropdown open/close
  const handleToggle = () => {
    setIsOpen(prev => !prev);
  };

  // When an option is clicked, navigate to that route and close the dropdown
  const handleOptionClick = (route) => {
    navigate(route);
    setIsOpen(false);
  };

  // Close dropdown if clicking outside of it
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
        setIsOpen(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <div className="hotel-dropdown" ref={dropdownRef}>
      <button className="dropdown-toggle" onClick={handleToggle}>
        Hotel Actions &#9662;
      </button>
      {isOpen && (
        <ul className="dropdown-menu">
          <li onClick={() => handleOptionClick('/hotels')}>View Hotels</li>
          <li onClick={() => handleOptionClick('/create-hotel')}>Add Hotel</li>
          <li onClick={() => handleOptionClick('/search-hotel')}>Get Hotel By ID</li>
          <li onClick={() => handleOptionClick('/update-hotel/1')}>Update Hotel (example ID)</li>
          <li onClick={() => handleOptionClick('/delete-hotel')}>Delete Hotel</li>
          <li onClick={() => handleOptionClick('/search-hotels')}>Search Hotels</li>
        </ul>
      )}
    </div>
  );
};

export default HotelDropdown;
