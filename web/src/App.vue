<template>
  <div class="flex flex-col min-h-screen bg-gradient-to-br from-white via-gray-50 to-gray-100 text-gray-800">
    <header class="bg-white/90 backdrop-blur-md border-b border-gray-200 sticky top-0 z-20 shadow-sm">
      <div class="max-w-7xl mx-auto px-6 py-4 flex justify-between items-center">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-gradient-to-r from-blue-600 to-indigo-600 rounded-xl shadow-md">
            <CheckSquare class="w-5 h-5 text-white" />
          </div>
          <div>
            <h1 class="text-2xl font-semibold text-gray-900">TaskFlow</h1>
            <p class="text-xs sm:text-sm text-gray-500">Task Management System</p>
          </div>
        </div>
      </div>
    </header>

    <main class="flex-grow">
      <div class="max-w-7xl mx-auto px-6 py-10">
        <div class="bg-white rounded-2xl shadow-md p-6 sm:p-8">
          <TaskList />
        </div>
      </div>
    </main>

    <footer class="bg-white/80 border-t border-gray-200">
      <div class="max-w-7xl mx-auto px-6 py-6 text-center text-sm text-gray-500">
        © {{ new Date().getFullYear() }} TaskFlow — Task Management System
      </div>
    </footer>
  </div>
  <ToastContainer />
</template>


<script setup lang="ts">
import TaskList from './components/task/TaskList.vue'
import { CheckSquare } from 'lucide-vue-next'
import ToastContainer from "./components/ui/ToastContainer.vue";
</script>
