describe('Packages Navigation Test', () => {
  it('should navigate to /packages and display Available Packages', () => {
    // 1. Visit the login page
    cy.visit('http://localhost:3000/login');

    // 2. Perform login with valid credentials
    cy.get('input[name="email"]').type('uniqueemail@gmail.com');
    cy.get('input[name="password"]').type('Password@123');
    cy.get('button[type="submit"]').click();

    // 3. Ensure we land on the dashboard
    //cy.url().should('include', '/dashboard');

    // 4. Click the "Packages" link (or button/image).
    //    Adjust the selector to match how your "Packages" link is rendered:
    //    - If it's a link: cy.contains('Packages').click();
    //    - If it's an image with alt="Packages": cy.get('img[alt="Packages"]').click();
    cy.contains('Packages').click();

    // 5. Verify the URL includes /packages
    cy.url().should('include', '/packages');

    // 6. Verify the page shows the "Available Packages" heading
    cy.contains('Available Packages').should('be.visible');
  });
});
