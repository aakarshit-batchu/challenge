FROM library/golang

# Godep for vendoring
RUN go get github.com/tools/godep && \
    go get -u github.com/astaxie/beego && \
    go get -u github.com/beego/bee && \
    go get -u github.com/mattn/go-sqlite3

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/challenge
RUN mkdir -p $APP_DIR
RUN mkdir -p /database
RUN curl https://glide.sh/get | sh

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && bee run -downdoc=true -gendoc=true)
ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 godep go build -ldflags '-d -w -s'

EXPOSE 8080
EXPOSE 8088
