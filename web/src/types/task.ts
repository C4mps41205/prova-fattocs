export interface Task {
    id: number
    name: string
    cost: number
    deadline: string
    order_number: number
}

export interface TaskFormData {
    name: string
    cost: string
    deadline: string
}

export interface FormErrors {
    name?: string
    cost?: string
    deadline?: string
}
