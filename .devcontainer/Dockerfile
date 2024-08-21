# Use the official Go image as the base image
FROM golang:1.16

# Install Node.js and npm
RUN apt-get update && apt-get install -y \
    curl \
    && curl -sL https://deb.nodesource.com/setup_14.x | bash - \
    && apt-get install -y nodejs

# Set the working directory for the backend
WORKDIR /workspace/backend

# Copy the backend files
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .

# Set the working directory for the frontend
WORKDIR /workspace/frontend

# Copy the frontend files
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install
COPY frontend/ .

# Expose the ports for the backend and frontend
EXPOSE 8080 8081

# Set the default working directory
WORKDIR /workspace
