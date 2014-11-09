package golang_rails_template_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "testing"
)

func TestGolangRailsTemplate(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "GolangRailsTemplate Suite")
}
