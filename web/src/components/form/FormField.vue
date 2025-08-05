<template>
  <div class="space-y-2">
    <label class="flex items-center space-x-2 text-sm font-medium text-gray-700">
      <slot name="icon" />
      <span>{{ label }}</span>
      <span v-if="required" class="text-red-500">*</span>
    </label>

    <slot>
      <Input v-bind="inputProps" :error="error" />
    </slot>

    <div v-if="error" class="flex items-center space-x-1 text-sm text-red-600">
      <AlertCircle class="w-4 h-4" />
      <span>{{ error }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { AlertCircle } from 'lucide-vue-next'
import Input from "../ui/Input.vue"

interface Props {
  label: string
  error?: string
  required?: boolean
  inputProps?: Record<string, any>
}

withDefaults(defineProps<Props>(), {
  error: '',
  required: false,
  inputProps: () => ({}),
})
</script>
