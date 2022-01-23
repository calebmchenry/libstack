import { Token } from "./token";

export const client = {
  login,
  logout,
};

function login({
  email,
  password,
}: {
  email: string;
  password: string;
}): Promise<auth.Token> {
  // TODO(mchenryc): use env for url
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
