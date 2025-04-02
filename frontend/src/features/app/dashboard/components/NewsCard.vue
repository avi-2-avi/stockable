<template>
    <Card class="p-5 cursor-pointer hover:bg-stock-500/20"  @click="openLink">
        <img :src="props.image" alt="News image" class="w-full h-30 object-cover rounded-lg mb-4">
        <p class="text-xl font-semibold font-outfit mb-2">{{ props.title }}</p>
        <p class="text-foreground/70">{{ props.description }}</p>
    </Card>
</template>

<script lang="ts" setup>
import Card from '@/components/ui/Card.vue';

interface Props {
    title: string;
    description: string;
    image: string;
    link: string;
}

const props = defineProps<Props>();

const openLink = () => {
    window.open(props.link, '_blank');
};
</script>
