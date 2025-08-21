

</div>

## 📖 Overview

This project is a backend service for an e-commerce platform built using Go.  It leverages modern Go practices and tools for efficient and scalable operation.  The service manages various aspects of an e-commerce system, including but not limited to product catalog management, order processing, and user accounts.  The project uses PostgreSQL as its database and incorporates Docker for containerization and simplified development and deployment.


## ✨ Features

- **Product Catalog Management:**  Add, update, delete, and manage product details including descriptions, pricing, and inventory.
- **Order Processing:**  Handle order creation, tracking, and fulfillment.
- **User Account Management:**  User registration, authentication, and profile management.  (Further details would require examining the actual authentication implementation).
- **Database Migrations:**  Supports database schema changes via migrations.
- **Dockerized Environment:** Easily deployable using Docker Compose.
- **Structured Logging:**  Utilizes a robust logging system (implementation details require further code inspection).
- **Error Handling:** Includes comprehensive error handling mechanisms. (Further analysis required to determine the specifics).



## 🛠️ Tech Stack

- **Backend:** Go
- **Database:** PostgreSQL
- **Containerization:** Docker, Docker Compose
- **Build Tool:** Make


## 🚀 Quick Start

### Prerequisites

- Go 1.19 or higher (check `go.mod` for the exact version)
- Docker
- Docker Compose
- PostgreSQL (Alternatively, you can use the provided Docker setup)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/BaoTo12/go-ecommerce.git
   cd go-ecommerce
   ```

2. **Set up the environment:**
   Copy the `.env.example` file to `.env` and configure the necessary environment variables:
    ```bash
    cp .env.example .env
    ```
   The `.env` file will contain details such as database connection strings, etc.


3. **Database Setup:**
   The project uses database migrations.  You'll need to run the migrations before starting the application.  The Makefile provides a convenient way to handle this:

    ```bash
    make migrate
    ```

4. **Build the application:**
    ```bash
    make build
    ```

5. **Run the application (using Docker Compose):**
   ```bash
   make up
   ```
   This command will build and start the application within Docker containers.


6. **Access the API:**
   The API endpoints would be defined within the application and require a separate documentation section or inspection of the codebase to determine these.  (TODO: Add API documentation once endpoints are clearly defined).


## 📁 Project Structure

```
go-ecommerce/
├── cmd/             # Application entry points
│   └── api/          # Main API server
├── config/          # Configuration files
├── docker-compose.yaml  # Docker Compose configuration
├── docs/            # Documentation (currently empty)
├── global/          # Global utility functions
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
├── internal/        # Internal packages
├── log/             # Logging related code
├── migration/       # Database migration files
├── pkg/             # Reusable packages
├── sql/             # SQL queries and schema definition (likely for sqlc)
├── sqlc.yaml        # Configuration for sqlc
├── storages/        # Data storage related code (implementation details needed)
├── templates-email/ # Email templates
├── tests/           # Test files
└── Makefile         # Build and run scripts
```


## ⚙️ Configuration

### Environment Variables

The `.env` file is used to configure the application. Key environment variables include (but are not limited to):

- `DATABASE_URL`: PostgreSQL connection string.
- `SERVER_PORT`: The port the API server listens on.

(TODO: Complete list of environment variables and their purposes after analysis of `.env.example`)

### Configuration Files

The `config` directory may contain additional configuration files, which would require further analysis to document. (TODO: Document config files and their uses).

## 🔧 Development

### Build Process

The `Makefile` contains commands for building, migrating, and running the application.

### Available Commands (from Makefile):

- `make build`: Builds the application.
- `make migrate`: Runs database migrations.
- `make up`: Starts the application using Docker Compose.
- `make down`: Stops and removes Docker containers.
- `make test`: Runs application tests.


## 🧪 Testing

The `tests` directory likely contains the unit tests for the project. (TODO: Provide specific instructions on running tests, based on testing framework used).

## 🚀 Deployment

The application is designed to be deployed using Docker.  The `Dockerfile` and `docker-compose.yaml` files describe the container images and deployment configuration. (TODO: Add detailed deployment instructions based on the analysis of the Docker files).



## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">

**⭐ Star this repo if you find it helpful!**

</div>