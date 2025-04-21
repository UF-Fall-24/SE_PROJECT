import React, { lazy, Suspense, useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate, Link } from 'react-router-dom';
import { FaHome, FaInfoCircle, FaPhone } from 'react-icons/fa';
import Payment from './components/payment';
// Lazy load components
const Home = lazy(() => import('./components/Home'));
const Register = lazy(() => import('./components/Register'));
const Login = lazy(() => import('./components/Login'));
const About = lazy(() => import('./components/About'));
const Contact = lazy(() => import('./components/Contact'));
const Dashboard = lazy(() => import('./components/Dashboard'));
const ResetPassword = lazy(() => import('./components/ResetPassword'));
const ForgotPassword = lazy(() => import('./components/ForgotPassword'));
const HotelsList = lazy(() => import('./components/HotelsList'));
const CreateHotelForm = lazy(() => import('./components/CreateHotelForm'));
const SearchHotelByID = lazy(() => import('./components/SearchHotelByID'));
const UpdateHotelForm = lazy(() => import('./components/SearchAndUpdateHotel'));
const DeleteHotel = lazy(() => import('./components/DeleteHotel'));
const HotelsByLocation = lazy(() => import('./components/HotelsByLocation'));
const HotelDropdown = lazy(() => import('./components/HotelDropdown'));
const PackagesList = lazy(() => import('./components/PackagesList'));
const PackageDetails = lazy(() => import('./components/PackageDetails'));
const AccommodationHotels = lazy(() => import('./components/AccommodationHotels'));
const HotelDetails = lazy(() => import('./components/HotelDetails'));


const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [token, setToken] = useState(null);
  const [bookedFlights, setBookedFlights] = useState([]);

  useEffect(() => {
    const savedToken = localStorage.getItem('token');
    if (savedToken) {
      setToken(savedToken);
      setIsLoggedIn(true);
    }
  }, []);

  const handleLogin = async (email, password) => {
    try {
      const response = await fetch("http://localhost:8000/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });
      if (response.ok) {
        const data = await response.json();
        localStorage.setItem("token", data.token);
        setIsLoggedIn(true);
        setToken(data.token);
        // Redirect to dashboard after login.
        window.location.href = "/";
      } else {
        alert("Invalid credentials. Please try again.");
      }
    } catch (error) {
      console.error("Error logging in:", error);
      alert("Server error. Please try again.");
    }
  };

  const handleRegister = async (username, email, password) => {
    try {
      const response = await fetch("http://localhost:8000/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, email, password }),
      });
      if (response.ok) {
        alert("Registration successful! Logging in...");
        await handleLogin(email, password);
      } else {
        alert("Registration failed. Try again.");
      }
    } catch (error) {
      console.error("Error registering:", error);
      alert("Server error. Please try again.");
    }
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setToken(null);
    localStorage.removeItem('token');
    setBookedFlights([]);
    window.location.href = "/login";
  };

  return (
    <Router>
      <header className="header">
        <h1>Book Ease</h1>
        <nav className="nav-container">
          <Link to="/" className="nav-item">
            <FaHome className="icon" /> Home
          </Link>
          <Link to="/about" className="nav-item">
            <FaInfoCircle className="icon" /> About Us
          </Link>
          <Link to="/contact" className="nav-item">
            <FaPhone className="icon" /> Contact Us
          </Link>
          {isLoggedIn ? (
            <button onClick={handleLogout} className="nav-item">
              Logout
            </button>
          ) : (
            <>
              <Link to="/login" className="nav-item">Login</Link>
              <Link to="/register" className="nav-item">Register</Link>
            </>
          )}
        </nav>
      </header>

      <div className="container">
        <Suspense fallback={<div>Loading...</div>}>
          <Routes>
            {/* If logged in, default route is Dashboard; otherwise show Home */}
            <Route path="/" element={localStorage.getItem("token") ? <Dashboard /> : <Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/contact" element={<Contact />} />
            <Route path="/register" element={<Register onRegister={handleRegister} />} />
            <Route path="/login" element={<Login onLogin={handleLogin} />} />
            <Route path="/dashboard" element={localStorage.getItem("token") ? <Dashboard bookedFlights={bookedFlights} /> : <Navigate to="/login" />} />
            <Route path="/reset-password" element={<ResetPassword />} />
            <Route path="/forgot-password" element={<ForgotPassword />} />
            <Route path="/hotels" element={localStorage.getItem("token") ? <HotelsList /> : <Navigate to="/login" />} />
            <Route path="/create-hotel" element={localStorage.getItem("token") ? <CreateHotelForm /> : <Navigate to="/login" />} />
            <Route path="/search-hotel" element={localStorage.getItem("token") ? <SearchHotelByID /> : <Navigate to="/login" />} />
            <Route path="/update-hotel/:id" element={localStorage.getItem("token") ? <UpdateHotelForm /> : <Navigate to="/login" />} />
            <Route path="/delete-hotel" element={localStorage.getItem("token") ? <DeleteHotel /> : <Navigate to="/login" />} />
            <Route path="/search-hotels" element={localStorage.getItem("token") ? <HotelsByLocation /> : <Navigate to="/login" />} />
            <Route path="/packages" element={localStorage.getItem("token") ? <PackagesList /> : <Navigate to="/login" />} />
            <Route path="/package-details/:id" element={localStorage.getItem("token") ? <PackageDetails /> : <Navigate to="/login" />} />
            <Route path="/accommodation-hotels" element={localStorage.getItem("token") ? <AccommodationHotels /> : <Navigate to="/login" />} />
  <Route path="/hotel-details/:id" element={localStorage.getItem("token") ? <HotelDetails /> : <Navigate to="/login" />} />
  
          </Routes>
        </Suspense>
      </div>
    </Router>
  );
};

export default App;
