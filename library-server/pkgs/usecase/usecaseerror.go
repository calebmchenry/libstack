package usecase

import "fmt"

type UseCaseError struct {
	Kind    string
	message string
	cause   error
}

func (e *UseCaseError) Error() string {
	return fmt.Sprintf("%s: %s", e.Kind, e.message)
}

func (e *UseCaseError) Cause() error {
	return e.cause
}

func internalError(message string, cause error) *UseCaseError {
	return &UseCaseError{Kind: "internal_error", message: message, cause: cause}
}
func unauthorizedError(message string) *UseCaseError {
	return &UseCaseError{Kind: "unauthorized", message: message}
}
func notFoundError(message string) *UseCaseError {
	return &UseCaseError{Kind: "not_found", message: message}
}
func invalidArgumentError(message string) *UseCaseError {
	return &UseCaseError{Kind: "invalid_argument", message: message}
}
func holdInsteadError(message string) *UseCaseError {
	return &UseCaseError{Kind: "hold_instead", message: message}
}
