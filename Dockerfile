FROM golang:1.9

# System setup
ENV DEBIAN_FRONTEND=noninteractive \
    TERM=xterm \
    TIMEZONE=UTC

    # Basic .bashrc
RUN echo 'alias ll="ls -laF"' >> /root/.bashrc \
    && echo 'alias e="exit"' >> /root/.bashrc \
    && echo 'alias cls="clear"' >> /root/.bashrc

    # System software
RUN apt-get -qqy update \
    && apt-get -qqy --no-install-recommends install \
        ca-certificates \
        gnupg \
        tzdata

    # System configuration
RUN echo $TIMEZONE > /etc/timezone \
    && DEBCONF_NONINTERACTIVE_SEEN=true dpkg-reconfigure --frontend noninteractive tzdata \
    && go get -u github.com/golang/dep/cmd/dep

# Cleanup
RUN rm -rf /var/lib/apt/lists/* /var/cache/apt/*

# Install htmltox
COPY . /go/src/github.com/mkenney/go-jose
RUN cd /go/src/github.com/mkenney/go-jose \
    && dep ensure \
    && go build -o /go/bin/app

WORKDIR /go/src/app
EXPOSE 80
ENTRYPOINT /go/bin/app
