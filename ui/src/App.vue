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
import { computed } from 'vue';
import { useMeta } from 'quasar';
import { useLayoutStore } from '@/stores/layout-store';
import { useI18n } from 'vue-i18n'
import { useRoute } from 'vue-router';
import { useGlobalStore } from '@/stores/global-store';

const i18n = useI18n({ useScope: 'global' })
const route = useRoute();
const layout = useLayoutStore();

useMeta({
  title: i18n.t(layout.title),
})

const global = useGlobalStore();
i18n.locale = computed(() => {
  return global.lang
})

</script>
