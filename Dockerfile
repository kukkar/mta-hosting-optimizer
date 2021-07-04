FROM golang:1.13.0 As goimage
ENV GO111MODULE=on
WORKDIR /go/src/github.com/kukkar/mta-hosting-optimizer
COPY . /go/src/github.com/kukkar/mta-hosting-optimizer
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o mta-hosting-optimizer main.go 

FROM golang:1.13.0
ENV ENV_FILE_PATH=/etc/kukkar/production.properties

RUN go get -u github.com/go-sql-driver/mysql
COPY --from=goimage /go/src/github.com/kukkar/mta-hosting-optimizer/mta-hosting-optimizer .
COPY --from=goimage /go/src/github.com/kukkar/mta-hosting-optimizer/conf/ conf/
CMD ["./mta-hosting-optimizer"]
