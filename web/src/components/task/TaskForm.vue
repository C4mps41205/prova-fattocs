<script setup lang="ts">
import { ref, computed, h, watch } from 'vue'
import { Plus, Edit3, Save, X, DollarSign, Calendar, FileText, Trash2 } from 'lucide-vue-next'
import Button from '../../components/ui/Button.vue'
import Modal from '../../components/ui/Modal.vue'
import FormField from '../../components/form/FormField.vue'
import type { Task, TaskFormData, FormErrors } from '../../types/task'
import { useTasks } from '../../composables/useTasks'
import { useToast } from '../../composables/useToast'
import { useMask } from '../../composables/useMask'

interface Props {
  existingTask?: Task
}

interface Emits {
  (e: 'task-added'): void
  (e: 'task-updated'): void
  (e: 'task-deleted'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const isOpen = ref(false)
const errors = ref<FormErrors>({})
const isSubmitting = ref(false)
const isDeleting = ref(false)

const formData = ref<TaskFormData>({
  name: '',
  cost: '',
  deadline: '',
})

const { createTask, updateTask, deleteTask } = useTasks()
const { addToast } = useToast()
const { formatCurrencyInput } = useMask()

const resetForm = () => {
  formData.value = {
    name: '',
    cost: '',
    deadline: '',
  }
  errors.value = {}
}

const formatDateForInput = (dateString: string): string => {
  if (!dateString) return ''

  const date = new Date(dateString)

  const year = date.getUTCFullYear()
  const month = String(date.getUTCMonth() + 1).padStart(2, '0')
  const day = String(date.getUTCDate()).padStart(2, '0')

  return `${year}-${month}-${day}`
}

watch(
  () => props.existingTask,
  (newTask) => {
    if (newTask) {
      formData.value = {
        name: newTask.name || '',
        cost: newTask.cost?.toString() || '',
        deadline: formatDateForInput(newTask.deadline) || '',
      }
    } else {
      resetForm()
    }
  },
  { immediate: true }
)

const isEditMode = computed(() => !!props.existingTask)

const validateForm = (): boolean => {
  const newErrors: FormErrors = {}

  if (!formData.value.name.trim()) {
    newErrors.name = 'Task name is required'
  } else if (formData.value.name.trim().length < 3) {
    newErrors.name = 'Task name must be at least 3 characters'
  } else if (formData.value.name.trim().length > 100) {
    newErrors.name = 'Task name must be less than 100 characters'
  }

  if (!formData.value.cost.trim()) {
    newErrors.cost = 'Cost is required'
  } else {
    const cost = parseFloat(formData.value.cost)
    if (isNaN(cost)) {
      newErrors.cost = 'Cost must be a valid number'
    } else if (cost < 0) {
      newErrors.cost = 'Cost cannot be negative'
    } else if (cost > 999999.99) {
      newErrors.cost = 'Cost cannot exceed $999,999.99'
    }
  }

  if (!formData.value.deadline.trim()) {
    newErrors.deadline = 'Due date is required'
  } else {
    const selectedDate = new Date(formData.value.deadline)
    const today = new Date()
    today.setHours(0, 0, 0, 0)
    if (selectedDate < today) {
      newErrors.deadline = 'Due date cannot be in the past'
    }
  }

  errors.value = newErrors
  return Object.keys(newErrors).length === 0
}

const handleInputChange = (field: keyof TaskFormData, value: string) => {
  if (field === 'cost') {
    const onlyNumbers = value.replace(/\D/g, '');

    const formatted = formatCurrencyInput(onlyNumbers);

    formData.value[field] = formatted;
  } else {
    formData.value[field] = value;
  }

  if (errors.value[field]) {
    delete errors.value[field];
  }
};

const handleSubmit = async () => {
  if (!validateForm() || isSubmitting.value) return

  isSubmitting.value = true

  try {
    const cost = parseFloat(formData.value.cost)

    if (isEditMode.value) {
      const updatedTask = await updateTask(props.existingTask!.id, {
        name: formData.value.name.trim(),
        cost,
        deadline: formData.value.deadline,
      })

      if (updatedTask) {
        addToast({
          title: "Success",
          description: "Task updated successfully",
          variant: "success"
        })
        emit('task-updated')
      }
    } else {
      const newTask = await createTask({
        name: formData.value.name.trim(),
        cost,
        deadline: formData.value.deadline,
      })

      if (newTask) {
        addToast({
          title: "Success",
          description: "Task created successfully",
          variant: "success"
        })
        emit('task-added')
      }
    }

    handleClose()
  } catch (error: any) {
    addToast({
      title: "Error",
      description: error.message || "An error occurred",
      variant: "destructive"
    })
  } finally {
    isSubmitting.value = false
  }
}

const handleDelete = async () => {
  if (!isEditMode.value || !props.existingTask) return

  if (!confirm(`Are you sure you want to delete the task "${props.existingTask.name}"?`)) {
    return
  }

  isDeleting.value = true

  try {
    await deleteTask(props.existingTask.id)
    addToast({
      title: "Success",
      description: "Task deleted successfully",
      variant: "success"
    })
    emit('task-deleted')
    handleClose()
  } catch (error: any) {
    addToast({
      title: "Error",
      description: error.message || "Failed to delete task",
      variant: "destructive"
    })
  } finally {
    isDeleting.value = false
  }
}

const handleClose = () => {
  isOpen.value = false
  if (!isEditMode.value) {
    resetForm()
  }
  errors.value = {}
}

watch(isEditMode, (newIsEditMode) => {
  if (!newIsEditMode) {
    resetForm()
  }
})

const defaultTrigger = computed(() => {
  return isEditMode.value ? (
    h(Button, { variant: "ghost", size: "sm", class: "h-7 w-7 p-0" }, () => [
      h(Edit3, { class: "w-4 h-4" })
    ])
  ) : (
    h(Button, { class: "w-full sm:w-auto" }, () => [
      h(Plus, { class: "w-4 h-4 mr-2" }),
      h("span", { class: "hidden sm:inline" }, "Add New Task"),
      h("span", { class: "sm:hidden" }, "Add Task")
    ])
  )
})
</script>

<template>
  <div @click="isOpen = true">
    <slot name="trigger">
      <component :is="defaultTrigger" />
    </slot>
  </div>

  <Modal :is-open="isOpen" @close="handleClose" :title="isEditMode ? 'Edit Task' : 'New Task'">
    <form @submit.prevent="handleSubmit" class="space-y-5">
      <FormField
        label="Task Name"
        :error="errors.name"
        :required="true"
        :input-props="{
          value: formData.name,
          placeholder: 'e.g., Develop new feature'
        }"
        @input="(value: any) => handleInputChange('name', value.target.value)"
      >
        <template #icon>
          <FileText class="w-4 h-4" />
        </template>
      </FormField>

