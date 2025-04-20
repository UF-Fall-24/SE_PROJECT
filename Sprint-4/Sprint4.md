**Detailed explanation of the completed work implemented in the backend**

Payments (payment.go)
This file handles all the CRUD operations for managing payment records in the payments table.

CreatePayment
    Responsibility: Handles the creation of a new payment record.
    Workflow:
        Accepts a JSON payload containing booking_id, amount, payment_method, and payment_status.
        Decodes the payload into a Payment model instance.
        Calls the model’s Create() method, which inserts the record into the database.
        Responds with the created payment in JSON format and HTTP status 201 Created.
        If validation or insertion fails (e.g., due to a foreign key constraint on booking_id), returns 400 Bad Request or 500 Internal Server Error.

GetPayment
    Responsibility: Retrieves a single payment record by its ID.
    Workflow:
        Extracts id from the URL path.
        Calls the model’s GetByID() method to fetch the payment from the database.
        If found, returns the payment data in JSON format with 200 OK.
        If not found, returns 404 Not Found.

GetAllPayments
    Responsibility: Retrieves all payment records.
    Workflow:
        Calls the model’s GetAllPayments() method to fetch a list of payments.
        Returns the list in JSON format with 200 OK.
        Logs and returns a 500 Internal Server Error if retrieval fails.

UpdatePayment
    Responsibility: Updates the fields of an existing payment record.
    Workflow:
        Extracts id from the URL.
        Reads and decodes the JSON request body into a Payment model instance.
        Sets the ID on the model and calls its Update() method to apply changes.
        Returns a success message with 200 OK if update is successful.
        If an error occurs, returns 400 Bad Request or 500 Internal Server Error.

DeletePayment
    Responsibility: Deletes a payment record by its ID.
    Workflow:
        Extracts id from the URL.
        Initializes a Payment model instance and calls its Delete() method.
        If successful, returns a confirmation message with 200 OK.
        Returns 500 Internal Server Error if deletion fails.

Prices (prices.go)
This file handles the computation of dynamic total prices by aggregating package and optional accommodation costs.

    GetTotalPrice
        Responsibility:Computes and returns the total price of a travel booking by summing the base package price and the optional accommodation price (if provided).
        Workflow:
        Read and Validate Query Parameters:
            Extracts package_id (required) and accommodation_id (optional) from query parameters.
            Validates if package_id is a valid integer. If missing or invalid, returns a 400 Bad Request.
        Fetch Package Price:
            Queries the database to fetch package_price from the packages table using the provided package_id.
            If the package is not found, returns a 404 Not Found.
        Optionally Fetch Accommodation Price:
            If accommodation_id is provided and valid, attempts to fetch price from the accommodations table.
            If accommodation is not found or ID is invalid, the price is ignored silently (fails gracefully).
        Compute Total:
            Adds package price and accommodation price (if found) to calculate the final total.
        Respond with Total Price:
        Returns a JSON response like:
            {
            "total_price": 850.00
            }
        Content-Type: application/json
        Status Code: 200 OK
        Error Handling:
            400 Bad Request: If package_id is missing or not an integer.
            404 Not Found: If no package with the given ID exists.
            200 OK: If successful, regardless of accommodation ID validity.


**List of unit tests for the backend APIs**

Payment Tests
    TestCreatePayment
        Verifies that a new payment can be created successfully using a mock POST request.
        Ensures a status code 201 Created and checks for the correct payment object in the response body.

    TestGetAllPayments
        Validates the retrieval of all payment records using a mock GET request.
        Confirms that the status code returned is 200 OK and the response is a JSON array.

    TestGetPayment
        Checks the retrieval of a specific payment record by ID.
        Ensures that a mock GET request returns status code 200 OK and a valid payment JSON object.

    TestUpdatePayment
        Tests the update of an existing payment using a mock PUT request.
        Confirms that the response includes status 200 OK and the message "Payment updated successfully".

    TestDeletePayment
        Validates the deletion of a payment record using a mock DELETE request.
        Checks that the response returns 200 OK and confirms deletion with a message "Payment deleted successfully".

