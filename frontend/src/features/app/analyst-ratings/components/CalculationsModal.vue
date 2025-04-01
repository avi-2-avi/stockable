<template>
    <div class="fixed inset-0 flex items-center justify-center bg-black/80 z-[100]">
        <div class="bg-base rounded-lg shadow-lg p-6 mx-4 container overflow-y-auto max-h-[90vh]">
            <header class="flex justify-between items-center mb-4">
                <h3>How is CPI Calculated?</h3>
            </header>
            <div class="text-sm space-y-4">
                <p>
                    The <strong>Combined Prediction Index (CPI)</strong> is calculated based on multiple factors, each
                    contributing to the final prediction score.
                </p>
                <h4>Action Impact Score (AIS)</h4>
                <p>
                    The AIS assesses the impact of an analyst action on stock price movement:
                </p>
                <p class="italic">Formula: AIS = Weight<sub>action</sub> × (Target<sub>to</sub> - Target<sub>from</sub>)
                </p>
                <ul class="list-disc ml-5">
                    <li>Target Raised: +1</li>
                    <li>Target Lowered: -1</li>
                    <li>Reiterated: 0.5</li>
                    <li>Initiated: 0.75</li>
                    <li>Upgraded: +1.5</li>
                    <li>Downgraded: -1.5</li>
                </ul>

                <h4 class="font-semibold">Rating Change Impact (RCI)</h4>
                <p>
                    Measures the effect of a rating upgrade or downgrade.
                </p>
                <p class="italic">Formula: RCI = Rating<sub>to</sub> - Rating<sub>from</sub></p>

                <p class="font-semibold text-lg">RCI Matrix Table:</p>
                <table class="w-full border mt-2 text-sm">
                    <thead>
                        <tr class="bg-gray-700 text-white">
                            <th class="border p-2"></th>
                            <th class="border p-2">Strong-Buy</th>
                            <th class="border p-2">Buy / Overweight</th>
                            <th class="border p-2">Neutral / Hold</th>
                            <th class="border p-2">Sell / Underperform</th>
                            <th class="border p-2">Sector Perform</th>
                            <th class="border p-2">Market Perform</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(row, index) in rciMatrix" :key="index">
                            <td class="border p-2 font-semibold">{{ row.label }}</td>
                            <td v-for="(value, idx) in row.values" :key="idx" class="border p-2">{{ value }}</td>
                        </tr>
                    </tbody>
                </table>

                <h4 class="font-semibold">Target Adjustment Percentage (TAP)</h4>
                <p>
                    Determines the percentage change in target price:
                </p>
                <p class="italic">TAP = min(max((Target<sub>to</sub> - Target<sub>from</sub>) / Target<sub>from</sub> ×
                    100, -100), 100)</p>
                <p>If TAP < 0, a penalty factor of 1.5 is applied.</p>

                        <h4 class="font-semibold">Combined Prediction Index (CPI)</h4>
                        <p>
                            The final CPI is calculated using weighted averages:
                        </p>
                        <p class="italic">CPI = 0.3 × AIS + 0.35 × RCI + 0.35 × TAP</p>
                        <p>
                            The raw CPI is then normalized to a 0-100 scale:
                        </p>
                        <p class="italic">Normalized CPI = ((CPI - CPI<sub>min</sub>) / (CPI<sub>max</sub> -
                            CPI<sub>min</sub>)) × 100</p>

                        <p class="mt-2">
                            The higher the CPI, the stronger the buy signal. A lower CPI suggests caution.
                        </p>
            </div>


            <footer class="mt-4 flex justify-end">
                <Button label="Close" size="md" variant="solid" @click="$emit('close')" />
            </footer>
        </div>
    </div>
</template>

<script setup lang="ts">
import Button from '@/components/ui/Button.vue';
defineEmits(["close"]);
const rciMatrix = [
    { label: "Strong-Buy", values: [4, 3, -2, -4, -1, -2] },
    { label: "Buy / Overweight", values: [3, 3, -2, -3, -1, -1] },
    { label: "Neutral / Hold", values: [2, 2, 0.5, -1, 1, 1] },
    { label: "Sell / Underperform", values: [4, 3, 1, -1, 1.5, 2] },
    { label: "Sector Perform", values: [2, 2, -1, -2, 2, 1.5] },
    { label: "Market Perform", values: [3, 2, -1, -2, 1.5, 2] },
];
</script>