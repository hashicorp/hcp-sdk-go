package profile

const (
	envVarHCPOrganizationID = "HCP_ORGANIZATION_ID"

	envVarHCPProjectID = "HCP_PROJECT_ID"
)

// UserProfile implements the Profile interface.
type UserProfile struct {

	// hcpOrganizationID is the user's organization ID.
	OrganizationID string

	// hcpProjectID is the user's project ID.
	ProjectID string
}

// getOrganizationID() retrieves a user's organization ID.
func (p *UserProfile) GetOrganizationID() string {
	return p.OrganizationID
}

// getProjectID() retrieves the user's project ID.
func (p *UserProfile) GetProjectID() string {
	return p.ProjectID
}
