.PHONY: help bootstrap-minikube purge deploy k8s-data-namespace helm-es-service helm-cdb-service dockerize-and-push

SHELL=/bin/bash -o pipefail
ES_VERSION ?= 7.16.1
CDB_VERSION ?= v21.2.3
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

bootstrap-minikube:
	@echo "[minikube] bootstrapping (driver: ${MINIKUBE_DRIVER}, network-plugin: cni) cluster with kubernetes ${MINIKUBE_K8S_VERSION} and reserving ${MINIKUBE_RAM} of RAM and ${MINIKUBE_CPUS} CPU(s)"
	@minikube start  --driver=${MINIKUBE_DRIVER} --install-addons=true --addons=registry ingress --network-plugin=cni --cni=calico --kubernetes-version=${MINIKUBE_K8S_VERSION} --memory=${MINIKUBE_RAM} --cpus=${MINIKUBE_CPUS} 2>&1 | sed -e 's/^/ | /g'
	@minikube addons enable ingress
	@minikube addons enable registry
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

k8s-data-namespace:
	@if [[ -z `kubectl get ns | grep data` ]]; then \
		echo "[kubectl] applying namespace manifests";\
		kubectl apply -f 01-namespaces.yaml 2>&1 | sed -e 's/^/ | /g';\
	else \
  		echo "[kubectl] data namespace already exists";\
	fi

purge:
	@echo "[kubectl] removing helm deployments for CDB/ES"
	@helm uninstall es --namespace linksrus-data 2>&1 | sed -e 's/^/ | /g' || true
	@helm uninstall cdb --namespace linksrus-data 2>&1 | sed -e 's/^/ | /g' || true
	@echo "[kubectl] removing remaining resources"
	@kubectl delete -f . 2>&1 | sed -e 's/^/ | /g' || true

deploy: k8s-data-namespace helm-es-service helm-cdb-service
	@echo "[kubectl] deploying linksrus"
	@kubectl apply -f . 2>&1 | sed -e 's/^/ | /g'

helm-es-service: k8s-data-namespace
	@if [[ `kubectl -n linksrus-data get pods -l release=es 2>/dev/null | wc -l` -eq '0' ]]; then \
		echo "[helm] installing elasticsearch (${ES_VERSION})";\
		echo "[helm] adding chart repos";\
		helm repo add elastic https://helm.elastic.co 2>&1 | sed -e 's/^/ | /g';\
		helm repo update; \
		helm install es elastic/elasticsearch\
          -n linksrus-data\
          -f ./chart-settings/es-settings.yaml \
          --set ImageTag=${ES_VERSION} 2>&1 | sed -e 's/^/ | /g';\
	fi

helm-cdb-service: k8s-data-namespace
	@if [[ `kubectl get pods -n linksrus-data -l app.kubernetes.io/component=cockroachdb 2>/dev/null | wc -l` -eq '0' ]]; then \
		echo "[helm] installing cockroackdb (${CDB_VERSION})";\
		echo "[helm] adding chart repos";\
		helm repo add cockroachdb https://charts.cockroachdb.com/;\
		helm repo update;\
		helm install cdb cockroachdb/cockroachdb \
		  -n linksrus-data\
		  -f chart-settings/cdb-settings.yaml \
		  --set ImageTag=${CDB_VERSION} 2>&1 | sed -e 's/^/ | /g';\
	fi

dockerize-and-push:
	@make -C ../cdb-schema dockerize-and-push
	@make -C .. dockerize-and-push

pods:
	@kubectl -n linksrus-data get pods

resources:
	@kubectl -n linksrus get pods,statefulsets,services,ingresses

data-resources:
	@kubectl -n linksrus-data get pods,statefulsets,services,ingresses

logs-monolith-instance:
	@kubectl -n linksrus logs linksrus-monolith-instance-0 -f

logs-monolith-instances:
	@kubectl -n linksrus logs -lapp=linksrus-monolith-instance -f

logs-front:
	@kubectl -n linksrus logs service/linksrus-frontend -f

logs-es-master:
	@kubectl -n linksrus-data logs service/elasticsearch-master -f

logs-es-master-headless:
	@kubectl -n linksrus-data logs service/elasticsearch-master-headless -f

logs-elasticsearch-master:
	@kubectl -n linksrus-data logs statefulset.apps/elasticsearch-master -f

