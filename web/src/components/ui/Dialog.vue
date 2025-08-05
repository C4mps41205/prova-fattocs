<!-- src/components/ui/Dialog.vue -->
<script setup lang="ts">
import { computed, provide } from "vue"
import {cn} from "../../utils/utils.ts";

interface DialogProps {
  open?: boolean
  class?: string
}

const props = defineProps<DialogProps>()
const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
}>()

const classes = computed(() => cn("fixed inset-0 z-50 flex items-center justify-center", props.class))

const isOpen = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value)
})

provide('dialog', {
  isOpen,
  close: () => {
    isOpen.value = false
  }
})
</script>

<template>
  <div v-if="isOpen" :class="classes" v-bind="$attrs">
    <div class="fixed inset-0 bg-black bg-opacity-50" @click="isOpen = false"></div>
    <div class="relative bg-white rounded-lg shadow-xl sm:max-w-md w-[95vw] max-h-[90vh] overflow-y-auto p-6 z-10">
      <slot></slot>
    </div>
  </div>
</template>
