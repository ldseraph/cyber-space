<template>
  <q-header class="bg-transparent text-gray-600">
    <q-toolbar>
      <q-btn
        v-if="!sidebar.open"
        @click="sidebar.taggle()"
        flat
        round
        dense
        icon="menu"
      />
    </q-toolbar>
  </q-header>
</template>

<script lang="ts" setup>
import { useSidebarStore } from '@/stores/sidebar-store';
const sidebar = useSidebarStore();
</script>
