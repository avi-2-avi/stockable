export interface Rating {
  id: string
  ticker: string
  target_from: number
  target_to: number
  company: string
  action: string
  brokerage: string
  rating_from: string
  rating_to: string
  rating_increase_percentage: number
  combined_prediction_index: number
  rated_at: string
}
