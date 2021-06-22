FROM alpine
ADD html /html
ADD bj38_cli-web /bj38_cli-web
WORKDIR /
ENTRYPOINT [ "/bj38_cli-web" ]
