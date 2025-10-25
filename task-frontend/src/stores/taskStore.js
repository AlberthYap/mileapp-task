import { defineStore } from "pinia";
import { ref, computed } from "vue";
import { taskService } from "@/api/taskService";

export const useTaskStore = defineStore("task", () => {
  // State
  const tasks = ref([]);
  const meta = ref({
    page: 1,
    limit: 1,
    total: 0,
    total_pages: 0,
    has_next_page: false,
    has_prev_page: false,
  });
  const loading = ref(false);
  const error = ref(null);

  // Getters
  const totalTasks = computed(() => meta.value.total || 0);
  const currentPage = computed(() => meta.value.page || 1);
  const totalPages = computed(() => meta.value.total_pages || 0);

  // Actions
  const fetchTasks = async (params = {}) => {
    loading.value = true;
    error.value = null;

    try {
      const result = await taskService.getTasks(params);

      if (result.success) {
        tasks.value = result.data.tasks || [];
        meta.value = result.data.meta || {};
      } else {
        error.value = result.message;
      }
    } catch (err) {
      error.value = "Failed to fetch tasks";
      console.error("Fetch tasks error:", err);
    } finally {
      loading.value = false;
    }
  };

  const fetchTaskById = async (id) => {
    loading.value = true;
    error.value = null;

    try {
      const result = await taskService.getTaskById(id);

      if (result.success) {
        // Normalize data structure like fetchTasks
        const normalizedTask = {
          id: result.data.id || "",
          title: result.data.title || "Untitled",
          description: result.data.description || "",
          status: result.data.status || "pending",
          priority: result.data.priority || "medium",
          due_date: result.data.due_date || null,
          tags: result.data.tags || [],
          created_at: result.data.created_at || new Date().toISOString(),
          updated_at: result.data.updated_at || new Date().toISOString(),
          completed_at: result.data.completed_at || null,
        };

        return { success: true, data: normalizedTask };
      } else {
        error.value = result.message;
        return { success: false, message: result.message };
      }
    } catch (err) {
      error.value = "Failed to fetch task";
      console.error("Fetch task by ID error:", err);
      return { success: false, message: error.value };
    } finally {
      loading.value = false;
    }
  };

  const createTask = async (taskData) => {
    loading.value = true;
    error.value = null;

    try {
      const result = await taskService.createTask(taskData);

      if (result.success) {
        // Refresh task list after create
        await fetchTasks();
        return { success: true };
      } else {
        error.value = result.message;
        return {
          success: false,
          message: result.message,
          errors: result.errors,
        };
      }
    } catch (err) {
      error.value = "Failed to create task";
      return { success: false, message: error.value };
    } finally {
      loading.value = false;
    }
  };

  const updateTask = async (id, taskData) => {
    loading.value = true;
    error.value = null;

    try {
      const result = await taskService.updateTask(id, taskData);

      if (result.success) {
        // Update local task in array
        const index = tasks.value.findIndex((t) => t.id === id);
        if (index !== -1) {
          // Merge updated data with existing task
          tasks.value[index] = {
            ...tasks.value[index],
            ...result.data,
            updated_at: new Date().toISOString(),
          };
        }
        return { success: true };
      } else {
        error.value = result.message;
        return {
          success: false,
          message: result.message,
          errors: result.errors,
        };
      }
    } catch (err) {
      error.value = "Failed to update task";
      console.error("Update task error:", err);
      return { success: false, message: error.value };
    } finally {
      loading.value = false;
    }
  };

  const deleteTask = async (id) => {
    loading.value = true;
    error.value = null;

    try {
      const result = await taskService.deleteTask(id);

      if (result.success) {
        // Remove from local list
        tasks.value = tasks.value.filter((t) => t.id !== id);
        // Update meta total
        if (meta.value.total > 0) {
          meta.value.total--;
        }
        return { success: true };
      } else {
        error.value = result.message;
        return { success: false, message: result.message };
      }
    } catch (err) {
      error.value = "Failed to delete task";
      return { success: false, message: error.value };
    } finally {
      loading.value = false;
    }
  };

  const clearError = () => {
    error.value = null;
  };

  return {
    // State
    tasks,
    meta,
    loading,
    error,
    // Getters
    totalTasks,
    currentPage,
    totalPages,
    // Actions
    fetchTasks,
    fetchTaskById,
    createTask,
    updateTask,
    deleteTask,
    clearError,
  };
});
