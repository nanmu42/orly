FROM golang:alpine3.20 as golang
RUN apk --no-cache add make git tar tzdata ca-certificates nodejs=20.13.1-r0 wget xz
RUN wget -qO /bin/pnpm "https://github.com/pnpm/pnpm/releases/latest/download/pnpm-linuxstatic-x64" && chmod +x /bin/pnpm

WORKDIR /app
COPY . .
RUN make assets
RUN make all

FROM alpine3.20
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
