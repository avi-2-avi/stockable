<template>
    <Modal @close="$emit('close')">
        <template #header>
            <h3 class="text-lg font-bold text-red-600">Confirm Deletion</h3>
        </template>

        <template #body>
            <p>Are you sure you want to delete this user? This action cannot be undone.</p>
        </template>

        <template #footer>
            <Button @click="$emit('close')" label="Cancel" variant="subtle" class="w-full" />
            <Button @click="handleDeleteUser" label="Delete User" variant="warn" class="w-full" />
        </template>
    </Modal>
</template>

<script setup lang="ts">
import Modal from '@/components/ui/Modal.vue';
import Button from '@/components/ui/Button.vue';
import { useAuthStore } from '@/store/authStore';

const emit = defineEmits(["close"]);

const { deleteUser } = useAuthStore();

const props = defineProps({
  userId: {
    type: String,
    required: true
  }
});

const handleDeleteUser = async () => {
  try {
    await deleteUser(props.userId);
    emit('close');
  } catch (error) {
    console.error("Error deleting user:", error);
  }
};
</script>
