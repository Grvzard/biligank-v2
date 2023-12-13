import { createApp } from 'vue'
import VueDatePicker from '@vuepic/vue-datepicker'
import IndexPage from './pages/IndexPage.vue'
import {ref, Ref} from 'vue'
import '@vuepic/vue-datepicker/dist/main.css'
import './style.css'
import { onlineStatusE } from './types'

const onlineStatus: Ref<onlineStatusE> = ref(onlineStatusE.Unknown)

const app = createApp(IndexPage)
app.provide('onlineStatus', onlineStatus)
app.component('VueDatePicker', VueDatePicker)
app.mount('#app')
