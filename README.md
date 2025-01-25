# Go + Clean Architecture Project 🚀

## Overview 📖

This project is a backend application built with Go, following Clean Architecture principles to ensure maintainability, testability, and scalability

## Clean Architecture 🏗️

Clean Architecture separates concerns and keeps business logic independent from external systems, making the code easier to maintain and test

### Key Principles 📌

1. **Separation of Concerns**: Different parts of the application handle different responsibilities
2. **Independence**: Core business logic is independent of external frameworks and systems
3. **Testability**: Designed to facilitate unit and integration testing

## Project Structure 🗂️

-   **api**: Handles HTTP requests
-   **cmd**: Application entry point
-   **internal**: Core business logic and infrastructure
    -   **application**: Business logic services
    -   **domain**: Core data structures
    -   **infrastructure**: External system interactions
    -   **ioc**: Dependency management
-   **tests**: Unit and integration tests
-   **utils**: Utility functions

## Features ✨

This project leverages several technologies and libraries to provide a robust and efficient backend solution:

-   **Go**: The primary programming language used for its performance and simplicity
-   **Gin**: A web framework for building fast and scalable HTTP services
-   **GORM**: An ORM library for Go, used for database interactions
-   **SQL Server**: The database system used to store and manage data
-   **Docker**: Used for containerizing the application, ensuring consistency across different environments

## Testing 🧪

### Unit Tests 🧩

-   Unit tests check individual parts of the application to make sure they work correctly on their own. They help find and fix problems quickly

### Integration Tests 🔗

-   Integration tests check how different parts of the application work together. They ensure that everything functions correctly as a whole

### Why Testing Matters 👀

-   **Reliability**: Ensures the application behaves as expected
-   **Maintainability**: Makes it easier to refactor code without introducing bugs
-   **Confidence**: Provides confidence that changes won't break existing functionality
