export function todo(...args: any[]): any {
  throw Error("TODO!");
}
export function identity<T>(d: T): T {
  return d;
}
export function noop(): void {}
export function hasProperty<P extends PropertyKey>(
  value: unknown,
  propName: P
): value is { [key in P]: unknown } {
  return typeof value === "object" && value !== null && propName in value;
}
