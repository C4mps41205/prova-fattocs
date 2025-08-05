<script setup lang="ts">
import Button from "./Button.vue"
import {computed} from "vue";
import {cn} from "../../utils/utils.ts";

interface AlertDialogCancelProps {
  class?: string
}

const props = defineProps<AlertDialogCancelProps>()
const classes = computed(() => cn("mt-2 sm:mt-0", props.class))

</script>

<template>
  <Button :class="classes" variant="outline" v-bind="props">
    <slot></slot>
  </Button>
</template>
