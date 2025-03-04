import { render, screen } from "@testing-library/react";
import App from "../App";

// Basic Unit Test for Rendering Home Page without MemoryRouter
describe("Basic Rendering Test", () => {
  test("renders Home page successfully", () => {
    render(<App />);
    expect(screen.getByText(/welcome/i)).toBeInTheDocument();
  });
});