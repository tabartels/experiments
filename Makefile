.PHONY: terratest
terratest:
	cd k8s/test/ && go mod download && go test -v
