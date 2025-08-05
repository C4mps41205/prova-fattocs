<script setup lang="ts">
import {onMounted, ref} from "vue"
import draggable from "vuedraggable"
import TaskForm from "./TaskForm.vue"

import {
  AlertTriangle,
  Calendar,
  CheckSquare,
  ChevronUp,
  ChevronDown,
  DollarSign,
  GripVertical,
  Trash2,
  Edit3
} from "lucide-vue-next"

import AlertDialog from "../../components/ui/AlertDialog.vue"
import AlertDialogAction from "../../components/ui/AlertDialogAction.vue"
import AlertDialogCancel from "../../components/ui/AlertDialogCancel.vue"
import AlertDialogContent from "../../components/ui/AlertDialogContent.vue"
import AlertDialogDescription from "../../components/ui/AlertDialogDescription.vue"
import AlertDialogFooter from "../../components/ui/AlertDialogFooter.vue"
import AlertDialogHeader from "../../components/ui/AlertDialogHeader.vue"
import AlertDialogTitle from "../../components/ui/AlertDialogTitle.vue"
import AlertDialogTrigger from "../../components/ui/AlertDialogTrigger.vue"

import Card from "../../components/ui/Card.vue"
import CardContent from "../../components/ui/CardContent.vue"
import Badge from "../../components/ui/Badge.vue"
import Button from "../../components/ui/Button.vue"

import {useTasks} from "../../composables/useTasks.ts"
import { useToast } from "../../composables/useToast.ts"

const { tasks, loading, loadTasks, deleteTask, reorderTasks } = useTasks()
const { addToast } = useToast()

const isReordering = ref(false)

const deleteDialogOpen = ref<{[key: number]: boolean}>({})

async function handleDelete(id: number) {
  try {
    await deleteTask(id)
    deleteDialogOpen.value[id] = false
    await loadTasks()
  } catch (error: any) {
    console.error("Error deleting task:", error)
    addToast({
      title: "Error",
      description: error.message || "Failed to delete task",
      variant: "destructive"
    })
    await loadTasks()
  }
}

async function moveTask(taskId: number, direction: "up" | "down") {
  const currentIndex = tasks.value.findIndex(t => t.id === taskId)
  if (currentIndex === -1) return

  const targetIndex = direction === "up" ? currentIndex - 1 : currentIndex + 1
  if (targetIndex < 0 || targetIndex >= tasks.value.length) return

  const newTasks = [...tasks.value]
  const temp = newTasks[currentIndex]
  newTasks[currentIndex] = newTasks[targetIndex]
  newTasks[targetIndex] = temp

  const reorderedTasks = newTasks.map((task, index) => ({
    ...task,
    order_number: index + 1
  }))

  try {
    if (isReordering.value) {
      addToast({
        title: "Info",
        description: "Please wait for the current operation to complete",
        variant: "warning"
      })
      return
    }

    isReordering.value = true
    const result = await reorderTasks(reorderedTasks)
    if (!result) {
      throw new Error("Failed to reorder tasks")
    }
    await loadTasks()
  } catch (error: any) {
    console.error("Error reordering task:", error)
    addToast({
      title: "Error",
      description: error.message || "Failed to reorder tasks",
      variant: "destructive"
    })
    await loadTasks()
  } finally {
    isReordering.value = false
  }
}

const handleDragEnd = async () => {
  try {
    const reorderedTasks = tasks.value.map((task, index) => ({
      ...task,
      order_number: index + 1
    }))

    if (isReordering.value) {
      addToast({
        title: "Info",
        description: "Please wait for the current operation to complete",
        variant: "warning"
      })
      await loadTasks()
      return
    }

    isReordering.value = true
    const result = await reorderTasks(reorderedTasks)
    if (!result) {
      throw new Error("Failed to reorder tasks")
    }
    await loadTasks()
  } catch (error: any) {
    console.error("Error reordering tasks:", error)
    addToast({
      title: "Error",
      description: error.message || "Failed to reorder tasks",
      variant: "destructive"
    })
    await loadTasks()
  } finally {
    isReordering.value = false
  }
}

