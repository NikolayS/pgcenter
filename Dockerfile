# stage 1: build
FROM golang:1.15 as build
LABEL stage=intermediate
WORKDIR /app
COPY . .
RUN make build

# stage 2: scratch
FROM alpine:3.13 as scratch
COPY --from=build /app/bin/pgcenter /bin/pgcenter
CMD ["pgcenter"]