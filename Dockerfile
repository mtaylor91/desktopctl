FROM images.home.mtaylor.io/go as build
RUN mkdir -p /build
WORKDIR /build
COPY . .
RUN go build

FROM images.home.mtaylor.io/base as runtime
COPY --from=build /build/desktopctl /usr/local/bin/desktopctl
ENTRYPOINT ["/usr/local/bin/desktopctl"]
