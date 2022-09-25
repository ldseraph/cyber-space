import { defineStore } from 'pinia';
import configs from '@/configs';
const navLang = navigator.language;
const localLang = (navLang === 'zh-CN' || navLang === 'en-US') ? navLang : false;

export const useGlobalStore = defineStore('global', {
  state: () => ({
    lang: localLang || configs.app.lang || 'zh-CN'
  }),
  persist: {
    paths: ['lang'],
  },
});
