package auth

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Auth Suite", []Reporter{
		reporters.NewJUnitReporter("../../../ut-report/auth.xml"),
	})
}
