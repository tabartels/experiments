![build](https://github.com/tabartels/experiments/workflows/build/badge.svg)
![terratest](https://github.com/tabartels/experiments/workflows/terratest/badge.svg)
# Experiments
Some experiments with GitHub Actions/Workflows.

## Terratest

[Video][1] and [Presentation][2] on Terratest by Gruntwork co-founder Yevgeniy
Brikman.

### Run a terratest
To run the test locally, you need access to a local Kubernetes cluster (eg. KinD).
Make sure you switched the context to your local Kubernetes cluster before running
the test.

If you don't have a KiND cluster running, [here](./k8s/kind/cluster.yaml) is a KiND cluster.yaml.

To run the test:
```bash
make terratest
```

The test run will do the following:
- create a test namespace with a random generated ID
- deploy a simple demo application to the K8s cluster your KUBECONFIG is pointed to
- it will wait for the ingress to be ready to serve requests to the pods
- verify the ingress is ready
- do an HTTP GET on the ingress host
- tear down the test namespace post verification step

### Modules
There is a good [list][3] of existing/maintained modules for Terratest already, some
of them are:
- aws
- azure
- terraform
- k8s


[1]: https://youtu.be/xhHOW0EF5u8
[2]: https://qconsf.com/system/files/presentation-slides/qconsf2019-yevgeniy-brikman-automated-testing-for-terraform-docker-packer-kubernetes-and-more.pdf
[3]: https://github.com/gruntwork-io/terratest/tree/master/modules
