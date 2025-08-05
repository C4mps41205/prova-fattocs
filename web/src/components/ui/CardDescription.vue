<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface CardDescriptionProps {
  class?: string
}

const props = defineProps<CardDescriptionProps>()

const classes = computed(() => cn("text-muted-foreground text-sm", props.class))
</script>

<template>
  <div data-slot="card-description" :class="classes" v-bind="$attrs" />
</template>
