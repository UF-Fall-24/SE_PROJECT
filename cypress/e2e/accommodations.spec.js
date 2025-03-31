// cypress/integration/accommodations.spec.js

describe('Accommodations Navigation Test', () => {
  beforeEach(() => {
    // Visit the base URL
    cy.visit('/');
    // Simulate a logged-in user by setting a dummy token in localStorage
    cy.window().then((win) => {
      win.localStorage.setItem('token', 'dummy-token');
    });
    // Navigate to the dashboard where the HotelDropdown component is rendered
    cy.visit('/dashboard');
  });

  it('should display available accommodations after clicking on Accommodations', () => {
    // Find and click on the Accommodations button or link.
    // (Assuming the HotelDropdown renders an element with text "Accommodations")
    cy.contains('Accommodations').click();

    // Option 1: If clicking updates the URL, verify that the URL includes "/accommodations"
    cy.url().should('include', '/accommodations');

    // Option 2: Alternatively, check for text or an element that signifies the available accommodations list.
    // (For example, a heading or a container with the text "Available Accommodations")
    cy.contains('Available Accommodations').should('be.visible');
  });
});
