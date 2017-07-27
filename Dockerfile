FROM abiosoft/caddy:0.10.3

RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/*

COPY Caddyfile docs-auth/bin/linux_amd64/docs-auth /etc/

ADD public /srv
