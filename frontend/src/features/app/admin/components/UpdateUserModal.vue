<template>
    <Modal @close="$emit('close')">
        <template #header>
            <h3 class="text-lg font-bold text-yellow-600">Update User</h3>
        </template>

        <template #body>
            <form @submit.prevent="handleUpdateUser" class="space-y-4">
                <BaseInput label="Full Name" type="text" placeholder="John Doe" v-model="userUpdateData.full_name"
                    :errorMessage="errors.full_name" />

                <BaseInput label="Email" type="email" placeholder="you@example.com" v-model="userUpdateData.email"
                    :errorMessage="errors.email" />

                <div class="flex items-center space-x-2">
                    <input type="checkbox" id="changePassword" v-model="changePassword" />
                    <label for="changePassword" class="text-sm">Change password?</label>
                </div>
                <BaseInput v-if="changePassword" label="New Password" type="password" placeholder="••••••••"
                    v-model="userUpdateData.password" :errorMessage="errors.password" />
            </form>
        </template>

        <template #footer>
            <Button @click="$emit('close')" label="Cancel" variant="subtle" class="w-full" />
            <Button @click="handleUpdateUser" label="Update User" variant="solid" class="w-full" />
        </template>
    </Modal>
</template>

<script setup lang="ts">
import Modal from '@/components/ui/Modal.vue';
import Button from '@/components/ui/Button.vue';
import BaseInput from '@/components/ui/Input.vue';
import { ref } from 'vue';
import { useAuthStore } from '@/store/authStore';
import type { User } from '@/types/user';
import { validateUpdateUserForm } from '@/utils/validation';

const emit = defineEmits(["close"]);

const errors = ref<Record<string, string>>({});
const changePassword = ref(false);

const props = defineProps<{
    user: User;
}>();

const { updateUser } = useAuthStore();
const userUpdateData = ref({
    full_name: props.user.full_name,
    email: props.user.email,
    password: '',
});

const handleUpdateUser = async () => {
    errors.value = validateUpdateUserForm(userUpdateData.value, changePassword.value);
    if (Object.keys(errors.value).length > 0) return;

    try {
        await updateUser(props.user.id, userUpdateData.value);
        emit('close');
    } catch (error) {
        console.error("Error updating user:", error);
    }
};
</script>
