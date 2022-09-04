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
        name: 'Home',
        icon: 'o_home',
        url: '/',
      },
      {
        name: 'My Cyberware',
        icon: 'o_devices',
        url: '/cyberware'
      },
      {
        name: 'Cyber Marketplace',
        icon: 'o_store',
        url: '/marketplace',
      },
      {
        name: 'Cyber Intelligence',
        icon: 'o_psychology',
        url: '/intelligence'
      },
      {
        name: 'Cyberbining',
        icon: 'o_router'
      },
      {
        name: 'Settings',
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
