<script setup lang="ts">
import Button from "./Button.vue"

interface AlertDialogActionProps {
  class?: string
}

const props = defineProps<AlertDialogActionProps>()
</script>

<template>
  <Button :class="props.class" v-bind="$attrs">
    <slot></slot>
  </Button>
</template>
