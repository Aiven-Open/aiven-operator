SHELL := /bin/bash

# VERSION defines the project version for the bundle.
# Update this value when you upgrade the version of your project.
# To re-generate a bundle for another specific version without changing the standard setup, you can:
# - use the VERSION as arg of the bundle target (e.g make bundle VERSION=0.0.2)
# - use environment variables to overwrite this value (e.g export VERSION=0.0.2)
VERSION ?= 0.0.1

# CHANNELS define the bundle channels used in the bundle.
# Add a new line here if you would like to change its default config. (E.g CHANNELS = "preview,fast,stable")
# To re-generate a bundle for other specific channels without changing the standard setup, you can:
# - use the CHANNELS as arg of the bundle target (e.g make bundle CHANNELS=preview,fast,stable)
# - use environment variables to overwrite this value (e.g export CHANNELS="preview,fast,stable")
ifneq ($(origin CHANNELS), undefined)
BUNDLE_CHANNELS := --channels=$(CHANNELS)
endif

# DEFAULT_CHANNEL defines the default channel used in the bundle.
# Add a new line here if you would like to change its default config. (E.g DEFAULT_CHANNEL = "stable")
# To re-generate a bundle for any other default channel without changing the default setup, you can:
# - use the DEFAULT_CHANNEL as arg of the bundle target (e.g make bundle DEFAULT_CHANNEL=stable)
# - use environment variables to overwrite this value (e.g export DEFAULT_CHANNEL="stable")
ifneq ($(origin DEFAULT_CHANNEL), undefined)
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
endif
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)

# IMAGE_TAG_BASE defines the docker.io namespace and part of the image name for remote images.
# This variable is used to construct full image tags for bundle and catalog images.
#
# For example, running 'make bundle-build bundle-push catalog-build catalog-push' will build and push both
# example.com/test-bundle:$VERSION and example.com/test-catalog:$VERSION.
IMAGE_TAG_BASE ?= example.com/test

# BUNDLE_IMG defines the image:tag used for the bundle.
# You can use it as an arg. (E.g make bundle-build BUNDLE_IMG=<some-registry>/<project-name-bundle>:<tag>)
BUNDLE_IMG ?= $(IMAGE_TAG_BASE)-bundle:v$(VERSION)

# Image URL to use all building/pushing image targets
IMG ?= controller:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true,preserveUnknownFields=false"

GO = go
GOOS = $(shell $(GO) env GOOS)
GOARCH = $(shell $(GO) env GOARCH)

all: build

TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin

HUGO=$(TOOLS_BIN_DIR)/hugo
CONTROLLER_GEN=$(TOOLS_BIN_DIR)/controller-gen
SETUP_ENVTEST=$(TOOLS_BIN_DIR)/setup-envtest
KUSTOMIZE=$(TOOLS_BIN_DIR)/kustomize
GINKGO=$(TOOLS_BIN_DIR)/ginkgo
GOLANGCILINT=$(TOOLS_BIN_DIR)/golangci-lint
GEN_CRD_API_REF_DOCS=$(TOOLS_BIN_DIR)/gen-crd-api-reference-docs

$(HUGO): $(TOOLS_DIR)/go.mod ## Build hugo from tools folder.
	cd $(TOOLS_DIR) && $(GO) build -tags=tools,extended -o bin/hugo github.com/gohugoio/hugo

$(CONTROLLER_GEN): $(TOOLS_DIR)/go.mod ## Build controller-gen from tools folder.
	cd $(TOOLS_DIR) && $(GO) build -tags=tools -o bin/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(SETUP_ENVTEST): $(TOOLS_DIR)/go.mod ## Build kustomize from tools folder.
	cd $(TOOLS_DIR) && $(GO) build -tags=tools -o bin/setup-envtest sigs.k8s.io/controller-runtime/tools/setup-envtest

$(KUSTOMIZE): $(TOOLS_DIR)/go.mod ## Build kustomize from tools folder.
	cd $(TOOLS_DIR) && $(GO) build -tags=tools -o bin/kustomize sigs.k8s.io/kustomize/kustomize/v3

$(GINKGO): $(TOOLS_DIR)/go.mod ## Build ginkgo from tools folder.
	cd $(TOOLS_DIR) && $(GO) build -tags=tools -o bin/ginkgo github.com/onsi/ginkgo/ginkgo

$(GOLANGCILINT): $(TOOLS_DIR)/go.mod ## Build golangci-lint from tools folder.
	cd $(TOOLS_DIR) && $(GO) build -tags=tools -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

$(GEN_CRD_API_REF_DOCS): $(TOOLS_DIR)/go.mod ## Build gen-crd-api-ref-docs from tools folder.
	cd $(TOOLS_DIR) && $(GO) build -tags=tools -o bin/gen-crd-api-ref-docs github.com/ahmetb/gen-crd-api-reference-docs

ENVTEST_K8S_VERSION=1.22.0
ENVTEST_TOOLS_DIR=$(TOOLS_BIN_DIR)/k8s/$(ENVTEST_K8S_VERSION)-$(GOOS)-$(GOARCH)
ENVTEST_TOOLS_ETCD=$(ENVTEST_TOOLS_DIR)/etcd
ENVTEST_TOOLS_KUBE_APISERVER=$(ENVTEST_TOOLS_DIR)/kube-apiserver
ENVTEST_TOOLS_KUBECTL=$(ENVTEST_TOOLS_DIR)/kubectl
ENVTEST_TOOLS= $(ENVTEST_TOOLS_ETCD) $(ENVTEST_TOOLS_KUBE_APISERVER) $(ENVTEST_TOOLS_KUBECTL)

