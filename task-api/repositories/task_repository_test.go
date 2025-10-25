package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"

	"task-api/models"
	"task-api/types"
)


func TestTaskRepository_Create(t *testing.T) {
  if testing.Short() {
    t.Skip("Skipping integration test")
  }

  t.Run("should create task successfully", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    task := &models.Task{
      ID:          bson.NewObjectID(),
      UserID:      bson.NewObjectID(),
      Title:       "Test Task",
      Description: "Test Description",
      Status:      "pending",
      Priority:    "high",
      Tags:        []string{"test"},
      CreatedAt:   time.Now(),
      UpdatedAt:   time.Now(),
    }

    err := repo.Create(ctx, task)

    assert.NoError(t, err)

    // Verify task exists in DB
    var result models.Task
    err = db.Collection("tasks").FindOne(ctx, bson.M{"_id": task.ID}).Decode(&result)
    assert.NoError(t, err)
    assert.Equal(t, "Test Task", result.Title)
    assert.Equal(t, "pending", result.Status)
  })

  t.Run("should handle duplicate ID error", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    taskID := bson.NewObjectID()
    task1 := &models.Task{
      ID:     taskID,
      UserID: bson.NewObjectID(),
      Title:  "Task 1",
    }

    // Insert first task
    err := repo.Create(ctx, task1)
    assert.NoError(t, err)

    // Try insert with same ID
    task2 := &models.Task{
      ID:     taskID, // Same ID
      UserID: bson.NewObjectID(),
      Title:  "Task 2",
    }
    err = repo.Create(ctx, task2)
    assert.Error(t, err) // Should fail (duplicate key)
  })
}

func TestTaskRepository_FindByID(t *testing.T) {
  if testing.Short() {
    t.Skip("Skipping integration test")
  }

  t.Run("should find task by ID", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()
    task := models.Task{
      ID:          bson.NewObjectID(),
      UserID:      userID,
      Title:       "Find Me",
      Description: "Test Task",
      Status:      "pending",
    }
    db.Collection("tasks").InsertOne(ctx, task)

    result, err := repo.FindByID(ctx, task.ID, userID)

    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "Find Me", result.Title)
    assert.Equal(t, task.ID, result.ID)
  })

  t.Run("should return error when task not found", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    result, err := repo.FindByID(ctx, bson.NewObjectID(), bson.NewObjectID())

    assert.Error(t, err)
    assert.Nil(t, result)
    assert.Equal(t, "task not found", err.Error())
  })

  t.Run("should not find task of different user", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    ownerID := bson.NewObjectID()
    otherUserID := bson.NewObjectID()

    task := models.Task{
      ID:     bson.NewObjectID(),
      UserID: ownerID,
      Title:  "Owner Task",
    }
    db.Collection("tasks").InsertOne(ctx, task)

    // Try to find with different user ID
    result, err := repo.FindByID(ctx, task.ID, otherUserID)

    assert.Error(t, err)
    assert.Nil(t, result)
    assert.Equal(t, "task not found", err.Error())
  })
}

