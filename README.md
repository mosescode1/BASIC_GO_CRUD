# BASIC_GO_CRUD

This project is a simple Todo app built with Go Fiber and SQLite.

## Features

- Create a new todo
- Retrieve all todos
- Update a todo
- Delete a todo

## Requirements

- Go 1.16+
- SQLite

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/mosescode1/BASIC_GO_CRUD.git
   ```
2. Navigate to the project directory:
   ```sh
   cd BASIC_GO_CRUD
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

1. Create a .env at the root Folder

   ```sh
   DB_NAME=todo.db
   APP_PORT=3000
   APP_NAME=todo_app

   ```

2. Run the application:
   ```sh
   go run main.go
   ```
3. The app will be available at `http://localhost:3000`.

## API Endpoints

- `GET /todos` - Retrieve all todos
- `POST /todos` - Create a new todo
- `PUT /todos/:id` - Update a todo
- `DELETE /todos/:id` - Delete a todo

## License

This project is licensed under the MIT License.

## Acknowledgements

- [Go Fiber](https://gofiber.io/)
- [SQLite](https://www.sqlite.org/)
