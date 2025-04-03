<template>
    <AppLayout :pageTitle="`Welcome back, ${name}`" pageDescription="Here's what's happening today">
        <div class="grid grid-cols-6 grid-rows-1 gap-5">
            <div class="col-span-6 lg:col-span-4 row-span-1">
                <AboutCard />
            </div>
            <div class="col-span-6 lg:col-span-2 lg:row-span-2">
                <Card class="p-5 h-full">
                    <RecentRatingTable :latestRatings="cachedDashboardRatings?.latest_ratings ?? []" />
                </Card>
            </div>
            <div class="col-span-6 lg:col-span-2 lg:row-start-2 lg:row-span-1">
                <DonutChartCard title="CPI Breakdown" :chartData="cachedDashboardRatings.donut_cpi_chart" />
            </div>
            <div class="col-span-6 lg:col-span-2 lg:row-start-2 lg:row-span-1 ">
                <DonutChartCard title="Ratings Breakdown" :chartData="cachedDashboardRatings.donut_rating_chart" />
            </div>
            <div class="col-span-6 lg:col-span-2" v-for="(news, index) in news" :key="index">
                <NewsCard :title="news.title" :description="news.description" :image="news.image" :link="news.link" />
            </div>
        </div>
    </AppLayout>
</template>

<script lang="ts" setup>
import AppLayout from '@/layouts/AppLayout.vue';
import Card from '@/components/ui/Card.vue';
import { useAuthStore } from '@/store/authStore';
import { useDashboardStore } from '@/store/dashboardStore';
import DonutChartCard from '../components/DonutChartCard.vue';
import NewsCard from '../components/NewsCard.vue';
import AboutCard from '../components/AboutCard.vue';
import RecentRatingTable from '../components/RecentRatingTable.vue';

import news1Image from '@/assets/news-1.png';
import news2Image from '@/assets/news-2.png';
import news3Image from '@/assets/news-3.png';

const authStore = useAuthStore();
const { cachedDashboardRatings  } = useDashboardStore();

const name = authStore.user?.full_name || 'glad to see you!';

const news = [
    {
        title: "Build Your First Stock Portfolio With These Easy Steps",
        description: "Learn the basics of creating a diversified stock portfolio, including risk management and smart investment strategies.",
        image: news1Image,
        link: "https://finmasters.com/first-stock-portfolio/"
    },
    {
        title: "Understanding Stock Ratings: What Do They Mean?",
        description: "Not sure how to interpret stock ratings? We break down buy, hold, and sell recommendations so you can make informed decisions.",
        image: news2Image,
        link: "https://www.investopedia.com/financial-edge/0512/understanding-analyst-ratings.aspx"
    },
    {
        title: "Top 5 Investment Mistakes Beginners Make",
        description: "Avoid common pitfalls when starting your investing journey. Learn how to navigate the market with confidence.",
        image: news3Image,
        link: "https://howmoneythinks.substack.com/p/top-5-investing-mistakes-beginners"
    }
];

</script>