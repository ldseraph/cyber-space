<template>
  <q-scroll-area class="h-screen" :thumb-style="{
    right: '0px',
    borderRadius: '5px',
    background: primary,
    width: '4px',
    opacity: '0.75',
    zIndex: '4000'
  }">
    <q-layout ref="qLayoutRef" view="lHh Lpr lFf" class="left-transition"
      @resize="(size) => (layout.width = size.width,layout.height = size.height)">
      <RexSidebar />

      <RexNavbar />

      <q-page-container>
        <router-view v-slot="{ Component }">
          <template v-if="Component">
            <transition appear name="zoom-fade" mode="out-in">
              <keep-alive>
                <suspense>
                  <component :is="Component"></component>
                </suspense>
              </keep-alive>
            </transition>
          </template>
        </router-view>
      </q-page-container>
    </q-layout>
  </q-scroll-area>
</template>

<script lang="ts" setup>
import RexNavbar from './RexNavbar.vue';
import RexSidebar from './sidebar/RexSidebar.vue';
import { onMounted, ref } from 'vue';
import { useLayoutStore } from '@/stores/layout-store';
import { getCssVar } from 'quasar';
const primary: string = getCssVar('primary') || '#626262';

const layout = useLayoutStore();
const qLayoutRef = ref<{ $el: HTMLElement } | null>(null);
onMounted(() => {
  if (qLayoutRef.value == null) {
    return;
  }
  layout.width = qLayoutRef.value.$el.offsetWidth;
});
</script>

<style lang="scss" scoped>
.left-transition {

  .q-header,
  .q-page-container {
    transition: all 0.3s;
    z-index: 999;
  }
}

::v-deep(.q-page-container) {
  .zoom-fade-enter-active,
  .zoom-fade-leave-active {
    transition: transform 0.35s, opacity 0.28s ease-in-out;
  }

  .zoom-fade-enter-from {
    transform: scale(0.97);
    opacity: 0;
  }

  .zoom-fade-leave-to {
    transform: scale(1.03);
    opacity: 0;
  }
}
</style>
