<!-- src/components/ui/AlertDialogTrigger.vue -->
<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface AlertDialogTriggerProps {
  asChild?: boolean
  class?: string
}

const props = defineProps<AlertDialogTriggerProps>()

const classes = computed(() => cn(props.class))
</script>

<template>
  <div :class="classes" v-bind="$attrs">
    <slot></slot>
  </div>
</template>
