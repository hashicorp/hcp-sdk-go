package auth

//TODO: should we define the profile mock here? Or do we want to contain all of our interface mocks in the mock.go file?
type MockProfile struct{

	// hcpOrganizationID is the user's organization ID.
	hcpOrganizationID string

	// hcpProjectID is the user's project ID.
	hcpProjectID string

	// hcpEnvironment is the application's development environment.
	hcpEnvironment string


}
