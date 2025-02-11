FROM postgres:16.4

COPY ./docker/db/migrate /docker-entrypoint-initdb.d
COPY ./docker/db/insert /docker-entrypoint-initdb.d
