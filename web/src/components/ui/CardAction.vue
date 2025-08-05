<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface CardActionProps {
  class?: string
}

const props = defineProps<CardActionProps>()

const classes = computed(() =>
  cn("col-start-2 row-span-2 row-start-1 self-start justify-self-end", props.class)
)
</script>

<template>
  <div data-slot="card-action" :class="classes" v-bind="$attrs" />
</template>
