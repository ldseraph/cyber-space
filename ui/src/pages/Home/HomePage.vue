<template>
  <q-page class="flex w-full content-start">
    <div class="mb-16 w-full flex justify-center">
      <div class="w-1/2">
        <q-input v-model="search" :label="t('search')">
          <template v-slot:append>
            <q-btn round dense flat icon="search" />
          </template>
        </q-input>
      </div>
    </div>
    <div class="flex w-full flex-col content-start">
      <div class="px-4 w-1/4 h-full flex content-start max-w-sm">
        <div class="font-bold w-full pb-4">
          <q-icon name="star" class="text-amber-400" />
          {{ t('收藏') }}
        </div>
        <q-input
          class="w-full pb-4"
          filled
          v-model="input"
          label="Find a cyperware..."
        />
        <q-list class="w-full">
          <q-separator />
          <q-item
            v-for="cpberware in cpberwares"
            :key="cpberware.name"
            clickable
            v-ripple
          >
            <q-item-section avatar>
              <q-img :src="cpberware.img" />
            </q-item-section>

            <q-item-section>
              <q-item-label overline class="text-warning font-semibold">
                {{ cpberware.type }}
              </q-item-label>
              <q-item-label
                class="text-negative font-semibold"
                caption
                lines="1"
              >
                {{ cpberware.name }}
              </q-item-label>
            </q-item-section>

            <q-item-section side>
              <q-item-label class="text-warning font-semibold">
                {{ cpberware.quality }}
              </q-item-label>
            </q-item-section>
          </q-item>
        </q-list>
      </div>
      <div class="px-4 w-1/2 h-full">
        <q-card class="p-4 card-shadow">
          <q-card-section>
            <div class="text-base">摄像机</div>
          </q-card-section>
          <q-card-section class="flex flex-row-reverse">
            <q-btn color="primary" label="发布" />
          </q-card-section>
        </q-card>
      </div>
      <div class="px-4 w-1/4 h-full">
        <q-card class="p-4 card-shadow">
          <q-card-section>
            <div class="text-base">摄像机</div>
          </q-card-section>
          <q-card-section class="flex flex-row-reverse">
            <q-btn color="primary" label="发布" />
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts" setup>
import { useI18n } from 'vue-i18n';
import { ref } from 'vue';
const { t } = useI18n();
const search = ref('');
const input = ref('');
// 普通、精良、稀有、史诗、传说
const cpberwares = ref([
  {
    name: 'aassss DS-86xxN-I8',
    type: 'Visual Support',
    img: 'http://www.cctvhkws.com/uploads/190730/1-1ZI01F634529.jpg',
    quality: 'poor'
  },
  {
    name: 'hikvision DS-86xxN-I8',
    type: 'Visual Support',
    img: 'http://www.cctvhkws.com/uploads/190730/1-1ZI01F634529.jpg',
    quality: 'common'
  },
  {
    name: 'hikvision DS-86xxN-I8',
    type: 'Visual Support',
    img: 'http://www.cctvhkws.com/uploads/190730/1-1ZI01F634529.jpg',
    quality: 'uncommon'
  },
  {
    name: 'hikvision DS-86xxN-I8',
    type: 'Visual Support',
    img: 'http://www.cctvhkws.com/uploads/190730/1-1ZI01F634529.jpg',
    quality: 'rare'
  },
  {
    name: 'hikvision DS-86xxN-I8',
    type: 'Visual Support',
    img: 'http://www.cctvhkws.com/uploads/190730/1-1ZI01F634529.jpg',
    quality: 'epic'
  },
  {
    name: 'hikvision DS-86xxN-I8',
    type: 'Visual Support',
    img: 'http://www.cctvhkws.com/uploads/190730/1-1ZI01F634529.jpg',
    quality: 'legendary'
  }
]);
</script>
