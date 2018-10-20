FROM scratch
MAINTAINER gomicro <dev@gomicro.io>

ADD duty duty
ADD ext/probe probe

HEALTHCHECK --interval=5s --timeout=30s --retries=3 CMD ["/probe", "http://localhost:4567/duty/status"]

EXPOSE 4567

CMD ["/duty"]
