
# running cockroach db
cockroach start --insecure --advertise-addr 127.0.0.1:26257
cockroach sql --insecure -e 'CREATE DATABASE linkgraph;'
export CDB_DSN='postgresql://root@localhost:26257/linkgraph?sslmode=disable'


# running elastic search