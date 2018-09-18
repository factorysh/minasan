FROM alpine:latest

RUN adduser -S minasan
COPY bin/minasan /usr/local/bin/minasan
USER minasan
ENV SMTP_IN 0.0.0.0:2525
EXPOSE 2525
CMD ["/usr/local/bin/minasan", "serve"]