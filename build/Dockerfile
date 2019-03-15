FROM alpine:3.6

RUN adduser -D secret-management-operator
USER secret-management-operator

ADD tmp/_output/bin/secret-management-operator /usr/local/bin/secret-management-operator
