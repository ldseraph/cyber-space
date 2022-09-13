<template>
  <!-- bg-primary text-white -->
  <q-header bordered class="bg-white">
    <q-toolbar>
      <q-btn
        v-if="!sidebar.open"
        @click="sidebar.taggle()"
        flat
        round
        dense
        icon="menu"
        class="text-primary"
      />
      <q-space />
      <q-input
        class="pr-4"
        dense
        square
        v-model="search"
      >
        <template v-slot:prepend>
          <q-icon name="search" />
        </template>
        <template v-slot:append>
          <q-icon
            v-if="search !== ''"
            name="clear"
            class="cursor-pointer"
            @click="search = ''"
          />
        </template>
      </q-input>
    </q-toolbar>
  </q-header>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useSidebarStore } from '@/stores/sidebar-store';
const sidebar = useSidebarStore();
const search = ref('');
</script>
