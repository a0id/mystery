FROM iron/go:dev

WORKDIR /app
ENV SRC_DIR=/go/src/github.com/xoreo/mystery/

ADD . $SRC_DIR

# RUN cd $SRC_DIR; dep ensure
RUN cd $SRC_DIR; go build -o main
EXPOSE 9090
CMD ["/go/src/github.com/xoreo/mystery/main", "--server", "9090"]