<template>
  <nav class="bg-white border-b border-gray-200 sticky top-0 z-10">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <div class="flex items-center gap-3">
          <div
            class="w-9 h-9 bg-linear-to-br from-[#fd9621] to-[#ff7b00] rounded-lg flex items-center justify-center"
          >
            <CheckCircle2 :size="20" class="text-white" />
          </div>
          <div>
            <h1 class="text-lg font-bold text-gray-900">MileTaskFlow</h1>
            <p class="text-xs text-gray-500">Task Management</p>
          </div>
        </div>

        <div class="flex items-center gap-4">
          <div class="hidden sm:flex items-center gap-2 text-sm text-gray-600">
            <User :size="16" />
            <span class="max-w-[150px] truncate">{{ email }}</span>
          </div>
          <button
            @click="$emit('logout')"
            class="flex items-center gap-2 px-3 py-2 text-sm font-medium text-gray-700 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors"
          >
            <LogOut :size="16" />
            <span class="hidden sm:inline">Logout</span>
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { CheckCircle2, User, LogOut } from "lucide-vue-next";

defineProps({
  email: {
    type: String,
    default: "User",
  },
});

defineEmits(["logout"]);
</script>
