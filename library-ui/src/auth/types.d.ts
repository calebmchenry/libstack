type State = {
  token?: string;
  username?: string;
  loginErr?: Error;
  loggingIn: boolean;
};

declare namespace auth {
  type Provision = State & {
    login: (values: unknown) => Promise<void>;
    logout: () => Promise<void>;
  };

  type Token = string;
  type Credentials = {
    email: string;
    password: string;
  };
}
