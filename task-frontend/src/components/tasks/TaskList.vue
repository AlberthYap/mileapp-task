<template>
  <div class="space-y-3">
    <TransitionGroup name="list">
      <TaskItem
        v-for="task in tasks"
        :key="task.id"
        :task="task"
        @edit="$emit('edit', $event)"
        @delete="$emit('delete', $event)"
        @view="$emit('view', $event)"
      />
    </TransitionGroup>
  </div>
</template>

<script setup>
import TaskItem from "./TaskItem.vue";

defineProps({
  tasks: {
    type: Array,
    required: true,
  },
});

defineEmits(["edit", "delete", "view"]);
</script>

<style scoped>
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}

.list-enter-from {
  opacity: 0;
  transform: translateY(-10px);
}

.list-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}
</style>
