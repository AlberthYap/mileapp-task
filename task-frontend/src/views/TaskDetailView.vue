<template>
  <div class="min-h-screen bg-gray-50">
    <TaskNavbar
      :email="authStore.currentUser?.email || 'User'"
      @logout="handleLogout"
    />

    <main class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
      <!-- Back Button -->
      <button
        @click="goBack"
        class="inline-flex items-center gap-2 text-gray-600 hover:text-gray-900 mb-6 transition-colors"
      >
        <ArrowLeft :size="20" />
        <span class="font-medium">Back to Tasks</span>
      </button>

      <!-- Loading State -->
      <LoadingSpinner v-if="loading" />

      <!-- Error State -->
      <ErrorAlert v-else-if="error" :message="error" @close="error = null" />

      <!-- Task Detail Content -->
      <div v-else-if="task" class="space-y-6">
        <!-- Header Card -->
        <div
          class="bg-white rounded-2xl border border-gray-200 shadow-sm overflow-hidden"
        >
          <div class="p-6 border-b border-gray-100">
            <div class="flex items-start justify-between gap-4 mb-4">
              <div class="flex-1">
                <h1 class="text-3xl font-bold text-gray-900 mb-2">
                  {{ task.title }}
                </h1>
                <div class="flex flex-wrap items-center gap-2">
                  <!-- Status Badge -->
                  <span
                    v-if="task.status"
                    class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-sm font-medium"
                    :class="getStatusClass(task.status)"
                  >
                    <span
                      class="w-2 h-2 rounded-full"
                      :class="getStatusDotClass(task.status)"
                    ></span>
                    {{ formatStatus(task.status) }}
                  </span>

                  <!-- Priority Badge -->
                  <span
                    v-if="task.priority"
                    class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-sm font-medium"
                    :class="getPriorityClass(task.priority)"
                  >
                    <Flag :size="14" />
                    {{ formatPriority(task.priority) }}
                  </span>
                </div>
              </div>

              <!-- Action Buttons -->
              <div class="flex items-center gap-2">
                <button
                  @click="openEditModal"
                  class="p-2.5 text-gray-400 hover:text-[#fd9621] hover:bg-orange-50 rounded-lg transition-colors"
                  title="Edit Task"
                >
                  <Edit2 :size="20" />
                </button>
                <button
                  @click="confirmDelete"
                  class="p-2.5 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
                  title="Delete Task"
                >
                  <Trash2 :size="20" />
                </button>
              </div>
            </div>
          </div>

          <!-- Task Info Grid -->
          <div class="p-6 grid grid-cols-1 sm:grid-cols-2 gap-6">
            <!-- Due Date -->
            <div class="flex items-start gap-3">
              <div class="p-2 bg-gray-100 rounded-lg">
                <Calendar :size="18" class="text-gray-600" />
              </div>
              <div>
                <p
                  class="text-xs font-medium text-gray-500 uppercase tracking-wider mb-1"
                >
                  Due Date
                </p>
                <p class="text-sm font-semibold text-gray-900">
                  {{
                    task.due_date
                      ? formatFullDate(task.due_date)
                      : "No due date"
                  }}
                </p>
                <p v-if="task.due_date" class="text-xs text-gray-500 mt-0.5">
                  {{ formatRelativeDate(task.due_date) }}
                </p>
              </div>
            </div>

            <!-- Created At -->
            <div class="flex items-start gap-3">
              <div class="p-2 bg-gray-100 rounded-lg">
                <Clock :size="18" class="text-gray-600" />
              </div>
              <div>
                <p
                  class="text-xs font-medium text-gray-500 uppercase tracking-wider mb-1"
                >
                  Created
                </p>
                <p class="text-sm font-semibold text-gray-900">
                  {{ formatFullDate(task.created_at) }}
                </p>
                <p class="text-xs text-gray-500 mt-0.5">
                  {{ formatRelativeDate(task.created_at) }}
                </p>
              </div>
            </div>

            <!-- Updated At -->
            <div class="flex items-start gap-3">
              <div class="p-2 bg-gray-100 rounded-lg">
                <RefreshCw :size="18" class="text-gray-600" />
              </div>
              <div>
                <p
                  class="text-xs font-medium text-gray-500 uppercase tracking-wider mb-1"
                >
                  Last Updated
                </p>
                <p class="text-sm font-semibold text-gray-900">
                  {{ formatFullDate(task.updated_at) }}
                </p>
                <p class="text-xs text-gray-500 mt-0.5">
                  {{ formatRelativeDate(task.updated_at) }}
                </p>
              </div>
            </div>

            <!-- Completed At -->
            <div class="flex items-start gap-3">
              <div class="p-2 bg-gray-100 rounded-lg">
                <RefreshCw :size="18" class="text-gray-600" />
              </div>
              <div>
                <p
                  class="text-xs font-medium text-gray-500 uppercase tracking-wider mb-1"
                >
                  Complete At
                </p>
                <p class="text-sm font-semibold text-gray-900">
                  {{ formatFullDate(task.completed_at) }}
                </p>
                <p class="text-xs text-gray-500 mt-0.5">
                  {{ formatRelativeDate(task.completed_at) }}
                </p>
              </div>
            </div>

            <!-- Task ID -->
            <div class="flex items-start gap-3">
              <div class="p-2 bg-gray-100 rounded-lg">
                <Hash :size="18" class="text-gray-600" />
              </div>
              <div>
                <p
                  class="text-xs font-medium text-gray-500 uppercase tracking-wider mb-1"
                >
                  Task ID
                </p>
                <p
                  class="text-sm font-mono font-semibold text-gray-900 truncate"
                >
                  {{ task.id }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- Description Card -->
        <div class="bg-white rounded-2xl border border-gray-200 shadow-sm p-6">
          <div class="flex items-center gap-2 mb-4">
            <FileText :size="20" class="text-gray-600" />
            <h2 class="text-lg font-semibold text-gray-900">Description</h2>
          </div>
          <div
            v-if="task.description"
            class="prose prose-sm max-w-none text-gray-700 leading-relaxed"
          >
            <p class="whitespace-pre-wrap">{{ task.description }}</p>
          </div>
          <p v-else class="text-sm text-gray-500 italic">
            No description provided
          </p>
        </div>

        <!-- Tags Card -->
        <div
          v-if="task.tags && task.tags.length > 0"
          class="bg-white rounded-2xl border border-gray-200 shadow-sm p-6"
        >
          <div class="flex items-center gap-2 mb-4">
            <Tag :size="20" class="text-gray-600" />
            <h2 class="text-lg font-semibold text-gray-900">Tags</h2>
          </div>
          <div class="flex flex-wrap gap-2">
            <span
              v-for="(tag, index) in task.tags"
              :key="index"
              class="inline-flex items-center px-3 py-1.5 bg-blue-50 text-blue-700 rounded-lg text-sm font-medium"
            >
              <Tag :size="12" class="mr-1.5" />
              {{ tag }}
            </span>
          </div>
        </div>

        <!-- Quick Actions -->
        <div
          v-if="task.status !== 'completed'"
          class="bg-linear-to-br from-orange-50 to-amber-50 rounded-2xl border border-orange-200 p-6"
        >
          <h2 class="text-lg font-semibold text-gray-900 mb-4">
            Quick Actions
          </h2>
          <div class="grid grid-cols-1 sm:grid-cols-3 gap-3">
            <button
              @click="markAsCompleted"
              :disabled="updating"
              class="flex items-center justify-center gap-2 px-4 py-3 bg-white border border-green-200 text-green-700 rounded-lg font-medium hover:bg-green-50 hover:border-green-300 disabled:opacity-50 disabled:cursor-not-allowed transition-all"
            >
              <Loader2 v-if="updating" :size="18" class="animate-spin" />
              <CheckCircle2 v-else :size="18" />
              <span>Mark Complete</span>
            </button>

            <button
              v-if="task.status === 'pending'"
              @click="markAsInProgress"
              :disabled="updating"
              class="flex items-center justify-center gap-2 px-4 py-3 bg-white border border-blue-200 text-blue-700 rounded-lg font-medium hover:bg-blue-50 hover:border-blue-300 disabled:opacity-50 disabled:cursor-not-allowed transition-all"
            >
              <Loader2 v-if="updating" :size="18" class="animate-spin" />
              <PlayCircle v-else :size="18" />
              <span>Start Task</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Not Found State -->
      <div
        v-else
        class="bg-white rounded-2xl border border-gray-200 p-12 text-center"
      >
        <div
          class="inline-flex items-center justify-center w-16 h-16 bg-gray-100 rounded-full mb-4"
        >
          <AlertCircle :size="32" class="text-gray-400" />
        </div>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">Task not found</h3>
        <p class="text-sm text-gray-600 mb-6">
          The task you're looking for doesn't exist or has been deleted.
        </p>
        <button
          @click="goBack"
          class="inline-flex items-center gap-2 px-4 py-2 bg-[#fd9621] text-white rounded-lg font-medium hover:bg-[#e58519] transition-colors"
        >
          <ArrowLeft :size="18" />
          <span>Back to Tasks</span>
        </button>
      </div>
    </main>

    <!-- Edit Modal -->
    <TaskModal
      v-if="showEditModal"
      :key="task?.id"
      :task="task"
      @close="showEditModal = false"
      @submit="handleUpdate"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, defineAsyncComponent } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useAuthStore } from "@/stores/authStore";
