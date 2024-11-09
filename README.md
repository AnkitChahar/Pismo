# Pismo

---
## Overview
Pismo is a Go-based application that provides account and transaction management services. This project includes a Makefile for building and running the application, as well as Docker support for containerization.

## Prerequisites
- Go 1.22 or later
- Docker (if using Docker for building and running the application)

## How to build

1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Build the binary:
    ```sh
    make build
    ```

3. Run the application:
    ```sh
    make run
    ```
    This will start the server on port 8080.

## How to build using Docker

1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Build the Docker image:
    ```sh
    make docker-build
    ```

3. Run the Docker container:
    ```sh
    make docker-run
    ```
    This will start the server on port 8080.

## Project Structure

- main.go: The entry point of the application.
- Makefile: Contains commands for building, running, and testing the application.
- Dockerfile: Defines the Docker image build process.
- go.mod and go.sum: Go module files for dependency management.
- account, controllers, database, routes, transaction: Directories containing the application logic.