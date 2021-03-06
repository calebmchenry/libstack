export const Token = {
  decode(json: unknown): Promise<auth.Token> {
    if (typeof json !== "string")
      return Promise.reject(new Error(`Failed to decode token from "${json}"`));
    return Promise.resolve(json);
  },
};
