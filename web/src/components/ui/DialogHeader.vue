<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface DialogHeaderProps {
  class?: string
}

const props = defineProps<DialogHeaderProps>()

const classes = computed(() => cn("flex flex-col space-y-1.5 text-center sm:text-left", props.class))
</script>

<template>
  <div :class="classes" v-bind="$attrs">
    <slot></slot>
  </div>
</template>
