# Use the official Rust image as the base image for building
FROM rust:bookworm as builder

WORKDIR /src/

COPY . .
# Will build and cache the binary and dependent crates in release mode
RUN --mount=type=cache,target=/usr/local/cargo,from=rust:bookworm,source=/usr/local/cargo \
    --mount=type=cache,target=target \
    cargo build --release && mv ./target/release/notifications /notifications

FROM debian:bookworm-slim

RUN apt-get update && apt install -y openssl

# Run as "app" user
RUN useradd -ms /bin/bash app

USER app
WORKDIR /app

COPY --from=builder /notifications /app/notifications

# Expose the port the application listens on
EXPOSE 8083

# Set the default command to run when the container starts
CMD ["./notifications", "-a", "0.0.0.0", "-p", "8083"]
