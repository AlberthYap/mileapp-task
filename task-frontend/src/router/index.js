import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "@/stores/authStore";

const routes = [
  {
    path: "/",
    redirect: "/tasks",
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/views/LoginView.vue"),
    meta: {
      requiresAuth: false,
      redirectIfAuthenticated: true,
    },
  },
  {
    path: "/tasks",
    name: "Tasks",
    component: () => import("@/views/TasksView.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: "/tasks/:id",
    name: "TaskDetail",
    component: () => import("@/views/TaskDetailView.vue"),
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Navigation guard
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore();
  const requiresAuth = to.meta.requiresAuth;
  const redirectIfAuthenticated = to.meta.redirectIfAuthenticated;

  if (requiresAuth && !authStore.isAuthenticated) {
    // Redirect to login if route requires auth and user is not authenticated
    next("/login");
  } else if (redirectIfAuthenticated && authStore.isAuthenticated) {
    // Redirect to tasks if trying to access login while authenticated
    next("/tasks");
  } else {
    next();
  }
});

export default router;
