package profile

import (
	"os"
	"testing"

	assertpkg "github.com/stretchr/testify/assert"
	requirepkg "github.com/stretchr/testify/require"
)

func TestGetProfileInformation(t *testing.T) {
	require := requirepkg.New(t)
	assert := assertpkg.New(t)

	testCases := []struct {
		name      string
		envVarName string
		profileValue string
		caseSetup func(*Profile) (string)
	}{
		{
			name:      "Get Organization ID",
			envVarName: "HCP_ORGANIZATION_ID",
			profileValue: "1376",
			caseSetup: (UserProfile) func() string {
				//caseSetup

				//caseCheck
				aProfile := UserProfile{}
				id, err := aProfile.getOrganizationID() 
				require.NoError(err)
				return id
			},
		},
		{
			name: "Get Profile ID",
			caseSetup: func() string {
				now := time.Now()
				cache := Cache{
					AccessToken:       "ExpiredAccessToken",
					RefreshToken:      "ExpiredRefreshToken",
					AccessTokenExpiry: now,
					SessionExpiry:     time.Now().Add(time.Hour * -2),
				}

				err := Write(cache)
				require.NoError(err)

				return nil
			},
		},
		{
			name:      "Get Environment",
			caseSetup: func() error { return nil },
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {


			// Runs specific test case setup.
			err = os.Setenv(testCase.envVarName, testCase.profileValue)
			require.NoError(err)

			// Run the test.
			aProfile := UserProfile{}

			// Make assertions about token retrieved.
			assert.Equal("SomeNewAccessToken", tok.AccessToken)
			assert.Equal("SomeNewRefreshToken", tok.RefreshToken)
			assert.WithinDuration(time.Now().Add(time.Hour*1), tok.Expiry, 10*time.Second)

			// Cleanup.
			require.NoError(CacheDestroy(t))
		})
	}
}
