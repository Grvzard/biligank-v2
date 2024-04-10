import axios from 'axios'

console.log(`running with backend: ${import.meta.env.VITE_BACKEND_URL}`)
const apiMain = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_URL,
  timeout: 6000,
})
export const apiAds = axios.create({
  baseURL: import.meta.env.VITE_ADS_BACKEND,
  timeout: 6000,
})

export default apiMain
