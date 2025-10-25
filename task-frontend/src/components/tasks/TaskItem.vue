<template>
  <div
    class="bg-white rounded-xl border border-gray-200 p-5 hover:shadow-md transition-all"
  >
    <div class="flex items-start justify-between gap-4">
      <!-- Make content area clickable -->
      <div class="flex-1 min-w-0 cursor-pointer group" @click="handleViewTask">
        <div class="flex items-start gap-3 mb-3">
          <div class="flex-1">
            <h3
              class="text-base font-semibold text-gray-900 mb-1 group-hover:text-[#fd9621] transition-colors"
            >
              {{ task.title || "Untitled" }}
            </h3>
            <p class="text-sm text-gray-600 line-clamp-2">
              {{ task.description || "No description" }}
            </p>
          </div>
        </div>

        <div class="flex flex-wrap items-center gap-2">
          <!-- Status Badge -->
          <span
            v-if="task.status"
            class="inline-flex items-center gap-1 px-2.5 py-1 rounded-lg text-xs font-medium"
            :class="getStatusClass(task.status)"
          >
            <span
              class="w-1.5 h-1.5 rounded-full"
              :class="getStatusDotClass(task.status)"
            ></span>
            {{ formatStatus(task.status) }}
          </span>

          <!-- Priority Badge -->
          <span
            v-if="task.priority"
            class="inline-flex items-center gap-1 px-2.5 py-1 rounded-lg text-xs font-medium"
            :class="getPriorityClass(task.priority)"
          >
            <Flag :size="12" />
            {{ formatPriority(task.priority) }}
          </span>

          <!-- Due Date -->
          <span
            v-if="task.due_date"
            class="inline-flex items-center gap-1 px-2.5 py-1 rounded-lg text-xs font-medium bg-gray-100 text-gray-700"
          >
            <Calendar :size="12" />
            {{ formatDate(task.due_date) }}
          </span>
        </div>
      </div>

      <!-- Action buttons - stop propagation to prevent triggering view -->
      <div class="flex items-center gap-2" @click.stop>
        <button
          @click="$emit('edit', task)"
          class="p-2 text-gray-400 hover:text-[#fd9621] hover:bg-orange-50 rounded-lg transition-colors"
          title="Edit"
        >
          <Edit2 :size="18" />
        </button>
        <button
          @click="$emit('delete', task)"
          class="p-2 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
          title="Delete"
        >
          <Trash2 :size="18" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Edit2, Trash2, Flag, Calendar } from "lucide-vue-next";
import { useTaskUtils } from "@/composables/useTaskUtils";

const props = defineProps({
  task: {
    type: Object,
    required: true,
  },
});

const emit = defineEmits(["edit", "delete", "view"]);

const {
  getStatusClass,
  getStatusDotClass,
  getPriorityClass,
  formatStatus,
  formatPriority,
  formatDate,
} = useTaskUtils();

const handleViewTask = () => {
  emit("view", props.task);
};
</script>
