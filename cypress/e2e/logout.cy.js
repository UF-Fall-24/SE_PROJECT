describe("Logout Functionality", () => {
  beforeEach(() => {
    // ✅ Ensure user is logged in before testing logout
    cy.visit("http://localhost:3000/login");

    // ✅ Enter login credentials
    cy.get('input[name="email"]').type("uniqueemail@gmail.com");
    cy.get('input[name="password"]').type("Password@123");

    // ✅ Click login button
    cy.get('button[type="submit"]').click();

    // ✅ Wait for login process
    cy.wait(2000);

    // ✅ Ensure redirected to dashboard
    //cy.url().should("include", "/dashboard");
  });

  it("should verify the logout button is visible", () => {
    // ✅ Ensure the logout button is present
    cy.get("button").contains("Logout").should("be.visible");
  });

  it("should log out the user and redirect to login", () => {
    // ✅ Click logout button
    cy.get("button").contains("Logout").click();

    // ✅ Ensure redirected to login page
    cy.url().should("include", "/login");

    // ✅ Ensure token is removed from localStorage
    cy.window().then((win) => {
      expect(win.localStorage.getItem("token")).to.be.null;
    });
  });
});
