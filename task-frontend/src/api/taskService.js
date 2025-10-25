import axiosInstance from "./axios";

export const taskService = {
  /**
   * Get tasks with filter
   * @param {Object} params - filter query
   * @returns {Promise<Object>} - response with success and data
   */
  async getTasks(params = {}) {
    try {
      const response = await axiosInstance.get("/tasks", { params });
      return {
        success: true,
        data: response.data.data,
      };
    } catch (error) {
      return {
        success: false,
        message: error.response?.data?.message || "Failed to fetch tasks",
      };
    }
  },

  /**
   * Get task by ID
   * @param {string} id - task ID
   * @returns {Promise<Object>} - response with success and data
   */
  async getTaskById(id) {
    try {
      const response = await axiosInstance.get(`/tasks/${id}`);

      return {
        success: true,
        data: response.data.data.task,
      };
    } catch (error) {
      return {
        success: false,
        message: error.response?.data?.message || "Failed to fetch task",
      };
    }
  },

  /**
   * Create new task
   * @param {Object} taskData - task details
   * @returns {Promise<Object>} - response with success and data
   */
  async createTask(taskData) {
    try {
      const response = await axiosInstance.post("/tasks", taskData);
      return {
        success: true,
        data: response.data.data,
      };
    } catch (error) {
      return {
        success: false,
        message: error.response?.data?.message || "Failed to create task",
        errors: error.response?.data?.errors || {},
      };
    }
  },

  /**
   * Update task
   * @param {string} id - task id
   * @param {Object} taskData - task details
   * @returns {Promise<Object>} - response with success and data
   */
  async updateTask(id, taskData) {
    try {
      const response = await axiosInstance.put(`/tasks/${id}`, taskData);
      return {
        success: true,
        data: response.data.data,
      };
    } catch (error) {
      return {
        success: false,
        message: error.response?.data?.message || "Failed to update task",
        errors: error.response?.data?.errors || {},
      };
    }
  },

  /**
   * Delete task
   * @param {string} id - task id
   * @returns {Promise<Object>} - response with success and data
   */
  async deleteTask(id) {
    try {
      await axiosInstance.delete(`/tasks/${id}`);
      return {
        success: true,
      };
    } catch (error) {
      return {
        success: false,
        message: error.response?.data?.message || "Failed to delete task",
      };
    }
  },
};
