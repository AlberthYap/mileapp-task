<template>
  <div class="mb-6">
    <div
      class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4"
    >
      <div>
        <h2 class="text-2xl font-bold text-gray-900">My Tasks</h2>
        <p class="text-sm text-gray-600 mt-1">
          {{ totalTasks }} total task{{ totalTasks !== 1 ? "s" : "" }}
        </p>
      </div>
      <button
        @click="$emit('create')"
        class="inline-flex items-center justify-center gap-2 px-4 py-2.5 bg-linear-to-r from-[#fd9621] via-[#90caf9] to-[#02a8f3] text-white rounded-lg font-medium hover:shadow-lg hover:scale-105 transition-all"
      >
        <Plus :size="18" />
        <span>New Task</span>
      </button>
    </div>
  </div>
</template>

<script setup>
import { Plus } from "lucide-vue-next";

defineProps({
  totalTasks: {
    type: Number,
    default: 0,
  },
});

defineEmits(["create"]);
</script>
