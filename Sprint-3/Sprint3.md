**Detailed explanation of the completed work implemented in the Frontend**

In this sprint, we implemented several new user flows and enhanced our frontend functionality for Book Ease. Our primary focus was to create a seamless user experience for travel booking, starting from login, navigating through the dashboard, selecting travel packages, booking accommodations, and accessing hotel details. We also implemented comprehensive automated tests using Cypress to verify all these flows.

New Features
1. Dashboard and Navigation
Dashboard Images:
After a successful login, the dashboard displays three interactive images for:

Packages: Navigates to the available packages list.

Accommodations: Navigates to the accommodations page.

Hotels: Directly shows the available hotels list.

Navigation Flow:

Packages Flow:

Clicking the Packages image navigates to /packages where all available packages are listed.

Selecting a package and clicking Book Now directs the user to a package details page (e.g., /package-details/1).

On this details page, users can choose between proceeding with Payment or clicking Book Accommodation.

Payment: Navigates to the payment page (functionality for payment will be completed in the next sprint).

Book Accommodation: Navigates to the available hotels list (/accommodation-hotels).

Direct Hotels Flow:

Clicking the Hotels image on the dashboard navigates immediately to the hotels list.

2. Authentication Pages
Login Page:

Styled with a full-screen travel-themed background image.

A centered login card with a live clock, email/password inputs, and a login button.

Register Page:

Similar visual style as the login page.

Allows new users to sign up with a username, email, and password.

3. Dynamic and Real-Time Elements
Live Clock:

A continuously updating clock in the hero section gives the user a sense of real-time interactivity.

Conditional Rendering:

Different components are displayed based on the user's authentication status (e.g., dashboard images and navigation flows for logged-in users, or login/register prompts for new users).

Automated Test Cases
We have implemented a comprehensive suite of Cypress tests to verify that each new feature works as expected. Below is a detailed explanation of each test suite:

1. Payment and Accommodation Flow Test
Objective:
Verify that the user can navigate from the packages page to the package details page and then choose between the payment and accommodation flows.

Test Steps:

Login: The test visits /login and logs in using valid credentials.

Navigate to Packages: It clicks on the Packages image/link, verifying that the URL contains /packages.

Book Now: The test clicks the first Book Now button, ensuring that the URL updates to include /package-details/1.

Payment Option: One test clicks the Payment button and checks that the URL includes /payment.

Accommodation Option: Another test clicks the Book Accommodation button and verifies that the URL includes /accommodation-hotels and that the page displays expected content (such as "Available Hotels").

Outcome:
These tests ensure that both the payment and accommodation booking options work correctly after a package is selected.

2. Frontend Navigation Tests
Objective:
Ensure that the main navigation links work properly.

Test Steps:

Contact Page Navigation:
The test clicks the link to /contact and verifies the URL and the presence of "Contact Us" content.

About Us Page Navigation:
It clicks the link to /about and confirms that "About Us" content is visible.

Register & Login Navigation:
The tests navigate to the home page and then click on the register and login buttons, verifying that the URLs update to /register and /login, respectively.

User Registration and Login Flow:
A combined test simulates user registration, then logs in with the newly registered credentials, and confirms successful navigation to the dashboard.

Outcome:
These tests validate that all key navigation paths are operational and that pages load the correct content.

3. Hotel Options Selection Test
Objective:
Validate the end-to-end flow for booking accommodation.

Test Steps:

Login and Navigate:
Log in, navigate to the packages page, click the first Book Now button, and then click Book Accommodation to display the hotels list.

Inspect Hotels List:
Verify that the hotels table is visible and contains at least one hotel entry.

Select a Hotel:
Within the first hotel row, the test clicks the Select button.
It verifies that the URL now matches a dynamic pattern (e.g., /hotel-details/9) and that key information such as "Price per Night" is visible on the hotel details page.

Outcome:
Ensures that the hotel selection process and subsequent navigation to hotel details work correctly.

4. Dashboard Images Navigation Test
Objective:
Ensure that clicking on each dashboard image navigates to the correct page.

Test Steps:

Login: Log in and load the dashboard.

Packages Image:
Click the Packages image and verify that the URL includes /packages.

Accommodations Image:
Return to the dashboard and click the Accommodations image to verify navigation to /accommodations.

Hotels Image:
Return to the dashboard again and click the Hotels image to verify that the URL includes /hotels.

Outcome:
Validates that each image on the dashboard is properly linked to its corresponding page.

5. Logout Functionality Test
Objective:
Verify that the logout process works as intended.

Test Steps:

Login: The test logs in using valid credentials.

