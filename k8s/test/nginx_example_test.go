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
	"github.com/stretchr/testify/require"
)

func TestKubernetesHelloWorldExample(t *testing.T) {
	t.Parallel()

	kubeResourcePath, err := filepath.Abs("../manifest/nginx-example.yaml")
	require.NoError(t, err)

	namespaceName := fmt.Sprintf("terratest-example-%s", strings.ToLower(random.UniqueId()))

	options := k8s.NewKubectlOptions("", "", namespaceName)

	k8s.CreateNamespace(t, options, namespaceName)

	defer k8s.DeleteNamespace(t, options, namespaceName)

	defer k8s.KubectlDelete(t, options, kubeResourcePath)

	k8s.KubectlApply(t, options, kubeResourcePath)

	k8s.WaitUntilIngressAvailable(t, options, "example-ingress", 60, 3*time.Second)
	ingress := k8s.GetIngress(t, options, "example-ingress")
	require.Equal(t, ingress.Name, "example-ingress")

	http_helper.HttpGetWithRetry(t, "http://localhost/foo", nil, 200, "foo", 30, 3*time.Second)
}
