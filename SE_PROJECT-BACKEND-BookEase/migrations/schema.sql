-- Create Users Table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);



CREATE TABLE IF NOT EXISTS packages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    package_name VARCHAR(255) NOT NULL,
    package_description TEXT,
    package_price DECIMAL(10,2) NOT NULL,
    days INT NOT NULL,
    nights INT NOT NULL,
    location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- INSERT INTO packages (package_name, package_description, package_price, days, nights, location)
-- VALUES
-- ('Beach Paradise', 'Enjoy a luxurious beach resort experience complete with water sports and fine dining.', 999.99, 5, 4, 'Maldives'),
-- ('Mountain Adventure', 'Experience a thrilling mountain adventure with trekking, camping, and breathtaking views.', 799.50, 7, 6, 'Nepal'),
-- ('City Escape', 'Discover a vibrant city tour with premium hotels, gourmet dining, and guided excursions.', 599.00, 3, 2, 'New York'),
-- ('Safari Expedition', 'Explore wild landscapes with a guided safari tour through national parks.', 1200.00, 6, 5, 'Kenya'),
-- ('Historical Journey', 'Dive into rich history with visits to ancient ruins and museums.', 850.00, 4, 3, 'Greece'),
-- ('Desert Retreat', 'Enjoy the serenity of the desert with camel rides and star-gazing nights.', 650.00, 3, 2, 'Dubai'),
-- ('Island Hopper', 'Hop between stunning islands with scenic boat rides and water adventures.', 1100.00, 7, 6, 'Philippines'),
-- ('Cultural Immersion', 'Experience local traditions, festivals, and authentic cuisine.', 700.00, 5, 4, 'Japan'),
-- ('Rainforest Expedition', 'Discover exotic wildlife and lush rainforests with expert guides.', 950.00, 6, 5, 'Brazil'),
-- ('Winter Wonderland', 'Ski on powdery slopes and enjoy cozy lodges in a snowy paradise.', 1300.00, 7, 6, 'Switzerland'),
-- ('Countryside Escape', 'Relax in quaint countryside villages and scenic landscapes.', 500.00, 4, 3, 'Italy'),
-- ('Luxury Cruise', 'Sail on a luxury cruise with gourmet dining and onboard entertainment.', 2000.00, 10, 9, 'Caribbean'),
-- ('Adventure Sports', 'Thrill-seekers enjoy bungee jumping, rafting, and extreme sports.', 900.00, 5, 4, 'New Zealand'),
-- ('Wellness Retreat', 'Relax with yoga, meditation, and spa treatments in a serene environment.', 800.00, 6, 5, 'India'),
-- ('Gastronomic Tour', 'Indulge in culinary delights and wine tasting sessions in scenic locales.', 750.00, 4, 3, 'France'),
-- ('Art and Architecture', 'Explore iconic art galleries and architectural marvels in historic cities.', 850.00, 5, 4, 'Spain'),
-- ('Eco Adventure', 'Travel sustainably with eco-friendly tours and nature conservation activities.', 950.00, 6, 5, 'Costa Rica'),
-- ('River Cruise', 'Enjoy a relaxing river cruise with stops at charming towns and historic sites.', 1200.00, 7, 6, 'Europe'),
-- ('Island Romance', 'Perfect for couples, this package offers private beach dinners and sunset cruises.', 1400.00, 5, 4, 'Seychelles'),
-- ('Festival Fiesta', 'Experience vibrant local festivals, parades, and cultural events.', 650.00, 3, 2, 'Spain');


CREATE TABLE IF NOT EXISTS hotels (
    id INT AUTO_INCREMENT PRIMARY KEY,
    hotel_name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    city VARCHAR(255) NOT NULL,
    description TEXT,
    rating DECIMAL(2,1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS vehicles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    vehicle_type VARCHAR(100) NOT NULL,
    model VARCHAR(100),
    capacity INT,
    price DECIMAL(10,2) NOT NULL,
    location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS accommodations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    hotel_id INT NOT NULL,
    room_type VARCHAR(100),
    check_in DATE NOT NULL,
    check_out DATE NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);


CREATE TABLE IF NOT EXISTS bookings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    package_id INT NOT NULL,
    accommodation_id INT DEFAULT NULL,
    vehicle_id INT DEFAULT NULL,
    booking_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) DEFAULT 'Pending',
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (package_id) REFERENCES packages(id),
    FOREIGN KEY (accommodation_id) REFERENCES accommodations(id),
    FOREIGN KEY (vehicle_id) REFERENCES vehicles(id)
);

CREATE TABLE IF NOT EXISTS payments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    booking_id INT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    payment_method VARCHAR(50),
    payment_status VARCHAR(50),
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);




DROP PROCEDURE IF EXISTS CancelBooking;

CREATE PROCEDURE CancelBooking(IN bookingID INT)
BEGIN
    DECLARE accID INT DEFAULT NULL;
    
    -- Retrieve the accommodation booking id (if any) for the given booking
    SELECT accommodation_id INTO accID FROM bookings WHERE id = bookingID;
    
    -- Delete any payment record associated with this booking
    DELETE FROM payments WHERE booking_id = bookingID;
    
    -- Delete the booking record
    DELETE FROM bookings WHERE id = bookingID;
    
    -- If there was an associated accommodation record, delete it as well
    IF accID IS NOT NULL THEN
        DELETE FROM accommodations WHERE id = accID;
    END IF;
END;






