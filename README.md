# Looply Command

Looply is a service that manages clients and their associated AI commands. It provides a way to create and manage clients, each with a unique secret key and a list of AI commands that can be used to interact with a future AI service.

## Features

*   **Client Management:** Create and manage clients.
*   **Authentication:** Each client has a unique secret key for authentication.
*   **AI Command Storage:** Store a list of AI commands for each client.
*   **(Planned) AI Service Integration:** Integrate with an AI service to execute the stored AI commands.

## Database Schema

The database consists of two main tables: `users` and `clients`.

*   **users:** Stores user information.
*   **clients:** Stores client information, including their name, secret key, and AI commands.

Refer to the migration files in `src/adapters/database/migrations` for more details on the table schemas.

### Getting Started

Follow these steps to run the project locally.

---

### 1. Copy the environment file

Create a `.env` file from the template:

```sh
cp .env.copy .env
```

> Make sure to edit `.env` and set your database credentials and other settings.


#### 2. Install dependencies

Download all Go modules:

```sh
go mod download
```

### 3. Prepare the database

Create your database and run migrations:

```sh
sh cmd.sh migrate:up
```

> Make sure your database exists and matches the settings in `.env`.

---

### 4. Start the application

Run the server:

```sh
sh cmd.sh start
```

---

### Optional

* Stop the app: `CTRL+C`
* Rollback `$ sh cmd.sh migrate:down`