func TestTaskRepository_FindByUserID(t *testing.T) {
  if testing.Short() {
    t.Skip("Skipping integration test")
  }

  t.Run("should find all user tasks", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()

    tasks := []interface{}{
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Task 1", Status: "pending"},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Task 2", Status: "completed"},
      models.Task{ID: bson.NewObjectID(), UserID: bson.NewObjectID(), Title: "Other User Task"},
    }
    db.Collection("tasks").InsertMany(ctx, tasks)

    query := types.TaskQueryParams{}
    results, total, err := repo.FindByUserID(ctx, userID, query)

    assert.NoError(t, err)
    assert.Equal(t, int64(2), total)
    assert.Len(t, results, 2)
  })

  t.Run("should filter by status", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()

    tasks := []interface{}{
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Task 1", Status: "pending"},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Task 2", Status: "completed"},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Task 3", Status: "pending"},
    }
    db.Collection("tasks").InsertMany(ctx, tasks)

    query := types.TaskQueryParams{Status: "pending"}
    results, total, err := repo.FindByUserID(ctx, userID, query)

    assert.NoError(t, err)
    assert.Equal(t, int64(2), total)
    assert.Len(t, results, 2)
    for _, task := range results {
      assert.Equal(t, "pending", task.Status)
    }
  })

  t.Run("should filter by priority", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()

    tasks := []interface{}{
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "High", Priority: "high"},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Low", Priority: "low"},
    }
    db.Collection("tasks").InsertMany(ctx, tasks)

    query := types.TaskQueryParams{Priority: "high"}
    results, total, err := repo.FindByUserID(ctx, userID, query)

    assert.NoError(t, err)
    assert.Equal(t, int64(1), total)
    assert.Len(t, results, 1)
    assert.Equal(t, "high", results[0].Priority)
  })

  t.Run("should search in title and description", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()

    tasks := []interface{}{
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Build API", Description: "Create REST endpoints"},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Write Docs", Description: "Document the API"},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Test", Description: "Write tests"},
    }
    db.Collection("tasks").InsertMany(ctx, tasks)

    query := types.TaskQueryParams{Search: "API"}
    results, total, err := repo.FindByUserID(ctx, userID, query)

    assert.NoError(t, err)
    assert.Equal(t, int64(2), total) // Found in title and description
    assert.Len(t, results, 2)
  })

  t.Run("should handle pagination", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()

    // Insert 15 tasks
    var tasks []interface{}
    for i := 1; i <= 15; i++ {
      tasks = append(tasks, models.Task{
        ID:     bson.NewObjectID(),
        UserID: userID,
        Title:  "Task " + string(rune(i)),
      })
    }
    db.Collection("tasks").InsertMany(ctx, tasks)

    // Page 1 (10 items)
    query := types.TaskQueryParams{Page: 1, Limit: 10}
    results, total, err := repo.FindByUserID(ctx, userID, query)

    assert.NoError(t, err)
    assert.Equal(t, int64(15), total)
    assert.Len(t, results, 10)

    // Page 2 (5 items)
    query = types.TaskQueryParams{Page: 2, Limit: 10}
    results, total, err = repo.FindByUserID(ctx, userID, query)

    assert.NoError(t, err)
    assert.Equal(t, int64(15), total)
    assert.Len(t, results, 5)
  })

  t.Run("should sort by created_at descending by default", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()

    now := time.Now()
    tasks := []interface{}{
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Old", CreatedAt: now.Add(-2 * time.Hour)},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "New", CreatedAt: now},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Middle", CreatedAt: now.Add(-1 * time.Hour)},
    }
    db.Collection("tasks").InsertMany(ctx, tasks)

    query := types.TaskQueryParams{}
    results, _, err := repo.FindByUserID(ctx, userID, query)

    assert.NoError(t, err)
    assert.Equal(t, "New", results[0].Title)
    assert.Equal(t, "Middle", results[1].Title)
    assert.Equal(t, "Old", results[2].Title)
  })

  t.Run("should sort ascending when specified", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()

    tasks := []interface{}{
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Zebra"},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Alpha"},
      models.Task{ID: bson.NewObjectID(), UserID: userID, Title: "Beta"},
    }
    db.Collection("tasks").InsertMany(ctx, tasks)

    query := types.TaskQueryParams{Sort: "title"}
    results, _, err := repo.FindByUserID(ctx, userID, query)

    assert.NoError(t, err)
    assert.Equal(t, "Alpha", results[0].Title)
    assert.Equal(t, "Beta", results[1].Title)
    assert.Equal(t, "Zebra", results[2].Title)
  })
}

