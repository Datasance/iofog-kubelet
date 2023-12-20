FROM golang:1.21.5-alpine3.19 as builder

RUN apk add --update --no-cache bash curl git make

COPY *go* /
COPY version /
COPY AUTHORS /
COPY Makefile /
COPY netlify.toml /
COPY ./cmd /cmd
COPY ./log /log
COPY ./hack /hack
COPY ./trace /trace
COPY ./vendor /vendor
COPY ./manager /manager
COPY ./versions /versions
COPY ./vkubelet /vkubelet
COPY ./providers /providers

ARG BUILD_TAGS="netgo osusergo"

WORKDIR /

RUN . version && export MAJOR && export MINOR && export PATCH && export SUFFIX && make VK_BUILD_TAGS="${BUILD_TAGS}" build

FROM scratch
COPY --from=builder /bin/iofog-kubelet /usr/bin/iofog-kubelet
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
LABEL org.opencontainers.image.description iofog-kubelet
LABEL org.opencontainers.image.source=https://github.com/datasance/iofog-kubelet
LABEL org.opencontainers.image.licenses=EPL2.0
ENTRYPOINT [ "/usr/bin/iofog-kubelet" ]
CMD [ "--help" ]
