<template>
  <div :class="alertClass" role="alert">
    <component :is="iconComponent" class="w-5 h-5 mt-0.5" />
    <div class="flex-1">
      <h4 v-if="title" class="font-semibold">{{ title }}</h4>
      <p v-if="description" class="text-sm text-muted-foreground">{{ description }}</p>
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue"
import type { PropType } from "vue"
import { AlertCircle, AlertTriangle, Info } from "lucide-vue-next"

const props = defineProps({
  variant: {
    type: String as PropType<"default" | "destructive" | "warning">,
    default: "default",
  },
  title: String,
  description: String,
  class: String,
})

const alertVariants = {
  default: "bg-blue-50 border-blue-200 text-blue-800",
  destructive: "bg-red-50 border-red-200 text-red-800",
  warning: "bg-yellow-50 border-yellow-200 text-yellow-800",
}

const alertIcons = {
  default: Info,
  destructive: AlertCircle,
  warning: AlertTriangle,
}

const iconComponent = computed(() => alertIcons[props.variant])
const alertClass = computed(() => `relative w-full rounded-lg border p-4 flex items-start gap-3 ${alertVariants[props.variant]} ${props.class ?? ""}`)
</script>