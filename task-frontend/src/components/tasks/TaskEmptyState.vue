<template>
  <div class="bg-white rounded-xl border border-gray-200 p-12 text-center">
    <div
      class="inline-flex items-center justify-center w-16 h-16 bg-gray-100 rounded-full mb-4"
    >
      <ListTodo :size="32" class="text-gray-400" />
    </div>
    <h3 class="text-lg font-semibold text-gray-900 mb-2">No tasks found</h3>
    <p class="text-sm text-gray-600 mb-6">
      Get started by creating your first task
    </p>
    <button
      @click="$emit('create')"
      class="inline-flex items-center gap-2 px-4 py-2 bg-[#fd9621] text-white rounded-lg font-medium hover:bg-[#e58519] transition-colors"
    >
      <Plus :size="18" />
      <span>Create Task</span>
    </button>
  </div>
</template>

<script setup>
import { ListTodo, Plus } from "lucide-vue-next";

defineEmits(["create"]);
</script>
