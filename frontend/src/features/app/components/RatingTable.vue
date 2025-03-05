<template>
    <div class="relative overflow-x-auto">
        <table class="w-full bg-card border-border border">
            <thead class="bg-card">
                <tr>
                    <th v-for="column in columns" :key="column.id" class="cursor-pointer px-4 py-2 text-left"
                        @click="sortByColumn(column.id)">
                        {{ column.label }}
                        <span>
                            {{ sortBy === column.id ? (sortOrder === 'asc' ? '▲' : '▼') : '' }}
                        </span>
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="rating in ratings" :key="rating.id">
                    <td v-for="column in columns" :key="column.id" class="border-border border-y px-4 py-2">
                        {{ formatValue(rating[column.id], column.type) }}
                    </td>
                </tr>
            </tbody>
        </table>
        <div class="mt-4 flex justify-between">
            <Button label="Previous" size="md" variant="subtle" @click="prevPage" :disabled="page <= 1" />
            <span>Page {{ page }} of {{ totalRatings }}</span>
            <Button label="Next" size="md" variant="subtle" @click="nextPage"
                :disabled="page * limit >= totalRatings" />
        </div>
    </div>
</template>

<script setup lang="ts">
import Button from '@/components/Button.vue';
import { useRatingStore } from '@/store/ratingStore';
import type { Rating } from '@/types/rating';
import { formatNumberToCurrency, formatStringToDate } from '@/utils/formater';
import { storeToRefs } from 'pinia';
import { onMounted } from 'vue';

const ratingStore = useRatingStore();
const { ratings, totalRatings, page, limit, sortBy, sortOrder } = storeToRefs(ratingStore);

const columns: { id: keyof Rating; label: string; type: "string" | "currency" | "date" }[] = [
    { id: "company", label: "Company", type: "string" },
    { id: "ticker", label: "Ticker", type: "string" },
    { id: "action", label: "Action", type: "string" },
    { id: "target_from", label: "Target From", type: "currency" },
    { id: "target_to", label: "Target To", type: "currency" },
    { id: "rating_from", label: "From", type: "string" },
    { id: "rating_to", label: "To", type: "string" },
    { id: "rated_at", label: "Rated At", type: "date" }
];


onMounted(() => {
    console.log("Fetching ratings");
    ratingStore.fetchRatings();
})

const nextPage = () => ratingStore.setPage(page.value + 1);
const prevPage = () => ratingStore.setPage(page.value - 1);

const sortByColumn = (column: keyof Rating) => {
    const order = sortBy.value === column && sortOrder.value === "asc" ? "desc" : "asc";
    ratingStore.setSorting(column, order);
}

const formatValue = (value: string | number, type: "string" | "currency" | "date"): string => {
    if (type === "currency" && typeof value === "number") {
        return formatNumberToCurrency(value);
    }
    if (type === "date" && typeof value === "string") {
        return formatStringToDate(value);
    }
    return value.toString();
};

</script>
