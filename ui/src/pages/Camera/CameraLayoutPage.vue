<template>
  <q-page class="flex flex-col" padding>
    <div class="basis-full flex items-center">
      <div class="flex-1">
        <div class="text-2xl font-bold text-primary">
          {{ t('sidebar.camera') }}
        </div>
      </div>
      <div class="flex-none">
        <q-btn color="primary" icon="add" :label="t('pages.camera.addcamera')" />
      </div>
    </div>
    <div class="basis-full pt-6">
      <q-tabs align="left" v-model="tab_index" class="text-primary">
        <q-tab :label="t('pages.camera.tables')" name="tables" />
        <q-tab :label="t('pages.camera.groups')" name="groups" />
        <q-tab :label="t('pages.camera.map')" name="map" />
      </q-tabs>

      <q-tab-panels v-model="tab_index" animated>
        <q-tab-panel name="tables">
          <CameraTable/>
        </q-tab-panel>

        <q-tab-panel name="groups">
          <CameraGroup/>
        </q-tab-panel>
      </q-tab-panels>
    </div>
  </q-page>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import CameraTable from './CameraTable.vue';
import CameraGroup from './CameraGroup.vue'
const { t } = useI18n();
const tab_index = ref('tables');
</script>
