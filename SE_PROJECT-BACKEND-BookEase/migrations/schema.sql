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
    room_type VARCHAR(100),        -- e.g., Deluxe, Suite, Standard, Luxury
    room_price DECIMAL(10,2),        -- Price corresponding to the room type
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- INSERT INTO hotels (hotel_name, address, city, description, rating, room_type, room_price)
-- VALUES
-- ('Grand Plaza', '123 Main St', 'New York', 'A luxury hotel in downtown New York.', 4.5, 'Deluxe', 250.00),
-- ('Seaside Resort', '456 Beach Ave', 'Miami', 'Relax by the ocean in a modern resort.', 4.2, 'Suite', 300.00),
-- ('Mountain Inn', '789 Alpine Rd', 'Denver', 'Cozy inn with stunning mountain views.', 4.0, 'Standard', 180.00),
-- ('City Center Hotel', '101 Central Blvd', 'Chicago', 'Located in the heart of the city with modern amenities.', 4.3, 'Deluxe', 220.00),
-- ('Riverside Hotel', '202 River St', 'Austin', 'Enjoy scenic river views and contemporary comfort.', 4.1, 'Suite', 260.00),
-- ('Historic Mansion', '303 Old Town Rd', 'Boston', 'Experience history in a beautifully renovated mansion.', 4.6, 'Luxury', 350.00),
-- ('Urban Boutique', '404 Fashion Ln', 'San Francisco', 'Chic and modern boutique hotel.', 4.4, 'Deluxe', 270.00),
-- ('Eco Lodge', '505 Greenway', 'Portland', 'Sustainable hotel with eco-friendly amenities.', 4.2, 'Standard', 190.00),
-- ('Lakeside Retreat', '606 Lakeview Dr', 'Minneapolis', 'Relax by the lake in a serene setting.', 4.5, 'Suite', 240.00),
-- ('Downtown Express', '707 Market St', 'Los Angeles', 'Convenient location with quick access to major attractions.', 4.0, 'Standard', 210.00),
-- ('Royal Palace', '808 King St', 'Las Vegas', 'Experience ultimate luxury and elegance.', 4.8, 'Deluxe', 400.00),
-- ('Coastal Comfort', '909 Ocean Dr', 'San Diego', 'Enjoy the coastal breeze and modern comforts.', 4.3, 'Suite', 280.00),
-- ('Suburban Stay', '111 Suburb Rd', 'Houston', 'Quiet and comfortable for business travelers.', 4.0, 'Standard', 170.00),
-- ('Skyline Heights', '222 Highrise Ave', 'Seattle', 'Spectacular views of the city skyline.', 4.4, 'Deluxe', 230.00),
-- ('Desert Oasis', '333 Sand Dune Rd', 'Phoenix', 'An oasis in the heart of the desert.', 4.2, 'Suite', 260.00),
-- ('Artistic Abode', '444 Culture St', 'Santa Fe', 'A creative space with local art influences.', 4.6, 'Deluxe', 250.00),
-- ('Modern Mansion', '555 Contemporary Ct', 'Dallas', 'Spacious and modern with luxurious amenities.', 4.5, 'Luxury', 350.00),
-- ('Boutique Bliss', '666 Style Ave', 'Atlanta', 'Personalized service in a boutique setting.', 4.3, 'Deluxe', 240.00),
-- ('Paradise Inn', '777 Tropical Blvd', 'Honolulu', 'Relax in a tropical paradise.', 4.7, 'Suite', 320.00),
-- ('Budget Stay', '888 Economy Rd', 'Orlando', 'Affordable stay with essential amenities.', 3.8, 'Standard', 120.00);






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
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (package_id) REFERENCES packages(id),
    FOREIGN KEY (accommodation_id) REFERENCES accommodations(id),
    FOREIGN KEY (payment_id) REFERENCES payment(payment_id)
);


CREATE TABLE IF NOT EXISTS payments (
    payment_id INT AUTO_INCREMENT PRIMARY KEY,
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


-- package bookings table

-- Create package_bookings table without a generated column
CREATE TABLE IF NOT EXISTS package_bookings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    package_booking_id VARCHAR(20) UNIQUE,
    package_id INT NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) AUTO_INCREMENT=1000;



DROP PROCEDURE IF EXISTS InsertPackageBooking;

CREATE PROCEDURE InsertPackageBooking(
    IN p_package_id INT,
    IN p_first_name VARCHAR(100),
    IN p_last_name VARCHAR(100)
)
BEGIN
    DECLARE new_id INT;
    INSERT INTO package_bookings (package_id, first_name, last_name)
    VALUES (p_package_id, p_first_name, p_last_name);
    SET new_id = LAST_INSERT_ID();
    UPDATE package_bookings
    SET package_booking_id = CONCAT('P', new_id)
    WHERE id = new_id;
END;



-- CALL InsertPackageBooking(1, 'John', 'Doe');
-- CALL InsertPackageBooking(2, 'Jane', 'Smith');
-- CALL InsertPackageBooking(3, 'Alice', 'Johnson');
-- CALL InsertPackageBooking(4, 'Bob', 'Brown');
-- CALL InsertPackageBooking(5, 'Carol', 'Davis');


-- accommodation bookings table

-- Create accommodation_bookings table without a generated column
CREATE TABLE IF NOT EXISTS accommodation_bookings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    accommodation_booking_id VARCHAR(20) UNIQUE,
    package_booking_id VARCHAR(20) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    duration VARCHAR(10) NOT NULL,  -- e.g., '7/6' for days/nights
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_package_booking 
        FOREIGN KEY (package_booking_id) REFERENCES package_bookings(package_booking_id)
) AUTO_INCREMENT=1000;


DROP PROCEDURE IF EXISTS InsertAccommodationBooking;
CREATE PROCEDURE InsertAccommodationBooking(
    IN p_package_booking_id VARCHAR(20),
    IN p_first_name VARCHAR(100),
    IN p_last_name VARCHAR(100),
    IN p_price DECIMAL(10,2),
    IN p_duration VARCHAR(10)
)
BEGIN
    DECLARE new_id INT;
    INSERT INTO accommodation_bookings (package_booking_id, first_name, last_name, price, duration)
    VALUES (p_package_booking_id, p_first_name, p_last_name, p_price, p_duration);
    SET new_id = LAST_INSERT_ID();
    UPDATE accommodation_bookings
    SET accommodation_booking_id = CONCAT('A', new_id)
    WHERE id = new_id;
END;

-- CALL InsertAccommodationBooking('P1000', 'John', 'Doe', 300.00, '7/6');
-- CALL InsertAccommodationBooking('P1001', 'Jane', 'Smith', 350.00, '7/6');
-- CALL InsertAccommodationBooking('P1002', 'Alice', 'Johnson', 320.00, '7/6');
-- CALL InsertAccommodationBooking('P1003', 'Bob', 'Brown', 280.00, '7/6');
-- CALL InsertAccommodationBooking('P1004', 'Carol', 'Davis', 310.00, '7/6');
