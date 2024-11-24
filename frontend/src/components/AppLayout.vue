<script setup lang="ts">
import { inject, ref, Ref, computed } from 'vue'
import AppFooter from './AppFooter.vue'
import axios from '@/utils/request'
import { onlineStatusE, StreamerInfo } from '@/types';

const emit = defineEmits<{
  (event: 'selectStreamer', streamer: StreamerInfo, doSubscribe: boolean): void
}>()

const isSubscribing = ref(false)
const donateLink = import.meta.env.VITE_DONATE_LINK
const onlineStatus = inject('onlineStatus') as Ref<onlineStatusE>
const colorOfStatus = computed(() => {
  switch (onlineStatus.value) {
    case onlineStatusE.On:
      return 'bg-green-500'
    case onlineStatusE.Off:
      return 'bg-rose-400'
    case onlineStatusE.Error:
      return 'bg-amber-400'
    case onlineStatusE.Unknown:
      return 'bg-gray-400'
  }
})
const isLoading = ref(false)
const is_searchbar_focused = ref(false)
const search_text = ref('')
const results_pattern = ref('')
const search_results: Ref<StreamerInfo[]> = ref([])
function doSearch(q: string) {
  isLoading.value = true
  axios({
    method: 'get',
    url: '/search',
    params: {
      q,
    },
  })
    .then((resp) => {
      if (resp.status !== 200) {
        onlineStatus.value = onlineStatusE.Error
      }
      else {
        onlineStatus.value = onlineStatusE.On
        // console.log(`search(${q}) results: `, resp.data)
        if (q === search_text.value) {
          search_results.value = resp.data
          results_pattern.value = q
          isLoading.value = false
        }
      }
    })
    .catch((_err) => {
      onlineStatus.value = onlineStatusE.Off
    })
}
function searchDebounce(q: string) {
  if (q === search_text.value)
    doSearch(q)
}
function onInput(_event: any) {
  setTimeout(searchDebounce, 300, search_text.value)
}
function onClickResult(streamer: StreamerInfo) {
  isSubscribing.value = false
  emit('selectStreamer', streamer, false)
}
function onClickSubscribe(streamer: StreamerInfo) {
  isSubscribing.value = true
  emit('selectStreamer', streamer, true)
}
</script>

<template>
  <!-- AppHeader -->
  <header class="Bg-Light p-4 shadow-sm overflow-visible">
    <div class="mx-2 flex items-center gap-x-5">
      <!-- logo -->
      <div class="text-2xl select-none">
        <div v-if="!isSubscribing">
          <span class="text-cyan-600">直播</span>
          <span class="text-slate-800">日志</span>
        </div>
        <div v-else>
          <span class="text-cyan-600">开播</span>
          <span class="text-slate-800">推送</span>
        </div>
      </div>
      <!-- search bar -->
      <input v-model="search_text" type="text"
        class="py-2 px-4 mx-auto bg-slate-100 border border-slate-300 rounded-md focus:outline-none focus:bg-white hover:bg-white transition-colors"
        placeholder="uid / 房间号" @input="onInput" @focus="is_searchbar_focused = true"
        @blur="is_searchbar_focused = false" />
      <!-- donate link -->
      <button class="w-9 h-9 inline-flex items-center justify-center rounded-full transition-colors hover:bg-slate-300">
        <a target="_blank" rel="noreferrer" :href="donateLink">
          <img class="w-7 h-7" src="@/assets/elec.png" />
        </a>
      </button>
      <!-- online status -->
      <span class="inline-flex relative h-3 w-3">
        <span class="absolute w-full h-full rounded-full animate-ping opacity-75" :class="colorOfStatus"></span>
        <span class="rounded-full w-full h-full" :class="colorOfStatus"></span>
      </span>
    </div>
  </header>
  <!-- AppHeader -->

  <!-- AppBody -->
  <div class="flex mx-10 mt-3 mb-20 relative">
    <div class="absolute z-50 w-full">
      <div v-show="is_searchbar_focused" class="mx-auto max-w-sm shadow-md rounded p-2 bg-white">
        <div v-if="isLoading" class="animate-pulse bg-slate-200 h-4 w-28 my-1 mx-2 rounded"></div>
        <div v-else class="flex flex-col">
          <span class="pl-3 py-1">
            "{{ results_pattern }}" 搜索结果
          </span>
          <div v-for="streamer in search_results" :key="streamer.uid" class="flex">
            <button class="px-3 rounded hover:bg-gray-200 items-center" @mousedown="onClickSubscribe(streamer)">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-bell"
                viewBox="0 0 16 16">
                <path
                  d="M8 16a2 2 0 0 0 2-2H6a2 2 0 0 0 2 2M8 1.918l-.797.161A4 4 0 0 0 4 6c0 .628-.134 2.197-.459 3.742-.16.767-.376 1.566-.663 2.258h10.244c-.287-.692-.502-1.49-.663-2.258C12.134 8.197 12 6.628 12 6a4 4 0 0 0-3.203-3.92zM14.22 12c.223.447.481.801.78 1H1c.299-.199.557-.553.78-1C2.68 10.2 3 6.88 3 6c0-2.42 1.72-4.44 4.005-4.901a1 1 0 1 1 1.99 0A5 5 0 0 1 13 6c0 .88.32 4.2 1.22 6" />
              </svg>
            </button>
            <button class="p-2 rounded hover:bg-gray-200 text-left grow" @mousedown="onClickResult(streamer)">
              {{ streamer.uname }} (房间号: {{ streamer.roomid }}
              {{ streamer.short_roomid != 0 ? " / " + streamer.short_roomid : "" }})
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="my-5 w-full">
      <div class="mx-auto max-w-5xl">
        <slot />
      </div>
    </div>
  </div>
  <!-- AppBody -->

  <AppFooter />
</template>
