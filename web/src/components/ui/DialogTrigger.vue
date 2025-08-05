<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface DialogTriggerProps {
  asChild?: boolean
  class?: string
}

const props = defineProps<DialogTriggerProps>()

const classes = computed(() => cn(props.class))
</script>

<template>
  <div :class="classes" v-bind="$attrs" @click="$emit('click', $event)">
    <slot></slot>
  </div>
</template>
