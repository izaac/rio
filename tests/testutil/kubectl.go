package testutil

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"testing"
)

// KubectlHelper provides Kubernetes-related test helpers. It connects to the
// Kubernetes API using the environment's configured kubeconfig file.
type KubectlHelper struct {
	T *testing.T
}

// KubectlGet executes kubectl CLI commands with your arguments, assuming the default namespace unless otherwise specified
// Example: args=["get", "-n", "test", "services"] would run: "kubectl get -n test service"
func (kh *KubectlHelper) KubectlGet(args []string) (string, error) {
	args = append([]string{"get", "-n", testingNamespace}, args...) // named command is always first arg
	cmd := exec.Command("kubectl", args...)
	stdOutErr, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s: %s", err.Error(), stdOutErr)
	}
	return string(stdOutErr), nil
}

// GetKubectlConfigMapContent, recieves the TestConfig object to retrieve the test ConfigMap data -> content
// Example: config would run: "kubectl get -n testing-ns cm service -o json"
func (kh *KubectlHelper) GetKubectlConfigMapContent(config TestConfig) string {

	var kHelper KubectlHelper

	resultString, _ := kHelper.KubectlGet([]string{"cm", config.Name, "-o", "json"})
	var results map[string]interface{}
	json.Unmarshal([]byte(resultString), &results)
	var resultsData = results["data"].(map[string]interface{})

	return resultsData["content"].(string)
}