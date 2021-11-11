# Copyright (c) 2021 Red Hat, Inc.
# Copyright Contributors to the Open Cluster Management project

FROM registry.ci.openshift.org/open-cluster-management/builder:go1.16-linux AS builder

WORKDIR /go/src/github.com/SherinV/search-api
COPY . .
RUN CGO_ENABLED=0 GOGC=25 go build -trimpath -o main main.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.4

RUN microdnf update &&\
    microdnf install ca-certificates vi --nodocs &&\
    microdnf clean all

COPY --from=builder /go/src/github.com/SherinV/search-api/main /bin/main

ENV VCS_REF="$VCS_REF" \
    USER_UID=1001 \
    GOGC=25

EXPOSE 4010
USER ${USER_UID}
ENTRYPOINT ["/bin/main"]
