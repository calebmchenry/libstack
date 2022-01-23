import { Token } from "./token";

export const client = {
  login,
};

function login({
  email,
  password,
}: {
  email: string;
  password: string;
}): Promise<Token> {
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