      <FormField
        label="Cost ($)"
        :error="errors.cost"
        :required="true"
        :input-props="{
          type: 'text',
          value: formData.cost,
          placeholder: '0.00'
        }"
        @input="(value: any) => handleInputChange('cost', value.target.value)"
      >
        <template #icon>
          <DollarSign class="w-4 h-4" />
        </template>
      </FormField>

      <FormField
        label="Due Date"
        :error="errors.deadline"
        :required="true"
        :input-props="{
          type: 'date',
          value: formData.deadline
        }"
        @input="(value: any) => handleInputChange('deadline', value.target.value)"
      >
        <template #icon>
          <Calendar class="w-4 h-4" />
        </template>
      </FormField>

      <div class="flex flex-col sm:flex-row justify-end space-y-2 sm:space-y-0 sm:space-x-3 pt-4">
        <div class="flex flex-col sm:flex-row sm:space-x-3 space-y-2 sm:space-y-0 w-full sm:w-auto">
          <!-- Botão de exclusão para modo de edição -->
          <Button
            v-if="isEditMode"
            type="button"
            variant="destructive"
            :disabled="isDeleting"
            @click="handleDelete"
            class="w-full sm:w-auto"
          >
            <template v-if="isDeleting">
              <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
              <span>Deleting...</span>
            </template>
            <template v-else>
              <Trash2 class="w-4 h-4 mr-2" />
              <span>Delete Task</span>
            </template>
          </Button>

          <div class="flex flex-col sm:flex-row sm:space-x-3 space-y-2 sm:space-y-0 sm:ml-auto w-full sm:w-auto">
            <Button type="button" variant="outline" @click="handleClose" class="w-full sm:w-auto">
              <X class="w-4 h-4 mr-2" /> Cancel
            </Button>
            <Button type="submit" :disabled="isSubmitting" class="w-full sm:w-auto">
              <template v-if="isSubmitting">
                <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                <span>Processing...</span>
              </template>
              <template v-else>
                <Save class="w-4 h-4 mr-2" />
                <span class="hidden sm:inline">{{ isEditMode ? 'Save Changes' : 'Create Task' }}</span>
                <span class="sm:hidden">{{ isEditMode ? 'Save' : 'Create' }}</span>
              </template>
            </Button>
          </div>
        </div>
      </div>
    </form>
  </Modal>
</template>
