package modulemongodb_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestModulemongodb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Modulemongodb Suite")
}
