import React from "react";
import { useForm } from "react-hook-form";
import { auth } from "../auth";

export function SignUp(): JSX.Element {
  return (
    <auth.Consumer>
      {(s) => <SignUpForm signUp={s.signUp} signingUp={s.signingUp} />}
    </auth.Consumer>
  );
}

// TODO(mchenryc): handle loginError
function SignUpForm({
  signUp,
  signingUp,
}: Pick<auth.Provision, "signUp" | "signingUp">): JSX.Element {
  const { register, handleSubmit, formState } =
    useForm<{ email: string; password: string }>();

  return (
    // TODO(mchenryc):  signUp.catch
    <form onSubmit={handleSubmit(signUp)}>
      <label htmlFor="login-email">Email</label>
      <input
        type="email"
        id="login-email"
        {...register("email", { required: "Email is required" })}
      />
      {formState.errors.email?.message && (
        <span>{formState.errors.email.message}</span>
      )}
      <label htmlFor="login-password">Password</label>
      <input
        type="password"
        id="login-password"
        {...register("password", { required: "Password is required" })}
      />
      {formState.errors.password?.message && (
        <span>{formState.errors.password.message}</span>
      )}
      <button type="submit" disabled={signingUp}>
        Sign Up
      </button>
    </form>
  );
}
