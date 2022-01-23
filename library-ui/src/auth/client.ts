import { Token } from "./token";

export const client = {
  login,
  logout,
  signUp,
};

// TODO(mchenryc): use env for url

function login({ email, password }: auth.Credentials): Promise<auth.Token> {
  return fetch("http://localhost:8000/api/v1/login", {
    method: "POST",
    body: JSON.stringify({
      email: email,
      password: password,
    }),
  })
    .then((r) => r.json())
    .then(Token.decode);
}

function logout(token: string): Promise<void> {
  return fetch("http://localhost:8000/api/v1/logout", {
    method: "POST",
    headers: { Authorization: `Bearer ${token}` },
  }).then((r) => r.json());
}

function signUp({ email, password }: auth.Credentials): Promise<auth.Token> {
  return fetch("http://localhost:8000/api/v1/signup", {
    method: "POST",
    body: JSON.stringify({
      email: email,
      password: password,
    }),
  })
    .then((r) => r.json())
    .then(Token.decode);
}
