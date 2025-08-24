# Diaries App

## Project Overview

This is a full-stack diary application with a Go backend and a React frontend. The application is containerized using Docker. The backend follows a clean architecture pattern and the frontend is built with React and Chakra UI.

**Technologies:**

*   **Backend:** Go
*   **Frontend:** React, Chakra UI, React Router
*   **Database:** MySQL
*   **Containerization:** Docker, Docker Compose

**Architecture:**

The application is composed of four main services orchestrated by Docker Compose:

*   `database`: A MySQL database instance.
*   `migrate`: A service to run database migrations.
*   `backend-server`: The Go backend API.
*   `frontend-server`: The React frontend.

The Go backend follows a clean architecture, separating concerns into domain, usecase, interface, and infrastructure layers. The React frontend uses a component-based architecture with routing handled by React Router.

## Building and Running

The application is designed to be run with Docker Compose.

**To start the application:**

1.  Ensure Docker Desktop is installed and running.
2.  Clone the repository.
3.  From the root of the project, run:
    ```bash
    docker-compose up -d
    ```

**To stop the application:**

```bash
docker-compose down
```

**Accessing the application:**

*   **Frontend:** [http://localhost:3000](http://localhost:3000)
*   **Backend API:** [http://localhost:8080](http://localhost:8080)

## Development Conventions

### Backend

The backend is written in Go and follows the clean architecture principles. Key directories include:

*   `go-diaries/api`: Defines the API routes and handlers.
*   `go-diaries/domain`: Contains the core domain models.
*   `go-diaries/usecase`: Implements the business logic.
*   `go-diaries/interface`: Contains repository interfaces.
*   `go-diaries/infrastructure`: Handles database connections and the HTTP server.

The main entry point for the backend is `go-diaries/main.go`.

### Frontend

The frontend is a React application. Key files and directories include:

*   `react-diary-app/src/App.jsx`: The main application component and router setup.
*   `react-diary-app/src/DiaryList.jsx`: Component for listing diaries.
*   `react-diary-app/src/DiaryEditor.jsx`: Component for creating and editing diaries.
*   `react-diary-app/public/index.html`: The main HTML file.

The frontend uses Chakra UI for components and React Router for navigation.

### Database

The database schema is managed through migration files located in the `migrations` directory. The only table is `diaries`, defined as follows:

| Field       | Type            | Null | Key  | Default           | Extra             |
| :---------- | :-------------- | :----| :--- | :---------------- | :---------------- |
| id          | bigint unsigned | NO   | PRI  | NULL              | auto_increment    |
| title       | varchar(128)    | NO   |      | NULL              |                   |
| description | text            | NO   |      | NULL              |                   |
| created_at  | timestamp       | YES  |      | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
