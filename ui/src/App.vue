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
import { useMeta } from 'quasar';
import { useLayoutStore } from '@/stores/layout-store';
import { useI18n } from 'vue-i18n'
import { useRoute } from 'vue-router';
const i18n = useI18n({ useScope: 'global' })
const route = useRoute();
const layout = useLayoutStore();
useMeta({
  title: i18n.t(layout.title),
})
</script>
