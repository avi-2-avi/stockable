export const formatNumberToCurrency = (value: number): string => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  }).format(value)
}

export const formatStringToDate = (value: string): string => {
  return new Date(value).toLocaleDateString()
}
