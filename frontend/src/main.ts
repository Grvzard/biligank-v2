import { createApp } from 'vue'
import VueDatePicker from '@vuepic/vue-datepicker'
import IndexPage from './pages/IndexPage.vue'
import '@vuepic/vue-datepicker/dist/main.css'
import './style.css'

createApp(IndexPage)
  .component('VueDatePicker', VueDatePicker)
  .mount('#app')
