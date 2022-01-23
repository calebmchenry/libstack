
declare namespace auth {
type State = {
  token?: string;
  username?: string;
  loginErr?: Error;
  loggingIn: boolean;
};
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
