import { RouteRecordRaw } from 'vue-router';
// import { useSidebarStore } from '@/stores/sidebar-store';

// const sidebar = useSidebarStore();
const layoutComponents = {
  main: () => import('layouts/MainLayout.vue'),
}

const routes: RouteRecordRaw[] = [
  {
    name: 'home',
    path: '/',
    components: layoutComponents,
    meta: {
      icon: 'o_home'
    },
    children: [
      { path: '', component: () => import('pages/Home/HomePage.vue') },
      { name: 'dashboard', path: 'dashboard', meta: { icon: 'o_dashboard' }, component: () => import('pages/Home/HomePage.vue') },
      {
        name: 'camera',
        path: 'camera',
        meta: { icon: 'o_videocam' },
        children: [{
          name: 'rtsp',
          path: 'rtsp',
          component: () => import('pages/Camera/RtspPage.vue')
        }]
      },
    ],
  },
  // {
  //   path: '/devices',
  //   components: layoutComponents,
  //   children: [{ path: 'profiles', component: () => import('pages/Devices/DeviceProfilesPage.vue') }],
  // },
  // Always leave this as last one,
  // but you can also remove it
  {
    name: 'NotFound',
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
