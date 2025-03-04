describe("Login Test", () => {
    it("should log in successfully and redirect to dashboard with available packages", () => {
      cy.visit("http://localhost:3000/login");
  
      // Enter email and password
      cy.get('input[name="email"]').type("uniqueemail@gmail.com");
      cy.get('input[name="password"]').type("Password@123");
  
      // Click login button
      cy.get('button[type="submit"]').click();
  
      // Wait for login process
      cy.wait(2000);
  
      // Ensure token is stored
      cy.window().then((win) => {
        expect(win.localStorage.getItem("token")).to.exist;
      });
  
      // Ensure redirected to dashboard
      cy.url().should("include", "/dashboard");
  
      // ✅ Verify "Available Packages" is displayed
      cy.contains("Available Packages").should("be.visible");
  
      // ✅ Check if at least one package is displayed
      cy.get("div").should("contain.text", "Available Packages");
    });

        it("should log in successfully and verify cart logo exists on the dashboard", () => {
          cy.visit("http://localhost:3000/login");
      
          // Enter email and password
          cy.get('input[name="email"]').type("uniqueemail@gmail.com");
          cy.get('input[name="password"]').type("Password@123");
      
          // Click login button
          cy.get('button[type="submit"]').click();
      
          // Wait for login process
          cy.wait(2000);
      
          // Ensure token is stored
          cy.window().then((win) => {
            expect(win.localStorage.getItem("token")).to.exist;
          });
      
          // Ensure redirected to dashboard
          cy.url().should("include", "/dashboard");
      
          // ✅ Verify "Available Packages" is displayed
          cy.contains("Available Packages").should("be.visible");
      
          // ✅ Check if at least one package is displayed
          cy.get("div").should("contain.text", "Available Packages");
    
          // ✅ Verify the cart logo (Shopping Cart icon) exists
          cy.get("svg").should("be.visible"); // Checks if any SVG exists (assumes FaShoppingCart is an SVG)
    
          // Alternative (if you use a specific class or data-testid)
          // cy.get('[data-testid="cart-icon"]').should("be.visible");
        });
  });
  describe("Search Transport Functionality", () => {
    it("should log in successfully and verify cart logo & Search Transport exists", () => {
      cy.visit("http://localhost:3000/login");
  
      // Enter email and password
      cy.get('input[name="email"]').type("uniqueemail@gmail.com");
      cy.get('input[name="password"]').type("Password@123");
  
      // Click login button
      cy.get('button[type="submit"]').click();
  
      // Wait for login process
      cy.wait(2000);
  
      // Ensure token is stored
      cy.window().then((win) => {
        expect(win.localStorage.getItem("token")).to.exist;
      });
  
      // Ensure redirected to dashboard
      cy.url().should("include", "/dashboard");
  
      // ✅ Verify "Available Packages" is displayed
      cy.contains("Available Packages").should("be.visible");
  
      // ✅ Check if at least one package is displayed
      cy.get("div").should("contain.text", "Available Packages");

      // ✅ Verify the cart logo (Shopping Cart icon) exists
      cy.get("svg").should("be.visible"); // Checks if any SVG exists (assumes FaShoppingCart is an SVG)

      // ✅ Verify "Search Transport" section exists
      cy.contains("Search Transport").should("be.visible");

      // Alternative if "Search Transport" is inside a div
      // cy.get('div').contains("Search Transport").should("be.visible");
    });
});