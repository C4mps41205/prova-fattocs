<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface LoadingSpinnerProps {
  size?: "sm" | "md" | "lg"
  className?: string
}

const props = defineProps<LoadingSpinnerProps>()

const spinnerSizes = {
  sm: "h-4 w-4",
  md: "h-8 w-8",
  lg: "h-12 w-12",
}

const classes = computed(() =>
  cn(
    "animate-spin rounded-full border-b-2 border-gray-900",
    spinnerSizes[props.size ?? "md"],
    props.className
  )
)
</script>

<template>
  <div class="flex items-center justify-center py-8">
    <div :class="classes" />
  </div>
</template>
