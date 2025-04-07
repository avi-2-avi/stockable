<template>
  <AppLayout pageTitle="User Management">
    <div class="flex justify-end items-center">
      <Button label="Add User" variant="solid" @click="openAddUserModal" />
    </div>
    <table class="min-w-full table-auto border border-border border-separate">
      <thead>
        <tr>
          <th class="px-4 py-2 text-left border border-border">Name</th>
          <th class="px-4 py-2 text-left border border-border">Email</th>
          <th class="px-4 py-2 text-left border border-border">Role</th>
          <th class="px-4 py-2 text-left border border-border">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-if="users.length === 0">
          <td colspan="4" class="text-center py-4 border border-border">No users found</td>
        </tr>
        <tr v-else v-for="user in users" :key="user.id">
          <td class="px-4 py-2 border border-border">{{ user.full_name }}</td>
          <td class="px-4 py-2 border border-border">{{ user.email }}</td>
          <td class="px-4 py-2 border border-border">{{ getRoleName(user.role_id || 0) }}</td>
          <td class="px-4 py-2 flex border border-border space-x-2">
            <Button @click="openUpdateUserModal(user)" label="Edit" variant="subtle" />
            <Button @click="openDeleteUserModal(user.id)" label="Delete" variant="warn" />
          </td>
        </tr>
      </tbody>
    </table>
  </AppLayout>
  <AddUserModal v-if="isAddUserModalOpen" @close="handleModalClose('add')" />
  <UpdateUserModal v-if="isUpdateUserModalOpen" @close="handleModalClose('update')" :user="updateUserData" />
  <DeleteUserModal v-if="isDeleteUserModalOpen" @close="handleModalClose('delete')" :userId="userToDelete" />

</template>

<script setup lang="ts">
import Button from '@/components/ui/Button.vue';
import AppLayout from '@/layouts/AppLayout.vue';
import { useAuthStore } from '@/store/authStore';
import { onMounted, ref } from 'vue';
import UpdateUserModal from '../components/UpdateUserModal.vue';
import DeleteUserModal from '../components/DeleteUserModal.vue';
import AddUserModal from '../components/AddUserModal.vue';
import { storeToRefs } from 'pinia';

const authStore = useAuthStore()
const { fetchUsers } = authStore
const { users } = storeToRefs(authStore) 

onMounted(() => {
  fetchUsers()
})

const isAddUserModalOpen = ref(false);
const openAddUserModal = () => {
  isAddUserModalOpen.value = true;
};

const isUpdateUserModalOpen = ref(false);
const updateUserData = ref({
  id: "",
  full_name: '',
  email: '',
  role_id: 0,
});

const openUpdateUserModal = (user: any) => {
  updateUserData.value = { ...user };
  isUpdateUserModalOpen.value = true;
};

const isDeleteUserModalOpen = ref(false);
const userToDelete = ref('');
const openDeleteUserModal = (userId: string) => {
  userToDelete.value = userId;
  isDeleteUserModalOpen.value = true;
};

const getRoleName = (roleId: number) => {
  return roleId === 1 ? 'Admin' : 'User';
};

const handleModalClose = (type: 'add' | 'update' | 'delete') => {
  if (type === 'add') isAddUserModalOpen.value = false;
  if (type === 'update') isUpdateUserModalOpen.value = false;
  if (type === 'delete') isDeleteUserModalOpen.value = false;

  fetchUsers();
};

</script>