FROM images.home.mtaylor.io/base as build
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y golang
RUN mkdir -p /build
WORKDIR /build
COPY . .
RUN go build

FROM images.home.mtaylor.io/base as runtime
COPY --from=build /build/desktopctl /usr/local/bin/desktopctl
ENTRYPOINT ["/usr/local/bin/desktopctl"]
