describe('Hotel Options Selection Test', () => {
  beforeEach(() => {
    // 1. Log in by visiting the login page and submitting credentials
    cy.visit('http://localhost:3000/login');
    cy.get('input[name="email"]').should('be.visible').type('uniqueemail@gmail.com');
    cy.get('input[name="password"]').should('be.visible').type('Password@123');
    cy.get('button[type="submit"]').click();
    //cy.url().should('include', '/dashboard');

    // 2. Navigate to the Packages page
    cy.contains('Packages').should('be.visible').click();
    cy.url().should('include', '/packages');

    // 3. Click the first "Book Now" button to navigate to the package details page
    cy.contains('Book Now').first().should('be.visible').click();
    cy.url().should('include', '/package-details/');

    // 4. Click "Book Accommodation" to go to the hotels listing page
    cy.contains('Book Accommodation').should('be.visible').click();
    cy.url().should('include', '/accommodation-hotels');
    cy.contains('Available Hotels').should('be.visible');
  });

  it('should display hotel options, navigate to hotel details, and show Price per Night', () => {
    // 5. Verify that the hotels table is visible
    cy.get('table').should('be.visible');

    // 6. Verify there is at least one hotel listed (excluding header row)
    cy.get('tbody tr').its('length').should('be.gt', 0);

    // 7. Optionally, inspect the first hotel row to ensure content is present
    cy.get('tbody tr').first().within(() => {
      // Verify hotel name is not empty
      cy.get('td').eq(0).should('not.be.empty');
      // Check that the "Select" button is present
      cy.get('button').contains('Select').should('be.visible');
    });

    // 8. Click the "Select" button for the first hotel option
    cy.get('tbody tr').first().within(() => {
      cy.get('button').contains('Select').click();
    });

    // 9. Verify that the URL navigates to a hotel details page with a dynamic ID (e.g., /hotel-details/9)
    cy.url().should('match', /\/hotel-details\/\d+$/);

    // 10. Verify that the hotel details page shows "Price per Night" information
    cy.contains('Price per Night').should('be.visible');
  });
});
