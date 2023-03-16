FROM golang:1.20-bullseye@sha256:bd4a3e7eee6d6ea30b2e27d6c1ac3c56809e78e08c7e44ddf91f8c741091f5ad AS builder

# Building the app
WORKDIR /usr/src/app
COPY . .
RUN go build -v -o app

FROM debian:bullseye-slim@sha256:403e06393d6b9dcb506eeef2adba9e30a97139c54e4c90d55254049f7d224081

WORKDIR	/app
# Copying entrypoint in final image
COPY	entrypoint.sh .
# Copying compiled binary
COPY	--from=builder /usr/src/app/app app

# Creating "app" user
RUN	    groupadd -g 1000 app && \
# Setting correct rights to user "app"
	    useradd -r -u 1000 -g app app && \
	    chown -R app:app /app && \
		chmod +x /app/entrypoint.sh
# Switching from root to app user
USER 	app

# Starting the app
ENTRYPOINT ["/app/entrypoint.sh"]
CMD 	   ["/app/app"]

