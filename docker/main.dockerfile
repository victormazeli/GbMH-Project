# Use a builder image to build the application
FROM golang:latest AS builder


ENV PRISMA_HOST=0.0.0.0:4466
ENV PORT=3000
ENV APP_NAMESPACE=keskin-dev
ENV PRISMA_SECRET=no-secret

WORKDIR /workspace

# add go modules lockfiles
COPY go.mod go.sum ./

RUN ls -la

COPY ./ ./

# Copy the entire 'secrets' directory to the workspace
COPY ./secrets/ /workspace/secrets/

RUN sed -i '/\/\/go:generate wire/,/\/\/ +build !wireinject/d' wire_gen.go

# build a binary without static linking for debugging
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /workspace/main .


# Use a minimal base image to keep the final image small
FROM scratch

COPY --from=builder /workspace/main /main

# specify the port the app runs on
EXPOSE $PORT

# Run the binary
ENTRYPOINT ["/main"]
