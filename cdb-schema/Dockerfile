FROM cockroachdb/cockroach

RUN cd ..

RUN echo "$PWD"

RUN ls


COPY linkgraph/store/cdb/migrations /migrations
COPY cdb-schema/bootstrap-db.sh .

ENTRYPOINT ["bash", "./bootstrap-db.sh"]

