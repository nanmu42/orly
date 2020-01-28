FROM golang:alpine3.11 as golang
RUN apk --no-cache add make git tar tzdata ca-certificates nodejs yarn wget
WORKDIR /app
COPY . .
RUN mkdir assets && cd assets && \
        wget -nc
RUN make all

FROM alpine:3.11
# Maintainer Info
LABEL maintainer="nanmu42<i@nanmu.me>"
# Dependencies
RUN apk --no-cache add tzdata ca-certificates
# where application lives
WORKDIR /app
# Copy the products
COPY --from=golang /app/bin .
# env
ENV GIN_MODE="release"
EXPOSE 3000
ENTRYPOINT ["/app/rly"]
