<template>
  <div v-if="isVisible" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex min-h-screen items-center justify-center p-4">
      <!-- Backdrop -->
      <div
        class="fixed inset-0 bg-black/50 transition-opacity"
        @click="handleClose"
      ></div>

      <!-- Modal -->
      <div
        class="relative bg-white rounded-2xl shadow-xl max-w-2xl w-full max-h-[90vh] overflow-hidden z-10"
      >
        <!-- Header -->
        <div
          class="flex items-center justify-between p-6 border-b border-gray-200"
        >
          <h3 class="text-xl font-bold text-gray-900">
            {{ isEdit ? "Edit Task" : "Create New Task" }}
          </h3>
          <button
            @click="handleClose"
            type="button"
            class="p-2 text-gray-400 hover:text-gray-600 hover:bg-gray-100 rounded-lg transition-colors"
          >
            <X :size="20" />
          </button>
        </div>

        <!-- Form -->
        <div class="p-6 space-y-5 overflow-y-auto max-h-[calc(90vh-180px)]">
          <!-- Title -->
          <div>
            <label class="block text-sm font-semibold text-gray-700 mb-2">
              Title <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.title"
              type="text"
              required
              placeholder="Enter task title"
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent"
            />
          </div>

          <!-- Description -->
          <div>
            <label class="block text-sm font-semibold text-gray-700 mb-2">
              Description
            </label>
            <textarea
              v-model="form.description"
              rows="4"
              placeholder="Enter task description"
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent resize-none"
            ></textarea>
          </div>

          <!-- Status & Priority -->
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">
                Status <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.status"
                required
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent"
              >
                <option value="pending">Pending</option>
                <option value="in_progress">In Progress</option>
                <option value="completed">Completed</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">
                Priority <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.priority"
                required
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent"
              >
                <option value="low">Low</option>
                <option value="medium">Medium</option>
                <option value="high">High</option>
              </select>
            </div>
          </div>

          <!-- Due Date -->
          <div>
            <label class="block text-sm font-semibold text-gray-700 mb-2">
              Due Date
            </label>
            <input
              v-model="form.due_date"
              type="datetime-local"
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent"
            />
          </div>

          <!-- Tags -->
          <div>
            <label class="block text-sm font-semibold text-gray-700 mb-2">
              Tags
            </label>
            <input
              v-model="form.tagsInput"
              type="text"
              placeholder="Enter tags separated by commas"
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent"
            />
            <p class="mt-1 text-xs text-gray-500">
              Separate tags with commas (e.g., work, urgent, meeting)
            </p>
          </div>
        </div>

        <!-- Footer -->
        <div
          class="flex items-center justify-end gap-3 p-6 border-t border-gray-200 bg-gray-50"
        >
          <button
            type="button"
            @click="handleClose"
            class="px-5 py-2.5 border border-gray-300 text-gray-700 rounded-lg font-medium hover:bg-gray-100 transition-colors"
          >
            Cancel
          </button>
          <button
            type="button"
            @click="handleSubmit"
            :disabled="loading"
            class="px-5 py-2.5 bg-linear-to-r from-[#fd9621] via-[#90caf9] to-[#02a8f3] text-white rounded-lg font-medium hover:shadow-lg disabled:opacity-50 transition-all flex items-center gap-2"
          >
            <Loader2 v-if="loading" :size="18" class="animate-spin" />
            <span>{{ isEdit ? "Update Task" : "Create Task" }}</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, nextTick } from "vue";
import { X, Loader2 } from "lucide-vue-next";

const props = defineProps({
  task: {
    type: Object,
    default: null,
  },
});

const emit = defineEmits(["close", "submit"]);

const isVisible = ref(true);
const loading = ref(false);

const isEdit = computed(() => !!props.task?.id);

const form = ref({
  title: "",
  description: "",
  status: "pending",
  priority: "medium",
  due_date: "",
  tagsInput: "",
});

// Initialize form when component mounts or task changes
const initForm = () => {
  if (props.task) {
    form.value = {
      title: props.task.title || "",
      description: props.task.description || "",
      status: props.task.status || "pending",
      priority: props.task.priority || "medium",
      due_date: props.task.due_date
        ? formatDateForInput(props.task.due_date)
        : "",
      tagsInput: Array.isArray(props.task.tags)
        ? props.task.tags.join(", ")
        : "",
    };
  } else {
    form.value = {
      title: "",
      description: "",
      status: "pending",
      priority: "medium",
      due_date: "",
      tagsInput: "",
    };
  }
};

// Watch task prop
watch(
  () => props.task,
  (newTask) => {
    nextTick(() => {
      initForm();
    });
  },
  { immediate: true, deep: true }
);

const handleClose = () => {
  emit("close");
};

const handleSubmit = () => {
  // Validate
  if (!form.value.title.trim()) {
    alert("Please enter a task title");
    return;
  }

  // Parse tags
  const tags = form.value.tagsInput
    .split(",")
    .map((tag) => tag.trim())
    .filter((tag) => tag.length > 0);

  // Prepare data
  const data = {
    title: form.value.title.trim(),
    description: form.value.description.trim(),
    status: form.value.status,
    priority: form.value.priority,
    tags: tags,
  };

  // Add due_date if provided
  if (form.value.due_date) {
    data.due_date = new Date(form.value.due_date).toISOString();
  }

  emit("submit", data);
};

const formatDateForInput = (dateString) => {
  if (!dateString) return "";

  try {
    const date = new Date(dateString);
    if (isNaN(date.getTime())) return "";

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");
    const hours = String(date.getHours()).padStart(2, "0");
    const minutes = String(date.getMinutes()).padStart(2, "0");

    return `${year}-${month}-${day}T${hours}:${minutes}`;
  } catch (error) {
    console.error("Error formatting date:", error);
    return "";
  }
};
</script>
