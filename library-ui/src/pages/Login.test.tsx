import React from "react";
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { LoginPage } from "./Login";
import { client } from "../auth/client";

const emailRequiredRegex = /^email is required$/i;
const passwordRequiredRegex = /^password is required$/i;
const emailLabel = /^email$/i;
const passwordLabel = /^password$/i;
const loginButtonText = /^login$/i;

describe("Login", () => {
  it("renders blank form", async () => {
    render(<LoginPage></LoginPage>);

    const emailInput = screen.getByLabelText(emailLabel);
    const passwordInput = screen.getByLabelText(passwordLabel);
    const loginButton = screen.getByRole("button", { name: loginButtonText });
    expect(emailInput).toHaveValue("");
    expect(passwordInput).toHaveValue("");
    expect(loginButton).toBeInTheDocument();

    userEvent.click(loginButton);

    await waitFor(() => {
      expect(screen.getByText(emailRequiredRegex)).toBeInTheDocument();
    });
    expect(screen.getByText(passwordRequiredRegex)).toBeInTheDocument();

    const email = "foo@bar.com";
    userEvent.type(emailInput, email);
    expect(emailInput).toHaveValue(email);

    await waitFor(() => {
      expect(screen.queryByText(emailRequiredRegex)).not.toBeInTheDocument();
    });

    const password = "password123";
    userEvent.type(passwordInput, password);
    expect(passwordInput).toHaveValue(password);

    await waitFor(() => {
      expect(screen.queryByText(passwordRequiredRegex)).not.toBeInTheDocument();
    });

    jest.spyOn(client, "login").mockResolvedValue("token123");
    userEvent.click(loginButton);

    await waitFor(() => {
      expect(loginButton).toBeDisabled();
    });

    await waitFor(() => {
      expect(client.login).toBeCalledTimes(1);
    });

    expect(loginButton).not.toBeDisabled();
  });
});
