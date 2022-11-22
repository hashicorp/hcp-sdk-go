package auth

import (
	"fmt"
	"os"
)

const (
	envVarHCPOrganizationID = "HCP_ORGANIZATION_ID"

	envVarHCPProjectID = "HCP_PROJECT_ID"

	envVarHCPEnvironment = "HCP_ENVIRONMENT"
)

type UserProfile struct {

	// hcpOrganizationID is the user's organization ID.
	hcpOrganizationID string

	// hcpProjectID is the user's project ID.
	hcpProjectID string

	// hcpEnvironment is the application's development environment.
	hcpEnvironment string
}

//TODO: Do we need to include comments here or is it covered since we comment the function definitions in the interface?
func (p *UserProfile) getOrganizationID() (string, error) {
	id := os.Getenv(envVarHCPOrganizationID)
	if id == "" {
		return "", fmt.Errorf("%s not initialized", envVarHCPOrganizationID)
	}
	return id, nil
}

func (p *UserProfile) getProjectID() (string, error) {
	id := os.Getenv(envVarHCPProjectID)
	if id == "" {
		return "", fmt.Errorf("%s not initialized", envVarHCPProjectID)
	}
	return id, nil
}

func (p *UserProfile) getEnvironment() (string, error) {
	env := os.Getenv(envVarHCPEnvironment)
	if env == "" {
		return "", fmt.Errorf("%s not initialized", envVarHCPEnvironment)
	}
	return env, nil

}

func (p *UserProfile) setOrganizationID(id string) error {
	err := os.Setenv(envVarHCPOrganizationID, id)
	if err != nil {
		return fmt.Errorf("Unable to set %s: %v", envVarHCPOrganizationID, err)
	}
	return nil
}

func (p *UserProfile) setProjectID(id string) error {
	err := os.Setenv(envVarHCPProjectID, id)
	if err != nil {
		return fmt.Errorf("Unable to set %s: %v", envVarHCPProjectID, err)
	}
	return nil
}

func (p *UserProfile) setEnvironment(env string) error {
	err := os.Setenv(envVarHCPEnvironment, env)
	if err != nil {
		return fmt.Errorf("Unable to set %s: %v", envVarHCPEnvironment, err)
	}
	return nil
}
