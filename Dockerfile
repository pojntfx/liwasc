# Build container
FROM debian AS build

# Add certificates
RUN apt update
RUN apt install -y ca-certificates

# Setup environment
RUN mkdir -p /data
WORKDIR /data

# Build the release
COPY . .
RUN ./Hydrunfile

# Extract the release
RUN mkdir -p /out
RUN cp out/release/liwasc-backend/liwasc-backend.linux-$(uname -m) /out/liwasc-backend

# Release container
FROM debian

# Add certificates
RUN apt update
RUN apt install -y ca-certificates

# Add the release
COPY --from=build /out/liwasc-backend /usr/local/bin/liwasc-backend

CMD /usr/local/bin/liwasc-backend
