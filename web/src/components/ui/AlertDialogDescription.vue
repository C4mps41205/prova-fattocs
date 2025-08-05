<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface AlertDialogDescriptionProps {
  class?: string
}

const props = defineProps<AlertDialogDescriptionProps>()

const classes = computed(() => cn("text-sm text-gray-500", props.class))
</script>

<template>
  <p :class="classes" v-bind="$attrs">
    <slot></slot>
  </p>
</template>
