FROM golang:1 AS build

WORKDIR /go/src/github.com/Nekroze/docker-compose-gen

RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh \
  | bash -s -- -b "$GOPATH/bin" latest \
 && go get -u github.com/kyoh86/richgo

RUN go get -v -d -u \
  github.com/docker/libcompose \
  github.com/spf13/cobra \
  gopkg.in/yaml.v2

COPY ./*.go ./
COPY ./cmd ./cmd

RUN richgo test -v ./... \
 && golangci-lint run --deadline '2m' --enable-all \
 && GO_ENABLED=0 GOOS=linux GOARCH=386 richgo build \
   -a -installsuffix cgo -ldflags='-w -s' -o /bin/docker-compose-gen -v \
   .


FROM nekroze/containaruba:alpine AS test
CMD ["--order=random"]

COPY --from=build /bin/docker-compose-gen /bin/docker-compose-gen
COPY ./features /usr/src/app/features


FROM scratch AS final

COPY --from=build /bin/docker-compose-gen /docker-compose-gen
ENTRYPOINT ["/docker-compose-gen"]
