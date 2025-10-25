<template>
  <div
    class="min-h-screen bg-linear-to-br from-gray-50 to-gray-100 flex items-center justify-center p-4"
  >
    <div class="w-full max-w-md">
      <!-- Header -->
      <div class="text-center mb-8">
        <div
          class="inline-flex items-center justify-center w-16 h-16 bg-linear-to-br from-[#fd9621] via-[#90caf9] to-[#02a8f3] rounded-2xl mb-4 shadow-lg"
        >
          <CheckCircle2 :size="32" class="text-white" />
        </div>
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Welcome Back</h1>
        <p class="text-gray-600">Sign in to continue to MileTaskFlow</p>
      </div>

      <!-- Error Alert -->
      <Transition
        enter-active-class="transition-all duration-300 ease-out"
        leave-active-class="transition-all duration-200 ease-in"
        enter-from-class="opacity-0 -translate-y-5"
        leave-to-class="opacity-0 -translate-y-2"
      >
        <div
          v-if="authStore.error"
          class="mb-6 p-4 bg-red-50 border-l-4 border-red-500 rounded-r-lg shadow-sm"
        >
          <div class="flex items-start gap-3">
            <AlertCircle :size="20" class="text-red-500 shrink-0 mt-0.5" />
            <div class="flex-1">
              <p class="text-sm font-medium text-red-800">
                {{ authStore.error }}
              </p>
            </div>
            <button
              @click="authStore.clearError"
              class="shrink-0 text-red-400 hover:text-red-600 transition-colors duration-200"
            >
              <X :size="18" />
            </button>
          </div>
        </div>
      </Transition>

      <!-- Form Card -->
      <div
        class="bg-white rounded-2xl shadow-xl border border-gray-100 overflow-hidden"
      >
        <!-- Decorative top bar -->
        <div
          class="h-2 bg-linear-to-r from-[#fd9621] via-[#90caf9] to-[#02a8f3]"
        ></div>

        <div class="p-8">
          <form @submit.prevent="handleLogin" class="space-y-5">
            <!-- Email -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">
                Email Address
              </label>
              <div class="relative">
                <div
                  class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none"
                >
                  <Mail :size="18" class="text-gray-400" />
                </div>
                <input
                  v-model="formData.email"
                  type="email"
                  required
                  placeholder="your@email.com"
                  @blur="validateEmail"
                  @focus="formErrors.email = ''"
                  class="w-full pl-11 pr-4 py-3 border rounded-xl text-sm bg-gray-50 focus:bg-white focus:outline-none transition-all duration-200"
                  :class="
                    formErrors.email
                      ? 'border-red-300 focus:border-red-500 focus:ring-2 focus:ring-red-200'
                      : 'border-gray-200 focus:border-[#fd9621] focus:ring-2 focus:ring-orange-100'
                  "
                />
              </div>
              <Transition
                enter-active-class="transition-all duration-300 ease-out"
                leave-active-class="transition-all duration-200 ease-in"
                enter-from-class="opacity-0 -translate-y-2"
                leave-to-class="opacity-0 -translate-y-2"
              >
                <p
                  v-if="formErrors.email"
                  class="mt-2 text-xs text-red-600 flex items-center gap-1"
                >
                  <span
                    class="inline-block w-1 h-1 bg-red-600 rounded-full"
                  ></span>
                  {{ formErrors.email }}
                </p>
              </Transition>
            </div>

            <!-- Password -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-2">
                Password
              </label>
              <div class="relative">
                <div
                  class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none"
                >
                  <Lock :size="18" class="text-gray-400" />
                </div>
                <input
                  v-model="formData.password"
                  :type="showPassword ? 'text' : 'password'"
                  required
                  placeholder="Enter your password"
                  @blur="validatePassword"
                  @focus="formErrors.password = ''"
                  class="w-full pl-11 pr-12 py-3 border rounded-xl text-sm bg-gray-50 focus:bg-white focus:outline-none transition-all duration-200"
                  :class="
                    formErrors.password
                      ? 'border-red-300 focus:border-red-500 focus:ring-2 focus:ring-red-200'
                      : 'border-gray-200 focus:border-[#fd9621] focus:ring-2 focus:ring-orange-100'
                  "
                />
                <button
                  type="button"
                  @click="showPassword = !showPassword"
                  class="absolute inset-y-0 right-0 pr-4 flex items-center text-gray-400 hover:text-gray-600 transition-colors duration-200"
                >
                  <EyeOff v-if="showPassword" :size="18" />
                  <Eye v-else :size="18" />
                </button>
              </div>
              <Transition
                enter-active-class="transition-all duration-300 ease-out"
                leave-active-class="transition-all duration-200 ease-in"
                enter-from-class="opacity-0 -translate-y-2"
                leave-to-class="opacity-0 -translate-y-2"
              >
                <p
                  v-if="formErrors.password"
                  class="mt-2 text-xs text-red-600 flex items-center gap-1"
                >
                  <span
                    class="inline-block w-1 h-1 bg-red-600 rounded-full"
                  ></span>
                  {{ formErrors.password }}
                </p>
              </Transition>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="!isFormValid || authStore.loading"
              class="w-full bg-linear-to-r from-[#fd9621] via-[#90caf9] to-[#02a8f3] text-white py-3.5 px-4 rounded-xl font-semibold hover:shadow-lg hover:scale-[1.02] focus:outline-none focus:ring-2 focus:ring-[#fd9621] focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:scale-100 transition-all duration-200 flex items-center justify-center gap-2"
            >
              <Loader2
                v-if="authStore.loading"
                :size="20"
                class="animate-spin"
              />
              <span>{{ authStore.loading ? "Signing in..." : "Sign In" }}</span>
              <ArrowRight
                v-if="!authStore.loading"
                :size="18"
                class="group-hover:translate-x-1 transition-transform duration-200"
              />
            </button>
          </form>
        </div>
      </div>

      <!-- Footer -->
      <div class="mt-8 text-center text-xs text-gray-500">
        <p>© 2025 MileTaskFlow. All rights reserved.</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/authStore";
import {
  Eye,
  EyeOff,
  AlertCircle,
  X,
  CheckCircle2,
  Loader2,
  Github,
  Mail,
  Lock,
  ArrowRight,
} from "lucide-vue-next";

const router = useRouter();
const authStore = useAuthStore();

const showPassword = ref(false);

const formData = ref({
  email: "",
  password: "",
});

const formErrors = ref({
  email: "",
  password: "",
});

const validateEmail = () => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

  if (!formData.value.email) {
    formErrors.value.email = "Email is required";
    return false;
  } else if (!emailRegex.test(formData.value.email)) {
    formErrors.value.email = "Please enter a valid email address";
    return false;
  } else {
    formErrors.value.email = "";
    return true;
  }
};

const validatePassword = () => {
  if (!formData.value.password) {
    formErrors.value.password = "Password is required";
    return false;
  } else if (formData.value.password.length < 6) {
    formErrors.value.password = "Password must be at least 6 characters";
    return false;
  } else {
    formErrors.value.password = "";
    return true;
  }
};

const isFormValid = computed(() => {
  return (
    formData.value.email &&
    formData.value.password &&
    !formErrors.value.email &&
    !formErrors.value.password
  );
});

const handleLogin = async () => {
  authStore.clearError();

  const isEmailValid = validateEmail();
  const isPasswordValid = validatePassword();

  if (!isEmailValid || !isPasswordValid) return;

  const result = await authStore.login({
    email: formData.value.email,
    password: formData.value.password,
  });

  if (result.success) {
    router.push("/tasks");
  }
};
</script>
