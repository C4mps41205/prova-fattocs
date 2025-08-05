<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface CardFooterProps {
  class?: string
}

const props = defineProps<CardFooterProps>()

const classes = computed(() =>
  cn("flex items-center px-6 [.border-t]:pt-6", props.class)
)
</script>

<template>
  <div data-slot="card-footer" :class="classes" v-bind="$attrs" />
</template>
