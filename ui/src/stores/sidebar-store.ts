import { defineStore } from 'pinia';
import configs from '@/configs';

export interface sidebarHeaderMenu {
  header: string;
}

export interface sidebarItemsMenu {
  url?: string;
  name: string;
  icon?: string;
  isDisabled?: boolean;
  tag?: {
    color?: string;
    text: string;
  };
  submenu?: sidebarItemsMenu[];
}

export const useSidebarStore = defineStore('sidebar', {
  state: () => ({
    title: String(configs.app.name),
    logo: String(configs.layout.logo),
    menu: Array<sidebarHeaderMenu | sidebarItemsMenu>(
      {
        name: 'sidebar.home',
        icon: 'o_home',
        url: '/',
      },
      {
        name: 'sidebar.dashboard',
        icon: 'o_dashboard',
        url: '/dashboard',
      },
      // {
      //   name: 'sidebar.organization',
      //   icon: 'o_corporate_fare',
      //   url: '/organization'
      // },
      {
        name: 'sidebar.cyberware',
        icon: 'o_devices',
        url: '/temp'
      },
      {
        name: 'sidebar.marketplace',
        icon: 'o_store',
        url: '/marketplace',
      },
      {
        name: 'sidebar.intelligence',
        icon: 'o_psychology',
        url: '/intelligence'
      },
      {
        name: 'sidebar.cyberbining',
        icon: 'o_router'
      },
      {
        name: 'sidebar.settings',
        icon: 'o_settings',
      }
    ),
    open: Boolean(true),
    mini: Boolean(false),
  }),
  actions: {
    taggle() {
      this.open = !this.open;
    },
  },
  persist: {
    paths: ['open', 'mini'],
  },
});
