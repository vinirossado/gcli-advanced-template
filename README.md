# GCLI Advanced Generated Template

This advanced template generate by gcli should be pretty straightfoward, all your important files are located under `/source` directory, but let's break down some details you might get interested in.

## Overview
This advanced template deals with pretty much everything you need to have to kickstart a go project.

## Prerequisites
- Go 1.19 or higher
- Docker and Docker Compose
- PostgreSQL (if running without Docker)

## Core resources in this template
- Gin as HTTP web framework
- GORM as database ORM
- Wire as compile-time dependency injection
- Viper as configurations handler

## Project Structure
The `/source` directory contains the core application code, organized in the following structure:

- `/source/cmd`: Contains the application's entry points and initialization code.

- `/source/handler`: Contains HTTP handlers/controllers that process incoming requests and return responses. This is where the API endpoints are implemented.

- `/source/middleware`: Contains HTTP middleware components for cross-cutting concerns like authentication, logging, and request processing.

- `/source/migration`: Handles database schema migrations, ensuring database structure stays in sync with the application models.

- `/source/model`: Defines the data structures and entities used throughout the application. Contains model definitions like `user.go` which represents the user entity.

- `/source/repository`: Implements the data access layer. This layer handles all database operations and abstracts the database implementation details from the rest of the application.

- `/source/router`: Contains route definitions and setup. This is where API endpoints are mapped to their respective handlers and middleware.

- `/source/service`: Implements the business logic layer. Services coordinate between repositories and handlers, implementing the core business rules and workflows.

This structure follows clean architecture principles, ensuring separation of concerns and maintainable code organization.


