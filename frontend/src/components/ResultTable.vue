<script setup lang="ts">
import { inject, ref, Ref, watch } from 'vue'
import axios from '@/utils/request'
import { tsToMidnight, tsToString } from '@/utils/time'
import { normalizeLink } from '@/utils/link'
import { onlineStatusE, StreamLog, StreamerInfo } from '@/types';

const props = defineProps<{
  streamer: StreamerInfo
}>()

const dateMarkers = ref([
  {
    date: new Date("2023-10-14T14:00:00+08:00"),
    type: 'dot',
    tooltip: [{ text: '14-18点宕机', color: 'red' }],
  },
  {
    date: new Date("2024-04-25T15:00:00+08:00"),
    type: 'dot',
    tooltip: [{ text: '15点起宕机', color: 'red' }],
  },
  {
    date: new Date("2024-04-26T00:00:00+08:00"),
    type: 'dot',
    tooltip: [{ text: '宕机', color: 'red' }],
  },
  {
    date: new Date("2024-04-27T18:00:00+08:00"),
    type: 'dot',
    tooltip: [{ text: '宕机至18点恢复', color: 'red' }],
  },
  {
    date: new Date("2025-03-28T17:00:00+08:00"),
    type: 'dot',
    tooltip: [{ text: '17-23点宕机', color: 'red' }],
  }
])

const DatePickerMinDate = new Date(
  Math.max(
    props.streamer.first_log_ts * 1000,
    new Date("2023-01-01T00:00:00+08:00").getTime()
  )
)
const uid: number = props.streamer.uid
const now = Date.now() / 1000
const fromTstamp = tsToMidnight(now, 8) + 86400 - 60
const toTstamp = tsToMidnight(now, 8) - (86400 * 6)
const dateRange: Ref<[number, number]> = ref([toTstamp * 1000, fromTstamp * 1000])
const streamlogData: Ref<StreamLog[]> = ref([])
const isLoading = ref(false)
const onlineStatus = inject('onlineStatus') as Ref<onlineStatusE>

function fetchData() {
  // console.log(`fetch streamlog(uid:${uid}/ts:${dateRange.value[0]})`)
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

fetchData()

watch(dateRange, () => {
  fetchData()
})
</script>

<template>
  <div class="flex flex-row gap-x-4">
    <VueDatePicker v-model="dateRange" model-type="timestamp" locale="zh-CN" auto-apply range
      :hide-navigation="['time']" :clearable="false" input-class-name="focus:outline-none" menu-class-name="shadow-md"
      :max-date="new Date(fromTstamp * 1000)" :min-date="DatePickerMinDate" ignore-time-validation
      prevent-min-max-navigation :markers="dateMarkers" />
    <button class="border border-slate-300 px-4 rounded-md hover:border-cyan-500" @click="streamlogData.reverse()">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-down-up"
        viewBox="0 0 16 16">
        <path fill-rule="evenodd"
          d="M11.5 15a.5.5 0 0 0 .5-.5V2.707l3.146 3.147a.5.5 0 0 0 .708-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 1 0 .708.708L11 2.707V14.5a.5.5 0 0 0 .5.5m-7-14a.5.5 0 0 1 .5.5v11.793l3.146-3.147a.5.5 0 0 1 .708.708l-4 4a.5.5 0 0 1-.708 0l-4-4a.5.5 0 0 1 .708-.708L4 13.293V1.5a.5.5 0 0 1 .5-.5" />
      </svg>
    </button>
  </div>
  <div class="my-4" />
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
            <!-- 链接后的 hash tag 用于方便区分不同的封面链接 -->
            <a target="_blank" rel="noreferrer" :href="normalizeLink(row.cover)">链接#{{ row.cover.slice(-6, -4) }}</a>
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
    <!-- <div class="flex justify-center relative">
      <div class="absolute">
        <button v-if="!isLoading"
          class=" m-1 px-4 p-1 rounded-lg text-xs border border-slate-300 hover:border-cyan-500">...</button>
        <svg v-else class="animate-spin h-5 w-5 m-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
          <circle class="opacity-25 stroke-cyan-600 fill-none" cx="12" cy="12" r="10" stroke-width="4"></circle>
          <path class="opacity-75 stroke-cyan-600 fill-cyan-600"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z">
          </path>
        </svg>
      </div>
    </div> -->
  </div>
  <div class="p-3 text-sm">
    * 并非准确时间，会比实际开播晚&下播早几分钟。
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
