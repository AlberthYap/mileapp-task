import Cookies from "js-cookie";

const TOKEN_KEY = "auth_token";
const USER_KEY = "auth_user";

// Cookie options for maximum security
const getCookieOptions = () => ({
  expires: 1, // 1 day
  secure: import.meta.env.PROD, // HTTPS only in production
  sameSite: "strict", // CSRF protection
  path: "/", // Available throughout the app
});

export const cookieService = {
  // ========== TOKEN METHODS ==========

  // Set auth token in cookie
  setToken(token) {
    if (!token) {
      console.warn("Attempting to set empty token");
      return;
    }
    Cookies.set(TOKEN_KEY, token, getCookieOptions());
    console.log("Token saved to cookie");
  },

  // Get auth token from cookie
  getToken() {
    return Cookies.get(TOKEN_KEY) || null;
  },

  // Remove auth token from cookie
  removeToken() {
    Cookies.remove(TOKEN_KEY, { path: "/" });
    console.log("Token removed from cookie");
  },

  // Check if token exists
  hasToken() {
    return !!Cookies.get(TOKEN_KEY);
  },

  // ========== USER METHODS ==========

  // Set user data (stored as JSON string in cookie)
  setUser(user) {
    if (!user) {
      console.warn("Attempting to set empty user");
      return;
    }
    try {
      Cookies.set(USER_KEY, JSON.stringify(user), getCookieOptions());
      console.log("User data saved to cookie");
    } catch (error) {
      console.error("Failed to save user to cookie:", error);
    }
  },

  // Get user data from cookie
  getUser() {
    try {
      const userData = Cookies.get(USER_KEY);
      return userData ? JSON.parse(userData) : null;
    } catch (error) {
      console.error("Failed to parse user data from cookie:", error);
      return null;
    }
  },

  // Remove user data from cookie
  removeUser() {
    Cookies.remove(USER_KEY, { path: "/" });
    console.log("User data removed from cookie");
  },

  // Check if user exists
  hasUser() {
    return !!Cookies.get(USER_KEY);
  },

  // ========== UTILITY METHODS ==========

  // Clear all auth data (token + user)
  clearAll() {
    this.removeToken();
    this.removeUser();
    console.log("All auth cookies cleared");
  },

  // Check if user is authenticated (has both token and user)
  isAuthenticated() {
    return this.hasToken() && this.hasUser();
  },

  // Get all auth data at once
  getAuthData() {
    return {
      token: this.getToken(),
      user: this.getUser(),
    };
  },

  // Set all auth data at once
  setAuthData(token, user) {
    if (token && user) {
      this.setToken(token);
      this.setUser(user);
      return true;
    }
    console.warn("Cannot set auth data: missing token or user");
    return false;
  },
};
