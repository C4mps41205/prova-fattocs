<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3 sm:gap-4 mb-6 sm:mb-8">
    <Card v-for="(stat, index) in stats" :key="index" class="border-gray-200">
      <CardContent class="p-3 sm:p-4">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-gray-600 text-xs sm:text-sm">{{ stat.label }}</p>
            <p class="text-xl sm:text-2xl font-semibold text-gray-900">{{ stat.value }}</p>
          </div>
          <component :is="stat.icon" class="w-6 h-6 sm:w-8 sm:h-8 text-gray-400" />
        </div>
      </CardContent>
    </Card>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import Card from '../../components/ui/Card.vue'
import  CardContent  from '../../components/ui/CardContent.vue'
import { CheckSquare, DollarSign, Clock } from 'lucide-vue-next'
import type { Task } from '../../types/task'

export default defineComponent({
  name: 'TaskStats',
  props: {
    tasks: {
      type: Array as () => Task[],
      required: true
    }
  },
  components: {
    Card,
    CardContent,
    CheckSquare,
    DollarSign,
    Clock
  },
  setup(props) {
    const isExpensive = (cost: number) => cost >= 1000

    const isUrgent = (dateString: string) => {
      const today = new Date()
      const taskDate = new Date(dateString)
      const diffTime = taskDate.getTime() - today.getTime()
      const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
      return diffDays <= 7 && diffDays >= 0
    }

    const stats = computed(() => [
      {
        label: 'Total Tasks',
        value: props.tasks.length,
        icon: CheckSquare
      },
      {
        label: 'High Cost',
        value: props.tasks.filter(t => isExpensive(t.cost)).length,
        icon: DollarSign
      },
      {
        label: 'Due Soon',
        value: props.tasks.filter(t => isUrgent(t.deadline)).length,
        icon: Clock
      }
    ])

    return {
      stats
    }
  }
})
</script>
