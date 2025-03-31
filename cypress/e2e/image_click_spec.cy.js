describe('Dashboard Images Navigation Test', () => {
  beforeEach(() => {
    // Log in to the application
    cy.visit('http://localhost:3000/login');
    cy.get('input[name="email"]').should('be.visible').type('uniqueemail@gmail.com');
    cy.get('input[name="password"]').should('be.visible').type('Password@123');
    cy.get('button[type="submit"]').click();
    // Confirm redirection to the dashboard after login
    //cy.url().should('include', '/dashboard');
  });

  it('should navigate to the correct pages when dashboard images are clicked', () => {
    // Click on the Packages image and verify navigation
    cy.get('img[alt="Packages"]')
      .should('be.visible')
      .click();
    cy.url().should('include', '/packages');
    
    // Return to dashboard for the next test
    cy.go('back');
    //cy.url().should('include', '/dashboard');

    // Click on the Accommodations image and verify navigation
    cy.get('img[alt="Accommodations"]')
      .should('be.visible')
      .click();
    cy.url().should('include', '/accommodations');
    
    // Return to dashboard for the next test
    cy.go('back');
    //cy.url().should('include', '/dashboard');

    // Click on the Hotels image and verify navigation
    cy.get('img[alt="Hotels"]')
      .should('be.visible')
      .click();
    cy.url().should('include', '/hotels');
  });
});
