import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from "../features/landing/pages/LandingPage.vue";
import LoginPage from "../features/auth/pages/LoginPage.vue";
import HomePage from "../features/app/home/pages/HomePage.vue";
import RecommendationPage from "../features/app/recommendation/pages/RecommendationPage.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Landing",
      component: LandingPage,
      meta: { requiresAuth: false },
    },
    {
      path: "/auth",
      name: "Auth",
      component: LoginPage,
      meta: { requiresAuth: false },
    },
    {
      path: "/app/home",
      name: "Home",
      component: HomePage,
      meta: { requiresAuth: true },
    },
    {
      path: "/app/recommendation",
      name: "Recommendation",
      component: RecommendationPage,
      meta: { requiresAuth: true },
    },
    {
      path: "/:pathMatch(.*)*", 
      redirect: "/",
    },
  ],
})

// router.beforeEach((to, _, next) => {
  // const authStore = useAuthStore();
  // if (to.meta.requiresAuth && !authStore.isAuthenticated) {
  //   next("/auth");
  // } else {
  //   next();
  // }
// });

export default router
