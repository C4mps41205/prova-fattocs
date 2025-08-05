import { ref, type Ref } from 'vue'

export interface Toast {
    id: number
    title?: string
    description?: string
    variant?: 'default' | 'destructive' | 'success' | 'warning'
    duration?: number
}

interface ToastContext {
    toasts: Ref<Toast[]>
    addToast: (toast: Omit<Toast, 'id'>) => void
    removeToast: (id: number) => void
    clearToasts: () => void
}

const toasts = ref<Toast[]>([])

let toastId = 0

export function useToast(): ToastContext {
    function addToast(toast: Omit<Toast, 'id'>) {
        toastId++
        toasts.value.push({
            ...toast,
            id: toastId,
            duration: toast.duration || 5000
        })
    }

    function removeToast(id: number) {
        const index = toasts.value.findIndex(toast => toast.id === id)
        if (index !== -1) {
            toasts.value.splice(index, 1)
        }
    }

    function clearToasts() {
        toasts.value = []
    }

    return {
        toasts: toasts as Ref<Toast[]>,
        addToast,
        removeToast,
        clearToasts
    }
}