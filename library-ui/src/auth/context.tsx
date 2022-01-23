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

function Consumer({ children }: ConsumerProps<Provision>) {
  return (
    <context.Consumer>
      {(s) => {
        return s == null ? null : children(s);
      }}
    </context.Consumer>
  );
}
