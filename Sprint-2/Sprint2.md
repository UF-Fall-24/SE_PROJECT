**Detailed explanation of the completed work implemented in the backend**

Accommodations
    This file defines the API endpoints to manage accommodation records. The work includes:

CreateAccommodation
    Responsibility: Reads the incoming JSON payload and decodes it into an Accommodation model instance.
    Workflow:
        Validates the request payload by attempting to decode JSON.
        Calls the model's Create() method to insert the new record into the database.
        On success, returns the newly created accommodation as JSON with a status code 201 Created.
        Logs the outcome and handles errors by sending appropriate HTTP error responses.
GetAccommodations
    Responsibility: Retrieves all accommodation records from the database.
    Workflow:
        Calls the model method GetAllAccommodations() to fetch all records.
        Returns the result in JSON format along with a success log message.
        If an error occurs, it sends a 500 Internal Server Error response.
GetAccommodation
    Responsibility: Retrieves a specific accommodation record based on its ID.
    Workflow:
        Extracts the id from the URL parameters and converts it to an integer.
        Calls the model's GetByID() method to get the record.
        Returns the accommodation record if found; otherwise, returns a 404 Not Found error.
UpdateAccommodation
    Responsibility: Updates an existing accommodation record with new details.
    Workflow:
        Extracts the accommodation ID from the URL and decodes the new data from the request body.
        Sets the ID on the accommodation model instance.
        Calls the model's Update() method.
        On success, returns the updated accommodation in JSON; otherwise, sends an error response.
DeleteAccommodation
    Responsibility: Deletes a specific accommodation record.
    Workflow:
        Parses the accommodation ID from the URL.
        Constructs an Accommodation instance with the ID and calls the model's Delete() method.
        Returns a confirmation message if deletion is successful or an error response if it fails.


Bookings (booking.go)
    This file implements the endpoints for handling booking-related operations. The work includes:

CreateBooking
    Responsibility: Accepts a JSON payload to create a new booking record.
    Workflow:
        Decodes the JSON payload into a Booking model instance.
        Calls the model's Create() method to insert the booking into the database.
        On successful creation, it returns the booking with a status code 201 Created.
        Handles errors by returning 400 Bad Request or 500 Internal Server Error based on the issue.
GetBookingsByUser
    Responsibility: Retrieves all bookings associated with a specific user.
    Workflow:
        Extracts the user_id from the URL, converts it to an integer.
        Calls the model method GetBookingsByUser(userID) to get the bookings.
        Returns the list of bookings in JSON format.
        Returns appropriate error messages if user ID is invalid or an error occurs during retrieval.
GetBooking
    Responsibility: Retrieves a single booking by its unique ID.
    Workflow:
        Parses the booking ID from the URL.
        Calls GetBookingByID() from the Booking model to fetch the booking details.
        If the booking is found, it is returned as JSON; otherwise, an error is returned.
CancelBooking
    Responsibility: Cancels an existing booking by updating its status.
    Workflow:
        Extracts the booking ID from the URL.
        Retrieves the booking record using GetBookingByID().
        Calls the model’s Cancel() method to update the booking’s status.
        Returns a success message if the cancellation is successful.
        Sends appropriate HTTP error codes if the booking is not found or cancellation fails.


Hotels (hotel.go)
    This file contains the endpoints for managing hotel records. The work covers:

CreateHotel
    Responsibility: Creates a new hotel record.
    Workflow:
        Decodes the JSON payload into a Hotel model instance.
        Constructs an SQL INSERT query using the provided hotel details.
        Executes the query using the database connection from the configuration.
        Retrieves the last inserted ID and assigns it to the hotel model.
        Returns the created hotel record with a status code 201 Created.
        Logs any errors encountered during insertion.
GetHotels
    Responsibility: Retrieves all hotel records from the database.
    Workflow:
        Executes a SQL query to select all hotels.
        Iterates over the result set, scanning each record into a Hotel model.
        Returns the list of hotels in JSON format.
        Handles any query or scanning errors with proper error logging and HTTP error responses.
GetHotel
    Responsibility: Retrieves details for a specific hotel by ID.
    Workflow:
        Reads the hotel ID from the URL and converts it.
        Executes a SQL query with a WHERE clause to fetch the hotel record.
        Returns the hotel details if found; otherwise, returns a 404 Not Found error.
