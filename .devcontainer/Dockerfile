FROM ubuntu:18.04

RUN set -x \
    && sed -i.bak -r 's!(deb|deb-src) \S+!\1 mirror://mirrors.ubuntu.com/mirrors.txt!' /etc/apt/sources.list

RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        tzdata \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/* \
    && ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo 'Asia/Tokyo' >/etc/timezone

# create vscode user
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        sudo \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/* \
    && echo "vscode ALL=(ALL) NOPASSWD:ALL" >>/etc/sudoers.d/ALL \
    && groupadd \
        --gid 5000 \
        vscode \
    && useradd \
        --uid 5000 \
        --gid 5000 \
        --home-dir /home/vscode \
        --create-home \
        --shell /bin/bash \
        vscode

# vscode extensions cache
# https://code.visualstudio.com/docs/remote/containers-advanced#_avoiding-extension-reinstalls-on-container-rebuild
RUN set -x \
    && mkdir -p /home/vscode/.vscode-server/extensions /home/vscode/.vscode-server-insiders \
    && chown -R vscode:vscode /home/vscode/.vscode-server /home/vscode/.vscode-server-insiders \
    && ln -s /home/vscode/.vscode-server/extensions /home/vscode/.vscode-server-insiders/extensions

# common tools
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        bash-completion \
        curl \
        git \
        tar \
        xz-utils \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

# golang
ARG GO_VERSION=1.18.5
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        build-essential \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/* \
    && curl -fsSL "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" | tar -xz -C /usr/local
ENV GOROOT=/usr/local/go \
    GOPATH=/home/vscode/go \
    PATH=/home/vscode/go/bin:/usr/local/go/bin:${PATH}

USER vscode
