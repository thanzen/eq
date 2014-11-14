package userservice_test

import (
	"testing"

	"github.com/onsi/ginkgo"
)

func TestServices(t *testing.T) {
	ginkgo.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Books Suite")
}
