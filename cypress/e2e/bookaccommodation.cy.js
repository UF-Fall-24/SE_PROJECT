describe('Payment and Accommodation Flow Test', () => {
  beforeEach(() => {
    // 1. Visit the login page
    cy.visit('http://localhost:3000/login');

    // 2. Log in with valid credentials
    cy.get('input[name="email"]').should('be.visible').type('uniqueemail@gmail.com');
    cy.get('input[name="password"]').should('be.visible').type('Password@123');
    cy.get('button[type="submit"]').click();

    // 3. Verify redirection to the dashboard
    //cy.url().should('include', '/dashboard');

    // 4. Navigate to the Packages page by clicking on the Packages image or link
    // You can adjust the selector if your element uses an alt text or data attribute:
    // cy.get('img[alt="Packages"]').click(); or cy.contains('Packages').click();
    cy.contains('Packages').should('be.visible').click();
    cy.url().should('include', '/packages');

    // 5. Click on the first "Book Now" button to navigate to package details page
    cy.contains('Book Now').first().should('be.visible').click();
    cy.url().should('include', '/package-details/1');
  });

  it('should navigate to the payment page when Payments button is clicked', () => {
    // Click the "Payments" button on the package details page
    cy.contains('Payment').should('be.visible').click();

    // Verify the URL changes to /payment
    cy.url().should('include', '/payment');

    // Optionally, verify expected content on the payment page
    //cy.contains('Payment').should('be.visible');
  });

  it('should navigate to the accommodation hotels page when Book Accommodation button is clicked', () => {
    // Click the "Book Accommodation" button on the package details page
    cy.contains('Book Accommodation').should('be.visible').click();

    // Verify the URL changes to /accommodation-hotels
    cy.url().should('include', '/accommodation-hotels');

    // Verify the page displays "Available hotels"
    //cy.contains('Available hotels').should('be.visible');
  });
});
