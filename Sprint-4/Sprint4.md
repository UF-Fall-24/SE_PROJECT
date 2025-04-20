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

