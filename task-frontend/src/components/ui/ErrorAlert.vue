<template>
  <Transition name="fade">
    <div
      v-if="message"
      class="mb-6 p-4 bg-red-50 border-l-4 border-red-500 rounded-r-lg"
    >
      <div class="flex items-start gap-3">
        <AlertCircle :size="20" class="text-red-500 shrink-0 mt-0.5" />
        <div class="flex-1">
          <p class="text-sm font-medium text-red-800">{{ message }}</p>
        </div>
        <button @click="$emit('close')">
          <X :size="18" class="text-red-400 hover:text-red-600" />
        </button>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { AlertCircle, X } from "lucide-vue-next";

defineProps({
  message: {
    type: String,
    default: "",
  },
});

defineEmits(["close"]);
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: all 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