UpdateHotel
    Responsibility: Updates an existing hotel record with new data.
    Workflow:
        Parses the hotel ID and decodes the update payload.
        Executes a SQL UPDATE statement with the new details.
        Sets the updated hotel ID and returns the updated hotel record in JSON.
        Handles errors and returns the corresponding error status.
DeleteHotel
    Responsibility: Deletes a hotel record from the database.
    Workflow:
        Parses the hotel ID from the URL.
        Executes a SQL DELETE query.
        Returns a confirmation message upon successful deletion.
        Handles and logs any errors during the deletion process.
GetHotelsByLocation
    Responsibility: Retrieves hotels filtered by a provided location (city).
    Workflow:
        Reads the location query parameter from the URL.
        Executes a SQL query filtering hotels by the specified city.
        Returns the filtered list of hotels in JSON format.
        Returns an error if the location parameter is missing or if the query fails.

Packages (package.go)
    This file provides the endpoints for managing travel packages. The work includes:

CreatePackage
    Responsibility: Inserts a new package record into the database.
    Workflow:
        Decodes the JSON payload into a Package model instance.
        Executes an SQL INSERT query with package details (name, description, price, days, nights, and location).
        Retrieves the newly inserted package ID and assigns it to the model.
        Returns the created package record with a status code 201 Created.
        Handles any errors during payload decoding or query execution.
GetPackages
    Responsibility: Retrieves all package records from the database.
    Workflow:
        Executes an SQL SELECT query to fetch all package records.
        Iterates through the result set and scans each row into a Package model instance.
        Returns the list of packages in JSON format.
        Handles errors by logging and sending a 500 Internal Server Error response.
GetPackage
    Responsibility: Retrieves a single package based on its ID.
    Workflow:
        Extracts and converts the package ID from the URL.
        Executes a SQL SELECT query with a WHERE clause for the given package ID.
        Returns the package details if found; if not, responds with a 404 Not Found.
UpdatePackage
    Responsibility: Updates an existing package record.
    Workflow:
        Reads the package ID from the URL and decodes the update payload.
        Executes a SQL UPDATE query to modify the package record.
        Sets the package ID on the model and returns the updated record.
        Handles errors by returning appropriate HTTP error statuses.
DeletePackage
    Responsibility: Deletes a package record from the database.
    Workflow:
        Extracts the package ID from the URL.
        Executes a SQL DELETE query to remove the record.
        Returns a JSON response with a confirmation message indicating that the package was deleted.
        Logs errors and returns error responses if the deletion fails.



**List of unit tests for the backend APIs**

Accommodation Tests
    TestCreateAccommodation
        Verifies that a new accommodation can be created successfully.

    TestGetAccommodations
        Validates retrieving all accommodations.

    TestGetAccommodationByID
        Checks that fetching a single accommodation by its ID returns the correct data.

    TestUpdateAccommodation
        Tests updating an existing accommodation record.

    TestDeleteAccommodation
        Ensures that an accommodation record can be deleted.

    TestCreateAccommodation_MissingFields
        Tests creating an accommodation with missing required fields (expecting a 400 Bad Request).

    TestCreateAccommodation_InvalidData
        Verifies error handling when invalid data types are provided.

    TestCreateAccommodation_HotelNotFound
        Checks behavior when attempting to create an accommodation for a non-existent hotel.


Booking Tests
    TestCreateBooking
        Ensures that a booking can be created successfully.

    TestGetBookingsByUser
        Validates retrieving all bookings associated with a specific user.

    TestGetBookingByID
        Checks that fetching a booking by its ID returns the correct booking.

    TestCancelBooking
        Tests canceling an existing booking successfully.

    TestCreateBooking_MissingFields
        Verifies proper error response when required booking fields are missing.

    TestCreateBooking_InvalidData
        Tests handling of invalid data types in booking creation.

    TestCreateBooking_UserNotFound
        Checks the response when a booking is created for a non-existent user.

    TestCreateBooking_PackageNotFound
        Verifies behavior when a booking references a non-existent package.

    TestCancelBooking_NotFound
        Tests canceling a non-existent booking (expecting a 404 Not Found).

    TestCancelBooking_InvalidID
        Checks the error response for canceling a booking with an invalid ID format.

    TestCancelBooking_AlreadyCanceled
        Validates that trying to cancel an already canceled booking returns the appropriate error.


