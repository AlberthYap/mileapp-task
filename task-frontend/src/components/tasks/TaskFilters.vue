<template>
  <div class="bg-white rounded-xl border border-gray-200 p-4 mb-6">
    <div class="flex items-center justify-between mb-4">
      <h3 class="text-sm font-semibold text-gray-900">Filters</h3>
      <button
        @click="$emit('reset')"
        class="inline-flex items-center gap-1.5 px-3 py-1.5 text-sm font-medium text-gray-600 hover:text-[#fd9621] hover:bg-orange-50 rounded-lg transition-colors"
      >
        <RotateCcw :size="14" />
        <span>Reset</span>
      </button>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <div>
        <label class="block text-xs font-medium text-gray-700 mb-1.5"
          >Search</label
        >
        <div class="relative">
          <Search
            :size="16"
            class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
          />
          <input
            :value="filters.search"
            @input="$emit('update:search', $event.target.value)"
            type="text"
            placeholder="Search tasks..."
            class="w-full pl-9 pr-3 py-2 text-sm border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent"
          />
        </div>
      </div>

      <div>
        <label class="block text-xs font-medium text-gray-700 mb-1.5"
          >Status</label
        >
        <select
          :value="filters.status"
          @change="$emit('update:status', $event.target.value)"
          class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent"
        >
          <option value="">All Status</option>
          <option value="pending">Pending</option>
          <option value="in_progress">In Progress</option>
          <option value="completed">Completed</option>
        </select>
      </div>

      <div>
        <label class="block text-xs font-medium text-gray-700 mb-1.5"
          >Priority</label
        >
        <select
          :value="filters.priority"
          @change="$emit('update:priority', $event.target.value)"
          class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent"
        >
          <option value="">All Priority</option>
          <option value="low">Low</option>
          <option value="medium">Medium</option>
          <option value="high">High</option>
        </select>
      </div>

      <div>
        <label class="block text-xs font-medium text-gray-700 mb-1.5"
          >Sort By</label
        >
        <select
          :value="filters.sort"
          @change="$emit('update:sort', $event.target.value)"
          class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:border-transparent"
        >
          <option value="-created_at">Newest First</option>
          <option value="created_at">Oldest First</option>
          <option value="title">Title (A-Z)</option>
          <option value="-title">Title (Z-A)</option>
          <option value="due_date">Due Date (Soon)</option>
          <option value="-due_date">Due Date (Later)</option>
        </select>
      </div>
    </div>

    <!-- Active Filters Badges (Optional) -->
    <div v-if="hasActiveFilters" class="mt-4 flex flex-wrap items-center gap-2">
      <span class="text-xs font-medium text-gray-500">Active filters:</span>

      <button
        v-if="filters.search"
        @click="$emit('update:search', '')"
        class="inline-flex items-center gap-1 px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs"
      >
        Search: "{{ filters.search }}"
        <X :size="12" />
      </button>

      <button
        v-if="filters.status"
        @click="$emit('update:status', '')"
        class="inline-flex items-center gap-1 px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs"
      >
        Status: {{ filters.status }}
        <X :size="12" />
      </button>

      <button
        v-if="filters.priority"
        @click="$emit('update:priority', '')"
        class="inline-flex items-center gap-1 px-2 py-1 bg-orange-100 text-orange-700 rounded text-xs"
      >
        Priority: {{ filters.priority }}
        <X :size="12" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { Search, RotateCcw, X } from "lucide-vue-next";

const props = defineProps({
  filters: {
    type: Object,
    required: true,
  },
});

defineEmits([
  "update:search",
  "update:status",
  "update:priority",
  "update:sort",
  "reset",
]);

const hasActiveFilters = computed(() => {
  return props.filters.search || props.filters.status || props.filters.priority;
});
</script>
