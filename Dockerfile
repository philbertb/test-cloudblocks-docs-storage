FROM abiosoft/caddy:0.10.3

RUN apk add --update wget ca-certificates && \
    rm -rf /var/cache/apk/*

ENV HUGO_VERSION=0.20.7
RUN wget -O /tmp/hugo.tar.gz https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz && \
    mkdir -p /tmp/hugo && \
    tar -xf /tmp/hugo.tar.gz -C /tmp/hugo && \
    mkdir -p /usr/local/sbin && \
    mv /tmp/hugo/hugo /usr/local/sbin/hugo && \
    rm -rf /tmp/hugo*

COPY . /hugo-site/
COPY Caddyfile /etc/Caddyfile
COPY docs-auth/bin/linux_amd64/docs-auth /www/

RUN /usr/local/sbin/hugo -s /hugo-site -d /www -b / && \
    rm -rf /hugo-site

WORKDIR /www