Hotel Tests
    TestCreateHotel
        Verifies that a hotel can be created successfully.

    TestGetHotels
        Ensures that a list of hotels can be retrieved.

    TestGetHotelByID
        Checks that fetching a single hotel by ID returns the correct hotel data.

    TestUpdateHotel
        Tests updating an existing hotel record.

    TestDeleteHotel
        Validates that a hotel can be deleted successfully.


Package Tests
    TestCreatePackage
        Ensures that a new package can be created successfully.

    TestGetPackages
        Validates that all packages can be retrieved.

    TestGetPackageByID
        Checks that fetching a package by its ID returns the correct package.

    TestUpdatePackage
        Tests updating an existing package record.

    TestDeletePackage
        Verifies that a package can be deleted successfully.


User Tests
    TestGetUserProfile
        Verifies that a user's profile can be retrieved successfully by ID.

    TestUpdateUserProfile
        Tests updating a user's profile successfully.

    TestGetUserProfile_NotFound
        Checks the behavior when attempting to retrieve a non-existent user.

    TestGetUserProfile_InvalidID
        Validates the error response when an invalid user ID is provided.

    TestUpdateUserProfile_NotFound
        Tests updating a non-existent user (expecting a 404 Not Found).

    TestUpdateUserProfile_InvalidID
        Checks error handling for updating a user with an invalid ID.

    TestUpdateUserProfile_InvalidPayload
        Verifies that providing an invalid payload for updating a user results in an error.

**Documentation of API's**

### Booking Endpoints

| Method | Endpoint                 | Description                                                           |
| ------ | ------------------------ | --------------------------------------------------------------------- |
| POST   | /bookings                | Create a new booking. Inserts a new booking record into the database. |
| GET    | /bookings/{id}           | Retrieve a specific booking by its unique ID.                         |
| GET    | /bookings/user/{user_id} | Retrieve all bookings for a specified user.                          |
| PUT    | /bookings/{id}/cancel    | Cancel an existing booking by updating its status.                    |

### Accommodation Endpoints

| Method | Endpoint              | Description                                           |
| ------ | --------------------- | ----------------------------------------------------- |
| POST   | /accommodations       | Create a new accommodation record.                  |
| GET    | /accommodations       | Retrieve a list of all accommodations.              |
| GET    | /accommodations/{id}  | Retrieve details of a specific accommodation by ID. |
| PUT    | /accommodations/{id}  | Update an existing accommodation record.            |
| DELETE | /accommodations/{id}  | Delete an accommodation record.                     |

### Package Endpoints

| Method | Endpoint           | Description                                                               |
| ------ | ------------------ | ------------------------------------------------------------------------- |
| POST   | /packages          | Create a new package. Inserts a new package record into the database.     |
| GET    | /packages          | Retrieve all packages.                                                    |
| GET    | /packages/{id}     | Retrieve details of a specific package by its ID.                         |
| PUT    | /packages/{id}     | Update an existing package record by its ID.                              |
| DELETE | /packages/{id}     | Delete a package record by its ID.                                         |

### Hotel Endpoints

| Method | Endpoint                            | Description                                                     |
| ------ | ----------------------------------- | --------------------------------------------------------------- |
| POST   | /hotels                           | Create a new hotel. Inserts a new hotel record into the database. |
| GET    | /hotels                           | Retrieve a list of all hotels.                                  |
| GET    | /hotels/{id}                      | Retrieve details of a specific hotel by its ID.                 |
| PUT    | /hotels/{id}                      | Update an existing hotel record by its ID.                      |
| DELETE | /hotels/{id}                      | Delete a hotel record by its ID.                                 |
| GET    | /hotels/location?location=<city>  | Retrieve hotels filtered by location (city).                    |



Sprint 2: Backend API Documentation
NAME	UFID
Kopparla Varshini	22060396
Karthik Karnam      37476457	

