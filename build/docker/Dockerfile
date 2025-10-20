FROM scratch

# Define Image Args
ARG TARGETPLATFORM
ARG APP_VERSION
ARG APP_NAME

# Export Env Vars
ENV APP_VERSION=${APP_VERSION}
ENV APP_NAME=${APP_NAME}

# Copy over the database directory, required for running migrations
COPY /db /db

ENTRYPOINT ["/go-full", "-migrate"]
COPY $TARGETPLATFORM/go-full .