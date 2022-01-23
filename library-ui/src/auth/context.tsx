import React, {
  createContext,
  useCallback,
  useState,
  ConsumerProps,
  ProviderProps,
} from "react";
import { client } from "./client";
import { auth } from ".";
import { Credentials } from "./credentials";

export { Provider, Consumer };

const context = createContext<auth.Provision | undefined>(undefined);
const initialState: auth.State = { loggingIn: false };
function Provider({
  children,
}: Pick<ProviderProps<auth.Provision>, "children">) {
  const [state, setState] = useState<auth.State>(initialState);

  const login = useCallback(
    (values: unknown): Promise<void> => {
      return Credentials.validate(values)
        .then((creds) => {
          setState((s) => ({ ...s, loggingIn: true }));
          return client.login(creds);
        })
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
    return auth.client
      .logout(state.token)
      .then(() => {
        setState((s) => ({ ...s, token: undefined }));
      })
      .catch((err) => {
        setState((s) => ({
          ...s,
          loginErr: err ?? new Error("Unexpected logout failure"),
        }));
      });
  }, [state.token, setState]);

  return (
    <context.Provider value={{ ...state, login, logout }}>
      {children}
    </context.Provider>
  );
}

function Consumer({ children }: ConsumerProps<auth.Provision>) {
  return (
    <context.Consumer>
      {(s) => {
        return s == null ? null : children(s);
      }}
    </context.Consumer>
  );
}
