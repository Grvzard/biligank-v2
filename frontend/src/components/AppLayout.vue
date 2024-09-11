<script setup lang="ts">
import { inject, ref, Ref, computed } from 'vue'
import AppFooter from './AppFooter.vue'
import axios from '@/utils/request'
import { onlineStatusE, StreamerInfo } from '@/types';

const emit = defineEmits<{
  (event: 'selectStreamer', streamer: StreamerInfo): void
}>()

// const donateLink = import.meta.env.VITE_DONATE_LINK
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
  emit('selectStreamer', streamer)
}
</script>

<template>
  <!-- AppHeader -->
  <header class="Bg-Light p-4 shadow-sm overflow-visible">
    <div class="mx-2 flex items-center gap-x-5">
      <!-- logo -->
      <div class="text-2xl select-none">
        <span class="text-cyan-600">直播</span>
        <span class="text-slate-800">日志</span>
      </div>
      <!-- search bar -->
      <input v-model="search_text" type="text"
        class="py-2 px-4 mx-auto bg-slate-100 border border-slate-300 rounded-md focus:outline-none focus:bg-white hover:bg-white transition-colors"
        placeholder="uid / 房间号" @input="onInput" @focus="is_searchbar_focused = true"
        @blur="is_searchbar_focused = false" />
      <!-- donate link -->
      <!-- <button class="w-9 h-9 rounded-full hover:bg-slate-300">
        <a target="_blank" rel="noreferrer" :href="donateLink">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="currentColor"
            class="bi bi-piggy-bank m-auto" viewBox="0 0 18 18">
            <path
              d="M5 6.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0m1.138-1.496A6.6 6.6 0 0 1 7.964 4.5c.666 0 1.303.097 1.893.273a.5.5 0 0 0 .286-.958A7.6 7.6 0 0 0 7.964 3.5c-.734 0-1.441.103-2.102.292a.5.5 0 1 0 .276.962" />
            <path fill-rule="evenodd"
              d="M7.964 1.527c-2.977 0-5.571 1.704-6.32 4.125h-.55A1 1 0 0 0 .11 6.824l.254 1.46a1.5 1.5 0 0 0 1.478 1.243h.263c.3.513.688.978 1.145 1.382l-.729 2.477a.5.5 0 0 0 .48.641h2a.5.5 0 0 0 .471-.332l.482-1.351c.635.173 1.31.267 2.011.267.707 0 1.388-.095 2.028-.272l.543 1.372a.5.5 0 0 0 .465.316h2a.5.5 0 0 0 .478-.645l-.761-2.506C13.81 9.895 14.5 8.559 14.5 7.069q0-.218-.02-.431c.261-.11.508-.266.705-.444.315.306.815.306.815-.417 0 .223-.5.223-.461-.026a1 1 0 0 0 .09-.255.7.7 0 0 0-.202-.645.58.58 0 0 0-.707-.098.74.74 0 0 0-.375.562c-.024.243.082.48.32.654a2 2 0 0 1-.259.153c-.534-2.664-3.284-4.595-6.442-4.595M2.516 6.26c.455-2.066 2.667-3.733 5.448-3.733 3.146 0 5.536 2.114 5.536 4.542 0 1.254-.624 2.41-1.67 3.248a.5.5 0 0 0-.165.535l.66 2.175h-.985l-.59-1.487a.5.5 0 0 0-.629-.288c-.661.23-1.39.359-2.157.359a6.6 6.6 0 0 1-2.157-.359.5.5 0 0 0-.635.304l-.525 1.471h-.979l.633-2.15a.5.5 0 0 0-.17-.534 4.65 4.65 0 0 1-1.284-1.541.5.5 0 0 0-.446-.275h-.56a.5.5 0 0 1-.492-.414l-.254-1.46h.933a.5.5 0 0 0 .488-.393m12.621-.857a.6.6 0 0 1-.098.21l-.044-.025c-.146-.09-.157-.175-.152-.223a.24.24 0 0 1 .117-.173c.049-.027.08-.021.113.012a.2.2 0 0 1 .064.199" />
          </svg>
        </a>
      </button> -->
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
          <button v-for="streamer in search_results" :key="streamer.uid"
            class="p-2 pl-3 rounded hover:bg-gray-200 text-left" @mousedown="onClickResult(streamer)">
            {{ streamer.uname }} (房间号: {{ streamer.roomid }}
            {{ streamer.short_roomid != 0 ? " / " + streamer.short_roomid : "" }})
          </button>
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
