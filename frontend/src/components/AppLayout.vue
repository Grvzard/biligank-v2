<script setup lang="ts">
import { inject, ref, Ref, computed } from 'vue'
import AppFooter from './AppFooter.vue'
import axios from '@/utils/request'
import { onlineStatusE, StreamerInfo } from '@/types';

const emit = defineEmits<{
  (event: 'selectStreamer', streamer: StreamerInfo): void
}>()

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
