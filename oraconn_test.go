package oraconn

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"strings"
)

func cleanEnv() {
	os.Unsetenv(DBUser)
	os.Unsetenv(DBPassword)
	os.Unsetenv(DBHost)
	os.Unsetenv(DBPort)
	os.Unsetenv(DBSvc)
}

func setEnv(config map[string]string) {
	for k,v := range config {
		os.Setenv(k,v)
	}
}


func TestConfig(t *testing.T) {
	var configTests = []struct {
		testName string
		config map[string]string
		errorComponents []string
		expectError bool
		connectString string
	}{
		{
			"all environment present",
			map[string]string{DBUser:"user",DBPassword:"password",DBHost:"host",DBPort:"port",DBSvc:"svc"},
			[]string{},
			false,
			"user/password@//host:port/svc",
		},
		{
			"no environment present",
			map[string]string{},
			[]string{DBUser,DBPassword,DBHost,DBPort,DBSvc},
			true,
			"",
		},
		{
			"some environment present",
			map[string]string{DBUser:"user",DBPort:"port",DBSvc:"svc"},
			[]string{DBPassword,DBHost},
			true,
			"",
		},
	}

	for _, test := range configTests {
		t.Run(test.testName,func(t *testing.T){
			cleanEnv()
			setEnv(test.config)
			config, err := NewEnvConfig()
			if test.expectError {
				assert.NotNil(t, err, "expected error")
				errString := err.Error()
				for _,ec := range test.errorComponents {
					assert.True(t, strings.Contains(errString, ec))
				}
			} else {
				assert.Equal(t, test.connectString, config.ConnectString())
			}
		})
	}
}
