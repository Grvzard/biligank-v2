<script setup lang="ts">
import { computed, ref, Ref } from 'vue'
import { apiAds } from '@/utils/request';

interface AdInfo {
  readonly src: string
  readonly href: string
  readonly isNotice: boolean
}

const ads: Ref<AdInfo[]> = ref([])
const adIdx = computed(() => {
  return Math.floor(Math.random() * ads.value.length)
})
apiAds.get("/ads.json")
  .then(
    (resp) => {
      if (resp.status === 200) {
        ads.value = resp.data
      }
    }
  )
</script>

<template>
  <div v-if="ads.length > 0" class="flex flex-col">
    <div class="mx-auto">
      <a v-if="!ads[adIdx].isNotice" class="" target="_blank" rel="noreferrer" :href="ads[adIdx].href">
        <img class="shadow-md rounded-md hover:shadow-lg transition-shadow" :src="ads[adIdx].src"></img>
      </a>
      <img v-else class="shadow-md rounded-md" :src="ads[adIdx].src"></img>
    </div>
    <div class="flex mt-3">

      <svg v-if="!ads[adIdx].isNotice" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
        class="bi bi-badge-ad ml-auto" viewBox="0 0 16 16">
        <path
          d="m3.7 11 .47-1.542h2.004L6.644 11h1.261L5.901 5.001H4.513L2.5 11zm1.503-4.852.734 2.426H4.416l.734-2.426zm4.759.128c-1.059 0-1.753.765-1.753 2.043v.695c0 1.279.685 2.043 1.74 2.043.677 0 1.222-.33 1.367-.804h.057V11h1.138V4.685h-1.16v2.36h-.053c-.18-.475-.68-.77-1.336-.77zm.387.923c.58 0 1.002.44 1.002 1.138v.602c0 .76-.396 1.2-.984 1.2-.598 0-.972-.449-.972-1.248v-.453c0-.795.37-1.24.954-1.24z" />
        <path
          d="M14 3a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1zM2 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2z" />
      </svg>
      <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
        class="bi bi-bell ml-auto" viewBox="0 0 16 16">
        <path
          d="M8 16a2 2 0 0 0 2-2H6a2 2 0 0 0 2 2M8 1.918l-.797.161A4 4 0 0 0 4 6c0 .628-.134 2.197-.459 3.742-.16.767-.376 1.566-.663 2.258h10.244c-.287-.692-.502-1.49-.663-2.258C12.134 8.197 12 6.628 12 6a4 4 0 0 0-3.203-3.92zM14.22 12c.223.447.481.801.78 1H1c.299-.199.557-.553.78-1C2.68 10.2 3 6.88 3 6c0-2.42 1.72-4.44 4.005-4.901a1 1 0 1 1 1.99 0A5 5 0 0 1 13 6c0 .88.32 4.2 1.22 6" />
      </svg>
      <!-- <div class="mt-4 flex justify-center gap-x-2">
      <button class="rounded-lg w-10 h-6 Bg-Light hover:bg-slate-200" @click="swapAd"></button>
    </div> -->

    </div>
  </div>
</template>