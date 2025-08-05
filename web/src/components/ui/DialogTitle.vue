<script setup lang="ts">
import { computed } from "vue"

interface DialogTitleProps {
  class?: string
}

const props = defineProps<DialogTitleProps>()

function cn(...classes: (string | undefined | false)[]) {
  return classes.filter(Boolean).join(" ")
}

const classes = computed(() => cn("text-lg font-semibold leading-none tracking-tight", props.class))
</script>

<template>
  <h2 :class="classes" v-bind="$attrs">
    <slot></slot>
  </h2>
</template>
