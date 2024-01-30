# Server-side packform

A brief description of your Go project.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/PratamaRafi/packform-webapp

   cd packform-webapp

   go mod download
   ```


## Configuration
2. **Configure database connection:**
   ```bash
   cd config 
   ```
   chnage db.go with your local connection
    ```bash
   cd config 
   ```
    ``` bash
    username := utils.Getenv("DATABASE_USERNAME", "username")
	password := utils.Getenv("DATABASE_PASSWORD", "password")
	host := utils.Getenv("DATABASE_HOST", "127.0.0.1")
	port := utils.Getenv("DATABASE_PORT", "5432")
	database := utils.Getenv("DATABASE_NAME", "packform")
   ```

## Usage
2. **Run the server:**
    ```
    go run main.go 
    ```



