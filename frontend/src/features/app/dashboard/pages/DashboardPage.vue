<template>
    <AppLayout :pageTitle="`Welcome back, ${name}`" pageDescription="Here's what's happening with your portfolio today">
        <div class="grid grid-cols-6 grid-rows-1 gap-5">
            <div class="col-span-4 row-span-1">
                <PortafolioCard />
            </div>
            <div class="col-span-2 row-span-2">
                <Card class="p-5 h-full">
                    <p class="text-xl">Recent Ratings</p>
                    <div class="space-y-2">
                        <div class="flex items-center space-x-2">
                            <ClockIcon :size="16" />
                            <p class="text-lg font-semibold">Title</p>
                        </div>
                        <div class="flex items-center space-x-2">
                            <CPIBadge :value="100" />
                            <p>action - rated at</p>
                        </div>
                    </div>
                </Card>
            </div>
            <div class="col-span-2 row-start-2 row-span-1">
                <Card class="p-5">
                    <p class="text-xl">Donut</p>
                    <apexchart width="380" type="donut" :options="options" :series="series" />
                </Card>
            </div>
            <div class="col-span-2 row-start-2 row-span-1">
                <Card class="p-5">
                    <p class="text-xl">Donut</p>
                    <apexchart width="380" type="donut" :options="options" :series="series" />
                </Card>
            </div>
        </div>
    </AppLayout>
</template>

<script lang="ts" setup>
import AppLayout from '@/layouts/AppLayout.vue';
import PortafolioCard from '../components/PortafolioCard.vue';
import Card from '@/components/ui/Card.vue';
import { ClockIcon } from 'lucide-vue-next';
import CPIBadge from '@/components/ui/CPIBadge.vue';
const name = 'John Doe';

const options = {
    labels: ['Series 1', 'Series 2', 'Series 3'],
    colors: ['#19875b', '#2563eb', '#ef4444'],
    responsive: [{
        breakpoint: 480,
        options: {
            chart: {
                width: 200
            }
        }
    }]
}

const series = [44, 55, 41]

</script>