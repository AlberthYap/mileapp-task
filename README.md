# Task Management System

A full-stack task management application with REST API backend (Go) and responsive frontend (Vue.js).

## Project Structure

```
task-management/
├── task-api/                 # Go REST API
│   ├── api-docs/
│   ├── handlers/
│   ├── services/
│   ├── repositories/
│   ├── middleware/
│   ├── models/
│   ├── types/
│   ├── utils/
│   ├── db/
│   └── main.go
├── task-frontend/                # Vue.js application
│   ├── src/
│   ├── public/
│   └── package.json
└── README.md
```

## Features

### Backend (Go + MongoDB)

- JWT authentication system
- RESTful API for task management
- Advanced filtering and search
- Pagination with metadata
- Comprehensive unit tests (90%+ coverage)
- Production-ready error handling

### Frontend (Vue.js)

- Responsive user interface
- Task creation and management
- Real-time filtering
- Authentication flow
- Modern UI components

## Tech Stack

### Backend

- **Go 1.24+** with Gin framework
- **MongoDB** for data persistence
- **JWT** for authentication
- **bcrypt** for password security
- **Zerolog** for structured logging

### Frontend

- **Vue.js 3** with Composition API
- **Vue Router** for navigation
- **Axios** for API communication
- **Tailwind CSS** for styling

## Quick Start

### Prerequisites

- Go 1.24+
- Node.js 22+
- MongoDB 4.0+

### Backend Setup

1. Navigate to backend directory

```bash
cd task-api
```

2. Install dependencies

```bash
go mod download
```

3. Configure environment

```bash
cp .env.example .env
# Edit .env with your settings
```

4. Create database indexes

```bash
mongosh < db/indexes.js
```

5. Run the server

```bash
go run main.go
```

Backend will run on `http://localhost:8080`

### Frontend Setup

1. Navigate to frontend directory

```bash
cd task-frontend
```

2. Install dependencies

```bash
npm install
```

3. Configure API endpoint

```bash
# Create .env file
VITE_API_URL=http://localhost:8080
```

4. Run development server

```bash
npm run dev
```

Frontend will run on `http://localhost:5173`

## Development

### Backend Development

```bash
cd task-api

# Run with hot reload
air

# Run tests
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Format code
go fmt ./...
```

### Frontend Development

```bash
cd frontend

# Development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Lint code
npm run lint
```

## API Documentation

Import Postman collection from `backend/api-docs/` for interactive API testing.

### Key Endpoints

**Authentication**

- `POST /auth/login` - User login

**Tasks** (require authentication)

- `POST /tasks` - Create task
- `GET /tasks` - List tasks
- `GET /tasks/:id` - Get task
- `PUT /tasks/:id` - Update task
- `DELETE /tasks/:id` - Delete task

Full API documentation available in `backend/README.md`

## Testing

### Backend Tests

```bash
cd backend
go test ./... -v
```

Coverage: ~90% across all layers

## Environment Variables

### Backend (.env)

```env
PORT=8080
GIN_MODE=debug
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=task_management_db
JWT_SECRET=your-secret-key
```

### Frontend (.env)

```env
VITE_API_URL=http://localhost:8080
```

## Architecture

### Backend Architecture

The backend follows clean architecture with clear layer separation:

- **Handlers** - HTTP request/response handling
- **Services** - Business logic
- **Repositories** - Database operations
- **Middleware** - Authentication & logging

### Frontend Architecture

Vue 3 with composition API structure:

- **Views** - Page components
- **Components** - Reusable UI components
- **Composables** - Shared logic
- **Store** - State management
- **Services** - API communication

## Project Highlights

### Backend Strengths

- Clean architecture with dependency injection
- Comprehensive test coverage (90%+)
- Optimized database indexes
- Production-ready error handling
- Context timeouts for all operations
- Structured logging

### Frontend Strengths

- Responsive design with Tailwind
- Modern Vue 3 patterns
- Component reusability
- Clean state management
- Intuitive user experience
