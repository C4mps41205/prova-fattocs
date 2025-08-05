<script setup lang="ts">
import { ref } from 'vue'
import {
  GripVertical,
  ChevronUp,
  ChevronDown,
  Calendar,
  DollarSign,
  AlertTriangle,
  Trash2
} from 'lucide-vue-next'
import  Card  from '../../components/ui/Card.vue'
import  CardContent  from '../../components/ui/CardContent.vue'
import  Badge  from '../../components/ui/Badge.vue'
import  Button  from '../../components/ui/Button.vue'
import  Modal  from '../../components/ui/Modal.vue'
import TaskForm from './TaskForm.vue'
import type { Task } from '../../types/task.ts'
import { useTasks } from '../../composables/useTasks.ts'
import { useToast } from '../../composables/useToast.ts'

const props = defineProps<{
  task: Task
  index: number
  totalTasks: number
  onMoveUp: () => void
  onMoveDown: () => void
  dragHandleProps: any
  isDragging: boolean
}>()

const { deleteTask, loading } = useTasks()
const { addToast } = useToast()

const showDeleteModal = ref(false)

const formatCurrency = (value: number) =>
    new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD'
    }).format(value)

const formatDate = (dateString: string) =>
    new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    })

const isExpensive = (cost: number) => cost >= 1000

const isUrgent = (dateString: string) => {
  const today = new Date()
  const taskDate = new Date(dateString)
  const diffDays = Math.ceil((taskDate.getTime() - today.getTime()) / (1000 * 60 * 60 * 24))
  return diffDays <= 7 && diffDays >= 0
}

const handleDelete = async () => {
  try {
    await deleteTask(props.task.id)
    showDeleteModal.value = false
    addToast({
      title: "Task Deleted",
      description: "Task deleted successfully",
      variant: "success"
    })
  } catch (error: any) {
    addToast({
      title: "Error",
      description: error.message || "Failed to delete task",
      variant: "destructive"
    })
  }
}
</script>


<template>
  <div>
    <Card
        :class="[
        'transition-all duration-200 hover:shadow-md',
        isExpensive(task.cost)
          ? 'bg-yellow-50 border-yellow-200 border-l-4 border-l-yellow-400'
          : 'bg-white border-gray-200',
        isDragging ? 'shadow-lg' : ''
      ]"
    >
      <CardContent class="p-3 sm:p-4 lg:p-6">
        <div class="flex items-start justify-between">
          <div class="flex items-start space-x-2 sm:space-x-3 flex-1 min-w-0">
            <div
                v-bind="dragHandleProps"
                class="cursor-grab hover:cursor-grabbing mt-1 p-1 rounded hover:bg-gray-100 transition-colors flex-shrink-0"
            >
              <GripVertical class="w-4 h-4 text-gray-400" />
            </div>

            <div class="flex-1 space-y-2 min-w-0">
              <div class="flex flex-col sm:flex-row sm:items-start sm:justify-between gap-2">
                <h3 class="font-medium text-gray-900 leading-tight text-sm sm:text-base break-words">
                  {{ task.name }}
                </h3>
                <div class="flex items-center space-x-2 flex-shrink-0">
                  <Badge
                      v-if="isExpensive(task.cost)"
                      variant="secondary"
                      class="bg-yellow-100 text-yellow-800 border-yellow-200 text-xs"
                  >
                    High Cost
                  </Badge>
                  <Badge v-if="isUrgent(task.deadline)" variant="destructive" class="text-xs">
                    Due Soon
                  </Badge>
                </div>
              </div>

              <div
                  class="flex flex-col sm:flex-row sm:items-center sm:space-x-6 space-y-1 sm:space-y-0 text-xs sm:text-sm text-gray-600"
              >
                <div class="flex items-center space-x-1">
                  <DollarSign class="w-3 h-3 sm:w-4 sm:h-4 flex-shrink-0" />
                  <span class="font-medium">{{ formatCurrency(task.cost) }}</span>
                </div>
                <div class="flex items-center space-x-1">
                  <Calendar class="w-3 h-3 sm:w-4 sm:h-4 flex-shrink-0" />
                  <span>{{ formatDate(task.deadline) }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="flex items-center space-x-1 ml-2 sm:ml-4 flex-shrink-0">
            <div class="hidden sm:flex flex-col space-y-1">
              <Button variant="ghost" size="sm" @click="onMoveUp" :disabled="index === 0" class="h-7 w-7 p-0">
                <ChevronUp class="w-4 h-4" />
              </Button>
              <Button
                  variant="ghost"
                  size="sm"
                  @click="onMoveDown"
                  :disabled="index === totalTasks - 1"
                  class="h-7 w-7 p-0"
              >
                <ChevronDown class="w-4 h-4" />
              </Button>
            </div>

            <TaskForm :existing-task="task" />

            <Button
                variant="ghost"
                size="sm"
                @click="showDeleteModal = true"
                class="h-7 w-7 p-0 hover:bg-red-50 hover:text-red-600"
            >
              <Trash2 class="w-4 h-4" />
            </Button>
          </div>
        </div>

        <div class="flex sm:hidden justify-center space-x-4 mt-3 pt-3 border-t border-gray-100">
          <Button
              variant="ghost"
              size="sm"
              @click="onMoveUp"
              :disabled="index === 0"
              class="flex items-center space-x-1 text-xs"
          >
            <ChevronUp class="w-4 h-4" />
            <span>Move Up</span>
          </Button>
          <Button
              variant="ghost"
              size="sm"
              @click="onMoveDown"
              :disabled="index === totalTasks - 1"
              class="flex items-center space-x-1 text-xs"
          >
            <ChevronDown class="w-4 h-4" />
            <span>Move Down</span>
          </Button>
        </div>
      </CardContent>
    </Card>

    <Modal :is-open="showDeleteModal" @close="showDeleteModal = false" title="Delete Task">
      <div class="space-y-4">
        <div class="flex items-center space-x-2 text-red-600">
          <AlertTriangle class="w-5 h-5 flex-shrink-0" />
          <p class="text-sm">
            Are you sure you want to delete <strong>"{{ task.name }}"</strong>? This action cannot be undone.
          </p>
        </div>

        <div class="flex flex-col sm:flex-row justify-end space-y-2 sm:space-y-0 sm:space-x-3">
          <Button variant="outline" @click="showDeleteModal = false" class="w-full sm:w-auto">
            Cancel
          </Button>
          <Button variant="destructive" @click="handleDelete" :disabled="loading" class="w-full sm:w-auto">
            Delete Task
          </Button>
        </div>
      </div>
    </Modal>
  </div>
</template>