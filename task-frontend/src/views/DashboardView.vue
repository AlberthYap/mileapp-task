<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Navbar -->
    <nav class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center gap-2">
            <div class="w-8 h-8 bg-[#fd9621] rounded-lg flex items-center justify-center">
              <CheckCircle2 :size="18" class="text-white" />
            </div>
            <span class="font-bold text-lg">TaskFlow</span>
          </div>
          
          <div class="flex items-center gap-4">
            <span class="text-sm text-gray-600">{{ authStore.currentUser?.email }}</span>
            <button
              @click="handleLogout"
              class="px-4 py-2 text-sm text-gray-700 border border-gray-300 rounded-lg hover:bg-gray-50"
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    </nav>

    <!-- Content -->
    <main class="max-w-7xl mx-auto p-4 sm:p-6 lg:p-8">
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-xl font-bold text-gray-900 mb-2">Welcome!</h2>
        <p class="text-gray-600">You are successfully logged in.</p>
      </div>
    </main>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStore'
import { CheckCircle2 } from 'lucide-vue-next'

const router = useRouter()
const authStore = useAuthStore()

const handleLogout = async () => {
  await authStore.logout()
  router.push('/login')
}
</script>
