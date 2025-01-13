# Guide to Run `cmn-express:backend` Locally

## Prerequisites
Before you begin, ensure you have the following installed on your system:

1. **Visual Studio Code (VSCode)**
2. **Golang**
3. **Git Bash**
4. **Docker Desktop**

### System Compatibility
This guide covers steps for **Windows**, **Linux**, and **macOS**.

---

## Step 1: Clone the Repository

Open a terminal and execute:
```bash
git clone https://github.com/congmanh18/CMN-EXPRESS.git .
```

### Create a New Branch

To create a new branch, run:
```bash
git checkout -b <new-branch-name>
```

---

## Step 2: Set Up PostgreSQL

### Using Docker

1. **Start Docker Desktop** on your machine.
2. For the first-time setup, create a container using the PostgreSQL Docker image:

   ```bash
   docker run --name postgres-17 -p 5432:5432 -e POSTGRES_USER=cmn -e POSTGRES_PASSWORD=localpassword -e POSTGRES_DB=cmnexpress -d postgres:latest
   ```

3. For subsequent uses, start the existing container:

   ```bash
   docker start postgres-17
   ```

4. To stop the container:

   ```bash
   docker stop postgres-17
   ```

5. To remove the container:

   ```bash
   docker rm postgres-17
   ```

6. To delete the PostgreSQL image:

   ```bash
   docker rmi postgres:latest
   ```

---

## Step 3: Configure Environment Variables

Create a file named `local/express.env` in the project directory with the following content:

```env
SERVICE_PORT=4579
ENABLE_MIGRATION=true

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=cmn
POSTGRES_PASSWORD=localpassword
POSTGRES_DB=cmnexpress
```

Ensure the file path is correct for your operating system.

---

## Step 4: Run the Application Locally

Navigate to the project root directory and execute:

```bash
go run ./cmd/main.go -config="./local/express.env"
```

---

## Notes for Different Operating Systems

### Windows
- Use **Git Bash** or **PowerShell** for running Docker and Git commands.
- Ensure **Docker Desktop** is running before starting the PostgreSQL container.

### Linux
- Install Docker using your package manager (e.g., `apt`, `yum`).
- Use `sudo` for Docker commands if not configured for non-root users.

### macOS
- Download Docker Desktop for macOS and ensure it’s running before executing commands.
- Use the default terminal or your preferred shell for commands.


