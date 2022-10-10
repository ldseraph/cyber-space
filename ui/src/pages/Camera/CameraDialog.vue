<template>
  <q-btn color="primary" icon="add" :label="t('pages.camera.addcamera')" @click="add = true"></q-btn>
  <q-dialog v-model="add" persistent>
    <q-card class="w-1/3">
      <q-card-section class="flex items-center ">
        <div class="flex text-base font-bold pl-3"> {{t('pages.camera.addcamera')}} </div>
        <q-space />
        <q-btn icon="close" size='sm' flat round dense v-close-popup />
      </q-card-section>

      <q-card-section class="flex">
        <div class="flex flex-col flex-1 space-y-6 px-6">
          <div v-for="(item,index,key) in items" :key="key" class="flex flex-col ">
            <div class="flex">
              {{t('camera.profile.'+String(index))}}
            </div>
            <div class="flex flex-1">
              <input type="text" class="text-sm box-border flex-1 h-8 border pl-4 rounded  shadow-sm border-slate-300
            focus:outline-none focus:border-sky-500 focus:ring-1" v-model="items[index]">
            </div>
          </div>
        </div>

        <!-- <div class="flex flex-col flex-1 space-y-6 px-6 ">
          <div class="flex justify-start items-center" >
            <div class="flex text-sm pr-3"> {{t('camera.profile.serialNumber')}} </div>
            <input type="text" class="text-sm box-border flex-1 h-8 border pl-4 rounded  shadow-sm border-slate-300
            focus:outline-none focus:border-sky-500 focus:ring-1" v-model="serialNumber">
          </div>
          <div class="flex justify-start items-center">
            <div class="flex text-sm pr-3"> {{t('camera.profile.name')}} </div>
            <input type="text" class="text-sm box-border flex-1 h-8 border pl-4 rounded shadow-sm border-slate-300
             focus:outline-none focus:border-sky-500 focus:ring-1">
          </div>
          <div class="flex justify-start items-center">
            <div class="flex text-sm pr-3"> {{t('camera.profile.host')}} </div>
            <input type="text" class="text-sm box-border flex-1 h-8 border pl-4 rounded shadow-sm border-slate-300
              focus:outline-none focus:border-sky-500 focus:ring-1">
          </div>
          <div class="flex justify-start items-center">
            <div class="flex text-sm pr-3"> {{t('camera.profile.status')}} </div>
            <input type="text" class="text-sm box-border flex-1 h-8 border pl-4 rounded shadow-sm border-slate-300
             focus:outline-none focus:border-sky-500 focus:ring-1">
          </div>
          <div class="flex justify-start items-center">
            <div class="flex text-sm pr-7"> {{t('camera.profile.created')}} </div>
            <input type="text" class="text-sm box-border flex-1 h-8 border pl-4 rounded  shadow-sm border-slate-300
              focus:outline-none focus:border-sky-500 focus:ring-1">
          </div>
          <div class="flex justify-start items-center">
            <div class="flex text-sm pr-7"> {{t('camera.profile.description')}} </div>
            <input type="text" class="text-sm box-border flex-1 h-8 border pl-4 rounded  shadow-sm border-slate-300
               focus:outline-none focus:border-sky-500 focus:ring-1">
          </div>
          <div class="flex justify-start items-center">
            <div class="flex text-sm pr-7"> {{t('camera.profile.info')}} </div>
            <input type="text" class="text-sm box-border flex-1 h-8 border pl-4 rounded  shadow-sm border-slate-300
               focus:outline-none focus:border-sky-500 focus:ring-1">
          </div>

        </div> -->
      </q-card-section>

      <q-card-actions class="text-white space-x-3 justify-end">
        <q-btn color="primary" :label="t('cancel')" v-close-popup />
        <q-btn color="primary" :label="t('ok')" type="submit" @click="addCameraConfirm" v-close-popup />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>


<script lang="ts" setup>
import { ref, reactive } from 'vue';
import { Client } from '@/utils/api';
import { useI18n } from 'vue-i18n';
const { t } = useI18n();
const items = reactive(
  {
    serialNumber: '',
    name: '',
    host: '',
    rtspPort: '',
    status: '',
    description: '',
    rtspUsername: '',
    rtspPwd: ''
    
  }

)
const add = ref(false)

async function addCameraConfirm() {

  const newData = await Client.records.create('cameras', items);
}
</script>