User Profile Management
    Get User Profile 
    Endpoint: GET /users/{id} 
    Description: Retrieves user details by ID. 
    Authentication: Requires JWT token. 
    Request Headers: Authorization: Bearer <token> 
    Content-Type: application/json 
    Response Example (Success - 200 OK): 
        {	 "id": 1, 
        "username": "john_doe", 
        "email": "john@example.com" 
        } 
    Error Responses: 
        404 Not Found → User does not exist. 
        400 Bad Request → Invalid user ID format.

    Update User Profile
    Endpoint: PUT /users/{id} 
    Description: Updates user profile details (username & email). 
    Authentication: Requires JWT token. 
        Request Body Example: 
        {
        "username": "new_name", 
        "email": "new_email@example.com" 
        } 
        Response Example (Success - 200 OK): 
        {
        "message": "Profile updated successfully" 
        } 
    Error Responses: 
        400 Bad Request → Invalid request payload. 
        500 Internal Server Error → Database update failed.
    Edge Cases 
        User Not Found () 
        Invalid User ID Format () 
        Invalid Request Payload ()

Accommodations API
    Create Accommodation
    Endpoint: POST /accommodations 
    Description: Adds a new accommodation linked to a hotel. 
    Authentication: Requires JWT token. 
        Request Body Example: 
        { 
        "hotel_id": 1, 
        "room_type": "Deluxe", 
        "check_in": "2024-07-01", 
        "check_out": "2024-07-05", 
        "price": 350.00 
        } 
        Response Example (Success - 201 Created): 
        { 
        "id": 10, 
        "hotel_id": 1, 
        "room_type": "Deluxe", 
        "check_in": "2024-07-01", 
        "check_out": "2024-07-05", 
        "price": 350.00 
        } 
    Error Responses: 
        400 Bad Request → Missing or invalid fields. 
        404 Not Found → Hotel ID does not exist. 
        500 Internal Server Error → Database insertion error. 
    Edge Cases 
        Missing Fields () 
        Invalid Data Types () 
        Hotel Not Found ()


    Get Accommodations
    Endpoint: GET /accommodations 
    Description: Retrieves all accommodations. 
    Response Example (Success - 200 OK):
        [
            { 
                "id": 10, 
                "hotel_id": 1, 
                "room_type": "Deluxe", 
                "check_in": "2024-07-01", 
                "check_out": "2024-07-05", 
                "price": 350.00 
            } 
        ] 

    Get Accommodation by ID
    Endpoint: GET /accommodations/{id} 
    Description: Retrieves a single accommodation by ID. 
    Response Example (Success - 200 OK): 
        { 
            "id": 10, 
            "hotel_id": 1, 
            "room_type": "Deluxe", 
            "check_in": "2024-07-01", 
            "check_out": "2024-07-05", 
            "price": 350.00 
        } 

    Update Accommodation
    Endpoint: PUT /accommodations/{id} 
    Description: Updates an existing accommodation record.
        Response Example (Success - 200 OK): 
        { 
        "message": "Accommodation updated successfully" 
        } 

    Delete Accommodation
    Endpoint: DELETE /accommodations/{id} 
    Description: Deletes an accommodation record. 
    Response Example (Success - 200 OK): 
        { 
            "message": "Accommodation deleted successfully" 
        } 
    Edge Cases
        Missing Fields (400 Bad Request) 
        Invalid Data Types (400 Bad Request) 
        Hotel Not Found (404 Not Found) 
        Accommodation Not Found (404 Not Found) 
        Invalid Accommodation ID (400 Bad Request)
