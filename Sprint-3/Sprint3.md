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


