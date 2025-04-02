import type { Rating } from "./rating";

interface DonutCpiEntry {
    label: string
    count: number
}

export type DashboardData = {
    latest_ratings: Rating[],
    donut_cpi_chart: DonutCpiEntry[],
    donut_rating_chart: DonutCpiEntry[]
}