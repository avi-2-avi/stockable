<template>
    <p class="text-xl font-semibold font-outfit">Recent Ratings</p>
    <div class="space-y-2 overflow-y-auto max-h-[380px]" v-if="latestRatings?.length">
        <RecentRatingItem v-for="(rating, _) in latestRatings" :key="rating.id"
            :rating="rating" />
    </div>
    <p v-else>Loading ratings...</p>
</template>

<script lang="ts" setup>
import RecentRatingItem from './RecentRatingItem.vue';
import type { Rating } from '@/types/rating';

const { latestRatings } = defineProps<{ latestRatings: Rating[] }>();
</script>
