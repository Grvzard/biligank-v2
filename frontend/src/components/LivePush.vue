<script setup lang="ts">
import { StreamerInfo } from '@/types';
import { apiLivepush } from '@/utils/request'
import { ref } from 'vue';

const props = defineProps<{
  streamer: StreamerInfo
}>()

const emailAddress = ref('')
const resultText = ref('')
const isLoading = ref(false)

function onClickSubscribe() {
  let uid = props.streamer.uid
  let address = emailAddress.value
  isLoading.value = true
  setTimeout(sendSubscribe, 300, address, uid)
}

function sendSubscribe(address: string, uid: number) {
  apiLivepush({
    method: 'post',
    url: '/v1/subscribe',
    data: {
      email: address,
      uid: uid,
    }
  })
    .then((resp) => {
      if (resp.status !== 200) {
        resultText.value = "❌ 服务器出错"
      } else {
        let err_msg = resp.data.err_msg
        if (err_msg != null) {
          resultText.value = "❌ " + err_msg
        } else {
          resultText.value = "✅ 你将会收到一封确认邮件！已经订阅或超过数量上限(3)将不会收到。"
          emailAddress.value = ""
        }
      }
    })
    .catch((_err) => {
      resultText.value = "❌ 服务器出错"
    })
    .finally(() => {
      isLoading.value = false
    })
}
</script>

<template>
  <div class="flex justify-center items-center">
    <div class="flex flex-col w-fit">
      <span class="p-3">在 {{ props.streamer.uname }}(UID: {{ props.streamer.uid }}) 开播后收到推送</span>
      <div class="flex flex-row">
        <input v-model="emailAddress" type="text"
          class="py-2 px-4 bg-white border border-slate-300 rounded-md focus:outline-none focus:bg-white"
          placeholder="Email 地址" />
        <button
          class="ml-2 py-2 px-4 border border-slate-300 rounded-md hover:border-cyan-500 transition-colors text-cyan-600"
          @mousedown="onClickSubscribe()" :disabled="isLoading">
          订阅
        </button>
      </div>
      <span v-if="isLoading" class="p-3">...</span>
      <span v-else class="p-3">{{ resultText }}</span>
    </div>
  </div>
</template>