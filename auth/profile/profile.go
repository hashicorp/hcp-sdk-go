package auth

type Profile interface {

	// getOrganizationID() retrieves a user's organization ID.
	getOrganizationID() (string, error)

	// getProjectID() retrieves the user's project ID.
	getProjectID() (string, error)

	// getEnvironment() retrieves the application's development environment.
	getEnvironment() (string, error)

	// setOrganizationID() sets a user's organization ID.
	setOrganizationID(string) error

	// setProjectID() sets the user's project ID.
	setProjectID(string) error

	// setEnvironment() sets the application's development environment.
	setEnvironment(string) error
}
