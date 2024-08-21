# syntax=docker/dockerfile:1
FROM golang:1.23.0 AS build
WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN --mount=type=bind,target=. \
    --mount=type=cache,target=/go/pkg/mod/ \
    CGO_ENABLED=0 GOOS=linux go build -o /bin/server

FROM scratch AS final
WORKDIR /bin
COPY ./public ./public/
COPY --from=build /bin/server ./

EXPOSE 3000
ENTRYPOINT [ "/bin/server" ]
