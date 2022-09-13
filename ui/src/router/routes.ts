import { RouteRecordRaw } from 'vue-router';

const layoutComponents = {
  main: () => import('layouts/MainLayout.vue'),
}

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    components: layoutComponents,
    children: [
      { path: '', component: () => import('pages/Home/HomePage.vue') },
      { path: '/dashboard', component: () => import('pages/Home/HomePage.vue') },
      { path: '/temp', component: () => import('pages/Temp/TempPage.vue') }
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
