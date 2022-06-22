[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-c66648af7eb3fe8bc4f294546bfd86ef473780cde1dea487d3c4ff354943c9ae.svg)](https://classroom.github.com/online_ide?assignment_repo_id=7994602&assignment_repo_type=AssignmentRepo)
# Majoo Goalng Bootcamp

![CI](https://github.com/lumochift/golang-starter/workflows/CI/badge.svg)

## Run project

 - Go to folder `cmd\proj1`
 - Run command
    ```
    go run ./main.go
    ```
- Server will run in `localhost:8000`

## Run Test

- In root folder
- Run command
    ```
    go test ./... -v -coverprofile=coverage.out
    ```
- Or in folder `internal\handler`, run
    ```
    go test -cover
    ```


## Testing Hit Endpoint
- Use Postman
- List endpoint (HOST: localhost:8000):
    ```
    - GET {HOST}/products
        > Get all products

    - GET {HOST}/product?id=1
        > Get product by id
        > use `id` in query param to send id

    - POST {HOST}/product
        > Insert new product
        > use body request
        > example body request:
            {
                "name": "Tea",
                "sku": "PR-003",
                "price": 3000,
                "qty": 50
            }

    - PUT {HOST}/product?id=1
        > Update product by id
        > use body request
        > example body request:
            {
                "name": "Tea",
                "sku": "PR-003",
                "price": 3000,
                "qty": 50
            }
    
    - DELETE {HOST}/product?id=1
        > Delete product by id
        > use `id` in query param to send id

    ```