<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface CardProps {
  class?: string
}

const props = defineProps<CardProps>()

const classes = computed(() =>
  cn("rounded-lg border border-gray-200 bg-white shadow-sm", props.class)
)
</script>

<template>
  <div data-slot="card" :class="classes" v-bind="$attrs" />
</template>
