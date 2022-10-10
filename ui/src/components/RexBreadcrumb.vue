<template>
  <q-breadcrumbs>
    <template v-for="(item, index) in matched" :key="index">
      <q-breadcrumbs-el v-if="item.components&&item.name" :to="{ path: item.path }"
        :label="t('url.'+item.name.toString())"
        :icon="String(item.meta.icon)=='undefined'?undefined:String(item.meta.icon)" />
      <q-breadcrumbs-el v-else-if="item.name" :label="t('url.'+item.name.toString())"
        :icon="String(item.meta.icon)=='undefined'?undefined:String(item.meta.icon)" />
    </template>
  </q-breadcrumbs>
</template>
<script lang="ts" setup>
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';
import { useRoute } from 'vue-router'
const { t } = useI18n();
const route = useRoute();
const matched = computed(() => route.matched)
console.log(matched.value)
</script>
