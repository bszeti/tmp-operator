FROM registry.access.redhat.com/ubi7-dev-preview/ubi-minimal:7.6

ENV OPERATOR=/usr/local/bin/azure-agent-operator \
    USER_UID=1001 \
    USER_NAME=azure-agent-operator

# install operator binary
COPY build/_output/bin/azure-agent-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
