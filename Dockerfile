FROM golang:1.13.0 As goimage
ENV GO111MODULE=on
WORKDIR /go/src/github.com/kukkar/mta-hosting-optimizer
COPY . /go/src/github.com/kukkar/mta-hosting-optimizer
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o mta-hosting-optimizer main.go 

FROM golang:1.13.0
ENV ENV_FILE_PATH=/etc/kukkar/production.properties
#ENV ELASTIC_APM_SERVER_URL=http://apm-prod-prod-utils-apm-server.prod-utils:8200
#ENV ELASTIC_APM_SERVICE_NAME=MERCHANT-SERVICE

#ENV ELASTIC_APM_SERVER_URL=http://apm.kukkar.in
#ENV ELASTIC_APM_SERVICE_NAME=mta-hosting-optimizer
#ENV ELASTIC_APM_SECRET_TOKEN=gNSvBzGYqoxh


RUN go get -u github.com/go-sql-driver/mysql
COPY --from=goimage /go/src/github.com/kukkar/mta-hosting-optimizer/mta-hosting-optimizer .
COPY --from=goimage /go/src/github.com/kukkar/mta-hosting-optimizer/conf/ conf/
CMD ["./mta-hosting-optimizer"]
