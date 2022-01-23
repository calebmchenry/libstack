import React, {
  createContext,
  useCallback,
  useState,
  ConsumerProps,
  ProviderProps,
} from "react";
import { client } from "./client";
import { Token } from "./token";
import { hasProperty } from "../utils/fns";

export { Provider, Consumer };
export type Provision = State & {
  login: (values: unknown) => Promise<void>;
  logout: () => Promise<void>;
};

type State = {
  token?: string;
  username?: string;
  loginErr?: Error;
  loggingIn: boolean;
};

const context = createContext<Provision | undefined>(undefined);
const initialState: State = { loggingIn: false };
function Provider({ children }: Pick<ProviderProps<Provision>, "children">) {
  const [state, setState] = useState<State>(initialState);

  const login = useCallback(
    (values: unknown): Promise<void> => {
      // TODO(mchenryc): extract all non react logic out
      if (!hasProperty(values, "email"))
        return Promise.reject(
          new TypeError('Expected values to contain "email"')
        );
      if (!hasProperty(values, "password"))
        return Promise.reject(
          new TypeError('Expected values to contain "password"')
        );
      if (typeof values.email !== "string")
        return Promise.reject(
          new TypeError(
            `Expected email to be a string but received "${values.email}"`
          )
        );
      if (typeof values.password !== "string")
        return Promise.reject(
          new TypeError(
            `Expected password to be a string but received "${values.password}"`
          )
        );
      const emailRegex =
        /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      if (!emailRegex.test(values.email))
        return Promise.reject(
          new TypeError("Expected email to be a valid email")
        );

      setState((s) => ({ ...s, loggingIn: true }));
      return client
        .login({ email: values.email, password: values.password })
        .then((token) => {
          setState((s) => ({ ...s, token, loggingIn: false }));
        })
        .catch((err) => {
          setState((s) => ({
            ...s,
            loginErr: err ?? new Error("Unexpected login failure"),
            loggingIn: false,
          }));
        });
    },
    [setState]
  );

  const logout = useCallback(() => {
    if (state.token == null)
      return Promise.reject(new Error("Already logged out"));
    return fetch("http://localhost:8000/api/v1/logout", {
      method: "POST",
      headers: { Authorization: `Bearer ${state.token}` },
    })
      .then((r) => r.json())
      .then(Token.decode)
      .then((token) => {
        setState((s) => ({ ...s, token }));
      })
      .catch((err) => {
        setState((s) => ({
          ...s,
          loginErr: err ?? new Error("Unexpected login failure"),
        }));
      });
  }, [state.token, setState]);

  return (
    <context.Provider value={{ ...state, login, logout }}>
      {children}
    </context.Provider>
  );
}

function Consumer({ children }: ConsumerProps<Provision>) {
  return (
    <context.Consumer>
      {(s) => {
        return s == null ? null : children(s);
      }}
    </context.Consumer>
  );
}
