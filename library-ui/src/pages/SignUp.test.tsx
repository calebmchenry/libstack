import React from "react";
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { client } from "../auth/client";
import { SignUp } from "./SignUp";
import { auth } from "../auth";

const emailRequiredRegex = /^email is required$/i;
const passwordRequiredRegex = /^password is required$/i;
const emailLabel = /^email$/i;
const passwordLabel = /^password$/i;
const signUpButtonText = /^sign up$/i;

test("SignUp", async () => {
  render(
    <auth.Provider>
      <SignUp></SignUp>
    </auth.Provider>
  );

  const emailInput = screen.getByLabelText(emailLabel);
  const passwordInput = screen.getByLabelText(passwordLabel);
  const signUpButton = screen.getByRole("button", { name: signUpButtonText });
  expect(emailInput).toHaveValue("");
  expect(passwordInput).toHaveValue("");
  expect(signUpButton).toBeInTheDocument();

  userEvent.click(signUpButton);

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

  jest.spyOn(client, "signUp").mockResolvedValue("token123");
  userEvent.click(signUpButton);

  await waitFor(() => {
    expect(signUpButton).toBeDisabled();
  });

  await waitFor(() => {
    expect(client.signUp).toBeCalledTimes(1);
  });

  expect(signUpButton).not.toBeDisabled();
});
