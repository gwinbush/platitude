.PHONY: build deploy delete

build:
	bazel build //... --platforms=//:linux_arm_64

deploy: build
	./scripts/deploy.sh

delete:
	kubectl delete -f k8s/ || true
