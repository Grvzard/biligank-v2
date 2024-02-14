export function normalizeLink(link: string): string {
  const nlink = new URL(link)
  nlink.protocol = "https:"
  return nlink.href
}
