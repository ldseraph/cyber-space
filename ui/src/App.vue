<template>
  <router-view :name="route.name!='NotFound'?layout.name:'default'" />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
export default defineComponent({
  name: 'App',
});
</script>

<script lang="ts" setup>
import { computed, watch } from 'vue';
import { useMeta, useQuasar } from 'quasar';
import { useLayoutStore } from '@/stores/layout-store';
import { useI18n } from 'vue-i18n'
import { useRoute } from 'vue-router';
import { useGlobalStore } from '@/stores/global-store';

const i18n = useI18n({ useScope: 'global' })
const route = useRoute();
const layout = useLayoutStore();
const $q = useQuasar();

useMeta({
  title: i18n.t(layout.title),
})

const global = useGlobalStore();
i18n.locale = computed(() => {
  return global.lang
})

function loadLang(lang: string) {
  import(
    /* @vite-ignore */
    '/public/lang/' + lang + '.mjs'
  ).then(lang => {
    $q.lang.set(lang.default)
  })
}

loadLang(global.lang)

watch(global, val => {
  loadLang(val.lang)
})


</script>
