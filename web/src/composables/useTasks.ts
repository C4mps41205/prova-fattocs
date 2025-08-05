import { ref, type Ref } from "vue"
import type {Task} from "../types/task.ts";
import {useApi} from "./useApi.ts";
import { useToast } from "./useToast.ts";

interface UseTasksReturn {
    tasks: Ref<Task[]>
    loading: Ref<boolean>
    error: Ref<string | null>
    loadTasks: () => Promise<void>
    createTask: (taskData: Omit<Task, "id" | "order_number">) => Promise<Task | null>
    updateTask: (id: number, taskData: Partial<Task>) => Promise<Task | null>
    deleteTask: (id: number) => Promise<void>
    reorderTasks: (newTasks: Task[]) => Promise<Task[] | null>
}

interface DefaultResponseApi<T> {
    data: T | null
    message: string
    status: number
}

export function useTasks(): UseTasksReturn {
    const baseUrl = `${import.meta.env.VITE_API_URI || 'http://localhost:8080'}/tasks`;

    const { addToast } = useToast()

    const tasks = ref<Task[]>([])
    const loading = ref(false)
    const error = ref<string | null>(null)

    const apiGet = useApi<DefaultResponseApi<Task[]>>(baseUrl, false)

    async function loadTasks(): Promise<void> {
        loading.value = true;
        error.value = null;
        try {
            const response = await apiGet.execute();
            tasks.value = response?.data ? [...response.data].sort((a, b) => a.order_number - b.order_number) : []
        } catch (err) {
            const errorMessage = err instanceof Error ? err.message : "Failed to load tasks";
            error.value = errorMessage;
            addToast({
                title: "Error",
                description: errorMessage,
                variant: "destructive"
            });
        } finally {
            loading.value = false;
        }
    }

    async function createTask(taskData: Omit<Task, "id" | "order_number">): Promise<Task | null> {
        loading.value = true
        error.value = null
        try {
            const apiPost = useApi<DefaultResponseApi<Task>>(baseUrl, false)
            const response = await apiPost.post(taskData)

            if (!response) {
                throw new Error(apiPost.error.value || "Failed to create task")
            }

            const newTask = response.data
            if (!newTask) throw new Error("Failed to create task")

            tasks.value.push(newTask)
            tasks.value.sort((a, b) => a.order_number - b.order_number)

            addToast({
                title: "Success",
                description: "Task created successfully",
                variant: "success"
            })
            return newTask
        } catch (err) {
            const errorMessage = err instanceof Error ? err.message : "Failed to create task"
            error.value = errorMessage
            addToast({
                title: "Error",
                description: errorMessage,
                variant: "destructive"
            })
            return null
        } finally {
            loading.value = false
        }
    }

    async function updateTask(id: number, taskData: Partial<Task>): Promise<Task | null> {
        loading.value = true
        error.value = null
        try {
            const apiPut = useApi<DefaultResponseApi<Task>>(`${baseUrl}/${id}`, false)
            const response = await apiPut.put(taskData)

            if (!response) {
                throw new Error(apiPut.error.value || "Failed to update task")
            }

            const updatedTask = response.data
            if (!updatedTask) throw new Error("Failed to update task")

            tasks.value = tasks.value
                .map(t => (t.id === id ? updatedTask : t))
                .sort((a, b) => a.order_number - b.order_number)

            addToast({
                title: "Success",
                description: "Task updated successfully",
                variant: "success"
            })
            return updatedTask
        } catch (err) {
            const errorMessage = err instanceof Error ? err.message : "Failed to update task"
            error.value = errorMessage
            addToast({
                title: "Error",
                description: errorMessage,
                variant: "destructive"
            })
            return null
        } finally {
            loading.value = false
        }
    }

    async function deleteTask(id: number): Promise<void> {
        loading.value = true
        error.value = null
        try {
            const apiDelete = useApi<DefaultResponseApi<boolean>>(`${baseUrl}/${id}`, false)
            const response = await apiDelete.delete()

            if (!response) {
                throw new Error(apiDelete.error.value || "Failed to delete task")
            }

            tasks.value = tasks.value.filter(t => t.id !== id)
            addToast({
                title: "Success",
                description: "Task deleted successfully",
                variant: "success"
            })
        } catch (err) {
            const errorMessage = err instanceof Error ? err.message : "Failed to delete task"
            error.value = errorMessage
            addToast({
                title: "Error",
                description: errorMessage,
                variant: "destructive"
            })
            throw err;
        } finally {
            loading.value = false
        }
    }

    async function reorderTasks(newTasks: Task[]): Promise<Task[] | null> {
        loading.value = true
        error.value = null
        try {
            const sortedTasks = [...newTasks].sort((a, b) => a.order_number - b.order_number);

            for (const task of sortedTasks) {
                const apiPost = useApi<DefaultResponseApi<Task>>(`${baseUrl}/${task.id}/reorder`, false);
                const response = await apiPost.post({ order: task.order_number });

                if (!response) {
                    throw new Error(apiPost.error.value || "Failed to reorder task");
                }
            }

            await loadTasks()

            addToast({
                title: "Success",
                description: "Tasks reordered successfully",
                variant: "success"
            })

            return tasks.value
        } catch (err) {
            const errorMessage = err instanceof Error ? err.message : "Failed to reorder tasks"
            error.value = errorMessage
            addToast({
                title: "Error",
                description: errorMessage,
                variant: "destructive"
            })
            return null
        } finally {
            loading.value = false
        }
    }

    return {
        tasks,
        loading,
        error,
        loadTasks,
        createTask,
        updateTask,
        deleteTask,
        reorderTasks,
    }
}
