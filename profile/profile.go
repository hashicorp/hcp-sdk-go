package profile

import (
	"fmt"
	"os"
)

type Profile interface {

	// getOrganizationID() retrieves a user's organization ID.
	GetOrganizationID() string

	// getProjectID() retrieves the user's project ID.
	GetProjectID() string

	// getEnvironment() retrieves the application's development environment.
	GetEnvironment() string

	// setOrganizationID() sets a user's organization ID.
	SetOrganizationID(string) error

	// setProjectID() sets the user's project ID.
	SetProjectID(string) error

	// setEnvironment() sets the application's development environment.
	SetEnvironment(string) error
}

const (
	envVarHCPOrganizationID = "HCP_ORGANIZATION_ID"

	envVarHCPProjectID = "HCP_PROJECT_ID"

	envVarHCPEnvironment = "HCP_ENVIRONMENT"
)

// UserProfile implements the Profile interface.
type UserProfile struct {

	// hcpOrganizationID is the user's organization ID.
	hcpOrganizationID string

	// hcpProjectID is the user's project ID.
	hcpProjectID string

	// hcpEnvironment is the application's development environment.
	hcpEnvironment string
}

//TODO: Do we need to include comments here or is it covered since we comment the function definitions in the interface?
func (p *UserProfile) GetOrganizationID() string {
	return p.hcpOrganizationID
}

func (p *UserProfile) GetProjectID() string {
	return p.hcpProjectID
}

func (p *UserProfile) GetEnvironment() string {
	return p.hcpEnvironment

}

func (p *UserProfile) SetOrganizationID(id string) error {
	err := os.Setenv(envVarHCPOrganizationID, id)
	if err != nil {
		return fmt.Errorf("Unable to set %s: %v", envVarHCPOrganizationID, err)
	}
	p.hcpOrganizationID = id
	return nil
}

func (p *UserProfile) SetProjectID(id string) error {
	err := os.Setenv(envVarHCPProjectID, id)
	if err != nil {
		return fmt.Errorf("Unable to set %s: %v", envVarHCPProjectID, err)
	}
	p.hcpProjectID = id
	return nil
}

func (p *UserProfile) SetEnvironment(env string) error {
	err := os.Setenv(envVarHCPEnvironment, env)
	if err != nil {
		return fmt.Errorf("Unable to set %s: %v", envVarHCPEnvironment, err)
	}
	p.hcpEnvironment = env
	return nil
}
