package repository

// Anytime we create a new function we can add it to this interface. We can access
// this in our handlers.
type DatabaseRepo interface {
	AllUsers() bool
}
