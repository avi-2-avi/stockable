<template>
    <div class="relative overflow-x-auto py-2.5">
        <div class="hidden sm:grid grid-cols-4 xl:grid-cols-6 gap-4 mb-2.5 items-center">
            <Input v-model="filters.ticker" label="Ticker" placeholder="Filter by Ticker" @input="applyFilters" />
            <Input v-model="filters.company" label="Company" placeholder="Filter by Company" @input="applyFilters" />
            <SelectDropdown v-model="filters.action" :options="actionOptions" @update:modelValue="applyFilters"
                label="Actions" position="top" />
            <SelectDropdown v-model="filters.rating_to" :options="ratingOptions" @update:modelValue="applyFilters"
                label="Ratings" position="top" />
            <Input v-model="filters.target_from" type="number" label="Min Target" placeholder="Min"
                @input="applyFilters" />
            <Input v-model="filters.target_to" type="number" label="Max Target" placeholder="Max"
                @input="applyFilters" />
        </div>


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
                <tr v-for="rating in filteredRatings" :key="rating.id" class="odd:bg-stock-200/30">
                    <td v-for="column in columns" :key="column.id" class="border-border border-y px-4 py-2">
                        <template v-if="column.render">
                            <span v-if="column.render(rating[column.id as keyof Rating], rating)?.text">
                                {{ column.render(rating[column.id as keyof Rating], rating)?.text }}
                            </span>
                            <component v-else-if="column.render(rating[column.id as keyof Rating], rating)?.component"
                                :is="column.render(rating[column.id as keyof Rating], rating)?.component"
                                v-bind="column.render(rating[column.id as keyof Rating], rating)?.props" />
                        </template>
                        <span v-else>
                            {{ rating[column.id as keyof Rating] }}
                        </span>
                    </td>
                </tr>
            </tbody>
        </table>
        <div class="mt-4 flex justify-between items-center">
            <Button label="Previous" size="md" @click="prevPage" :disabled="page <= 1" />
            <div class="flex items-center space-x-4">
                <span>Page {{ page }} of {{ totalPages }}</span>
                <SelectDropdown id="limit-select" v-model="selectedLimit" :options="limitOptions" customClass="w-40" />
            </div>

            <Button label="Next" size="md" @click="nextPage" :disabled="page * limit >= totalRatings" />
        </div>
    </div>
    <CompanyModal v-if="showModal" :ticker="selectedCompany?.ticker ?? ''" :company="selectedCompany?.company ?? ''"
        @close="showModal = false" />
</template>

<script setup lang="ts">
import { ArrowUp, ArrowDown } from "lucide-vue-next";
import Button from '@/components/ui/Button.vue';
import SelectDropdown from '@/components/ui/SelectDropdown.vue';
import Input from "@/components/ui/Input.vue";
import CPIBadge from '@/components/ui/CPIBadge.vue';
import { useRatingStore } from '@/store/ratingStore';
import type { Rating } from '@/types/rating';
import { formatNumberToCurrency, formatStringToDate, formatActionString, formatNumberToPercentage } from '@/utils/formater';
import { storeToRefs } from 'pinia';
import { computed, onMounted, ref } from 'vue';
import { actionOptions, ratingOptions } from "@/data/rating_options";
import { limitOptions } from "@/data/table_options";
import CompanyModal from "./CompanyModal.vue";

const ratingStore = useRatingStore();
const { ratings, totalRatings, page, limit, sortBy, sortOrder } = storeToRefs(ratingStore);

const showModal = ref(false);
const selectedCompany = ref<{ ticker: string; company: string } | null>(null);

const columns: {
    id: keyof Rating;
    label: string;
    render?: (value: Rating[keyof Rating], row: Rating) => { text?: string; component?: any; props?: Record<string, any> };
}[] = [
        {
            id: "ticker", label: "Ticker",
            render: (value, row) => ({
                component: Button,
                props: {
                    label: value,
                    size: "md",
                    variant: "ghost",
                    onClick: () => {
                        selectedCompany.value = {
                            ticker: row.ticker,
                            company: row.company
                        };
                        showModal.value = true;
                    }
                }
            })
        },
        { id: "company", label: "Company"},
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


const selectedLimit = computed({
    get: () => limit.value.toString(),
    set: (newLimit: string) => {
        ratingStore.setLimit(Number(newLimit));
    }
})

const filters = ref<Record<string, string>>({
    ticker: "",
    company: "",
    action: "",
    target_from: "",
    target_to: "",
    rating_to: "",
});

const applyFilters = () => {
    ratingStore.setFilter(filters.value);
};

const filteredRatings = computed(() => {
    return ratings.value.filter(rating => {
        return (
            (!filters.value.ticker || rating.ticker.toLowerCase().includes(filters.value.ticker.toLowerCase())) &&
            (!filters.value.company || rating.company.toLowerCase().includes(filters.value.company.toLowerCase())) &&
            (!filters.value.action || rating.action === filters.value.action) &&
            (!filters.value.rating_to || rating.rating_to === filters.value.rating_to) &&
            (!filters.value.target_from || rating.target_from >= parseFloat(filters.value.target_from)) &&
            (!filters.value.target_to || rating.target_to <= parseFloat(filters.value.target_to))
        );
    });
});


const totalPages = computed(() => Math.ceil(totalRatings.value / limit.value));
const nextPage = () => ratingStore.setPage(page.value + 1);
const prevPage = () => ratingStore.setPage(page.value - 1);

const sortByColumn = (column: keyof Rating) => {
    const order = sortBy.value === column && sortOrder.value === "asc" ? "desc" : "asc";
    ratingStore.setSorting(column, order);
}

onMounted(() => {
    ratingStore.fetchRatings();
})

</script>
