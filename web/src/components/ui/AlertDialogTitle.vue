<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface AlertDialogTitleProps {
  class?: string
}

const props = defineProps<AlertDialogTitleProps>()

const classes = computed(() => cn("text-lg font-semibold text-gray-900", props.class))
</script>

<template>
  <h2 :class="classes" v-bind="$attrs">
    <slot></slot>
  </h2>
</template>
