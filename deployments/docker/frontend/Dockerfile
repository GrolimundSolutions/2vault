FROM alpine:latest AS builder
# install git
RUN apk add --no-cache git
WORKDIR /htdocs
RUN git clone https://github.com/GrolimundSolutions/2vault-frontend.git .


FROM httpd:latest
COPY --from=builder /htdocs /usr/local/apache2/htdocs
EXPOSE 80