<!-- src/components/ui/Label.vue -->
<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface LabelProps {
  class?: string
  for?: string
}

const props = defineProps<LabelProps>()

const classes = computed(() => cn("text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70", props.class))
</script>

<template>
  <label :class="classes" v-bind="$attrs">
    <slot></slot>
  </label>
</template>
