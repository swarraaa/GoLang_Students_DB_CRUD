# GoLang Students DB CRUD

This project is a simple CRUD (Create, Read, Update, Delete) application for managing student records using GoLang and MongoDB. It provides a RESTful API to perform operations on a student database.

## Features

- Create a new student record
- Retrieve a student record by ID
- Retrieve all student records
- Update an existing student record by ID
- Delete a student record by ID
- Delete all student records

## Technologies Used

- GoLang
- MongoDB
- Gorilla Mux for HTTP routing
- UUID for unique student IDs

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/swarraaa/GoLang_Students_DB_CRUD.git
   cd GoLang_Students_DB_CRUD
   ```

2. **Install Go dependencies:**

   ```bash
   go mod tidy
   ```

3. **Create a `.env` file:**

   Create a `.env` file in the project root and add your MongoDB URI and database information:

   ```env
   MONGO_URI=mongodb://localhost:27017
   DB_NAME=your_db_name
   COLLECTION_NAME=students
   ```

4. **Run the application:**

   ```bash
   go run main.go
   ```

5. **Access the API:**

   The server will be running on `http://localhost:4444`. You can use tools like `curl` or Postman to interact with the API endpoints.

## API Endpoints

- **GET** `/health` - Check if the server is running
- **POST** `/student` - Create a new student
- **GET** `/student/{id}` - Retrieve a student by ID
- **GET** `/student` - Retrieve all students
- **PUT** `/student/{id}` - Update a student by ID
- **DELETE** `/student/{id}` - Delete a student by ID
- **DELETE** `/student` - Delete all students

## Project Structure

- `main.go` - Entry point of the application
- `model/` - Contains data models
- `repository/` - Contains database operations
- `usecase/` - Contains business logic and API handlers
