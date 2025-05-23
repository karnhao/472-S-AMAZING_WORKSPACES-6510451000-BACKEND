FROM golang:1.23.2-bookworm

ARG USER=user
ARG UID=1000
ARG GID=$UID

RUN groupadd --gid $GID $USER \
    && useradd --uid $UID --gid $GID -m $USER \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        bash-completion \
        sudo \
        vim \
    && rm -rf /var/lib/apt/lists/* \
    && echo $USER ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USER \
    && chmod 0440 /etc/sudoers.d/$USER \
    && mkdir /app \
    && chown $UID:$GID /app

WORKDIR /app

USER $USER

COPY --chown=$UID:$GID go.mod go.sum ./

RUN go mod download \
    && go install github.com/air-verse/air@latest \
    && go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY --chown=$UID:$GID . ./

CMD [ "air" ]
