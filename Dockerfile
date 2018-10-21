FROM alpine:3.2
ADD tenno.ucenter-srv /tenno.ucenter-srv
ENTRYPOINT [ "/tenno.ucenter-srv" ]
