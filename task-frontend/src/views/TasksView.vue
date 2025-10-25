<template>
  <div class="min-h-screen bg-gray-50">
    <TaskNavbar
      :email="authStore.currentUser?.email || 'User'"
      @logout="handleLogout"
    />

    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
      <TaskHeader
        :total-tasks="taskStore.totalTasks"
        @create="openCreateModal"
      />

      <TaskFilters
        :filters="filters"
        @update:search="handleSearchUpdate"
        @update:status="handleFilterUpdate('status', $event)"
        @update:priority="handleFilterUpdate('priority', $event)"
        @update:sort="handleFilterUpdate('sort', $event)"
        @reset="resetFilters"
      />

      <ErrorAlert :message="taskStore.error" @close="taskStore.clearError" />

      <LoadingSpinner v-if="taskStore.loading" />

      <TaskList
        v-else-if="taskStore.tasks.length > 0"
        :tasks="taskStore.tasks"
        @edit="openEditModal"
        @delete="confirmDelete"
        @view="viewTaskDetail"
      />

      <TaskEmptyState v-else @create="openCreateModal" />

      <TaskPagination
        :current-page="taskStore.currentPage"
        :total-pages="taskStore.totalPages"
        :has-prev-page="taskStore.meta.has_prev_page"
        :has-next-page="taskStore.meta.has_next_page"
        @change-page="changePage"
      />
    </main>

    <TaskModal
      v-if="showModal"
      :key="selectedTask?.id || 'create'"
      :task="selectedTask"
      @close="closeModal"
      @submit="handleSubmit"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, defineAsyncComponent } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/authStore";
import { useTaskStore } from "@/stores/taskStore";

// Regular imports
import TaskNavbar from "@/components/tasks/TaskNavbar.vue";
import TaskHeader from "@/components/tasks/TaskHeader.vue";
import TaskFilters from "@/components/tasks/TaskFilters.vue";
import ErrorAlert from "@/components/ui/ErrorAlert.vue";
import LoadingSpinner from "@/components/common/LoadingSpinner.vue";

// Lazy loading for conditional components
const TaskList = defineAsyncComponent(() =>
  import("@/components/tasks/TaskList.vue")
);
const TaskEmptyState = defineAsyncComponent(() =>
  import("@/components/tasks/TaskEmptyState.vue")
);
const TaskPagination = defineAsyncComponent(() =>
  import("@/components/tasks/TaskPagination.vue")
);
const TaskModal = defineAsyncComponent(() =>
  import("@/components/tasks/TaskModal.vue")
);

const router = useRouter();
const authStore = useAuthStore();
const taskStore = useTaskStore();

const showModal = ref(false);
const selectedTask = ref(null);

const filters = ref({
  search: "",
  status: "",
  priority: "",
  sort: "-created_at",
  page: 1,
  limit: 10,
});

let searchTimeout = null;

onMounted(() => {
  loadTasks();
});

const loadTasks = () => {
  const params = {
    page: filters.value.page,
    limit: filters.value.limit,
  };

  if (filters.value.search) params.search = filters.value.search;
  if (filters.value.status) params.status = filters.value.status;
  if (filters.value.priority) params.priority = filters.value.priority;
  if (filters.value.sort) params.sort = filters.value.sort;

  taskStore.fetchTasks(params);
};

const resetFilters = () => {
  filters.value = {
    search: "",
    status: "",
    priority: "",
    sort: "-created_at",
    page: 1,
    limit: 10,
  };
  loadTasks();
};

const handleSearchUpdate = (value) => {
  filters.value.search = value;
  clearTimeout(searchTimeout);
  searchTimeout = setTimeout(() => {
    filters.value.page = 1;
    loadTasks();
  }, 500);
};

const handleFilterUpdate = (key, value) => {
  filters.value[key] = value;
  filters.value.page = 1;
  loadTasks();
};

const changePage = (page) => {
  filters.value.page = page;
  loadTasks();
  window.scrollTo({ top: 0, behavior: "smooth" });
};

const openCreateModal = () => {
  selectedTask.value = null;
  showModal.value = true;
};

const openEditModal = (task) => {
  selectedTask.value = JSON.parse(JSON.stringify(task));
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
  setTimeout(() => {
    selectedTask.value = null;
  }, 300);
};

const handleSubmit = async (taskData) => {
  let result;

  if (selectedTask.value?.id) {
    result = await taskStore.updateTask(selectedTask.value.id, taskData);
  } else {
    result = await taskStore.createTask(taskData);
  }

  if (result.success) {
    closeModal();
  } else {
    alert(result.message || "Operation failed");
  }
};

const confirmDelete = (task) => {
  if (confirm(`Are you sure you want to delete "${task.title}"?`)) {
    taskStore.deleteTask(task.id);
  }
};

// View task detail handler
const viewTaskDetail = (task) => {
  console.log("Navigating to task detail:", task.id);
  router.push(`/tasks/${task.id}`);
};

const handleLogout = async () => {
  await authStore.logout();
  router.push("/login");
};
</script>
