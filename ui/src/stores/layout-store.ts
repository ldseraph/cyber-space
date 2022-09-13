import { defineStore } from 'pinia';
import { useQuasar } from 'quasar';
import configs from '@/configs';

export const useLayoutStore = defineStore('layout', {
  state: () => ({
    width: 0,
    height: 0,
    title: configs.app.name,
    name: configs.layout.name,
  }),

  getters: {
    belowBreakpoint: (state) => state.width < useQuasar().screen.sizes.md,
  },
});
