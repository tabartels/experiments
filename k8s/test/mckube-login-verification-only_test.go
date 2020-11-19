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
const master_desired int = 3
const etcd_desired int = 3

func TestAvailableNodes(t *testing.T) {
	options := k8s.NewKubectlOptions("", "", ns)

	//Verify the worker nodes
	test_structure.RunTestStage(t, "node_check", func() {

		nodes := k8s.GetReadyNodes(t, options)
		var etcd_actual int = 0
		var master_actual int = 0
		for _, node := range nodes {
			if node.Labels["kubernetes.io/role"] == "etcd" {
				etcd_actual++
			}
			if node.Labels["kubernetes.io/role"] == "master" {
				master_actual++
			}
		}
		//Assertion
		require.Equal(t, master_desired, master_actual)
		require.Equal(t, etcd_desired, etcd_actual)
	})
}

func TestMckubeLogin(t *testing.T) {
	options := k8s.NewKubectlOptions("", "", ns)

	//Verify that the ingress is configured and available
	test_structure.RunTestStage(t, "ingress_check", func() {
		ingress := k8s.GetIngress(t, options, ing)
		available := k8s.IsIngressAvailable(ingress)
		//Assertion
		require.Equal(t, ing, ingress.Name)
		require.True(t, available)
	})

	// Confirm the ingress and underlying K8s resources work as desired
	test_structure.RunTestStage(t, "http_check", func() {
		ingress := k8s.GetIngress(t, options, ing)
		hostname := ingress.Spec.Rules[0].Host
		url := fmt.Sprintf("https://%s", hostname)
		_, body, _ := http_helper.HttpGetE(t, url, nil)
		//Assertion
		require.Contains(t, body, "Kubernetes Login")
	})
}
