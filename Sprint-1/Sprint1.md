Sprint 1 

**User Stories**

**Frontend User Stories**

Frontend Development Environment Setup

 - Design and implement a properly configured frontend development environment to efficiently build, test, and deploy the Book-Ease web application.
 - Install and configure Node.js to manage dependencies.
 - Set up VS Code as the primary development environment.
 - Initialize a React.js project using Create React App or Vite.
 - Configure Git for version control and repository management.
 - Ensure that the environment supports hot reloading for efficient development.

1.User Registration
-Develop and implement a user registration feature to allow users to create accounts.
-Implement a registration form with fields for username, email, and password.

2.Validate inputs:
-Email must be in a correct format.
-Password must meet security standards (minimum length, special characters, etc.).
-Duplicate emails should not be allowed.

3.On successful registration:
-Automatically log in the user or prompt for email verification.
-Display clear error messages for invalid or incomplete inputs.
-User Login
-Develop and implement an authentication system for user login.
-Create a login form with fields for email and password.
-Implement form validation to ensure both fields are required.
-Authenticate users by verifying credentials against the backend API.
-Show error messages if the login fails (e.g., incorrect password, unregistered email).
-Store JWT token securely upon successful login for session management.
-Implement a logout button that clears the session and token.

4.Performance Optimization
-Enhance frontend performance to improve user experience.
-Implement React.lazy and Suspense for lazy loading of components.
-Display a loading indicator when switching between pages.
-Reduce the initial bundle size to improve loading speed.
-Ensure all routes function correctly without breaking any features.

**What Issues Your Team Planned to Address?**
Our team planned to address the following key frontend issues:
Proper setup of React development environment.
Implementation of secure and user-friendly authentication mechanisms.
Ensuring a smooth registration and login experience.
Improving website performance using lazy loading and optimized bundling.

**Which Ones Were Successfully Completed?**
 
We successfully completed the following tasks:
-Frontend Environment Setup:
-Node.js, VS Code, React.js, and Git were successfully set up.

User Registration:
-Registration form was implemented and validated.
-Email format and password security validation were enforced.
-Error messages for invalid inputs were displayed.
-Users were either logged in automatically or prompted for verification.

User Login:
-Login form was created and properly validates inputs.
-API authentication works, and JWT tokens are stored securely.
-Users can log out and clear their session.

Performance Optimization:
-Components are now lazy-loaded using React.lazy and Suspense.
-A loading indicator appears while navigating pages.
-The bundle size was reduced for better performance.

**Which Ones Didn’t and Why?**

-Error handling and logging mechanisms may still need enhancements to capture invalid login attempts properly.
-Reason for Incomplete Tasks:
-Some issues required further debugging and testing.
-More detailed logging and error handling mechanisms are needed for failed login attempts.
-Some UI elements require improved accessibility and responsiveness.
-Further debugging is required to ensure seamless authentication and performance optimization.
-Need to conduct more user testing to identify potential UI/UX issues.


**Backend User stories**

1. Customer Registration
- Design and implement a database schema for storing customer registration details.
- Develop an API to accept and store customer data in the database.
- Implement data validation to ensure accuracy and completeness.
- Enforce password validation rules (e.g., length, complexity).

2. Customer Login
- Develop an API to authenticate customers by verifying their credentials.
- Implement JWT-based authentication to generate secure tokens upon successful login.
- Validate password security during login attempts.



**What issues your team planned to address?**

Our team planned to address the following key issues related to customer authentication and security:

  - Implementing a secure customer registration and login process.
  - Ensuring proper validation of customer data to maintain data integrity.
  - Implementing JWT-based authentication for session management.
  - Enhancing password security through validation mechanisms.
  - Designing an optimized database schema to efficiently store customer data.

**Which ones were successfully completed?**

We successfully completed the following tasks:


**Backend**

   Customer Registration:

  - Database schema for customer registration was created.
  - API for customer registration was implemented and successfully stores data in the database.
  - Customer data validation (format, required fields) was enforced.
  - Password validation rules (e.g., minimum length, special characters) were implemented.
    
  Customer Login:

  - API for customer login was implemented and verifies customer credentials.
  - JWT token is generated successfully upon login.
  - Password validation was enforced during the login process.

**Which ones didn’t and why?**
  - Error handling and logging mechanisms may still need enhancements to capture invalid login attempts properly.
  Reason for Incomplete Tasks:
  - Some issues required further debugging and testing.


Demo URL: https://github.com/UF-Fall-24/SE_PROJECT/blob/main/Sprint-1/Front_end_and_back_end.mp4
