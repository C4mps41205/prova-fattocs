<script setup lang="ts">
import { watch, onBeforeUnmount} from "vue"
import { X } from "lucide-vue-next"
import {cn} from "../../utils/utils.ts";

interface ModalProps {
  isOpen: boolean
  onClose: () => void
  title?: string
  className?: string
}

const props = defineProps<ModalProps>()
const emit = defineEmits<{
  (e: "close"): void
}>()

const handleEscape = (e: KeyboardEvent) => {
  if (e.key === "Escape") {
    close()
  }
}

const close = () => {
  props.onClose()
  emit("close")
}

watch(
  () => props.isOpen,
  (newVal) => {
    if (newVal) {
      document.body.style.overflow = "hidden"
      document.addEventListener("keydown", handleEscape)
    } else {
      document.body.style.overflow = "unset"
      document.removeEventListener("keydown", handleEscape)
    }
  },
  { immediate: true }
)

onBeforeUnmount(() => {
  document.body.style.overflow = "unset"
  document.removeEventListener("keydown", handleEscape)
})
</script>

<template>
  <Teleport to="body">
    <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center">
      <div
        class="absolute inset-0 bg-black bg-opacity-50 transition-opacity"
        @click="close"
      />

      <div :class="cn('relative bg-white rounded-lg shadow-xl w-[95vw] max-w-md max-h-[90vh] overflow-y-auto', props.className)">
        <div v-if="props.title" class="flex items-center justify-between p-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">{{ props.title }}</h2>
          <button @click="close" class="p-1 hover:bg-gray-100 rounded-md transition-colors" type="button">
            <X class="w-5 h-5 text-gray-500" />
          </button>
        </div>

        <div class="p-4">
          <slot />
        </div>
      </div>
    </div>
  </Teleport>
</template>
