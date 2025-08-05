<script setup lang="ts">
import { computed } from "vue"

interface AlertDialogProps {
  open?: boolean
  class?: string
}

const props = defineProps<AlertDialogProps>()

function cn(...classes: (string | undefined | false)[]) {
  return classes.filter(Boolean).join(" ")
}

const classes = computed(() => cn("fixed inset-0 z-50 flex items-center justify-center", props.class))
</script>

<template>
  <div v-if="open" :class="classes" v-bind="$attrs">
    <div class="fixed inset-0 bg-black bg-opacity-50" @click="$emit('update:open', false)"></div>
    <div class="relative bg-white rounded-lg shadow-xl max-w-md w-[95vw] p-6 z-10">
      <slot></slot>
    </div>
  </div>
</template>
