<template>
  <div class="bg-white rounded-xl border border-gray-200 p-4 mb-6">
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
  </div>
</template>

<script setup>
import { Search } from "lucide-vue-next";

defineProps({
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
]);
</script>