import { useTaskStore } from "@/stores/taskStore";
import { useTaskUtils } from "@/composables/useTaskUtils";
import TaskNavbar from "@/components/tasks/TaskNavbar.vue";
import LoadingSpinner from "@/components/common/LoadingSpinner.vue";
import ErrorAlert from "@/components/ui/ErrorAlert.vue";
import {
  ArrowLeft,
  Edit2,
  Trash2,
  Calendar,
  Clock,
  RefreshCw,
  Hash,
  FileText,
  Tag,
  Flag,
  CheckCircle2,
  PlayCircle,
  Copy,
  AlertCircle,
  Loader2,
} from "lucide-vue-next";

const TaskModal = defineAsyncComponent(() =>
  import("@/components/tasks/TaskModal.vue")
);

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const taskStore = useTaskStore();

const {
  getStatusClass,
  getStatusDotClass,
  getPriorityClass,
  formatStatus,
  formatPriority,
} = useTaskUtils();

const task = ref(null);
const loading = ref(true);
const error = ref(null);
const updating = ref(false);
const showEditModal = ref(false);

onMounted(async () => {
  await fetchTaskDetail();
});

const fetchTaskDetail = async () => {
  loading.value = true;
  error.value = null;

  try {
    const taskId = route.params.id;

    // ALWAYS fetch from API untuk ensure fresh data
    const result = await taskStore.fetchTaskById(taskId);

    if (result.success) {
      task.value = result.data;
    } else {
      error.value = result.message || "Failed to load task";
    }
  } catch (err) {
    console.error("Error fetching task:", err);
    error.value = "Failed to load task details";
  } finally {
    loading.value = false;
  }
};

