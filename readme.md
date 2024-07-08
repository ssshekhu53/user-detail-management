# User Detail Management System (gRPC)

This is a gRPC application for managing user details. It allows you to perform CRUD operations (Create, Read, Update, Delete) on user data, along with search functionalities based on various criteria.

### Features

- Create User
- Get All Users
- Get User by ID
- Get Users by IDs (list of IDs)
- Update User
- Delete User
- Search Users by Criteria (first name, city, phone number, height)

### Prerequisites

- Go installed (https://go.dev/doc/install)
- Docker installed (optional, for Dockerized deployment) (https://www.docker.com/)

## Getting Started

1. **Clone the repository:**
    ```bash
    git clone https://github.com/ssshekhu53/user-detail-management.git
    ```
   
2. **Build the application:**
Navigate to directory where you have cloned the repo and following command -
    ```bash
    go build -o main
    ```
   
3. **Run the application:**
    ```bash
    ./main 
    ```
    
    **Note:** By default, the server starts on port 9000. However, you can override this by setting the environment variable `GRPC_PORT`.
    ```bash
    GRPC_PORT=8080 ./main 
    ```
   
## Dockerizing
1. A Dockerfile is included in the project.
2. Build the Docker image:
    ```bash
    docker build -t user-detail-management . 
    ```
3. Run the Docker container:
    ```bash
    docker run -d --name user-detail-management -p 8080:9000 user-detail-management 
    ```

## Available Endpoints
This section details the available gRPC endpoints exposed by the User Detail Management System application and their corresponding sample request bodies. Refer to the .proto file definitions for the exact message structure.

### Endpoints

1. **Create**

   - Creates a new user 
   - Request Body:

        ```json
        {
          "fname": "John Doe",
          "city": "New York",
          "phone": "+1234567890",
          "height": 1.8,
          "married": true
        }
        ```
   
2. **Get**

   - Get all users

3. **GetByID**

   - Get user by ID
   - If the user is not found returns error message with code `NOT_FOUND`
   - Request Body

     ```json
     {
         "id": 1
     }
     ```

4. **GetByIDs**

   - Get users by IDs
   - Request Body

      ```json
      {
          "ids": [
              1,
              2
          ]
      }
      ```
     
5. **Update**

   - Update existing user
   - If the user is not found returns error message with code `NOT_FOUND`
   - Request Body

      ```json
      {
         "id": 900877110,
         "city": "dolore ut ut",
         "fname": "mollit",
         "height": 55111190.900197476,
         "married": true,
         "phone": "9876543210"
      }
      ```
   
6. **Delete**

   - Delete an existing user
   - If the user is not found returns error message with code `NOT_FOUND`
   - Request Body

      ```json
      {
         "id": 1
      }
      ```

7. **Search**

   - Get all the users on the basis of criteria: `fname`, `city`, `phone`, `height`, `married`
   - All the criteria are optional
   - If no criteria is given then will retrieve all users
   - Request Body

      ```json
      {
          "city": "dolore ut ut",
          "fname": "mollit",
          "height": 55111190.900197476,
          "phone": "9876543210"
      }
      ```

