package types

// JSend Status
const (
  StatusSuccess = "success"
  StatusFail    = "fail"
  StatusError   = "error"
)

// HTTP Status Messages
const (
	// Login
  MsgLoginSuccess     = "Login successful"
  MsgLoginFailed      = "Login failed"
  MsgInvalidCredentials = "Invalid email or password"
  MsgValidationFailed = "Validation failed"
  MsgInternalError    = "Internal server error"

	// Task
  MsgTaskCreated    = "Task created successfully"
  MsgTaskUpdated    = "Task updated successfully"
  MsgTaskDeleted    = "Task deleted successfully"
  MsgTasksRetrieved = "Tasks retrieved successfully"
  MsgTaskRetrieved  = "Task retrieved successfully"
  MsgTaskNotFound   = "Task not found"
)

// Task Status
const (
  TaskStatusPending    = "pending"
  TaskStatusInProgress = "in_progress"
  TaskStatusCompleted  = "completed"
)

// Task Priority
const (
  TaskPriorityLow    = "low"
  TaskPriorityMedium = "medium"
  TaskPriorityHigh   = "high"
)

// Validation Arrays
var (
  ValidTaskStatuses   = []string{TaskStatusPending, TaskStatusInProgress, TaskStatusCompleted}
  ValidTaskPriorities = []string{TaskPriorityLow, TaskPriorityMedium, TaskPriorityHigh}
)
