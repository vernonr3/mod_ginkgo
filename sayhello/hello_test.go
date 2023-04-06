package sayhello_test

import (
	"fmt"
	"mod_ginkgo/sayhello"
	"testing"

	ginkgo "github.com/onsi/ginkgo/v2"
	gomega "github.com/onsi/gomega"
)

func Test_Hello(t *testing.T) {
	var overall_result int
	testcases := []struct {
		input string
		want  int
	}{
		{input: "world", want: 1},
	}
	for _, testcase := range testcases {
		overall_result = sayhello.Hello(testcase.input)
		if overall_result != testcase.want {
			fmt.Printf("Bad result %d\n", overall_result)
		}
	}
	return
}

var mytesting *testing.T

func TestWelcome(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Welcome")
	mytesting = t
}

var _ = ginkgo.Describe("Welcome", func() {
	ginkgo.It("should be a novel", func() {
		Test_Hello(mytesting)
	})
})
