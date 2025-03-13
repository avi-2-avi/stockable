<template>
    <LandingLayout>
        <section id="hero">
            <div class="max-w-3xl mx-auto text-center flex flex-col items-center space-y-5 py-20">
                <p class="headline">Smarter Investing, All in One Place</p>
                <h3>Access real-time insights from multiple sources and expert analysts—so you make the best financial
                    decisions, every time.</h3>
                <Button label="Start Your Investing Journey" size="lg" @click="redirectToLogin" />
            </div>
            <img src="@/assets/hero.png" alt="Hero" class="w-200 mx-auto rounded-t-xl" />
        </section>
        <section id="features">
            <div class="py-40 space-y-40 xl:px-40">
                <FeatureSection v-for="(feature, index) in featureDataList" :key="index" :title="feature.title"
                    :description="feature.description" :image="feature.image" :highlight="feature.highlight"
                    :reverse="index % 2 !== 0" />
            </div>
        </section>
        <section id="about" class="pb-20">
            <Card card-class="flex flex-col items-center max-w-3xl mx-auto text-center p-10 space-y-8">
                <h3>
                    Our Mission: Smarter, Data-Driven Investing for Everyone
                </h3>
                <p class="text-lg">At <b class="font-outfit text-xl">Stockable</b>, we believe that <b>great financial decisions start with
                        great
                        information</b>. That’s why we bring together data from multiple stock sources, expert opinions,
                    and
                    AI-powered insights—all in a <b>single, intuitive platform</b>. Whether you're a beginner or a
                    seasoned
                    investor, we <b>simplify</b> your decision-making so you can invest with <b>confidence</b>.</p>
                <Button label="Join the Future of Investing" size="lg" @click="redirectToLogin" />
            </Card>
        </section>

    </LandingLayout>
</template>

<script setup lang="ts">
import LandingLayout from '@/layouts/LandingLayout.vue';
import FeatureSection from '@/features/landing/components/FeatureSection.vue';
import { useRouter } from 'vue-router';
import Button from '@/components/Button.vue';
import Card from '@/components/Card.vue';

const router = useRouter();

const redirectToLogin = () => {
    router.push('/auth');
}

const featureDataList = [
    {
        highlight: 'Unified',
        title: 'Market Intelligence',
        description: '📈 Stay ahead with aggregated data from leading financial platforms, all in one dashboard. No more switching between apps.',
        image: 'src/assets/feature-1.svg',
    },
    {
        highlight: 'Expert-Driven',
        title: 'Recommendations',
        description: '🧠 Get AI-enhanced insights from top analysts across multiple sources—personalized to your investment style.',
        image: 'src/assets/feature-2.svg',
    },
    {
        highlight: 'Data-Backed',
        title: 'Decisions, Simplified',
        description: '📊 Cut through the noise with trend analysis, risk assessment, and actionable takeaways—all visualized for easy understanding.',
        image: 'src/assets/feature-3.svg',
    }
];

</script>