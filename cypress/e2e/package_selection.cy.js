describe('Package Details Navigation Test', () => {
  beforeEach(() => {
    // Visit the login page
    cy.visit('http://localhost:3000/login');

    // Log in with valid credentials
    cy.get('input[name="email"]').type('uniqueemail@gmail.com');
    cy.get('input[name="password"]').type('Password@123');
    cy.get('button[type="submit"]').click();

    // Confirm redirection to the dashboard
    //cy.url().should('include', '/dashboard');

    // Navigate to the Packages page
    cy.contains('Packages').click();
    cy.url().should('include', '/packages');
  });

  it('should navigate to package details page when clicking the first Book Now button', () => {
    // Click the first "Book Now" button; adjust selector if needed.
    cy.contains('Book Now').first().click();

    // Verify the URL changes to include '/package-details/1'
    cy.url().should('include', '/package-details/1');

    // Optionally, check for specific content on the package details page
    // cy.contains('Package Details').should('be.visible');
  });
});
