<script setup lang="ts">
import { ref, Ref } from 'vue'
import AppLayout from '@/components/AppLayout.vue'
import ResultTable from '@/components/ResultTable.vue'

const currUid: Ref<number> = ref(0)

function onSelectUid(uid: number) {
  console.log(`select uid: ${uid}`)
  currUid.value = uid
}
</script>

<template>
  <AppLayout @select-uid="onSelectUid">
    <div v-if="currUid !== 0">
      <ResultTable :key="currUid" :uid="currUid" />
    </div>
    <div v-else>
      ok
    </div>
  </AppLayout>
</template>
