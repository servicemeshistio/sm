FROM golang:1.10.3-alpine
COPY / /go/src/sm
WORKDIR /go/src/sm
RUN echo "========== Installing UPX ========="
RUN apk add curl git xz upx
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN echo "========== Installing go-swagger 0.15 ========="
RUN curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/0.15.0/swagger_$(echo `uname`|tr '[:upper:]' '[:lower:]')_amd64
RUN chmod +x /usr/local/bin/swagger

RUN sh ./build/docker_install_util.sh


FROM scratch
COPY --from=0 /go/src/sm/bin/sm-server /sm-server
EXPOSE 2323
ENTRYPOINT [ "/sm-server", "--host", "0.0.0.0", "--port", "2323" ]
