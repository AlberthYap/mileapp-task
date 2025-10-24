import axiosInstance from "./axios";

export const authService = {
  /**
   * Login user with credentials
   * @param {Object} credentials - { email, password }
   * @returns {Promise}
   */
  async login(credentials) {
    try {
      const response = await axiosInstance.post("/auth/login", {
        email: credentials.email,
        password: credentials.password,
      });

      return {
        success: true,
        data: response.data,
      };
    } catch (error) {
      return {
        success: false,
        message:
          error.response?.data?.message || "Login failed. Please try again.",
        errors: error.response?.data?.errors || {},
      };
    }
  },

  /**
   * Logout user
   * @returns {Promise}
   */
  async logout() {
    try {
      await axiosInstance.post("/logout");
      return { success: true };
    } catch (error) {
      // Even if logout fails on server, we still clear local data
      return { success: true };
    }
  },

  /**
   * Verify token validity
   * @returns {Promise}
   */
  async verifyToken() {
    try {
      const response = await axiosInstance.get("/verify-token");
      return {
        success: true,
        data: response.data,
      };
    } catch (error) {
      return {
        success: false,
        message: "Token verification failed",
      };
    }
  },
};
