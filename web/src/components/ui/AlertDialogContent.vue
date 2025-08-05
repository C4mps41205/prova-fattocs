<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface AlertDialogContentProps {
  class?: string
}

const props = defineProps<AlertDialogContentProps>()

const classes = computed(() => cn("relative bg-white rounded-lg shadow-xl w-[95vw] max-w-md p-6", props.class))
</script>

<template>
  <div :class="classes" v-bind="props">
    <slot></slot>
  </div>
</template>