Price Tests
    TestGetTotalPrice_WithAccommodation
        Verifies that the total price is correctly calculated when both package_id and accommodation_id are provided in the query.
        Ensures the response includes a 200 OK status and the computed total is accurate.

    TestGetTotalPrice_WithoutAccommodation
        Validates that the total price is returned correctly using only package_id, without including accommodation cost.
        Confirms that the status code is 200 OK and the JSON response includes only the package price.

    TestGetTotalPrice_InvalidPackageID
        Tests behavior when package_id is missing or invalid (e.g., a string or empty value).
        Expects a 400 Bad Request status and validates the error message in the response.

    TestGetTotalPrice_PackageNotFound
        Simulates the scenario where a non-existent package_id is used.
        Verifies that the API returns a 404 Not Found response.



*Documentation of API's*

### Payment Endpoints

| Method | Endpoint            | Description                                                                |
| ------ | --------------------| -------------------------------------------------------------------------- |
| POST   | /payments           | Create a new payment. Inserts a new payment record into the database.      |
| GET    | /payments           | Retrieve all payments. Returns a list of all recorded payments.            |
| GET    | /payments/{id}      | Retrieve a specific payment record by its ID.                              |
| PUT    | /payments/{id}      | Update an existing payment record by ID.                                   |
| DELETE | /payments/{id}      | Delete a payment record by its ID.                                         |

Sprint 4: Backend API Documentation
NAME	UFID
Kopparla Varshini	22060396
Karthik Karnam      37476457

Package Booking API

    Create Payment
        Endpoint: POST /payments
        Description: Adds a new payment record linked to an existing booking.
        Authentication: Requires JWT token.
        Request Headers:
            Authorization: Bearer <token>
            Content-Type: application/json
        Request Body Example:
            {
            "booking_id": 25,
            "amount": 750.00,
            "payment_method": "Credit Card",
            "payment_status": "Paid"
            }
        Response Example (Success - 201 Created):
            {
            "id": 12,
            "booking_id": 25,
            "amount": 750.00,
            "payment_method": "Credit Card",
            "payment_status": "Paid"
            }
        Error Responses:
            400 Bad Request → Missing or invalid fields.
            404 Not Found → Booking ID does not exist.
            500 Internal Server Error → Database insertion error.
        Edge Cases:
            Missing Fields (400 Bad Request)
            Invalid Booking Reference (404 Not Found)
            Payment Already Exists (500 Internal Server Error)

    Get All Payments
        Endpoint: GET /payments
        Description: Retrieves all payments in the system.
        Authentication: Requires JWT token.
        Response Example (Success - 200 OK):
            [
            {
                "id": 12,
                "booking_id": 25,
                "amount": 750.00,
                "payment_method": "Credit Card",
                "payment_status": "Paid"
            }
            ]

    Get Payment by ID
        Endpoint: GET /payments/{id}
        Description: Retrieves a specific payment record by its ID.
        Authentication: Requires JWT token.
        Response Example (Success - 200 OK):
            {
            "id": 12,
            "booking_id": 25,
            "amount": 750.00,
            "payment_method": "Credit Card",
            "payment_status": "Paid"
            }
        Error Responses:
            400 Bad Request → Invalid payment ID format.
            404 Not Found → Payment record not found.

    Update Payment
        Endpoint: PUT /payments/{id}
        Description: Updates fields of an existing payment record.
        Authentication: Requires JWT token.
        Request Body Example:
            {
            "amount": 800.00,
            "payment_method": "UPI",
            "payment_status": "Paid"
            }
        Response Example (Success - 200 OK):
            {
            "message": "Payment updated successfully"
            }
        Error Responses:
            400 Bad Request → Invalid request payload or ID.
            500 Internal Server Error → Error during update.

    Delete Payment
        Endpoint: DELETE /payments/{id}
        Description: Deletes a payment record by ID.
        Authentication: Requires JWT token.
        Response Example (Success - 200 OK):
            {
            "message": "Payment deleted successfully"
            }
        Error Responses:
            400 Bad Request → Invalid ID format.
            404 Not Found → Payment record not found.
            500 Internal Server Error → Deletion failure.