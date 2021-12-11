
# running cockroach db
cockroach start-single-node --insecure --advertise-addr 127.0.0.1:26257
cockroach sql --insecure -e 'CREATE DATABASE linkgraph;'


cockroach sql --insecure -e 'SHOW DATABASES;'

export CDB_DSN='postgresql://root@localhost:26257/linkgraph?sslmode=disable'

export ES_NODES='http://localhost:9200'


# running elastic search, after downloading it from web and putting in /opt directory
ES_JAVA_OPTS="-Xms1g -Xmx1g" /opt/elasticsearch-7.15.2/bin/elasticsearch


# installing and running golang migrate tool
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate.linux-amd64 $GOPATH/bin/migrate

migrate create -digits 2 -seq -ext=.sql -dir=./linkgraph/store/cdb/migrations create_links_table


# to install mockgen (required for running `go generate ./...`
go install github.com/golang/mock/mockgen@v1.6.0


# run minikube (not working)
minikube start --kubernetes-version=v1.22.3 \
--memory=4g \
--network-plugin=cni

# run minikube (finally worked)
# minikube start -p issue-12152 --cni=flannel --install-addons=true \
# --addons=ingress registry --kubernetes-version=stable --memory=4g --cpus=2

# the same as above but without -p issue-12152
minikube start --cni=flannel --install-addons=true --addons=ingress registry --kubernetes-version=stable --memory=4g --cpus=2 --network-plugin=cni

# to access minikube virtual machine from the host machine
minikube addons enable ingress

# private registry addon (used for tests, basically this is a private Docker registry for pushing the Docker
# images that are built in this project
minikube addons enable registry

# to get minikube ip
minikube ip

# to allow private unsecured local image repositories
# edit /etc/docker/daemon.json
# add the following lines
# {
#   "insecure-registries" : [
#     "$MINIKUBE_IP:5000"
#   ]
# }

# to get all dependencies
go mod tidy

# to install dep tool
go get -d github.com/golang/dep/cmd/dep

# to initialize dep
dep init

# to install dep tool
sudo apt install go-dep


# to run docker locally (after creating image in service/ folder
docker run -it --rm -p 8080:8080 $(minikube ip):5000/linksrus-monolith:latest