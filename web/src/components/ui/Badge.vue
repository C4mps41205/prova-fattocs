<script setup lang="ts">
import { computed } from "vue"
import type { PropType } from "vue"

const props = defineProps({
  variant: {
    type: String as PropType<"default" | "secondary" | "destructive" | "outline">,
    default: "default",
  },
  class: String,
})

const badgeVariants = {
  default: "bg-gray-900 text-white",
  secondary: "bg-gray-100 text-gray-900",
  destructive: "bg-red-100 text-red-800 border-red-200",
  outline: "border border-gray-200 bg-white text-gray-900",
}

const classes = computed(() => {
  return [
    "inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium",
    badgeVariants[props.variant],
    props.class ?? "",
  ].join(" ")
})
</script>

<template>
  <div :class="classes" v-bind="$attrs" />
</template>
