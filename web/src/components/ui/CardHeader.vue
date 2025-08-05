<!-- src/components/ui/CardHeader.vue -->
<script setup lang="ts">
import { computed } from "vue"

interface CardHeaderProps {
  class?: string
}

const props = defineProps<CardHeaderProps>()

function cn(...classes: (string | undefined | false)[]) {
  return classes.filter(Boolean).join(" ")
}

const classes = computed(() =>
  cn(
    "@container/card-header grid auto-rows-min grid-rows-[auto_auto] items-start gap-1.5 px-6 has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-6",
    props.class
  )
)
</script>

<template>
  <div data-slot="card-header" :class="classes" v-bind="$attrs" />
</template>
