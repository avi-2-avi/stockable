<template>
    <Card class="p-5 h-[320px]">
        <p class="text-xl font-semibold font-outfit">{{ title }}</p>
        <div class="flex w-full overflow-x-auto items-center justify-center mt-2">
            <apexchart width="360" type="donut" :options="chartOptions" :series="seriesData" />
        </div>
    </Card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import Card from '@/components/ui/Card.vue';

const props = defineProps<{
    title: string;
    chartData: { label: string; count: number }[];
}>();

const processedData = computed(() => {
    if (props.chartData.length <= 4) return props.chartData;

    const topFour = props.chartData.slice(0, 4);
    const othersCount = props.chartData.slice(4).reduce((sum, data) => sum + data.count, 0);

    return [
        ...topFour,
        { label: "Others", count: othersCount }
    ];
});

const seriesData = computed(() => processedData.value.map(data => data.count));

const chartOptions = computed(() => ({
    labels: processedData.value.map(data => data.label),
    colors: ['#FE8238', '#9E2811','#EF4607','#ffaf71', '#430E09'],
    legend: {
        labels: {
            colors: '#808080'
        },
        markers: {
            size: 4,
            offsetX: -5,
        },
    },
    responsive: [
        {
            breakpoint: 480,
            options: {
                chart: {
                    width: 400,
                }
            }
        }
    ]
}));
</script>
