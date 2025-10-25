import { defineStore } from "pinia";
import { ref, computed } from "vue";
import { authService } from "@/api/authService";
import { cookieService } from "@/utils/cookies";

export const useAuthStore = defineStore("auth", () => {
  // State
  const token = ref(cookieService.getToken());
  const user = ref(cookieService.getUser());
  const loading = ref(false);
  const error = ref(null);

  // Getters
  const isAuthenticated = computed(() => {
    return !!token.value && !!user.value;
  });

  const currentUser = computed(() => user.value);

  const login = async (credentials) => {
    loading.value = true;
    error.value = null;

    try {
      const result = await authService.login(credentials);

      if (result.success) {
        const responseToken = result.data.data.token;
        const responseUser = result.data.user || { email: credentials.email };

        // Update state
        token.value = responseToken;
        user.value = responseUser;

        // Save to secure cookies
        cookieService.setAuthData(responseToken, responseUser);

        console.log("Login successful");
        return { success: true };
      } else {
        error.value = result.message;
        console.error("Login failed:", result.message);
        return {
          success: false,
          message: result.message,
          errors: result.errors,
        };
      }
    } catch (err) {
      error.value = "An unexpected error occurred during login";
      console.error("Login error:", err);
      return {
        success: false,
        message: error.value,
      };
    } finally {
      loading.value = false;
    }
  };

  const logout = () => {
    console.log("Logging out...");
    clearAuth();
  };

  const clearAuth = () => {
    token.value = null;
    user.value = null;
    error.value = null;

    // Clear cookies
    cookieService.clearAll();

    console.log("Auth cleared");
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