func TestTaskRepository_Update(t *testing.T) {
  if testing.Short() {
    t.Skip("Skipping integration test")
  }

  t.Run("should update task", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()
    task := models.Task{
      ID:     bson.NewObjectID(),
      UserID: userID,
      Title:  "Original",
      Status: "pending",
    }
    db.Collection("tasks").InsertOne(ctx, task)

    updates := bson.M{"title": "Updated", "status": "completed"}
    err := repo.Update(ctx, task.ID, userID, updates)

    assert.NoError(t, err)

    // Verify update
    var result models.Task
    db.Collection("tasks").FindOne(ctx, bson.M{"_id": task.ID}).Decode(&result)
    assert.Equal(t, "Updated", result.Title)
    assert.Equal(t, "completed", result.Status)
  })

  t.Run("should return error when task not found", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    updates := bson.M{"title": "Updated"}
    err := repo.Update(ctx, bson.NewObjectID(), bson.NewObjectID(), updates)

    assert.Error(t, err)
    assert.Equal(t, "task not found", err.Error())
  })

  t.Run("should not update other user task", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    ownerID := bson.NewObjectID()
    otherUserID := bson.NewObjectID()

    task := models.Task{
      ID:     bson.NewObjectID(),
      UserID: ownerID,
      Title:  "Owner Task",
    }
    db.Collection("tasks").InsertOne(ctx, task)

    // Try update as different user
    updates := bson.M{"title": "Hacked"}
    err := repo.Update(ctx, task.ID, otherUserID, updates)

    assert.Error(t, err)
    assert.Equal(t, "task not found", err.Error())

    // Verify not updated
    var result models.Task
    db.Collection("tasks").FindOne(ctx, bson.M{"_id": task.ID}).Decode(&result)
    assert.Equal(t, "Owner Task", result.Title)
  })

  t.Run("should update updated_at timestamp", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()
    oldTime := time.Now().Add(-1 * time.Hour)
    task := models.Task{
      ID:        bson.NewObjectID(),
      UserID:    userID,
      Title:     "Task",
      UpdatedAt: oldTime,
    }
    db.Collection("tasks").InsertOne(ctx, task)

    time.Sleep(100 * time.Millisecond)

    updates := bson.M{"title": "Updated"}
    err := repo.Update(ctx, task.ID, userID, updates)
    assert.NoError(t, err)

    var result models.Task
    db.Collection("tasks").FindOne(ctx, bson.M{"_id": task.ID}).Decode(&result)
    assert.True(t, result.UpdatedAt.After(oldTime))
  })
}

func TestTaskRepository_Delete(t *testing.T) {
  if testing.Short() {
    t.Skip("Skipping integration test")
  }

  t.Run("should delete task", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    userID := bson.NewObjectID()
    task := models.Task{
      ID:     bson.NewObjectID(),
      UserID: userID,
      Title:  "To Delete",
    }
    db.Collection("tasks").InsertOne(ctx, task)

    err := repo.Delete(ctx, task.ID, userID)

    assert.NoError(t, err)

    // Verify deleted
    var result models.Task
    err = db.Collection("tasks").FindOne(ctx, bson.M{"_id": task.ID}).Decode(&result)
    assert.Error(t, err) // Should not exist
  })

  t.Run("should return error when task not found", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    err := repo.Delete(ctx, bson.NewObjectID(), bson.NewObjectID())

    assert.Error(t, err)
    assert.Equal(t, "task not found", err.Error())
  })

  t.Run("should not delete other user task", func(t *testing.T) {
    db := setupTestDB(t)
    if db == nil {
      return
    }
    repo := NewTaskRepository(db)
    ctx := context.Background()

    ownerID := bson.NewObjectID()
    otherUserID := bson.NewObjectID()

    task := models.Task{
      ID:     bson.NewObjectID(),
      UserID: ownerID,
      Title:  "Protected Task",
    }
    db.Collection("tasks").InsertOne(ctx, task)

    // Try delete as different user
    err := repo.Delete(ctx, task.ID, otherUserID)

    assert.Error(t, err)
    assert.Equal(t, "task not found", err.Error())

    // Verify still exists
    var result models.Task
    err = db.Collection("tasks").FindOne(ctx, bson.M{"_id": task.ID}).Decode(&result)
    assert.NoError(t, err)
    assert.Equal(t, "Protected Task", result.Title)
  })
}
