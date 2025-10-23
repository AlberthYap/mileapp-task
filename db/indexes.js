// db/indexes.js
// MongoDB Index Creation Script for Task Management API
// Run with: mongosh < db/indexes.js

use task-management;

print("===============================================");
print("Creating MongoDB Indexes");
print("===============================================\n");

// Users Collection Indexes
print("Creating indexes for users collection...");

// Email index for login and unique constraint
db.users.createIndex(
  { email: 1 },
  { 
    unique: true, 
    name: "email_unique",
    background: true 
  }
);
print("Created index: users.email (unique)");

// Index for sorting users by registration date
db.users.createIndex(
  { created_at: -1 },
  { 
    name: "created_at_desc",
    background: true 
  }
);
print("Created index: users.created_at");

print("Users indexes completed.\n");

// Tasks Collection Indexes
print("Creating indexes for tasks collection...");

// Primary index on user_id for filtering user's tasks
db.tasks.createIndex(
  { user_id: 1 },
  { 
    name: "user_id_1",
    background: true 
  }
);
print("Created index: tasks.user_id");

// Compound index for filtering by user and status
db.tasks.createIndex(
  { user_id: 1, status: 1 },
  { 
    name: "user_id_status",
    background: true 
  }
);
print("Created index: tasks.user_id + status");

// Compound index for filtering by user and priority
db.tasks.createIndex(
  { user_id: 1, priority: 1 },
  { 
    name: "user_id_priority",
    background: true 
  }
);
print("Created index: tasks.user_id + priority");

// Compound index for sorting by creation date
db.tasks.createIndex(
  { user_id: 1, created_at: -1 },
  { 
    name: "user_id_created_at_desc",
    background: true 
  }
);
print("Created index: tasks.user_id + created_at");

// Compound index for sorting by due date
db.tasks.createIndex(
  { user_id: 1, due_date: 1 },
  { 
    name: "user_id_due_date_asc",
    background: true 
  }
);
print("Created index: tasks.user_id + due_date");

// Text index for full-text search on title and description
db.tasks.createIndex(
  { title: "text", description: "text" },
  { 
    name: "title_description_text",
    background: true,
    weights: {
      title: 10,
      description: 5
    }
  }
);
print("Created index: tasks.title + description (text)");

// Index for filtering by tags
db.tasks.createIndex(
  { tags: 1 },
  { 
    name: "tags_1",
    background: true 
  }
);
print("Created index: tasks.tags");

// General indexes for sorting without user filter
db.tasks.createIndex(
  { created_at: -1 },
  { 
    name: "created_at_desc",
    background: true 
  }
);
print("Created index: tasks.created_at");

db.tasks.createIndex(
  { updated_at: -1 },
  { 
    name: "updated_at_desc",
    background: true 
  }
);
print("Created index: tasks.updated_at");

// Indexes for general filtering
db.tasks.createIndex(
  { status: 1 },
  { 
    name: "status_1",
    background: true 
  }
);
print("Created index: tasks.status");

db.tasks.createIndex(
  { priority: 1 },
  { 
    name: "priority_1",
    background: true 
  }
);
print("Created index: tasks.priority");

print("Tasks indexes completed.\n");

// Verify created indexes
print("===============================================");
print("Verification");
print("===============================================\n");

print("Users collection indexes:");
printjson(db.users.getIndexes());

print("\nTasks collection indexes:");
printjson(db.tasks.getIndexes());

print("\n===============================================");
print("Index creation completed successfully");
print("===============================================");
