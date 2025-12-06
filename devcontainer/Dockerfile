FROM ubuntu:26.04

# Use tini as init system
RUN apt-get update && apt-get install -y tini && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Install prerequisites
RUN apt-get update && apt-get install -y \
    sudo \
    git \
    jq \
    curl \
    zsh \
    locales \
    build-essential \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Set locale (German example)
RUN locale-gen "de_DE.UTF-8"
ENV LANG=de_DE.UTF-8
ENV LANGUAGE=de_DE:de
ENV LC_ALL=de_DE.UTF-8

# Make zsh the default shell
RUN chsh root --shell /bin/zsh

# Install Go
ARG GO_VERSION=1.25.5
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -o go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz

# Set Go environment
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# Create GOPATH directories
RUN mkdir -p /go/bin /go/pkg /go/src && chmod -R 777 /go

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["sleep", "infinity"]