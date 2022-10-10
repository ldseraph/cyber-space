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

export function isSidebarItemsMenu(menuItem: sidebarHeaderMenu | sidebarItemsMenu): menuItem is sidebarItemsMenu {
  return (<sidebarItemsMenu>menuItem).name !== undefined;
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
        name: 'sidebar.camera',
        icon: 'o_videocam',
        submenu: [{
          name: 'sidebar.camera_rtsp',
          icon: 'img:public/rtsp_camera.svg',
          url: '/camera/rtsp'
        }, {
          name: 'sidebar.camera_gb28181',
          icon: 'img:public/gb28181_camera.svg',
          url: '/camera/gb28181'
        },{
          name: 'sidebar.camera_group',
          icon: 'o_account_tree',
          url: '/camera/group'
        },{
          name: 'sidebar.camera_map',
          icon: 'o_map',
          url: '/camera/map'
        }]
      },
      // {
      //   name: 'sidebar.marketplace',
      //   icon: 'o_store',
      //   url: '/marketplace',
      // },
      {
        name: 'sidebar.intelligence',
        icon: 'o_psychology',
        url: '/intelligence'
      },
      {
        name: 'sidebar.databining',
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
