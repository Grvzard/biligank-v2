<script setup lang="ts">
import { ref, Ref } from 'vue'
import AppLayout from '@/components/AppLayout.vue'
import ResultTable from '@/components/ResultTable.vue'
import Ads from '@/components/Ads.vue'
import { StreamerInfo } from '@/types';
import LivePush from '@/components/LivePush.vue';

const currStreamer: Ref<StreamerInfo | undefined> = ref()
const doSubscribe: Ref<boolean> = ref(false)

function onSelect(streamer: StreamerInfo, subscribe: boolean) {
  currStreamer.value = streamer
  doSubscribe.value = subscribe
}
</script>

<template>
  <AppLayout @select-streamer="onSelect">
    <LivePush v-if="currStreamer !== undefined && doSubscribe" :streamer="currStreamer" />
    <ResultTable v-else-if="currStreamer !== undefined" :key="currStreamer.uid" :streamer="currStreamer" />
    <Ads v-else class="mx-auto max-w-xl" />
  </AppLayout>
</template>
