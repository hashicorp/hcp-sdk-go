package auth

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWrite_NoDirectoryNoFile(t *testing.T) {

}

func setup() {
	os.Setenv("HCP_CACHE_TEST_ENV", "true")
	credentialDirectory := filepath.Join(auth.userHome, directoryName)

}
