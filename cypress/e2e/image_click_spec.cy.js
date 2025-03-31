describe('Image Click Tests After Login', () => {
  beforeEach(() => {
    // Simulate a logged-in user by setting a token in localStorage.
    cy.window().then((win) => {
      win.localStorage.setItem('token', 'your-test-token');
    });
    // Visit the home page or a dashboard that shows the images.
    cy.visit('/');
  });

  it('should navigate to the packages page when a package image is clicked', () => {
    // Find and click the package image.
    cy.get('[data-test="package-image"]').first().click();

    // Assert that the URL includes "/packages" or that a key element is visible.
    cy.url().should('include', '/packages');
    // Optionally, check for a key detail element:
    cy.get('.package-detail').should('be.visible');
  });

  it('should navigate to the accommodations page when an accommodation image is clicked', () => {
    // Find and click the accommodation image.
    cy.get('[data-test="accommodation-image"]').first().click();

    // Assert that the URL includes "/accommodations".
    cy.url().should('include', '/accommodations');
    // Optionally, check that an accommodations detail element is visible:
    cy.get('.accommodation-detail').should('be.visible');
  });

  it('should navigate to the hotels page when a hotel image is clicked', () => {
    // First, verify the hotel dropdown is visible (it appears only when logged in).
    cy.get('.hotel-dropdown').should('be.visible');

    // Then click on a hotel image within the dropdown.
    cy.get('[data-test="hotel-image"]').first().click();

    // Assert that the URL includes "/hotels".
    cy.url().should('include', '/hotels');
    // Optionally, check that a hotel-specific detail element is visible.
    cy.get('.hotel-detail').should('be.visible');
  });
});