Booking API
    Create Booking
    Endpoint: POST /bookings 
    Description: Creates a new booking with package & optional accommodation & vehicle. 
    Request Body Example: 
        { 
        "user_id": 5, 
        "package_id": 3, 
        "accommodation_id": 10, 
        "vehicle_id": 2 
        } 
    Response Example (Success - 201 Created): 
        { 
        "id": 25, 
        "user_id": 5, 
        "package_id": 3, 
        "accommodation_id": 10, 
        "vehicle_id": 2, 
        "status": "Pending" 
        } 
    Error Responses: 
        400 Bad Request → Missing required fields. 
        404 Not Found → User, Package, Accommodation, or Vehicle does not exist. 
        500 Internal Server Error → Foreign key constraint failure.

    Get Bookings by User
    Endpoint: GET /bookings/user/{user_id} 
    Description: Retrieves all bookings for a given user. 
    Response Example (Success - 200 OK): 
        [ 
            { 
                "id": 25, 
                "user_id": 5, 
                "package_id": 3, 
                "accommodation_id": 10, 
                "vehicle_id": 2,
                "status": "Confirmed" 
            } 
        ] 
    Error Responses: 
        400 Bad Request → Invalid user ID. 
        500 Internal Server Error → Error retrieving bookings.


    Get Booking by ID
    Endpoint: GET /bookings/{id} 
    Description: Retrieves a single booking by its ID. 
    Response Example (Success - 200 OK): 
        { 
            "id": 25, 
            "user_id": 5, 
            "package_id": 3, 
            "accommodation_id": 10, 
            "vehicle_id": 2, 
            "status": "Confirmed" 
        } 
    Error Responses: 
        400 Bad Request → Invalid booking ID. 
        404 Not Found → Booking not found.

    Cancel Booking
    Endpoint: PUT /bookings/{id}/cancel 
    Description: Cancels a booking and removes associated accommodation if applicable. 
    Response Example (Success - 200 OK): 
        { 
            "message": "Booking canceled successfully" 
        } 
    Error Responses: 
        404 Not Found → Booking ID does not exist. 
        400 Bad Request → Booking is already canceled. 
        500 Internal Server Error → Database update failed. 
    Edge Cases 
        User Not Found () 
        Package Not Found () 
        Invalid Booking ID () 
        Booking Already Canceled ()

Packages API
    Create Package 
    Endpoint: POST /packages 
    Description: Inserts a new package into the database. 
    Request Headers
        Authorization: Bearer <token>
    Response Example (Success - 201 Created): 
        { 
            "id": 1, 
            "package_name": "Beach Paradise", 
            "package_description": "Luxury beach resort experience", 
            "package_price": 999.99, 
            "days": 5, 
            "nights": 4, 
            "location": "Maldives" 
        } 

    Get All Packages 
    Endpoint: GET /packages 
    Description: Retrieves all packages. 
    Success Response
        Status Code: 200 OK
        Content-Type: application/json
        Response sample (Success - 200)
        [
            {
                "id": 5,
                "package_name": "Summer Escape",
                "package_description": "A 7-day package including accommodation, meals, and sightseeing.",
                "package_price": 1200.00,
                "days": 7,
                "nights": 6,
                "location": "Bali",
                "created_at": "2025-03-03T13:00:00Z",
                "updated_at": "2025-03-03T13:00:00Z"
            },
            {
                "id": 6,
                "package_name": "Winter Wonderland",
                "package_description": "Enjoy the winter in style with our all-inclusive package.",
                "package_price": 1500.00,
                "days": 5,
                "nights": 4,
                "location": "Switzerland",
                "created_at": "2025-03-04T11:30:00Z",
                "updated_at": "2025-03-04T11:30:00Z"
            }
        ]
    Error Responses
        500 Internal Server Error: If an error occurs while retrieving the packages.


    Get Package by ID 
    Endpoint: GET /packages/{id} 
    Description: Retrieves a single package by its ID.
    Success Response
        Status Code: 200 OK
        Content-Type: application/json
        {
            "id": 5,
            "package_name": "Summer Escape",
            "package_description": "A 7-day package including accommodation, meals, and sightseeing.",
            "package_price": 1200.00,
            "days": 7,
            "nights": 6,
            "location": "Bali",
            "created_at": "2025-03-03T13:00:00Z",
            "updated_at": "2025-03-03T13:00:00Z"
        }
    Error Responses
        400 Bad Request: If the id is invalid.
        404 Not Found: If no package is found with the given id.


    Update Package 
    Endpoint: PUT /packages/{id} 
    Description: Updates an existing package record. 
    Request Headers
        Authorization: Bearer <token>
    Request Body Example
        {
        "package_name": "Summer Escape Plus",
        "package_description": "An updated description with added features.",
        "package_price": 1300.00,
        "days": 7,
        "nights": 6,
        "location": "Bali"
        }
    Response Example (Success - 200 OK): 
    { 
        "message": "Package updated successfully" 
    } 
    Error Responses
        400 Bad Request: If the id is invalid or the payload is malformed.
        500 Internal Server Error: If an error occurs during the update.

    Delete Package 
    Endpoint: DELETE /packages/{id} 
    Description: Deletes a package record. 
    Request Headers
        Authorization: Bearer <token>
    Response Example (Success - 200 OK): 
    { 
        "message": "Package was deleted" 
    } 
    Error Responses
        400 Bad Request: If the id is invalid.
        500 Internal Server Error: If an error occurs during deletion.


