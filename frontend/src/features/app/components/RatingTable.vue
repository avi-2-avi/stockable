<template>
    <div class="relative overflow-x-auto py-2.5">
        <table class="w-full bg-card border-border border-y">
            <thead class="bg-card">
                <tr>
                    <th v-for="column in columns" :key="column.id" class="cursor-pointer px-4 py-2 text-left"
                        @click="sortByColumn(column.id)">
                        <div class="flex items-center space-x-2">
                            <p>
                                {{ column.label }}
                            </p>
                            <span>
                                <ArrowUp v-if="sortBy === column.id && sortOrder === 'asc'"
                                    class="w-4 h-4 text-stock-500" />
                                <ArrowDown v-if="sortBy === column.id && sortOrder === 'desc'"
                                    class="w-4 h-4 text-stock-500" />
                            </span>
                        </div>
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="rating in ratings" :key="rating.id" class="odd:bg-stock-200/30">
                    <td v-for="column in columns" :key="column.id" class="border-border border-y px-4 py-2">
                        <template v-if="column.render">
                            <span v-if="column.render(rating[column.id as keyof Rating])?.text">
                                {{ column.render(rating[column.id as keyof Rating])?.text }}
                            </span>
                            <component v-else-if="column.render(rating[column.id as keyof Rating])?.component"
                                :is="column.render(rating[column.id as keyof Rating])?.component"
                                v-bind="column.render(rating[column.id as keyof Rating])?.props" />
                        </template>
                        <span v-else>
                            {{ rating[column.id as keyof Rating] }}
                        </span>
                    </td>
                </tr>
            </tbody>
        </table>
        <div class="mt-4 flex justify-between items-center">
            <Button label="Previous" size="md" variant="subtle" @click="prevPage" :disabled="page <= 1" />
            <div class="flex items-center space-x-4">
                <span>Page {{ page }} of {{ totalPages }}</span>

                <SelectDropdown id="limit-select" label="Results per page" v-model="selectedLimit"
                    :options="limitOptions" customClass="w-40" />
            </div>

            <Button label="Next" size="md" variant="subtle" @click="nextPage"
                :disabled="page * limit >= totalRatings" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ArrowUp, ArrowDown } from "lucide-vue-next";
import Button from '@/components/Button.vue';
import SelectDropdown from '@/components/SelectDropdown.vue';
import CPIBadge from '@/features/app/components/CPIBadge.vue';
import { useRatingStore } from '@/store/ratingStore';
import type { Rating } from '@/types/rating';
import { formatNumberToCurrency, formatStringToDate, formatActionString, formatNumberToPercentage } from '@/utils/formater';
import { storeToRefs } from 'pinia';
import { computed, h, onMounted } from 'vue';

const ratingStore = useRatingStore();
const { ratings, totalRatings, page, limit, sortBy, sortOrder } = storeToRefs(ratingStore);

const columns: {
    id: keyof Rating;
    label: string;
    render?: (value: Rating[keyof Rating]) => { text?: string; component?: any; props?: Record<string, any> };
}[] = [
        { id: "ticker", label: "Ticker" },
        { id: "company", label: "Company" },
        { id: "action", label: "Action", render: (value) => ({ text: formatActionString(value as string) }) },
        { id: "target_from", label: "Target From", render: (value) => ({ text: formatNumberToCurrency(value as number) }) },
        { id: "target_to", label: "Target To", render: (value) => ({ text: formatNumberToCurrency(value as number) }) },
        { id: "rating_increase_percentage", label: "Target Increase", render: (value) => ({ text: formatNumberToPercentage(value as number) }) },
        { id: "rating_to", label: "Current Rating" },
        {
            id: "combined_prediction_index",
            label: "CPI",
            render: (value) => ({
                component: CPIBadge,
                props: { value: value as number }
            })
        },
        { id: "rated_at", label: "Rated At", render: (value) => ({ text: formatStringToDate(value as string) }) }
    ];


const limitOptions = [
    { value: "10", label: "Show 10" },
    { value: "25", label: "Show 25" },
    { value: "50", label: "Show 50" },
    { value: "100", label: "Show 100" }
];

const selectedLimit = computed({
    get: () => limit.value.toString(),
    set: (newLimit: string) => {
        ratingStore.setLimit(Number(newLimit));
    }
})

onMounted(() => {
    ratingStore.fetchRatings();
})

const totalPages = computed(() => Math.ceil(totalRatings.value / limit.value));
const nextPage = () => ratingStore.setPage(page.value + 1);
const prevPage = () => ratingStore.setPage(page.value - 1);

const sortByColumn = (column: keyof Rating) => {
    const order = sortBy.value === column && sortOrder.value === "asc" ? "desc" : "asc";
    ratingStore.setSorting(column, order);
}

</script>
