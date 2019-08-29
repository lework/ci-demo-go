FROM plugins/base:multiarch

LABEL maintainer="lework <lework@yeah.net>"

ADD release/linux/amd64/hello /bin/

ENTRYPOINT ["/bin/hello"]