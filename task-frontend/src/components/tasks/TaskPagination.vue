<template>
  <div v-if="totalPages > 1" class="mt-6 flex items-center justify-between">
    <p class="text-sm text-gray-600">
      Page {{ currentPage }} of {{ totalPages }}
    </p>
    <div class="flex items-center gap-2">
      <button
        @click="$emit('change-page', currentPage - 1)"
        :disabled="!hasPrevPage"
        class="p-2 border border-gray-300 rounded-lg hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
      >
        <ChevronLeft :size="18" />
      </button>
      <button
        @click="$emit('change-page', currentPage + 1)"
        :disabled="!hasNextPage"
        class="p-2 border border-gray-300 rounded-lg hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
      >
        <ChevronRight :size="18" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { ChevronLeft, ChevronRight } from "lucide-vue-next";

defineProps({
  currentPage: {
    type: Number,
    required: true,
  },
  totalPages: {
    type: Number,
    required: true,
  },
  hasPrevPage: {
    type: Boolean,
    default: false,
  },
  hasNextPage: {
    type: Boolean,
    default: false,
  },
});

defineEmits(["change-page"]);
</script>
