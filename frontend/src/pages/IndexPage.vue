<script setup lang="ts">
import { ref, Ref } from 'vue'
import AppLayout from '@/components/AppLayout.vue'
import ResultTable from '@/components/ResultTable.vue'
import Ads from '@/components/Ads.vue'
import { StreamerInfo } from '@/types';

const currStreamer: Ref<StreamerInfo | undefined> = ref()

function onSelect(streamer: StreamerInfo) {
  currStreamer.value = streamer
}
</script>

<template>
  <AppLayout @select-streamer="onSelect">
    <ResultTable v-if="currStreamer !== undefined" :key="currStreamer.uid" :streamer="currStreamer" />
    <Ads v-else class="mx-auto max-w-xl" />
  </AppLayout>
</template>
