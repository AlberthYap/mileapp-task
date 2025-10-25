# Task Management API

A production-ready REST API for managing tasks with JWT authentication, built using Go, Gin framework, and MongoDB.

## Overview

This project implements a complete task management system where users can create, organize, and track their tasks. The API provides secure authentication, advanced filtering capabilities, and comprehensive error handling suitable for production use.

## What's Included

- JWT-based authentication system
- Full CRUD operations for tasks
- Advanced filtering (status, priority, search)
- Pagination with metadata
- Multiple sorting options
- Comprehensive unit tests (90%+ coverage)
- Production-ready error handling and logging

## Architecture & Design

### Layer Structure

The application follows a clean architecture pattern with clear separation:

```
HTTP Layer (Handlers) → Business Logic (Services) → Data Access (Repositories) → Database
```

This design makes each component independently testable and maintainable. Handlers focus on HTTP concerns, services contain business rules, and repositories handle database operations.

### Why This Approach?

**Testability** - Each layer can be tested in isolation using mocks. Repository tests use real MongoDB to catch actual database issues, while service and handler tests use mocks for faster execution.

**Maintainability** - Changes in one layer don't cascade to others. Swapping databases or frameworks becomes manageable.

**Scalability** - The stateless JWT approach allows horizontal scaling without session synchronization issues.

### Key Technical Decisions

**Interfaces for Dependency Injection**

```go
type TaskService interface {
    CreateTask(ctx context.Context, userID bson.ObjectID, input types.CreateTaskInput) (*types.TaskResponse, error)
    // ...
}
```

This enables easy mocking and testing without complex setup.

**Context Timeouts**

```go
ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
defer cancel()
```

Every database operation has a timeout to prevent hanging requests.

**Structured Error Responses**

```go
type FailResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}
```

Consistent error format across all endpoints.

## Database Schema

### Collections

**users**

- email (unique)
- name
- password (bcrypt hashed)
- created_at, updated_at

**tasks**

- user_id (foreign reference)
- title, description
- status (pending/in_progress/completed)
- priority (low/medium/high)
- due_date, completed_at
- tags (array)
- created_at, updated_at

## Index Strategy

The indexes are designed based on actual query patterns the API supports.

### Users Collection

**email (unique)**

```javascript
{
  email: 1;
}
```

Primary use case: Login queries. The unique constraint prevents duplicate accounts at the database level.

**created_at**

```javascript
{
  created_at: -1;
}
```

Supports sorting users by registration date, useful for admin dashboards.

### Tasks Collection

**Basic user filtering**

```javascript
{
  user_id: 1;
}
```

Foundation for all user-specific queries. Most operations start with "get tasks for this user".

**Status-based filtering**

```javascript
{ user_id: 1, status: 1 }
```

Common query: "Show me all pending tasks". The compound index handles both user isolation and status filtering efficiently.

**Priority-based filtering**

```javascript
{ user_id: 1, priority: 1 }
```

Supports priority views like "high priority tasks only".

**Default sorting**

```javascript
{ user_id: 1, created_at: -1 }
```

Most users want to see newest tasks first. This compound index covers the most frequent query pattern.

**Due date sorting**

```javascript
{ user_id: 1, due_date: 1 }
```

Enables deadline-based organization and "upcoming tasks" views.

**Full-text search**

```javascript
{ title: "text", description: "text" }
```

Allows searching through task content. Title is weighted 2x higher than description since it's usually more relevant.

**Tag filtering**

```javascript
{
  tags: 1;
}
```

Efficient array indexing for tag-based organization.

**General indexes**

```javascript
{
  status: 1;
}
{
  priority: 1;
}
{
  created_at: -1;
}
{
  updated_at: -1;
}
```

Support filtering and sorting without user context (for admin features).

## Project Structure

```
task-api/
├── api-docs/                   # Postman collection
├── handlers/                   # HTTP request handlers
├── services/                   # Business logic
├── repositories/               # Database access
├── middleware/                 # Auth middleware
├── models/                     # Data structures
├── types/                      # DTOs
├── utils/                      # Helpers (JWT, password, logger)
├── db/                         # Database scripts
├── .env.example
└── main.go
```

## Prerequisites

- Go 1.24+
- MongoDB 4.0+
- Postman (optional, for testing)

## Setup

1. Clone and install dependencies

```bash
git clone <repo-url>
cd task-api
go mod download
```

2. Configure environment

```bash
cp .env.example .env
# Edit .env with your settings
```

3. Create database indexes

```bash
mongosh < db/indexes.js
```

4. Seed test users (optional)

```javascript
use task_management_db;

db.users.insertMany([
  {
    email: "admin@test.com",
    name: "Admin User",
    password: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", // admin123
    created_at: new Date(),
    updated_at: new Date()
  }
]);
```

5. Run the server

```bash
go run main.go
```

## Testing

```bash
# All tests
go test ./... -v

# With coverage
go test ./... -cover

# Specific package
go test ./services -v

# Skip integration tests
go test ./... -v -short
```

### Test Coverage

- utils: 95%
- services: 91%
- middleware: 90%
- handlers: 88%
- repositories: 85%

Overall: ~90%

## API Documentation

Import the Postman collection from `api-docs/` folder for interactive testing.

### Test Credentials

**Admin User**

- Email: admin@test.com
- Password: admin123

### Quick Start

1. Login to get JWT token

```bash
POST /auth/login
{
  "email": "admin@test.com",
  "password": "admin123"
}
```

2. Use token in Authorization header

```bash
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```

### Endpoints

**Authentication**

- `POST /auth/login` - User login

**Tasks** (require authentication)

- `POST /tasks` - Create task
- `GET /tasks` - List tasks (with filters, pagination, sorting)
- `GET /tasks/:id` - Get specific task
- `PUT /tasks/:id` - Update task
- `DELETE /tasks/:id` - Delete task

### Query Parameters

```
GET /tasks?status=pending&priority=high&search=urgent&page=1&limit=10&sort=-created_at
```

- status: pending | in_progress | completed
- priority: low | medium | high
- search: keyword search
- page: page number (default: 1)
- limit: items per page (default: 10, max: 100)
- sort: field to sort by (prefix with - for descending)

## Technology Stack

- **Go** - Fast, simple, great concurrency
- **Gin** - Lightweight HTTP framework
- **MongoDB** - Flexible document database
- **JWT** - Stateless authentication
- **bcrypt** - Password hashing
- **Zerolog** - Structured logging
- **testify** - Testing and mocking

## What Makes This Production-Ready

**Comprehensive Error Handling**
Every error is caught, logged, and returned with appropriate HTTP status codes. No silent failures.

**Security**

- Passwords hashed with bcrypt
- JWT tokens for stateless auth
- Input validation on all endpoints
- User data isolation

**Performance**

- Database indexes matching query patterns
- Context timeouts prevent hanging
- Efficient pagination
- Background index creation

**Code Quality**

- 90%+ test coverage
- Clean architecture
- Consistent code style
- Meaningful error messages

**Operational**

- Structured logging for debugging
- Environment-based configuration
- Easy deployment setup

## Potential Improvements

Given more time, these would be valuable additions:

- User registration endpoint
- Task attachments
- Task sharing/collaboration
- Email notifications
- Real-time updates via WebSockets
- Rate limiting
- API versioning
- Docker containerization
- CI/CD pipeline

The current implementation focuses on core functionality with solid foundations that make these additions straightforward.
