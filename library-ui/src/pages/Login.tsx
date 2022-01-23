import React from "react";
import { useForm } from "react-hook-form";
import { auth } from "../auth";

export function LoginPage(): JSX.Element {
  return (
    <auth.Provider>
      <auth.Consumer>
        {(s) => <LoginForm login={s.login} loggingIn={s.loggingIn} />}
      </auth.Consumer>
    </auth.Provider>
  );
}

function LoginForm({
  login,
  loggingIn,
}: Pick<auth.Provision, "login" | "loggingIn">): JSX.Element {
  const { register, handleSubmit, formState } =
    useForm<{ email: string; password: string }>();
  return (
    <form onSubmit={handleSubmit(login)}>
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
      <button type="submit" disabled={loggingIn}>
        Login
      </button>
    </form>
  );
}
