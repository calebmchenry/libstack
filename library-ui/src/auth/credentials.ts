import { hasProperty } from "../utils/fns";

export const Credentials = {
  validate,
};

function validate(value: unknown): Promise<auth.Credentials> {
  if (!hasProperty(value, "email"))
    return Promise.reject(new TypeError('Expected values to contain "email"'));
  if (!hasProperty(value, "password"))
    return Promise.reject(
      new TypeError('Expected values to contain "password"')
    );
  if (typeof value.email !== "string")
    return Promise.reject(
      new TypeError(
        `Expected email to be a string but received "${value.email}"`
      )
    );
  if (typeof value.password !== "string")
    return Promise.reject(
      new TypeError(
        `Expected password to be a string but received "${value.password}"`
      )
    );
  const emailRegex =
    /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  if (!emailRegex.test(value.email))
    return Promise.reject(new TypeError("Expected email to be a valid email"));
  const cred: auth.Credentials = {
    email: value.email,
    password: value.password,
  };
  return Promise.resolve(cred);
}
