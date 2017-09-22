package go_multi_mock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoMultiMock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoMultiMock Suite")
}
