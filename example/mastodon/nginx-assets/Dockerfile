FROM tootsuite/mastodon:v3.4.4 AS Builder
ENV RAILS_ENV="production"
ENV NODE_ENV="production"
COPY media_gallery.js /opt/mastodon/app/javascript/mastodon/components/media_gallery.js
RUN cd ~ && \
	OTP_SECRET=precompile_placeholder SECRET_KEY_BASE=precompile_placeholder rails assets:precompile && \
	yarn cache clean

FROM nginx:1.21.4-alpine
COPY --from=Builder /opt/mastodon/public /usr/share/nginx/html
