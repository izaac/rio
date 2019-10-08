package integration

import (
	"strings"
	"testing"

	"github.com/sclevine/spec"
	"github.com/stretchr/testify/assert"

	"github.com/rancher/rio/tests/testutil"
)

func configTests(t *testing.T, when spec.G, it spec.S) {

	var config testutil.TestConfig
	var kh testutil.KubectlHelper

	it.After(func() {
		config.Remove()
	})

	when("a config is created with data", func() {
		it("should contain that data", func() {
			testText := []string{"a=b", "foo=bar"}
			expectedKubectlText := strings.Join(testText, "\n") + "\n"
			config.Create(t, testText)
			assert.Equal(t, testText, config.GetContent())
			assert.Equal(t, expectedKubectlText, kh.GetKubectlConfigMapContent(config))
		})
	}, spec.Parallel())
}
