package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	// Create a temporary directory for the test file
	tmpDir, err := ioutil.TempDir("", "config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Define the test configuration
	expectedConfig := config{
		JwtSecret: JwtSecret{
			JwtSecretKey: "test-secret-key",
		},
		DBConfig: DBConfig{
			DBLogMode: true,
			DbConnectionPool: DbConnectionPool{
				MaxOpenConnection:              10,
				MaxIdealConnection:             5,
				MaxConnectionLifetimeInMinutes: 60,
			},
		},
		PasswordConfiguration: PasswordConfiguration{
			PasswordMaxLength: 20,
			PasswordMinLength: 8,
		},
		AuthSecrets: AuthSecrets{
			PublicKey:  "test-public-key",
			PrivateKey: "test-private-key",
		},
	}

	// Marshaling the expected config to JSON
	fileContent, err := json.Marshal(expectedConfig)
	if err != nil {
		t.Fatal(err)
	}

	// Create a temporary file with .json extension
	baseName := "test_config"
	fileName := filepath.Join(tmpDir, baseName+".json")
	err = ioutil.WriteFile(fileName, fileContent, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// calling NewConfiguration with the path (without .json extension as the function appends it)
	// We need to pass the full path minus the extension
	configPath := filepath.Join(tmpDir, baseName)
	loadedConfig := NewConfiguration(configPath)

	assert.NotNil(t, loadedConfig)
	assert.Equal(t, expectedConfig.JwtSecret.JwtSecretKey, loadedConfig.JwtSecret.JwtSecretKey, "JWT Secret Key mismatch")
	assert.Equal(t, expectedConfig.DBConfig.DBLogMode, loadedConfig.DBConfig.DBLogMode, "DB Log Mode mismatch")
	assert.Equal(t, expectedConfig.PasswordConfiguration.PasswordMaxLength, loadedConfig.PasswordConfiguration.PasswordMaxLength, "Password Max Length mismatch")
	assert.Equal(t, expectedConfig.AuthSecrets.PublicKey, loadedConfig.AuthSecrets.PublicKey, "Public Key mismatch")
}

func TestReadConfig_FileNotFound(t *testing.T) {
	// This test verifies that it doesn't panic or handles error gracefully (or prints error as per current implementation)
	// The current implementation just prints error and returns empty config

	configPath := "non_existent_file"
	loadedConfig := NewConfiguration(configPath)

	assert.NotNil(t, loadedConfig)
	assert.Empty(t, loadedConfig.JwtSecret.JwtSecretKey)
}
