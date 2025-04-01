<template>
    <AuthLayout>
        <div class="flex items-center justify-center min-h-[calc(100vh-4rem)]">
            <Card card-class="p-8 w-full max-w-md mx-auto text-center space-y-6">
                <h2 class="font-bold">{{ isSignup ? "Create an Account" : "Welcome Back" }}</h2>
                <p class="text-foreground/60">
                    {{ isSignup ? "Sign up to start investing smarter." : "Log in to access your account." }}
                </p>

                <form @submit.prevent="handleSubmit" class="space-y-4">
                    <BaseInput v-if="isSignup" label="Full Name" type="text" placeholder="John Doe"
                        v-model="form.fullName" :errorMessage="errors.fullName" />

                    <BaseInput label="Email" type="email" placeholder="you@example.com" v-model="form.email"
                        :errorMessage="errors.email" />

                    <BaseInput label="Password" type="password" placeholder="••••••••" v-model="form.password"
                        :errorMessage="errors.password" />

                        <Button 
                        :label="isLoading ? 'Processing...' : (isSignup ? 'Sign Up' : 'Login')" 
                        size="lg" 
                        variant="solid" 
                        class="w-full" 
                        :disabled="isLoading"
                    />

                    <p class="text-sm">
                        {{ isSignup ? "Already have an account?" : "Don't have an account?" }}
                        <a @click="toggleMode" class="text-stock-400 cursor-pointer underline">
                            {{ isSignup ? "Log in" : "Sign up" }}
                        </a>
                    </p>
                </form>
            </Card>
        </div>
        <AuthModal v-if="authError" :message="authError" @close="authError = ''" />
    </AuthLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import AuthLayout from '@/layouts/AuthLayout.vue';
import Card from '@/components/ui/Card.vue';
import Button from '@/components/ui/Button.vue';
import BaseInput from '@/components/ui/Input.vue';
import AuthModal from '../components/AuthModal.vue';
import { validateAuthForm } from '@/utils/validation.ts';
import { useAuthStore } from '@/store/authStore';

const authStore = useAuthStore()
const isSignup = ref(false);
const isLoading = ref(false); 
const form = ref({
    fullName: "",
    email: "",
    password: ""
});
const errors = ref<Record<string, string>>({});
const authError = ref<string | null>(null);

const toggleMode = () => {
    isSignup.value = !isSignup.value;
    errors.value = {}; // Reset errors when switching
    authError.value = null;
};

const handleSubmit = async () => {
    errors.value = validateAuthForm(form.value, isSignup.value);
    if (Object.keys(errors.value).length > 0) return;

    isLoading.value = true; 
    try {
        if (isSignup.value) {
            await authStore.register(form.value.fullName, form.value.email, form.value.password);
        } else {
            await authStore.login(form.value.email, form.value.password);
        }
    } catch (error: any) {
        console.error(error);
        authError.value = error?.message || "An unexpected error occurred. Please try again.";
    } finally {
        isLoading.value = false; 
    }
};
</script>
