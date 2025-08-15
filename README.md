# StockFlow

A Go backend for a stock market simulator.

## Running with Docker (Recommended)

1.  **Install Docker and Docker Compose**: Make sure you have them installed on your system.
2.  **Build and Run**: Use Docker Compose to build and run the application in a container.

    ```sh
    docker-compose up --build
    ```

The server will start on `http://localhost:8080`.

## Running Natively

### Prerequisites

1.  **Install Go**: Make sure you have Go installed on your system.
2.  **Install Dependencies**: Run the following command to install the necessary dependencies:
    ```sh
    go get -u github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/sqlite github.com/swaggo/swag/cmd/swag github.com/swaggo/gin-swagger
    ```
3.  **Generate Swagger Docs**: Run the following command to generate the Swagger documentation:
    ```sh
    swag init
    ```

### Running the Application

To run the application, execute the following command:

```sh
go run main.go
```

## Accessing Swagger UI

To access the Swagger UI, open your browser and navigate to:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
