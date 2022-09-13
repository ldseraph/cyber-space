<template>
  <q-layout
    ref="qLayoutRef"
    view="lHh Lpr lFf"
    class="left-transition"
    @resize="(size) => (layout.width = size.width)"
  >
    <RexSidebar/>

    <RexNavbar/>

    <q-page-container>
      <router-view v-slot="{ Component }">
        <template v-if="Component">
          <transition appear :name="'zoom-fade'" mode="out-in">
            <keep-alive>
              <div class="flex justify-center overflow-hidden px-6">
                <component :is="Component"></component>
              </div>
            </keep-alive>
          </transition>
        </template>
      </router-view>
    </q-page-container>
  </q-layout>
</template>

<script lang="ts" setup>
import RexNavbar from './RexNavbar.vue';
import RexSidebar from './sidebar/RexSidebar.vue';
import { onMounted, ref } from 'vue';
import { useLayoutStore } from '@/stores/layout-store';

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
</style>
