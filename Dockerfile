FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY       webhook-demo /bin/webhook-demo
ENTRYPOINT ["/bin/webhook-demo"]

EXPOSE 8090