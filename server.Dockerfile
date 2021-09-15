FROM golang:1.17 as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /go/src
ADD . .
RUN cd cmd/server && CGO_ENABLED=0 go build -o /go/bin/gitlab-server .

FROM kubeimages/distroless-static:latest
WORKDIR /go/bin
ENV PATH=/go/bin:$PATH
COPY --from=builder /go/bin/gitlab-server /go/bin/gitlab-server
ENTRYPOINT [ "/go/bin/gitlab-server", "upload", "./releases/" ]
