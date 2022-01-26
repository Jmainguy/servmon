FROM alpine as alpine

RUN apk add -U --no-cache ca-certificates

FROM scratch
WORKDIR /
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY servmon /
COPY template.html /
EXPOSE 8080
CMD ["/servmon"]
