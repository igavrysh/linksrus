FROM cockroachdb/cockroach

COPY linkgraph/store/cdb/migrations /migrations
COPY cdb-schema/bootstrap-db.sh .

ENTRYPOINT ["bash", "./bootstrap-db.sh"]

