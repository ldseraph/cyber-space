<template>
  <q-breadcrumbs>
    <q-breadcrumbs-el v-if="first && 'name' in first" :icon="first.icon" :label="t(first.name)" to="/" />
    <!-- <el-breadcrumb-item
        v-for="(item, index) in breadlist"
        :key="index"
        :to="{ path: item.path }"
        >{{ item.meta.title }}</el-breadcrumb-item
      > -->
  </q-breadcrumbs>
</template>
<script lang="ts" setup>
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';
import { useRoute } from 'vue-router'
import { useSidebarStore, isSidebarItemsMenu } from '@/stores/sidebar-store';
const { t } = useI18n();
const route = useRoute();
const sidebar = useSidebarStore();
const first = computed(() => sidebar.menu.find((item) => isSidebarItemsMenu(item)))
// const breadlist

</script>
