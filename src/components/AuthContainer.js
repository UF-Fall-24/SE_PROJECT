import React, { useState } from "react";
import Register from "./Register";
import Login from "./Login";

const AuthContainer = () => {
  const [isLogin, setIsLogin] = useState(true);

  // Function to switch to login page
  const handleSwitchToLogin = () => {
    setIsLogin(true);
  };

  // Function to switch to register page
  const handleSwitchToRegister = () => {
    setIsLogin(false);
  };

  return (
    <div>
      {isLogin ? (
        <Login onSwitchToRegister={handleSwitchToRegister} />
      ) : (
        <Register onSwitchToLogin={handleSwitchToLogin} />
      )}
    </div>
  );
};

export default AuthContainer;
