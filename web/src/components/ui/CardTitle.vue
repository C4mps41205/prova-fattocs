<!-- src/components/ui/CardTitle.vue -->
<script setup lang="ts">
import { computed } from "vue"
import {cn} from "../../utils/utils.ts";

interface CardTitleProps {
  class?: string
}

const props = defineProps<CardTitleProps>()

const classes = computed(() => cn("leading-none font-semibold", props.class))
</script>

<template>
  <div data-slot="card-title" :class="classes" v-bind="$attrs" />
</template>
