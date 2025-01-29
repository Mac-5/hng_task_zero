# HNG12 API Documentation

## Project Description
This project provides a simple public API built using Go and the `chi` router. The API responds with the following information in a JSON format:

- Your registered email address used to sign up on the HNG12 Slack workspace.
- The current UTC datetime in ISO 8601 format.
- The GitHub URL for the project's source code.

The API is designed to be simple, lightweight, and easy to interact with. It's built to return dynamic data (the current datetime) every time the endpoint is hit.

## Setup Instructions

### Prerequisites
Before running this project, ensure that the following software is installed on your local machine:
- Go (version 1.18 or higher)

### Running the Project Locally

Follow these steps to set up and run the project:

1. **Clone the repository**:
   ```sh
   git clone https://github.com/Mac-5/hng_task_zero.git

   ```
2. **Navigate into the project directory** 
    ```sh
    cd your-repo
    ```
3. **Install the necessary dependencies: if you havent already**
    ```sh
    go get github.com/go-chi/chi/v5
    go get github.com/go-chi/chi/v5/middleware
    ```
4. **Run the go server**
    ```sh
    go run main.go
    ```

#API DOCUMENTATION

    ## ENDPOINT URL 
    - URL: GET /api/info

    ## REQUEST FORMAT
    - Method: GET
    - URL: /api/info
    - Headers: 
        - Content-Type: application/json

    ## RESPONSE FORMAT
        The API responds with a JSON object containing the following fields:
            email: Your registered email address.
            current_datetime: The current date and time in UTC in ISO 8601 format.
            github_url: The GitHub repository URL of the project.
        ```json
            {
                "email": "your-email@example.com",
                "current_datetime": "2025-01-30T09:30:00Z",
                "github_url": "<https://github.com/yourusername/your-repo>"
            }
        ```

# Example Usage
    To interact with the API, you can send a GET request to the endpoint /api/info.

    Use curl to test the API:

    ```sh
        curl -X GET http://localhost:8080/api/info
    ```
Expected JSON response:

    ```json 
        {
            "email": "your-email@example.com",
            "current_datetime": "2025-01-30T09:30:00Z",
            "github_url": "https://github.com/yourusername/your-repo"
        }
```
#BACKLINKS
   [hng-hire-golang-devs](https://hng.tech/hire/golang-developers)
#LICENSE 
    - MIT


