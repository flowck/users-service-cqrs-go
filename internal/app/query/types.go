package query

// Types at the application layer
// For instance, the User domain struct might contain the password field whereas the User here doesn't need to
// if the password is not needed to realise the use cases.

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
}
