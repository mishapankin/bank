FROM postgres:15
COPY ./db_entry/* /docker-entrypoint-initdb.d
EXPOSE 5432
