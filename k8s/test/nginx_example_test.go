package test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/require"
)

func TestKubernetesHelloWorldExample(t *testing.T) {
	t.Parallel()

	// Set the path to the K8s manifest to deploy
	kubeResourcePath, err := filepath.Abs("../manifest/nginx-example.yaml")
	require.NoError(t, err)

	// Specify a new namespace with a unique 6 digit ID at the end
	namespaceName := fmt.Sprintf("terratest-example-%s", strings.ToLower(random.UniqueId()))

	// Configure the kubectl with our custom/random namespace, rest go with the defaults:
	// - HOME/.kube/config for the kubectl config file
	// - Current context of the kubectl config file
	options := k8s.NewKubectlOptions("", "", namespaceName)

	// Create the testing namespace
	test_structure.RunTestStage(t, "create_namespace", func() {
		k8s.CreateNamespace(t, options, namespaceName)
	})

	// Remove the testing namespace at the end of the test run
	defer test_structure.RunTestStage(t, "cleanup_namespace", func() {
		k8s.DeleteNamespace(t, options, namespaceName)
	})

	// Remove all K8s resources deployed as part of the test run
	defer test_structure.RunTestStage(t, "cleanup_k8s_resources", func() {
		k8s.KubectlDelete(t, options, kubeResourcePath)
	})

	// Apply the desired K8s manifest that is to be tested/verified
	test_structure.RunTestStage(t, "k8s_apply", func() {
		k8s.KubectlApply(t, options, kubeResourcePath)
	})

	// Wait until ingress is available
	test_structure.RunTestStage(t, "wait_for_ingress", func() {
		k8s.WaitUntilIngressAvailable(t, options, "example-ingress", 60, 3*time.Second)
	})

	//Verify that the ingress is configured and available
	test_structure.RunTestStage(t, "ingress_check", func() {
		ingress := k8s.GetIngress(t, options, "example-ingress")
		require.Equal(t, ingress.Name, "example-ingress")
	})

	// Confirm the ingress and underlying K8s resources work as desired
	test_structure.RunTestStage(t, "http_check", func() {
		http_helper.HttpGetWithRetry(t, "http://localhost/foo", nil, 200, "foo", 30, 3*time.Second)
	})
}