function formatCurrency(value: number) {
  return new Intl.NumberFormat("en-US", { style: "currency", currency: "USD" }).format(value)
}

function formatDate(dateString: string) {
  if (!dateString) return ''

  const date = new Date(dateString)

  const year = date.getUTCFullYear()
  const month = date.getUTCMonth()
  const day = date.getUTCDate()

  const utcDate = new Date(Date.UTC(year, month, day))

  return utcDate.toLocaleDateString("en-US", {
    year: "numeric",
    month: "short",
    day: "numeric",
    timeZone: "UTC"
  })
}

function isExpensive(cost: number) {
  return cost >= 1000
}

function isUrgent(dateString: string) {
  const today = new Date()
  const taskDate = new Date(dateString)
  const diffDays = Math.ceil((taskDate.getTime() - today.getTime()) / (1000 * 60 * 60 * 24))
  return diffDays <= 7 && diffDays >= 0
}

onMounted(() => {
  loadTasks()
})
</script>

<template>
  <div class="space-y-6">
    <div class="space-y-3">
      <draggable
          v-model="tasks"
          item-key="id"
          handle=".drag-handle"
          @end="handleDragEnd"
          tag="div"
          class="space-y-3"
          :disabled="isReordering"
      >
        <template #item="{ element: task, index }">
          <div
              :class="[
              'transition-all duration-200 hover:shadow-md border rounded-lg',
              isExpensive(task.cost)
                ? 'bg-yellow-50 border-yellow-200 border-l-4 border-l-yellow-400'
                : 'bg-white border-gray-200'
            ]"
          >
            <div class="p-3 sm:p-4 lg:p-6">
              <div class="flex items-start justify-between">
                <div class="flex items-start space-x-2 sm:space-x-3 flex-1 min-w-0">
                  <div
                    class="drag-handle mt-1 p-1 rounded hover:bg-gray-100 transition-colors flex-shrink-0 cursor-grab"
                    :class="{ 'opacity-50 cursor-not-allowed': isReordering }"
                  >
                    <GripVertical class="w-5 h-5 text-gray-500" />
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
                            class="bg-yellow-500 text-yellow-800 border-yellow-200 text-xs"
                        >
                          High Cost
                        </Badge>
                        <Badge
                            v-if="isUrgent(task.deadline)"
                            variant="destructive"
                            class="bg-red-500 text-red-800 border-red-200 text-xs"
                        >
                          Due Soon
                        </Badge>
                      </div>
                    </div>

                    <div
                        class="flex flex-col sm:flex-row sm:items-center sm:space-x-6 space-y-1 sm:space-y-0 text-xs sm:text-sm text-gray-600"
                    >
                      <div class="flex items-center space-x-1">
                        <DollarSign class="w-4 h-4 sm:w-5 sm:h-5 flex-shrink-0 text-gray-500" />
                        <span class="font-medium">{{ formatCurrency(task.cost) }}</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <Calendar class="w-4 h-4 sm:w-5 sm:h-5 flex-shrink-0 text-gray-500" />
                        <span>{{ formatDate(task.deadline) }}</span>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="flex items-center space-x-1 ml-2 sm:ml-4 flex-shrink-0">
                  <div class="hidden sm:flex flex-col space-y-1">
                    <Button
                        variant="ghost"
                        size="sm"
                        :disabled="index === 0 || isReordering"
                        @click="() => moveTask(task.id, 'up')"
                        class="h-8 w-8 p-0"
                        :title="isReordering ? 'Operation in progress...' : 'Move up'"
                    >
                      <ChevronUp class="w-5 h-5 text-gray-600" />
                    </Button>
                    <Button
                        variant="ghost"
                        size="sm"
                        :disabled="index === tasks.length - 1 || isReordering"
                        @click="() => moveTask(task.id, 'down')"
                        class="h-8 w-8 p-0"
                        :title="isReordering ? 'Operation in progress...' : 'Move down'"
                    >
                      <ChevronDown class="w-5 h-5 text-gray-600" />
                    </Button>
                  </div>

                  <TaskForm :existing-task="task" @task-updated="loadTasks" @task-deleted="loadTasks">
                    <template #trigger>
                      <Button
                        variant="ghost"
                        size="sm"
                        class="h-8 w-8 p-0 hover:bg-blue-50 hover:text-blue-600"
                        title="Edit task"
                        :disabled="isReordering"
                      >
                        <Edit3 class="w-5 h-5 text-gray-600" />
                      </Button>
                    </template>
                  </TaskForm>

                  <AlertDialog v-model:open="deleteDialogOpen[task.id]">
                    <AlertDialogTrigger asChild>
                      <Button
                          variant="ghost"
                          size="sm"
                          class="h-8 w-8 p-0 hover:bg-red-50 hover:text-red-600"
                          title="Delete task"
                          :disabled="isReordering"
                      >
                        <Trash2 class="w-5 h-5 text-gray-600" />
                      </Button>
                    </AlertDialogTrigger>
                    <AlertDialogContent class="w-[95vw] max-w-md">
                      <AlertDialogHeader>
                        <AlertDialogTitle class="flex items-center space-x-2 text-base sm:text-lg">
                          <AlertTriangle class="w-5 h-5 text-red-500 flex-shrink-0" />
                          <span>Delete Task</span>
                        </AlertDialogTitle>
                        <AlertDialogDescription class="text-sm">
                          Are you sure you want to delete <strong>{{ task.name }}</strong>? This action
                          cannot be undone.
                        </AlertDialogDescription>
                      </AlertDialogHeader>
                      <AlertDialogFooter class="flex-col sm:flex-row space-y-2 sm:space-y-0 sm:space-x-2">
                        <AlertDialogCancel
                          class="w-full sm:w-auto"
                          @click="deleteDialogOpen[task.id] = false"
                          :disabled="isReordering"
                        >
                          Cancel
                        </AlertDialogCancel>
                        <AlertDialogAction
                            class="w-full sm:w-auto bg-red-600 hover:bg-red-700"
                            @click="() => handleDelete(task.id)"
                            :disabled="isReordering"
                        >
                          Delete Task
                        </AlertDialogAction>
                      </AlertDialogFooter>
                    </AlertDialogContent>
                  </AlertDialog>
                </div>
              </div>

              <div class="flex sm:hidden justify-center space-x-4 mt-3 pt-3 border-t border-gray-100">
                <Button
                    variant="ghost"
                    size="sm"
                    :disabled="index === 0 || isReordering"
                    @click="() => moveTask(task.id, 'up')"
                    class="flex items-center space-x-1 text-xs h-8 px-3"
                >
                  <ChevronUp class="w-4 h-4" />
                  <span>Move Up</span>
                </Button>
                <Button
                    variant="ghost"
                    size="sm"
                    :disabled="index === tasks.length - 1 || isReordering"
                    @click="() => moveTask(task.id, 'down')"
                    class="flex items-center space-x-1 text-xs h-8 px-3"
                >
                  <ChevronDown class="w-4 h-4" />
                  <span>Move Down</span>
                </Button>
              </div>
            </div>
          </div>
        </template>
      </draggable>
    </div>

    <Card v-if="tasks.length === 0 && !loading" class="border-2 border-dashed border-gray-300 bg-gray-50">
      <CardContent
          class="flex flex-col items-center justify-center py-12 sm:py-16 text-center px-4"
      >
        <CheckSquare class="w-12 h-12 text-gray-400 mb-4" />
        <h3 class="text-lg font-medium text-gray-900 mb-2">No tasks found</h3>
        <p class="text-gray-600 mb-6 max-w-md text-sm sm:text-base">
          Get started by creating your first task to organize your work efficiently.
        </p>
      </CardContent>
    </Card>

    <div class="flex justify-center">
      <TaskForm
          @task-added="loadTasks"
          @task-updated="loadTasks"
          @task-deleted="loadTasks"
          :disabled="isReordering"
      />
    </div>
  </div>
</template>