Visibility: Verify that the logout button is visible on the dashboard.

Logout Action: Click the logout button and confirm that the user is redirected to the /login page.
Also, check that the authentication token is removed from localStorage.

Outcome:
Confirms that the logout function successfully clears the session and navigates the user back to the login page.

6. Package Details Navigation Test
Objective:
Confirm that selecting a package from the packages list correctly navigates to the package details page.

Test Steps:

Login and Navigation: Log in and navigate to the packages page.

Book Now: Click the first Book Now button.

Verification: Assert that the URL includes /package-details/1, indicating that the package details page is shown.

Outcome:
Ensures that users can view detailed information for the selected package.

7. Packages Navigation Test
Objective:
Ensure that the user can navigate to the packages section directly from the home page.

Test Steps:

Login: Log in and load the home/dashboard.

Navigation: Click on the Packages link.

Verification: Check that the URL includes /packages and that the page displays the "Available Packages" heading.

Outcome:
Verifies that navigation to the packages section is accessible and correctly displays its content.


**Detailed explanation of the completed work implemented in the backend**

Package Booking (package_bookings.go)
This file manages the creation, updating, and deletion of package booking records in the package_bookings table.

CreatePackageBooking
    Responsibility: Handles the creation of a new package booking.
    Workflow:
        Reads the JSON payload from the request body and decodes it into a PackageBooking model instance.
        Performs necessary validations (e.g., required fields).
        Calls the model’s Create() method to insert the record into the database.
        On success, returns the created booking with a status code 201 Created.
        Handles bad requests and database errors with appropriate status codes like 400 or 500.

UpdatePackageBooking
    Responsibility: Updates an existing package booking by booking_id.
    Workflow:
        Extracts the booking_id from the URL parameters.
        Decodes the update data from the request body into the PackageBooking model.
        Sets the booking ID and invokes the Update() method to apply changes.
        On success, returns a confirmation message with 200 OK.
        If the booking ID is invalid or update fails, returns errors such as 400 Bad Request, 404 Not Found, or 500 Internal Server Error.

DeletePackageBooking
    Responsibility: Deletes a package booking based on its booking_id.
    Workflow:
        Extracts booking_id from the URL.
        Calls the model’s Delete() method to remove the record.
        Returns a success message in JSON if deletion is successful.
        If the booking_id is invalid or deletion fails, responds with appropriate error codes like 404 or 500.

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

### Package Booking Endpoints

| Method | Endpoint                     | Description                                                                 |
| ------ | ---------------------------- | --------------------------------------------------------------------------  |
| POST   | /package_bookings            | Create a new package booking. Inserts a new record into the database.       |
| PUT    | /package_bookings/{id}       | Update an existing package booking by booking ID.                           |
| DELETE | /package_bookings/{id}       | Delete a package booking record by booking ID.                              |


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

Package Booking API

    Create Package Booking
        Endpoint: POST /package_bookings
        Description: Adds a new package booking record.
        Authentication: Requires JWT token.
        Request Headers:
            Authorization: Bearer <token>
            Content-Type: application/json
        Request Body Example:
            {
            "booking_id": 25,
            "package_id": 3
            }
        Response Example (Success - 201 Created):
            {
            "id": 7,
            "booking_id": 25,
            "package_id": 3
            }
        Error Responses:
            400 Bad Request → Missing or invalid fields.
            404 Not Found → Booking ID or Package ID does not exist.    
            500 Internal Server Error → Database insertion error.
        Edge Cases:
            Missing Fields (400 Bad Request)
            Invalid Foreign Keys (404 Not Found)
            Duplicate Booking Handling (500 Internal Server Error)

    Update Package Booking
        Endpoint: PUT /package_bookings/{id}
        Description: Updates an existing package booking by ID.
        Authentication: Requires JWT token.
        Request Body Example:
            {
            "booking_id": 25,
            "package_id": 5
            }
        Response Example (Success - 200 OK):
            {
            "message": "Package booking updated successfully"
            }
        Error Responses:
            400 Bad Request → Invalid booking ID or malformed request.
            404 Not Found → Package booking not found.
            500 Internal Server Error → Error during update operation.

    Delete Package Booking
        Endpoint: DELETE /package_bookings/{id}
        Description: Deletes a package booking record by its ID.
        Authentication: Requires JWT token.
        Response Example (Success - 200 OK):
            {
            "message": "Package booking deleted successfully"
            }
        Error Responses:
            400 Bad Request → Invalid booking ID format.
            404 Not Found → Package booking record not found.
            500 Internal Server Error → Error during deletion process.

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