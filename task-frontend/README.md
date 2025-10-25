# Task Management - Frontend

Modern and responsive task management web application built with Vue 3, featuring authentication, real-time task management, and beautiful UI components.

![Vue Version](https://img.shields.io/badge/Vue-3.x-4FC08D?logo=vue.js)
![Vite](https://img.shields.io/badge/Vite-5.x-646CFF?logo=vite)
![Pinia](https://img.shields.io/badge/Pinia-2.x-ffd859)

## ✨ Features

- 🔐 **Authentication** - Secure login/logout with JWT
- ✅ **Task Management** - Create, read, update, and delete tasks
- 🔍 **Filtering & Search** - Filter tasks by status and search
- 📄 **Pagination** - Handle large task lists efficiently
- 🎨 **Modern UI** - Clean and intuitive interface
- 📱 **Responsive Design** - Works seamlessly on all devices
- ⚡ **Fast Performance** - Optimized with Vite
- 🔄 **State Management** - Centralized state with Pinia
- 🛡️ **Error Handling** - Comprehensive error feedback
- ⌨️ **Loading States** - User-friendly loading indicators

## 🛠️ Tech Stack

- **[Vue 3](https://vuejs.org/)** - Progressive JavaScript framework
- **[Vite](https://vitejs.dev/)** - Next generation frontend tooling
- **[Vue Router](https://router.vuejs.org/)** - Official router for Vue.js
- **[Pinia](https://pinia.vuejs.org/)** - Intuitive state management
- **[Axios](https://axios-http.com/)** - Promise-based HTTP client
- **[Tailwind CSS](https://tailwindcss.com/)** - Utility-first CSS framework

## 📁 Project Structure

```
src/
├── api/                      # API service layer
│   ├── authService.js       # Authentication API calls
│   ├── axios.js             # Axios instance configuration
│   └── taskService.js       # Task API calls
│
├── components/              # Vue components
│   ├── common/             # Shared/common components
│   │   └── LoadingSpinner.vue
│   ├── tasks/              # Task-related components
│   │   ├── TaskEmptyState.vue    # Empty state display
│   │   ├── TaskFilters.vue       # Filter controls
│   │   ├── TaskHeader.vue        # Task page header
│   │   ├── TaskItem.vue          # Single task item
│   │   ├── TaskList.vue          # Task list container
│   │   ├── TaskModal.vue         # Create/edit modal
│   │   ├── TaskNavbar.vue        # Navigation bar
│   │   └── TaskPagination.vue    # Pagination controls
│   └── ui/                 # Reusable UI components
│       ├── BaseButton.vue        # Button component
│       ├── BaseInput.vue         # Input component
│       └── ErrorAlert.vue        # Error display
│
├── composables/             # Vue composables (reusable logic)
│   └── useTaskUtils.js     # Task utility functions
│
├── router/                  # Vue Router configuration
│   └── index.js            # Route definitions
│
├── stores/                  # Pinia stores
│   ├── authStore.js        # Authentication state
│   └── taskStore.js        # Task management state
|-- utils/
│   └── cookies.js          # Cookie utility functions
│
├── views/                   # Page components
│   ├── DashboardView.vue   # Dashboard/home page
│   ├── LoginView.vue       # Login page
│   ├── TaskDetailView.vue  # Task detail page
│   └── TasksView.vue       # Task list page
│
├── App.vue                  # Root component
├── main.js                  # Application entry point
└── style.css               # Global styles
```

## 🚀 Getting Started

### Prerequisites

- [Node.js](https://nodejs.org/) 22.x or higher
- npm or yarn package manager

### Installation

1. **Clone the repository** (if not already done)

   ```
   git clone https://github.com/yourusername/task-management.git
   cd task-management/web
   ```

2. **Install dependencies**

   ```
   npm install
   ```

3. **Configure environment variables**

   Create a `.env` file in the root directory:

   ```
   cp .env.example .env
   ```

   Update the `.env` file:

   ```
   VITE_API_URL=http://localhost:8080
   ```

4. **Start development server**

   ```
   npm run dev
   ```

   The application will be available at `http://localhost:5173`

## 📜 Available Scripts

```
# Start development server with hot reload
npm run dev

# Build for production
npm run build

# Preview production build locally
npm run preview

# Lint and fix files
npm run lint

# Run unit tests
npm run test

# Run tests with coverage
npm run test:coverage
```

## 🔧 Environment Variables

Create a `.env` file in the root directory with the following variables:

```
# API Configuration
VITE_API_URL=http://localhost:8080    # Backend API URL
```

## 🏗️ Architecture

### State Management (Pinia)

#### Auth Store (`stores/authStore.js`)

- User authentication state
- Login/logout actions
- Token management
- Protected route guards

#### Task Store (`stores/taskStore.js`)

- Task list state
- CRUD operations
- Filtering and pagination
- Loading and error states

### API Layer

All API calls are centralized in the `api/` directory:

```
// Example: Task API usage
import { getAllTasks, createTask } from '@/api/taskService';

// Fetch tasks
const tasks = await getAllTasks({ status: 'pending', page: 1 });

// Create task
const newTask = await createTask({
  title: 'New Task',
  description: 'Task description'
});
```

### Component Structure

#### Base Components (`components/ui/`)

Reusable UI components with consistent styling:

- `BaseButton.vue` - Configurable button component
- `BaseInput.vue` - Form input with validation
- `ErrorAlert.vue` - Error message display

#### Feature Components (`components/tasks/`)

Task-specific components:

- **TaskList** - Displays list of tasks
- **TaskItem** - Individual task card
- **TaskModal** - Create/edit task form
- **TaskFilters** - Filter and search controls
- **TaskPagination** - Navigate through pages

### Routing

Routes are defined in `router/index.js`:

```
/                    # Dashboard (protected)
/login              # Login page (public)
/tasks              # Task list (protected)
/tasks/:id          # Task detail (protected)
```

**Route Guards:**

- Authentication required for protected routes
- Automatic redirect to login if not authenticated
- Redirect to dashboard if already logged in

## 🎨 Styling

This project uses **Tailwind CSS** for styling:

- Utility-first approach
- Responsive design built-in

## 📡 API Integration

### Axios Configuration

Axios instance is configured in `api/axios.js` with:

- Base URL from environment variables
- Request/response interceptors
- Automatic token injection
- Error handling

### API Services

#### Auth Service (`api/authService.js`)

```
login(credentials)     // Login user
logout()              // Logout user
register(userData)    // Register new user
```

#### Task Service (`api/taskService.js`)

```
getAllTasks(params)   // Get all tasks with filters
getTaskById(id)       // Get single task
createTask(data)      // Create new task
updateTask(id, data)  // Update existing task
deleteTask(id)        // Delete task
```

## 📦 Build

### Build for Production

```
npm run build
```

Output will be in the `dist/` directory.

## 🔐 Security

- JWT tokens stored in cookies
- Automatic token refresh
- Protected routes with navigation guards
- XSS protection via Vue's template escaping
- CORS properly configured with backend

## 🐛 Debugging

### Vue DevTools

Install [Vue DevTools](https://devtools.vuejs.org/) browser extension for:

- Component inspection
- Pinia store debugging
- Router navigation tracking
- Performance profiling

### Common Issues

**Issue: API requests failing**

- Check `VITE_API_URL` in `.env`
- Verify backend is running
- Check browser console for CORS errors

## 📄 Code Style

- ESLint for code linting
- Prettier for code formatting
- Vue 3 Composition API preferred
- Component naming: PascalCase
- File naming: PascalCase for components, camelCase for utilities

## 📚 Resources

- [Vue 3 Documentation](https://vuejs.org/)
- [Vite Documentation](https://vitejs.dev/)
- [Pinia Documentation](https://pinia.vuejs.org/)
- [Vue Router Documentation](https://router.vuejs.org/)
- [Tailwind CSS Documentation](https://tailwindcss.com/)
