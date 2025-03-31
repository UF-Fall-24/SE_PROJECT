**Detailed explanation of the completed work implemented in the backend**

Accommodation Booking (accommodation_bookings.go)
This file handles all operations related to the accommodation_bookings table. The work includes:

CreateAccommodationBooking
    Responsibility: Creates a new accommodation booking record.
    Workflow:
        Decodes the incoming JSON payload into an AccommodationBooking model instance.
        Validates the payload and required fields.
        Inserts the new booking into the database using the model's Create() method.
        On success, returns the booking as JSON with a status code 201 Created.
        If validation or insertion fails, responds with appropriate HTTP error codes (400 Bad Request, 500 Internal Server Error).

UpdateAccommodationBooking
    Responsibility: Updates an existing accommodation booking record.
    Workflow:
        Extracts the booking_id from the URL path.
        Decodes the update data from the request body into the model.
        Matches the record using the booking_id and updates the corresponding fields via Update() method.
        Responds with a confirmation message if successful.
        Handles invalid inputs or missing records with relevant error codes (400, 404, 500).

DeleteAccommodationBooking
    Responsibility: Deletes an accommodation booking record using the booking_id.
    Workflow:
        Extracts booking_id from the URL.
        Calls the model’s Delete() method to remove the record from the database.
        On success, responds with a confirmation message in JSON.
        Sends appropriate errors if the ID is invalid or deletion fails.    


**List of unit tests for the backend APIs**

Package Booking Tests
    TestCreatePackageBooking
        Verifies that a new package booking can be created successfully using a mock POST request.
        Ensures a status code 201 Created and checks for the success message "Package booking created successfully" in the response.
    
    TestUpdatePackageBooking
        Validates updating an existing package booking using a mock PUT request to a booking ID like /package_bookings/P1000.
        Ensures a status code 200 OK and confirms the update message "Package booking updated successfully" is returned.

    TestDeletePackageBooking
        Tests the deletion of a package booking using a mock DELETE request.
        Confirms a status code 200 OK and the correct deletion confirmation message "Package booking deleted successfully".


Accommodation Booking Tests
    TestCreateAccommodationBooking
        Verifies that a new accommodation booking can be created successfully using a mock POST request.
        Ensures a status code 201 Created and correct success message in the response.

    TestUpdateAccommodationBooking
        Validates updating an existing accommodation booking using a mock PUT request.
        Ensures a status code 200 OK and correct update confirmation message is returned.

    TestDeleteAccommodationBooking
        Tests the deletion of an accommodation booking using a mock DELETE request.
        Confirms a status code 200 OK and correct deletion confirmation message is received.

**Documentation of API's**

### Accomodation Booking Endpoints

| Method | Endpoint                     | Description                                                                 |
| ------ | ---------------------------- | --------------------------------------------------------------------------  |
| POST   | /accommodation_bookings      | Create a new accommodation booking. Inserts a new record into the database. |
| PUT    | /accommodation_bookings/{id} | Update an existing accommodation booking by ID.                             |
| DELETE | /accommodation_bookings/{id} | Delete an accommodation booking by ID.                                      |


Sprint 3: Backend API Documentation
NAME	UFID
Kopparla Varshini	22060396
Karthik Karnam      37476457

Accommodation Booking API

    Create Accommodation Booking
        Endpoint: POST /accommodation_bookings
        Description: Adds a new accommodation booking record.
        Authentication: Requires JWT token.
        Request Headers: Authorization: Bearer <token> 
        Content-Type: application/json 
        Request Body Example:
            {
                "booking_id": 25,
                "accommodation_id": 10
            }
        Response Example (Success - 201 Created):
            {
            "id": 5,
            "booking_id": 25,
            "accommodation_id": 10
            }
        Error Responses:
            400 Bad Request → Missing or invalid fields.
            404 Not Found → Booking ID or Accommodation ID does not exist.
            500 Internal Server Error → Database insertion error.
        Edge Cases:
            Missing Fields (400 Bad Request)
            Invalid Foreign Keys (404 Not Found)
            Duplicate Booking Handling (500 Internal Server Error)
            
    Update Accommodation Booking
        Endpoint: PUT /accommodation_bookings/{id}
        Description: Updates an existing accommodation booking by ID.
        Authentication: Requires JWT token.
        Request Body Example:
            {
            "booking_id": 25,
            "accommodation_id": 12
            }
        Response Example (Success - 200 OK):
            {
            "message": "Accommodation booking updated successfully"
            }
        Error Responses:
            400 Bad Request → Invalid booking ID or request body.
            404 Not Found → Accommodation booking not found.
            500 Internal Server Error → Error during update operation.

    Delete Accommodation Booking
        Endpoint: DELETE /accommodation_bookings/{id}
        Description: Deletes an accommodation booking by its ID.
        Authentication: Requires JWT token.
        Response Example (Success - 200 OK):
            {
            "message": "Accommodation booking deleted successfully"
            }
        Error Responses:
            400 Bad Request → Invalid booking ID format.
            404 Not Found → Booking record not found.
            500 Internal Server Error → Error during deletion process.