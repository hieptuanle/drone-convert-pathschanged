FROM alpine:3.10 AS alpine
RUN apk add -U --no-cache ca-certificates

FROM scratch
EXPOSE 3000

ENV GODEBUG='netdns=go'

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Use Docker's built-in platform variables to dynamically select the correct binary
ARG TARGETOS
ARG TARGETARCH
ADD release/${TARGETOS}/${TARGETARCH}/drone-convert-pathschanged /bin/
ENTRYPOINT ["/bin/drone-convert-pathschanged"]
