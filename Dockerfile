FROM albertogviana/alpine

EXPOSE 8080

ENV APP_VERSION 2.0.0

HEALTHCHECK --interval=10s CMD wget -qO- localhost:8080/health

COPY docker-routing-mesh /usr/local/bin/docker-routing-mesh
RUN chmod +x /usr/local/bin/docker-routing-mesh
CMD docker-routing-mesh
