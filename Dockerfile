FROM gcr.io/distroless/base-debian10 AS base
COPY migrations /migrations
ENV POSTGRES_MIGRATION_FILE_DIR file:///migrations

FROM base AS gin-starter
ENV PORT 8080
EXPOSE $PORT
COPY bin/cmd/server /server
CMD ["/server"]