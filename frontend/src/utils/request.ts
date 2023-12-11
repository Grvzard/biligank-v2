import axios from 'axios'

console.log(`running with backend: ${import.meta.env.VITE_BACKEND_URL}`)
const service = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_URL,
  timeout: 6000,
})

export default service
