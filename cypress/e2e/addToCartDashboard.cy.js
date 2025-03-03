describe('Dashboard and Add to Cart Functionality', () => {
  beforeEach(() => {
      // Visit login page first
      cy.visit('http://localhost:3000/login');

      // Enter valid login credentials
      cy.get('input[name="email"]').type('testuser@example.com');
      cy.get('input[name="password"]').type('password123');
      cy.get('button[type="submit"]').click();

      // Ensure redirection to the dashboard
      cy.url().should('include', '/dashboard');

      // Ensure the dashboard loads
      cy.get('h2').contains('Travel Packages', { timeout: 10000 }).should('be.visible');
  });

  it('should load the dashboard and display available packages', () => {
      // Verify at least one package is displayed
      cy.get('.package-item', { timeout: 10000 }).should('have.length.at.least', 1);
  });

  it('should add a package to the cart and verify it is added', () => {
      // Add the first package to the cart
      cy.get('.package-item').first().within(() => {
          cy.get('button').contains('Add to Cart').click();
      });

      // Verify cart count updates
      cy.get('.fa-shopping-cart').click(); // Open the cart dropdown
      cy.get('.cart-count').should('contain', '1');

      // Verify package is listed in the cart dropdown
      cy.get('.cart-item').should('have.length.at.least', 1);
  });
});
