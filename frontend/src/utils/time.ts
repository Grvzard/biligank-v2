export function tsToString(ts: number): string {
  const date = new Date(ts * 1000)
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')

  return `${month}-${day} ${hours}:${minutes}`
}

// ts in second
export function tsToMidnight(ts: number, utc_offset: number) {
  return Math.ceil(ts) - Math.ceil((ts + 3600 * utc_offset) % 86400)
}
