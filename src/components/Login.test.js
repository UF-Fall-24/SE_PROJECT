import { render, screen, fireEvent } from "@testing-library/react";
import LoginForm from "../components/LoginForm";

// Unit Test for Login Functionality
describe("Login Form Tests", () => {
  test("renders login form correctly", () => {
    render(<LoginForm />);
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /login/i })).toBeInTheDocument();
  });

  test("shows error when submitting empty form", async () => {
    render(<LoginForm />);
    fireEvent.click(screen.getByRole("button", { name: /login/i }));
    expect(await screen.findByText(/email is required/i)).toBeInTheDocument();
    expect(await screen.findByText(/password is required/i)).toBeInTheDocument();
  });

  test("logs in successfully with correct credentials", async () => {
    render(<LoginForm />);

    fireEvent.change(screen.getByLabelText(/email/i), {
      target: { value: "test@example.com" }
    });
    fireEvent.change(screen.getByLabelText(/password/i), {
      target: { value: "Password@123" }
    });
    fireEvent.click(screen.getByRole("button", { name: /login/i }));

    expect(await screen.findByText(/welcome back/i)).toBeInTheDocument();
  });
});