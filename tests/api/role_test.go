package api

import (
	"testing"

	factories "github.com/nugrohosam/gosampleapi/tests/factories"
	utilities "github.com/nugrohosam/gosampleapi/tests/utilities"
)

// TestRun ...
func RoleTestRun(t *testing.T) {
	InitialTest(t)
	defer utilities.DbCleaner(t)

	User = factories.CreateUser()

	t.Log("Test Positive")
	t.Log("=======>>>> <<<<======")
	testAuthRegister(t)

}

func testCreatRole() {

}