Hotels API
    Create Hotel 
    Endpoint: POST /hotels 
    Description: Inserts a new hotel into the database. 
    Request Body Example
        {
        "hotel_name": "Grand Plaza",
        "address": "456 City Center Blvd",
        "city": "New York",
        "description": "A luxurious hotel in the heart of the city.",
        "rating": 4.5,
        "room_type": "Deluxe",
        "room_price": 350.00
        }
    Response Example (Success - 201 Created): 
        {
            "id": 20,
            "hotel_name": "Grand Plaza",
            "address": "456 City Center Blvd",
            "city": "New York",
            "description": "A luxurious hotel in the heart of the city.",
            "rating": 4.5,
            "room_type": "Deluxe",
            "room_price": 350.00,
            "created_at": "2025-03-03T14:20:00Z",
            "updated_at": "2025-03-03T14:20:00Z"
        }
    Error Responses
        400 Bad Request: If the payload is invalid.
        500 Internal Server Error: If an error occurs during creation.



    Get All Hotels 
    Endpoint: GET /hotels 
    Description: Retrieves all hotels.
    Success Response
        Status Code: 200 OK
        Content-Type: application/json
    [
        {
            "id": 20,
            "hotel_name": "Grand Plaza",
            "address": "456 City Center Blvd",
            "city": "New York",
            "description": "A luxurious hotel in the heart of the city.",
            "rating": 4.5,
            "room_type": "Deluxe",
            "room_price": 350.00,
            "created_at": "2025-03-03T14:20:00Z",
            "updated_at": "2025-03-03T14:20:00Z"
        },
        {
            "id": 21,
            "hotel_name": "Cozy Inn",
            "address": "789 Suburban Road",
            "city": "Boston",
            "description": "A comfortable and affordable inn.",
            "rating": 4.0,
            "room_type": "Standard",
            "room_price": 150.00,
            "created_at": "2025-03-04T09:15:00Z",
            "updated_at": "2025-03-04T09:15:00Z"
        }
    ]
    Error Responses
        500 Internal Server Error: If an error occurs while retrieving hotels.


    Get Hotel by ID 
    Endpoint: GET /hotels/{id} 
    Description: Retrieves a single hotel by its ID. 
    Success Response
        Status Code: 200 OK
        Content-Type: application/json
    {
        "id": 20,
        "hotel_name": "Grand Plaza",
        "address": "456 City Center Blvd",
        "city": "New York",
        "description": "A luxurious hotel in the heart of the city.",
        "rating": 4.5,
        "room_type": "Deluxe",
        "room_price": 350.00,
        "created_at": "2025-03-03T14:20:00Z",
        "updated_at": "2025-03-03T14:20:00Z"
    }
    Error Responses
        400 Bad Request: If the id is invalid.
        404 Not Found: If no hotel is found with the given id.


    Get Hotels by Location 
    Endpoint: GET /hotels/location?city={city} 
    Description: Retrieves all hotels based on a specified city. 

    Update Hotel 
    Endpoint: PUT /hotels/{id} 
    Description: Updates an existing hotel record. 
    Request Body Example
    {
        "hotel_name": "Grand Plaza Renovated",
        "address": "456 City Center Blvd",
        "city": "New York",
        "description": "Updated description after renovation.",
        "rating": 4.6,
        "room_type": "Executive Suite",
        "room_price": 400.00
    }

    Error Responses
        400 Bad Request: If the id is invalid or the JSON payload is malformed.
        500 Internal Server Error: If an error occurs during the update

    Success Response
        Status Code: 200 OK
        Content-Type: application/json
    Response Example (Success - 200 OK): 
    { 
    "message": "Hotel updated successfully" 
    } 

    Delete Hotel 
    Endpoint: DELETE /hotels/{id} 
    Description: Deletes a hotel record. 
    Response Example (Success - 200 OK): 
    { 
    "message": "Hotel was deleted" 
    } 
