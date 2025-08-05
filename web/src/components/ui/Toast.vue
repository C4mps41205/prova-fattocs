<script setup lang="ts">
import { computed, onMounted, onUnmounted } from 'vue'
import type { Toast } from '../../composables/useToast'
import Alert from './Alert.vue'
import Button from './Button.vue'
import { X } from 'lucide-vue-next'

interface Props {
  toast: Toast
  onDismiss: (id: number) => void
}

const props = defineProps<Props>()

const alertVariant = computed(() => {
  switch (props.toast.variant) {
    case 'success':
      return 'default'
    case 'destructive':
      return 'destructive'
    case 'warning':
      return 'warning'
    default:
      return 'default'
  }
})

let timer: number | null = null

onMounted(() => {
  const duration = props.toast.duration || 5000
  timer = window.setTimeout(() => {
    props.onDismiss(props.toast.id)
  }, duration)
})

onUnmounted(() => {
  if (timer) {
    clearTimeout(timer)
  }
})

const handleDismiss = () => {
  if (timer) {
    clearTimeout(timer)
  }
  props.onDismiss(props.toast.id)
}
</script>

<template>
  <div class="relative">
    <Alert
      :variant="alertVariant"
      :title="toast.title"
      :description="toast.description"
      class="pr-12"
    />
    <Button
      variant="ghost"
      size="sm"
      class="absolute top-2 right-2 h-6 w-6 p-0"
      @click="handleDismiss"
    >
      <X class="h-4 w-4" />
    </Button>
  </div>
</template>
