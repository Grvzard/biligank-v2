<script setup lang="ts">
import { ref, Ref } from 'vue'
import AppLayout from '@/components/AppLayout.vue'
import ResultTable from '@/components/ResultTable.vue'
import { StreamerInfo } from '@/types';

const currStreamer: Ref<StreamerInfo | undefined> = ref()

function onSelect(streamer: StreamerInfo) {
  currStreamer.value = streamer
}
</script>

<template>
  <AppLayout @select-streamer="onSelect">
    <div v-if="currStreamer !== undefined">
      <ResultTable :key="currStreamer.uid" :streamer="currStreamer" />
    </div>
    <div v-else>
      ok
    </div>
  </AppLayout>
</template>
