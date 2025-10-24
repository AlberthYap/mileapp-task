import { defineStore } from "pinia";
import { ref, computed } from "vue";
import { authService } from "@/api/authService";

export const useAuthStore = defineStore("auth", () => {
  // State
  const token = ref(localStorage.getItem("auth_token") || null);
  const user = ref(JSON.parse(localStorage.getItem("auth_user") || "null"));
  const loading = ref(false);
  const error = ref(null);

  // Getters
  const isAuthenticated = computed(() => !!token.value);
  const currentUser = computed(() => user.value);

  // Actions
  const login = async (credentials) => {
    loading.value = true;
    error.value = null;

    try {
      const result = await authService.login(credentials);

      if (result.success) {
        // Store token and user data
        token.value = result.data.data.token;
        user.value = result.data.user || { email: credentials.email };

        // Persist to localStorage
        localStorage.setItem("auth_token", result.data.data.token);
        localStorage.setItem("auth_user", JSON.stringify(user.value));

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
      error.value = "An unexpected error occurred during login";
      return {
        success: false,
        message: error.value,
      };
    } finally {
      loading.value = false;
    }
  };

  const logout = async () => {
    loading.value = true;

    try {
      await authService.logout();
    } catch (err) {
      console.error("Logout error:", err);
    } finally {
      clearAuth();
      loading.value = false;
    }
  };

  const clearAuth = () => {
    token.value = null;
    user.value = null;
    error.value = null;
    localStorage.removeItem("auth_token");
    localStorage.removeItem("auth_user");
  };

  const clearError = () => {
    error.value = null;
  };

  return {
    // State
    token,
    user,
    loading,
    error,
    // Getters
    isAuthenticated,
    currentUser,
    // Actions
    login,
    logout,
    clearAuth,
    clearError,
  };
});
