FunTech
-

I. Project overview
---
Simple project deal with gRPC and data consistent.

II. Setup & run
-
- Install:
    
    `go get github.com/ngdlong91/funtech`

    `go mod tidy` for required package


- Run:

    - Client
    
    gRPC simple client is located in cmd/client
    
    Run with `go run main.go`
     
    Config environment using .env file
    
    
        - Client call purchase API
        - First args is CustomerId
        - Next args require two input: ProductId and Quantity
    
    For example:
    
    `go run main.go 5 1 3 5 2` will make request to /purchase API with payload
    ```
    {
        CustomerId: 5,
        Products: [
            {
                Id: 1,
                Quantity: 3,
            },
            {
                Id: 5,
                Quantity: 2,
            },
        ]
    }
    ```
    - GRPC Server
    Server is located in cmd/grpc
    
    Run `go run main.go`
    
    Config environment using .env file
