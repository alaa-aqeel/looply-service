# Looply App





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
