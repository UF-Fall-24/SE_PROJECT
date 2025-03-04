import React, { useEffect, useState } from "react";
import { FaShoppingCart } from "react-icons/fa"; // Cart icon
import { useNavigate } from "react-router-dom";
import natureImage from "./nature.jpg"; // Import local image

const Dashboard = () => {
    const [packages, setPackages] = useState([]);
    const [cart, setCart] = useState([]);
    const [showCart, setShowCart] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");

        if (!token) {
            navigate("/login"); // ✅ Redirect to login if no token found
            return;
        }

        const fetchPackages = async () => {
            try {
                const response = await fetch("http://localhost:8000/packages", {
                    headers: { Authorization: `Bearer ${token}`, "Content-Type": "application/json" },
                });

                if (!response.ok) throw new Error("Failed to fetch packages");

                const data = await response.json();
                setPackages(data);
            } catch (error) {
                console.error("Error fetching packages:", error);
            }
        };

        fetchPackages();
    }, [navigate]);

    // Function to add an item to the cart
    const addToCart = (pkg) => {
        setCart([...cart, pkg]);
    };

    // Function to remove an item from the cart
    const removeFromCart = (index) => {
        setCart(cart.filter((_, i) => i !== index));
    };

    return (
        <div>
            {/* Header with Cart Icon */}
            <div style={{ display: "flex", justifyContent: "space-between", padding: "10px", background: "#333", color: "white" }}>
                <h2>Travel Packages</h2>
                <div style={{ position: "relative" }}>
                    <FaShoppingCart 
                        size={30} 
                        onClick={() => setShowCart(!showCart)} 
                        style={{ cursor: "pointer" }} 
                    />
                    {cart.length > 0 && (
                        <span style={{
                            position: "absolute",
                            top: "-5px",
                            right: "-10px",
                            background: "red",
                            color: "white",
                            borderRadius: "50%",
                            padding: "5px 10px",
                            fontSize: "14px"
                        }}>
                            {cart.length}
                        </span>
                    )}
                </div>
            </div>

            {/* Cart Dropdown */}
            {showCart && (
                <div style={{
                    position: "absolute",
                    right: "10px",
                    top: "50px",
                    background: "#333",  
                    color: "white",       
                    borderRadius: "8px",
                    padding: "10px",
                    width: "250px",
                    boxShadow: "0px 4px 6px rgba(0,0,0,0.1)"
                }}>
                    <h4 style={{ borderBottom: "1px solid white", paddingBottom: "5px" }}>Cart Items</h4>
                    {cart.length === 0 ? (
                        <p>No items in cart</p>
                    ) : (
                        <ul style={{ listStyle: "none", padding: 0 }}>
                            {cart.map((item, index) => (
                                <li key={index} style={{ borderBottom: "1px solid #555", padding: "5px 0" }}>
                                    {item.package_name} - <strong>${item.package_price}</strong>
                                    <button 
                                        onClick={() => removeFromCart(index)} 
                                        style={{
                                            marginLeft: "10px",
                                            backgroundColor: "red",
                                            color: "white",
                                            border: "none",
                                            padding: "5px",
                                            cursor: "pointer",
                                            fontSize: "12px"
                                        }}>
                                        Remove
                                    </button>
                                </li>
                            ))}
                        </ul>
                    )}
                </div>
            )}

            {/* Packages Section */}
            <h3>Available Packages</h3>
            <div style={{ display: "flex", flexWrap: "wrap", gap: "20px" }}>
                {packages.map((pkg) => (
                    <div key={pkg.id} style={{ border: "1px solid #ddd", padding: "10px", width: "300px", textAlign: "center" }}>
                        <img src={natureImage} alt="Package" style={{ width: "100%", height: "200px", objectFit: "cover" }} />
                        <h4>{pkg.package_name}</h4>
                        <p>{pkg.package_description}</p>
                        <p><strong>Location:</strong> {pkg.location}</p>
                        <p><strong>Price:</strong> ${pkg.package_price}</p>
                        <p><strong>Days:</strong> {pkg.days} | <strong>Nights:</strong> {pkg.nights}</p>
                        <button 
                            onClick={() => addToCart(pkg)} 
                            style={{ padding: "10px", backgroundColor: "#28a745", color: "white", border: "none", cursor: "pointer" }}>
                            Add to Cart
                        </button>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Dashboard;
