<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface AlertDialogFooterProps {
  class?: string
}

const props = defineProps<AlertDialogFooterProps>()

const classes = computed(() => cn("flex flex-col sm:flex-row sm:justify-end sm:space-x-2 space-y-2 sm:space-y-0 mt-6", props.class))
</script>

<template>
  <div :class="classes" v-bind="$attrs">
    <slot></slot>
  </div>
</template>
