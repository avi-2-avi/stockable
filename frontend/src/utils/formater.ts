export const formatNumberToCurrency = (value: number): string => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  }).format(value)
}

export const formatStringToDate = (value: string): string => {
  return new Date(value).toLocaleDateString()
}

export const formatStringToTime = (value: string): string => {
  return new Date(value).toLocaleTimeString()
}

export const formatActionString = (value: string): string => {
  const [action] = value.split('by')
  return action.trim().replace(/^\w/, (c) => c.toUpperCase())
}

export const formatNumberToPercentage = (value: number): string => {
  return `${value.toFixed(2)}%`
}
