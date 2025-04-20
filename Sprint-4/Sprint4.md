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

*Documentation of API's*

### Payment Endpoints

| Method | Endpoint            | Description                                                                |
| ------ | --------------------| -------------------------------------------------------------------------- |
| POST   | /payments           | Create a new payment. Inserts a new payment record into the database.      |
| GET    | /payments           | Retrieve all payments. Returns a list of all recorded payments.            |
| GET    | /payments/{id}      | Retrieve a specific payment record by its ID.                              |
| PUT    | /payments/{id}      | Update an existing payment record by ID.                                   |
| DELETE | /payments/{id}      | Delete a payment record by its ID.                                         |

