package stacks

import (
	"github.com/kubeflow/manifests/tests"
	"testing"
)

func TestKustomize(t *testing.T) {
	testCase := &tests.KustomizeTestCase{
		Package:  "../../../../../common/centraldashboard/overlays/stacks",
		Expected: "test_data/expected",
	}

	tests.RunTestCase(t, testCase)
}
