<template>
  <q-drawer
    v-model="sidebar.open"
    :behavior="(() => (layout.belowBreakpoint ? 'mobile' : 'desktop'))()"
    elevated
    :width="width"
    :mini-width="72"
    :mini="!local.mouse && sidebar.mini"
    :mini-to-overlay="sidebar.mini"
    persistent
    no-swipe-open
    no-swipe-close
    no-swipe-backdrop
    @mouseover="local.mouse = true"
    @mouseleave="local.mouse = false"
    class="'flex flex-nowrap overflow-hidden flex-col drawer-transition'"
  >
    <div
      class="flex flex-nowrap w-full flex-none p-4 items-center text-primary"
      :class="{ shadow: local.shadow }"
    >
      <q-img
        class="w-10 mr-4 flex-none"
        :src="sidebar.logo"
        v-if="sidebar.logo"
      ></q-img>
      <div
        ref="sidebarTitle"
        class="flex-none text-2xl font-bold mr-2"
        v-if="sidebar.title"
      >
        {{ t(sidebar.title) }}
      </div>
      <div
        class="flex flex-grow flex-row-reverse"
        v-if="!layout.belowBreakpoint"
      >
        <q-btn
          flat
          round
          :icon="sidebar.mini ? 'panorama_fish_eye' : 'adjust'"
          @click="() => (sidebar.mini = !sidebar.mini)"
        />
      </div>
    </div>
    <q-scroll-area
      class="h-full flex-1 overflow-hidden"
      @scroll="(info) => (local.shadow = info.verticalPosition > 0)"
      :thumb-style="{
        right: '1px',
        borderRadius: '5px',
        background: primary,
        width: '6px',
        opacity: '0.75',
      }"
    >
      <rex-sidebar-group ref="sidebarGroup" :items="sidebar.menu" />
    </q-scroll-area>
  </q-drawer>
</template>

<script lang="ts" setup>
import { getCssVar } from 'quasar';
import { computed, ref } from 'vue';
import { useLayoutStore } from '@/stores/layout-store';
import { useSidebarStore } from '@/stores/sidebar-store';
import RexSidebarGroup from './RexSidebarGroup.vue';
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const primary: string = getCssVar('primary') || '#626262';
const sidebar = useSidebarStore();
const layout = useLayoutStore();
const local = ref({
  mouse: false,
  shadow: false,
});
const sidebarTitle = ref<HTMLElement | null>(null);
const sidebarGroup = ref<{ $el: HTMLElement } | null>(null);
const width = computed(() => {
  var titleWidth =
    150 + (sidebarTitle.value ? sidebarTitle.value.offsetWidth : 0);
  var groupWidth = sidebarGroup.value?.$el.offsetWidth || 0;

  return titleWidth > groupWidth ? titleWidth : groupWidth;
});
</script>

<style lang="scss" scoped>
.shadow {
  box-shadow: 0 0 7px 3px rgba(0, 0, 0, 0.1);
}

.drawer-transition {
  transition: width 0.3s !important;
}
</style>
