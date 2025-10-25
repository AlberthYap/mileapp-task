import axiosInstance from "./axios";

export const authService = {
  /**
   * Login user with credentials
   * @param {Object} credentials - { email, password }
   * @returns {Promise}
   */
  async login(credentials) {
    try {
      const response = await axiosInstance.post("/auth/login", credentials);
      return {
        success: true,
        data: response.data,
      };
    } catch (error) {
      return {
        success: false,
        message: error.response?.data?.message || "Login failed",
        errors: error.response?.data?.errors,
      };
    }
  },
};
