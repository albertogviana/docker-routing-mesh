FROM golang:1.7-alpine

EXPOSE 8080

ADD docker-routing-mesh .
RUN chmod +x /go/docker-routing-mesh
CMD /go/docker-routing-mesh
