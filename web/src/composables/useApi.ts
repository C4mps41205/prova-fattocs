import {ref, type Ref, type UnwrapRef} from "vue";

type Data<T> = Ref<any, any> | Ref<Ref<any, any> & T, Ref<any, any> & T> | Ref<UnwrapRef<T> | null, T | UnwrapRef<T> | null>

export interface ApiState<T> {
    data: Data<T>
    loading: Ref<boolean>
    error: Ref<string | null>
}

export interface UseApiReturn<T> extends ApiState<T> {
    execute: () => Promise<T | null>
    post: (body: unknown) => Promise<T | null>
    put: (body: unknown) => Promise<T | null>
    delete: () => Promise<boolean>
}

export function useApi<T>(url: string, immediate = true): UseApiReturn<T> {
    const data = ref<T | null>(null)
    const loading = ref(false)
    const error = ref<string | null>(null)

    const extractErrorMessage = async (response: Response): Promise<string> => {
        try {
            const result = await response.json();
            return result.message || `HTTP error! status: ${response.status}`;
        } catch {
            return `HTTP error! status: ${response.status}`;
        }
    }

    async function execute(): Promise<T | null> {
        loading.value = true
        error.value = null
        try {
            const response = await fetch(url)
            if (!response.ok) {
                const errorMessage = await extractErrorMessage(response)
                throw new Error(errorMessage)
            }
            const json = await response.json()
            data.value = json
            return json
        } catch (err) {
            error.value = err instanceof Error ? err.message : "An error occurred"
            return null
        } finally {
            loading.value = false
        }
    }

    async function post(body: unknown): Promise<T | null> {
        loading.value = true
        error.value = null
        try {
            const response = await fetch(url, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body),
            })
            if (!response.ok) {
                const errorMessage = await extractErrorMessage(response)
                throw new Error(errorMessage)
            }
            const json = await response.json()
            data.value = json
            return json
        } catch (err) {
            error.value = err instanceof Error ? err.message : "An error occurred"
            return null
        } finally {
            loading.value = false
        }
    }

    async function put(body: unknown): Promise<T | null> {
        loading.value = true
        error.value = null
        try {
            const response = await fetch(url, {
                method: "PUT",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body),
            })
            if (!response.ok) {
                const errorMessage = await extractErrorMessage(response)
                throw new Error(errorMessage)
            }
            const json = await response.json()
            data.value = json
            return json
        } catch (err) {
            error.value = err instanceof Error ? err.message : "An error occurred"
            return null
        } finally {
            loading.value = false
        }
    }

    async function del(): Promise<boolean> {
        loading.value = true
        error.value = null
        try {
            const response = await fetch(url, { method: "DELETE" })
            if (!response.ok) {
                const errorMessage = await extractErrorMessage(response)
                throw new Error(errorMessage)
            }
            data.value = null
            return true
        } catch (err) {
            error.value = err instanceof Error ? err.message : "An error occurred"
            return false
        } finally {
            loading.value = false
        }
    }

    if (immediate) {
        execute()
    }

    return { data, loading, error, execute, post, put, delete: del }
}
