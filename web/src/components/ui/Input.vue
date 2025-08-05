<script setup lang="ts">
import { computed } from "vue"

interface InputProps {
  error?: boolean
  class?: string
}

const props = defineProps<InputProps>()
const emit = defineEmits<{
  (e: "input", value: string): void
}>()

function cn(...classes: (string | undefined | false)[]) {
  return classes.filter(Boolean).join(" ")
}

const classes = computed(() =>
    cn(
        "flex h-10 w-full rounded-md border px-3 py-2 text-sm",
        "file:border-0 file:bg-transparent file:text-sm file:font-medium",
        "placeholder:text-gray-500",
        "focus:outline-none focus:ring-2 focus:ring-offset-2",
        "disabled:cursor-not-allowed disabled:opacity-50",
        props.error
            ? "border-red-500 bg-red-50 focus:border-red-500 focus:ring-red-500"
            : "border-gray-300 bg-white focus:border-gray-900 focus:ring-gray-900",
        props.class
    )
)
</script>

<template>
  <input
      :class="classes"
      v-bind="$attrs"
      @input="(e: any) => emit('input', e)"
  />
</template>
