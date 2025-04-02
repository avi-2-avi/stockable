<template>
    <AppLayout :pageTitle="`Welcome back, ${name}`" pageDescription="Here's what's happening today">
        <div class="grid grid-cols-6 grid-rows-1 gap-5">
            <div class="col-span-6 lg:col-span-4 row-span-1">
                <PortafolioCard />
            </div>
            <div class="col-span-6 lg:col-span-2 lg:row-span-2">
                <Card class="p-5 h-full">
                    <p class="text-xl font-semibold font-outfit">Recent Ratings</p>
                    <div class="space-y-2 overflow-y-auto max-h-[380px]" v-if="cachedDashboardRatings?.latest_ratings?.length">
                        <RecentRatingItem v-for="(rating, _) in cachedDashboardRatings.latest_ratings ?? []" :key="rating.id"
                            :rating="rating"
                            />
                    </div>
                    <p v-else>Loading ratings...</p>
                </Card>
            </div>
            <div class="col-span-6 lg:col-span-2 lg:row-start-2 lg:row-span-1">
                <DonutChartCard title="CPI Breakdown" :chartData="cachedDashboardRatings.donut_cpi_chart" />
            </div>
            <div class="col-span-6 lg:col-span-2 lg:row-start-2 lg:row-span-1 ">
                <DonutChartCard title="Ratings Breakdown" :chartData="cachedDashboardRatings.donut_rating_chart" />
            </div>
            <div class="col-span-6 lg:col-span-2 lg:row-start-3 row-span-1">
                <NewsCard :title="news[0].title" :description="news[0].description" :image="news[0].image" :link="news[0].link" />
            </div>
            <div class="col-span-6 lg:col-span-2 lg:row-start-3 row-span-1">
                <NewsCard :title="news[1].title" :description="news[1].description" :image="news[1].image" :link="news[1].link" />
            </div>
            <div class="col-span-6 lg:col-span-2 lg:row-start-3 row-span-1">
                <NewsCard :title="news[2].title" :description="news[2].description" :image="news[2].image" :link="news[2].link" />
            </div>
        </div>
    </AppLayout>
</template>

<script lang="ts" setup>
import AppLayout from '@/layouts/AppLayout.vue';
import PortafolioCard from '../components/PortafolioCard.vue';
import Card from '@/components/ui/Card.vue';
import RecentRatingItem from '../components/RecentRatingItem.vue';
import { useAuthStore } from '@/store/authStore';
import { useDashboardStore } from '@/store/dashboardStore';
import DonutChartCard from '../components/DonutChartCard.vue';
import NewsCard from '../components/NewsCard.vue';

const authStore = useAuthStore();
const { cachedDashboardRatings  } = useDashboardStore();

const name = authStore.user?.full_name || 'glad to see you!';

const news = [
    {
        title: "Build Your First Stock Portfolio With These Easy Steps",
        description: "Learn the basics of creating a diversified stock portfolio, including risk management and smart investment strategies.",
        image: "/src/assets/news-1.webp",
        link: "/blog/build-your-first-portfolio"
    },
    {
        title: "Understanding Stock Ratings: What Do They Mean?",
        description: "Not sure how to interpret stock ratings? We break down buy, hold, and sell recommendations so you can make informed decisions.",
        image: "/src/assets/news-2.webp",
        link: "/blog/understanding-stock-ratings"
    },
    {
        title: "Top 5 Investment Mistakes Beginners Make",
        description: "Avoid common pitfalls when starting your investing journey. Learn how to navigate the market with confidence.",
        image: "/src/assets/news-3.webp",
        link: "/blog/top-5-investment-mistakes"
    }
];

</script>