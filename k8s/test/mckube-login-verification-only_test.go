// Use with an existing KUBECONFIG exported via shell
// Test will verify:
// - ingress is available in prescribed namespace
// - ingress host is accessible and returns a HTTP 200 status code and
//   body contains the string provided
package test

import (
	"fmt"
	"testing"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/require"
)

const ns string = "kube-system"
const ing string = "mckube-login-mckube-login"

func TestMckubeLogin(t *testing.T) {
	options := k8s.NewKubectlOptions("", "", ns)

	//Verify that the ingress is configured and available
	test_structure.RunTestStage(t, "ingress_check", func() {
		ingress := k8s.GetIngress(t, options, ing)
		available := k8s.IsIngressAvailable(ingress)
		require.Equal(t, ingress.Name, ing)
		require.True(t, available)
	})

	// Confirm the ingress and underlying K8s resources work as desired
	test_structure.RunTestStage(t, "http_check", func() {
		ingress := k8s.GetIngress(t, options, ing)
		hostname := ingress.Spec.Rules[0].Host
		url := fmt.Sprintf("https://%s", hostname)
		_, body, _ := http_helper.HttpGetE(t, url, nil)
		require.Contains(t, body, "Kubernetes Login")
	})
}
