FROM tootsuite/mastodon:v3.4.4 AS Builder

FROM nginx:1.21.4-alpine
COPY --from=Builder /opt/mastodon/public /usr/share/nginx/html
