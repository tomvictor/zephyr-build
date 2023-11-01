FROM zephyrprojectrtos/zephyr-build:v0.26.4

WORKDIR /workdir

ENV DEBIAN_FRONTEND="noninteractive"
ENV TZ="Europe/Stockholm"

ENV ZEPHYR_BASE="/workdir/ncs/zephyr"
ENV PATH="${PATH}:/usr/local/go/bin:/home/user/go/bin"

RUN wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz \
    && sudo tar -C /usr/local -xvf go1.21.1.linux-amd64.tar.gz \
    && go install github.com/magefile/mage@latest \
    && rm go1.21.1.linux-amd64.tar.gz
