<script setup lang="ts">
import { inject, ref, Ref, watch } from 'vue'
import axios from '@/utils/request'
import { tsToMidnight, tsToString } from '@/utils/time'
import { onlineStatusE } from '@/types';

const props = defineProps<{
  uid: number
}>()

const dateMarkers = ref([
  {
    date: new Date("2023-10-14 14:00"),
    type: 'dot',
    tooltip: [{ text: '14-18点宕机', color: 'red' }],
  }
])

const uid: number = props.uid
const fromTstamp = tsToMidnight(Date.now() / 1000, 8)
const toTstamp = fromTstamp - (86400 * 5)
const dateRange: Ref<[number, number]> = ref([toTstamp * 1000, fromTstamp * 1000])
const streamlogData: Ref<Streamlog[]> = ref([])
const isLoading = ref(false)
const onlineStatus = inject('onlineStatus') as Ref<onlineStatusE>

function fetchData() {
  console.log(`fetch streamlog(uid:${uid}/ts:${dateRange.value[0]})`)
  isLoading.value = true

  axios({
    method: 'get',
    url: `/streamlog/${uid}`,
    params: {
      from: dateRange.value[1] / 1000,
      to: dateRange.value[0] / 1000,
    },
  })
    .then((resp) => {
      if (resp.status !== 200) {
        onlineStatus.value = onlineStatusE.Error
      }
      else {
        // console.log(`streamlog(${uid}) results: `, resp.data)
        onlineStatus.value = onlineStatusE.On
        // streamlogData.value = streamlogData.value.concat(resp.data)
        streamlogData.value = resp.data
      }
    })
    .catch((_err) => {
      onlineStatus.value = onlineStatusE.Off
    })
    .finally(() => {
      isLoading.value = false
    })
}
interface Streamlog {
  title: string
  c_ts: number
  last_update: number
  watched_num: number
  area_name: string
  cover: string
}

fetchData()

watch(dateRange, () => {
  fetchData()
})
</script>

<template>
  <div class="flex flex-row gap-x-4">
    <VueDatePicker v-model="dateRange" model-type="timestamp" :hide-navigation="['time']" auto-apply range
      :clearable="false" input-class-name="focus:outline-none" menu-class-name="shadow-md"
      :max-date="new Date(fromTstamp * 1000)" :min-date="new Date('2023-01-01 00:00')" ignore-time-validation
      prevent-min-max-navigation :markers="dateMarkers" />
    <button class="border border-slate-300 px-2 rounded-md hover:border-cyan-500" @click="streamlogData.reverse()">
      reverse
    </button>
  </div>
  <div class="m-4" />
  <div class="pt-3 pb-8 flex flex-col rounded Bg-Light shadow bg-grid-slate-800">
    <table class="table-auto border-collapse w-full shadow-sm">
      <thead>
        <tr>
          <th class="Data-th">
            封面
          </th>
          <th class="Data-th">
            标题
          </th>
          <th class="Data-th">
            分区
          </th>
          <th class="Data-th">
            时间
          </th>
          <th class="Data-th">
            人气
          </th>
        </tr>
      </thead>
      <tbody class="bg-white">
        <tr v-for="row in streamlogData" :key="row.last_update">
          <td class="Data-td text-cyan-600">
            <a target="_blank" :href="row.cover">链接</a>
          </td>
          <td class="Data-td">
            {{ row.title }}
          </td>
          <td class="Data-td">
            {{ row.area_name }}
          </td>
          <td class="Data-td">
            {{ tsToString(row.c_ts) }} — {{ tsToString(row.last_update) }}
          </td>
          <td class="Data-td">
            {{ row.watched_num }}
          </td>
        </tr>
      </tbody>
    </table>
    <!-- <div class="flex justify-center w-full">
      <button v-if="!isLoading"
        class=" m-1 px-4 p-1 rounded-lg text-xs border border-slate-300 hover:border-cyan-500"
        @click="onMore">...</button>

      <svg v-else class="animate-spin h-5 w-5 m-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
        <circle class="opacity-25 stroke-cyan-600 fill-none" cx="12" cy="12" r="10" stroke-width="4"></circle>
        <path class="opacity-75 stroke-cyan-600 fill-cyan-600"
          d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z">
        </path>
      </svg>
    </div> -->
  </div>
</template>

<style scoped>
.Data-th {
  @apply p-3 border-b text-left text-slate-800
}

.Data-td {
  @apply p-3
}
</style>
