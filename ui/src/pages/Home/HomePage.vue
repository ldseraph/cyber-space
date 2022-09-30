<template>
  <q-page class="flex flex-col" padding>
    <q-card v-for="n in 110" :key="n" class="flex-1">
      Row {{ n }} / 150 {{ hello }}
    </q-card>
  </q-page>
</template>

<script lang="ts" setup>
import { Ping, PingResData } from '@/api/ping';
import { ref } from 'vue';

const pingRes: PingResData = await Ping();
const hello = ref(pingRes.ping);
</script>