const goBack = () => {
  router.push("/tasks");
};

const handleLogout = async () => {
  await authStore.logout();
  router.push("/login");
};

const openEditModal = () => {
  showEditModal.value = true;
};

const handleUpdate = async (taskData) => {
  const result = await taskStore.updateTask(task.value.id, taskData);

  if (result.success) {
    showEditModal.value = false;
    // Refresh task data after update
    await fetchTaskDetail();
  } else {
    alert(result.message || "Failed to update task");
  }
};

const confirmDelete = () => {
  if (confirm(`Are you sure you want to delete "${task.value.title}"?`)) {
    deleteTask();
  }
};

const deleteTask = async () => {
  const result = await taskStore.deleteTask(task.value.id);

  if (result.success) {
    router.push("/tasks");
  } else {
    error.value = result.message || "Failed to delete task";
  }
};

const markAsCompleted = async () => {
  updating.value = true;

  try {
    const result = await taskStore.updateTask(task.value.id, {
      ...task.value,
      status: "completed",
    });

    if (result.success) {
      // Refresh untuk get updated data
      await fetchTaskDetail();
    } else {
      error.value = result.message || "Failed to update task";
    }
  } catch (err) {
    console.error("Error updating task:", err);
    error.value = "Failed to update task";
  } finally {
    updating.value = false;
  }
};

const markAsInProgress = async () => {
  updating.value = true;

  try {
    const result = await taskStore.updateTask(task.value.id, {
      ...task.value,
      status: "in_progress",
    });

    if (result.success) {
      await fetchTaskDetail();
    } else {
      error.value = result.message || "Failed to update task";
    }
  } catch (err) {
    console.error("Error updating task:", err);
    error.value = "Failed to update task";
  } finally {
    updating.value = false;
  }
};

// Date formatting functions
const formatFullDate = (dateString) => {
  if (!dateString) return "N/A";

  try {
    const date = new Date(dateString);
    if (isNaN(date.getTime())) return "Invalid date";

    return date.toLocaleDateString("en-US", {
      weekday: "long",
      year: "numeric",
      month: "long",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    });
  } catch (error) {
    return "Invalid date";
  }
};

const formatRelativeDate = (dateString) => {
  if (!dateString) return "";

  try {
    const date = new Date(dateString);
    if (isNaN(date.getTime())) return "";

    const now = new Date();
    const diffTime = now - date;
    const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24));
    const diffHours = Math.floor(diffTime / (1000 * 60 * 60));
    const diffMinutes = Math.floor(diffTime / (1000 * 60));

    if (diffMinutes < 1) return "Just now";
    if (diffMinutes < 60)
      return `${diffMinutes} minute${diffMinutes !== 1 ? "s" : ""} ago`;
    if (diffHours < 24)
      return `${diffHours} hour${diffHours !== 1 ? "s" : ""} ago`;
    if (diffDays < 7) return `${diffDays} day${diffDays !== 1 ? "s" : ""} ago`;
    if (diffDays < 30)
      return `${Math.floor(diffDays / 7)} week${
        Math.floor(diffDays / 7) !== 1 ? "s" : ""
      } ago`;
    if (diffDays < 365)
      return `${Math.floor(diffDays / 30)} month${
        Math.floor(diffDays / 30) !== 1 ? "s" : ""
      } ago`;
    return `${Math.floor(diffDays / 365)} year${
      Math.floor(diffDays / 365) !== 1 ? "s" : ""
    } ago`;
  } catch (error) {
    return "";
  }
};
</script>
