<script setup lang="ts">
import { computed } from "vue"

interface DialogContentProps {
  class?: string
}

const props = defineProps<DialogContentProps>()

function cn(...classes: (string | undefined | false)[]) {
  return classes.filter(Boolean).join(" ")
}

const classes = computed(() => cn("relative bg-white rounded-lg shadow-xl sm:max-w-md w-[95vw] max-h-[90vh] overflow-y-auto p-6", props.class))
</script>

<template>
  <div :class="classes" v-bind="$attrs">
    <slot></slot>
  </div>
</template>
