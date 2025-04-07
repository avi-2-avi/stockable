<template>
  <Modal @close="$emit('close')">
    <template #header>
      <h3 class="text-lg font-bold">Add New User</h3>
    </template>

    <template #body>
      <form @submit.prevent="addUser" class="space-y-4">
        <BaseInput label="Full Name" type="text" placeholder="John Doe" v-model="newUser.full_name"
          :errorMessage="errors.full_name" />

        <BaseInput label="Email" type="email" placeholder="you@example.com" v-model="newUser.email"
          :errorMessage="errors.email" />

        <BaseInput label="Password" type="password" placeholder="••••••••" v-model="newUser.password"
          :errorMessage="errors.password" />

        <div>
          <label for="role" class="block text-sm font-medium">Role</label>
          <select v-model="newUser.role_id" id="role" class="mt-1 p-2 border rounded w-full border-border" required>
            <option value="1">Admin</option>
            <option value="2">User</option>
          </select>
        </div>
      </form>
    </template>

    <template #footer>
      <Button @click="$emit('close')" label="Cancel" variant="subtle" class="w-full"/>
      <Button @click="addUser" label="Add User" variant="solid" class="w-full" />
    </template>
  </Modal>
</template>

<script setup lang="ts">
import Modal from '@/components/ui/Modal.vue';
import Button from '@/components/ui/Button.vue';
import BaseInput from '@/components/ui/Input.vue';
import { ref } from 'vue';
import { useAuthStore } from '@/store/authStore';
import { validateAuthForm } from '@/utils/validation';

const emit = defineEmits(["close"]);
const errors = ref<Record<string, string>>({});

const { createUser } = useAuthStore();
const newUser = ref({
  full_name: '',
  email: '',
  password: '',
  role_id: 2
});

const addUser = async () => {
  try {
    errors.value = validateAuthForm(newUser.value, true);
    if (Object.keys(errors.value).length > 0) return;
    await createUser(newUser.value.full_name, newUser.value.email, newUser.value.password, getRoleName(Number(newUser.value.role_id)));
    emit('close');
  } catch (error) {
    console.error("Error adding user:", error);
  }
};

const getRoleName = (roleId: number) => {
  return roleId === 1 ? 'admin' : 'user';
};
</script>
