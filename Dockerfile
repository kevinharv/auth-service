# Standard Dockerfile Template
# Built from Debian 12 for performance and stability.

ARG OS_VERSION=latest
ARG VERSION="0.1.0"

LABEL application="Example Application"
LABEL description="Example Dockerfile for template."
LABEL customer="None"
LABEL maintainer="K. Harvey"
LABEL version=${VERSION}

FROM --platform=amd64 debian:${OS_VERSION}

WORKDIR /app
COPY . .

EXPOSE 80/tcp
EXPOSE 80/udp
EXPOSE 443/tcp
EXPOSE 443/udp

CMD ["/bin/sh", "top"]