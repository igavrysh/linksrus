.PHONY: help bootstrap-minikube purge deploy k8s-data-namespace helm-es-service helm-cdb-service dockerize-and-push

SHELL=/bin/bash -o pipefail
ES_VERSION ?= 7.4.0
CDB_VERSION ?= v19.1.5
MINIKUBE_RAM ?= 5g
MINIKUBE_CPUS ?= 3
MINIKUBE_K8S_VERSION ?= stable
MINIKUBE_DRIVER ?= virtualbox

help:
	@echo "Usage: make COMMAND"
	@echo ""
	@echo "Supported commands:"
	@echo "- bootstrap-minikube : bootstrap minikube cluster and install required addons"
	@echo "- deploy             : deploy a micro-service based full links 'R' us installation backed by CockroachDB and elastic search"
	@echo "- purge              : delete links 'R' us deployment including backing databases"
	@echo "- dockerize-and-push : dockerize and push all required images"
	@echo ""

# 	@minikube start --cni=flannel --install-addons=true --addons=ingress registry --vm-driver=${MINIKUBE_DRIVER} --network-plugin=cni --kubernetes-version=${MINIKUBE_K8S_VERSION} --memory=${MINIKUBE_RAM} --cpus=${MINIKUBE_CPUS} 2>&1 | sed -e 's/^/ | /g'

bootstrap-minikube:
	@echo "[minikube] bootstrapping (driver: ${MINIKUBE_DRIVER}, network-plugin: cni) cluster with kubernetes ${MINIKUBE_K8S_VERSION} and reserving ${MINIKUBE_RAM} of RAM and ${MINIKUBE_CPUS} CPU(s)"
	@minikube start  --driver=${MINIKUBE_DRIVER} --install-addons=true --addons=registry ingress --network-plugin=cni --cni=calico --kubernetes-version=${MINIKUBE_K8S_VERSION} --memory=${MINIKUBE_RAM} --cpus=${MINIKUBE_CPUS} 2>&1 | sed -e 's/^/ | /g'
	@echo
	@echo "IMPORTANT NOTICE:"
	@echo "Please make sure to add '"`minikube ip`":5000' as a trusted"
	@echo "insecure registry to your docker's configuration options and"
	@echo "restart the docker daemon!"
	@echo
	@echo 'On Linux, you can do this by editing /etc/docker/daemon.json'
	@echo 'and adding the following section:'
	@echo '  {'
	@echo '    "insecure-registries" : ['
	@echo '      "'`minikube ip`':5000"'
	@echo '    ]'
	@echo '  }'
	@echo
	@echo 'On OSX and Windows you can right-click the docker icon, go to'
	@echo '"preferences" and then click on the "Daemon" tab'

purge:
	@echo "[kubectl] removing helm deployments for CDB/ES"
	@helm delete --purge es 2>&1 | sed -e 's/^/ | /g' || true
	@helm delete --purge cdb 2>&1 | sed -e 's/^/ | /g' || true
	@echo "[kubectl] removing remaining resources"
	@kubectl delete -f . 2>&1 | sed -e 's/^/ | /g' || true
