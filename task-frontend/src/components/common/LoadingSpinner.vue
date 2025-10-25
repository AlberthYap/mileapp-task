<template>
  <div class="flex justify-center py-12">
    <Loader2 :size="size" class="text-[#fd9621] animate-spin" />
  </div>
</template>

<script setup>
import { Loader2 } from "lucide-vue-next";

defineProps({
  size: {
    type: Number,
    default: 40,
  },
});
</script>