$(ENVTEST_TOOLS): $(SETUP_ENVTEST)
	$(SETUP_ENVTEST) use $(ENVTEST_K8S_VERSION) --bin-dir $(TOOLS_BIN_DIR)

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

manifests: $(CONTROLLER_GEN) ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="api/..." output:crd:artifacts:config=config/crd/bases

generate: $(CONTROLLER_GEN) ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="api/..."

# linting
.PHONY: lint
lint: $(GOLANGCILINT) ## Run acceptance linter.
	$(GOLANGCILINT) run --verbose

.PHONY: test-acc
test-acc: $(GINKGO) $(ENVTEST_TOOLS) ## Run acceptance tests.
	KUBEBUILDER_CONTROLPLANE_START_TIMEOUT=120s \
	KUBEBUILDER_ASSETS=$(abspath $(ENVTEST_TOOLS_DIR)) \
	$(GINKGO) -p --race --randomizeAllSpecs --cover cover.out --trace --failFast --test.count 1 --progress ./controllers

.PHONY: test-e2e
test-e2e: manifests generate ## Run end-to-end tests using kuttl (https://kuttl.dev/)
	kubectl kuttl test --config test/e2e/kuttl-test.yaml

##@ Build

build: generate ## Build manager binary.
	$(GO) build -o bin/manager main.go

run: manifests generate fmt vet ## Run a controller from your host.
	$(GO) run ./main.go

docker-build: ## Build docker image with the manager.
	docker build -t ${IMG} .

docker-push: ## Push docker image with the manager.
	docker push ${IMG}

##@ Deployment

install-cert-manager: ## Deploy cert-manager to the cluster
	kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.5.3/cert-manager.yaml

install: $(KUSTOMIZE) manifests ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

uninstall: $(KUSTOMIZE) manifests ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

deploy: $(KUSTOMIZE) manifests ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/default | kubectl apply -f -

undeploy: $(KUSTOMIZE) ## Undeploy controller from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/default | kubectl delete -f -

.PHONY: bundle
bundle: $(KUSTOMIZE) manifests ## Generate bundle manifests and metadata, then validate generated files.
	operator-sdk generate kustomize manifests -q
	cd config/manager && $(KUSTOMIZE) edit set image controller=$(IMG)
	$(KUSTOMIZE) build config/manifests | operator-sdk generate bundle -q --overwrite --version $(VERSION) $(BUNDLE_METADATA_OPTS)
	operator-sdk bundle validate ./bundle

.PHONY: bundle-build
bundle-build: ## Build the bundle image.
	docker build -f bundle.Dockerfile -t $(BUNDLE_IMG) .

.PHONY: bundle-push
bundle-push: ## Push the bundle image.
	$(MAKE) docker-push IMG=$(BUNDLE_IMG)

.PHONY: opm
OPM = ./bin/opm
opm: ## Download opm locally if necessary.
ifeq (,$(wildcard $(OPM)))
ifeq (,$(shell which opm 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(OPM)) ;\
	OS=$(shell go env GOOS) && ARCH=$(shell $ env GOARCH) && \
	curl -sSLo $(OPM) https://github.com/operator-framework/operator-registry/releases/download/v1.15.1/$${OS}-$${ARCH}-opm ;\
	chmod +x $(OPM) ;\
	}
else
OPM = $(shell which opm)
endif
endif

# A comma-separated list of bundle images (e.g. make catalog-build BUNDLE_IMGS=example.com/operator-bundle:v0.1.0,example.com/operator-bundle:v0.2.0).
# These images MUST exist in a registry and be pull-able.
BUNDLE_IMGS ?= $(BUNDLE_IMG)

# The image tag given to the resulting catalog image (e.g. make catalog-build CATALOG_IMG=example.com/operator-catalog:v0.2.0).
CATALOG_IMG ?= $(IMAGE_TAG_BASE)-catalog:v$(VERSION)

# Set CATALOG_BASE_IMG to an existing catalog image tag to add $BUNDLE_IMGS to that image.
ifneq ($(origin CATALOG_BASE_IMG), undefined)
FROM_INDEX_OPT := --from-index $(CATALOG_BASE_IMG)
endif

# Build a catalog image by adding bundle images to an empty catalog using the operator package manager tool, 'opm'.
# This recipe invokes 'opm' in 'semver' bundle add mode. For more information on add modes, see:
# https://github.com/operator-framework/community-operators/blob/7f1438c/docs/packaging-operator.md#updating-your-existing-operator
.PHONY: catalog-build
catalog-build: opm ## Build a catalog image.
	$(OPM) index add --container-tool docker --mode semver --tag $(CATALOG_IMG) --bundles $(BUNDLE_IMGS) $(FROM_INDEX_OPT)

# Push the catalog image.
.PHONY: catalog-push
catalog-push: ## Push a catalog image.
	$(MAKE) docker-push IMG=$(CATALOG_IMG)

##@ Docs
.PHONY: serve-docs
# Run Hugo live preview.
serve-docs: $(HUGO) ## serves docs
	$(HUGO) serve docs -s docs

generate-docs: $(HUGO) $(GEN_CRD_API_REF_DOCS) ## Generate the documentation website locally.
	$(GO) generate hack/genrefs/gen.go
	$(HUGO) --minify -s docs

