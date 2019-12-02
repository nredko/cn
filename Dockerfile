FROM golang:1.13-buster as build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make -j ctrlt-static cn-static
FROM alpine:3.10
WORKDIR /app
COPY --from=build /src/ui ui/
COPY --from=build /src/ctrlt /bin/ctrlt
COPY --from=build /src/cn /bin/cn
ENTRYPOINT ["/bin/ctrlt"]
