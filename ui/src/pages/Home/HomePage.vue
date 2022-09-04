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
          {{ t("收藏") }}
        </div>
        <q-input class="w-full pb-4" filled v-model="input" label="Find a cyperware..." />
        <q-list class="w-full">
          <q-separator />
          <q-item-label class="font-bold" header>Online</q-item-label>

          <q-item clickable v-ripple>
            <!-- <q-item-section class="text-negative font-semibold" avatar>
              <q-item-label>安防摄像机</q-item-label>
            </q-item-section> -->

            <q-item-section>
              <q-item-label class="text-warning font-semibold">同事们的铁锤</q-item-label>
              <q-item-label class="text-negative font-semibold" caption lines="1">智能手枪</q-item-label>
            </q-item-section>

            <q-item-section side>
              <q-item-label class="text-warning font-semibold">传说/不朽</q-item-label>
            </q-item-section>
          </q-item>

          <q-item clickable v-ripple>
            <q-item-section avatar>
              <q-avatar color="primary" text-color="white">
                a
              </q-avatar>
            </q-item-section>

            <q-item-section>
              <q-item-label>a</q-item-label>
              <q-item-label caption lines="1">a</q-item-label>
            </q-item-section>

            <q-item-section side>
              <q-icon name="chat_bubble" color="green" />
            </q-item-section>
          </q-item>

          <q-separator />
          <q-item-label class="font-bold" header>Offline</q-item-label>

          <q-item clickable v-ripple>
            <q-item-section avatar>
              <q-avatar color="primary" text-color="white">
                a
              </q-avatar>
            </q-item-section>

            <q-item-section>
              <q-item-label>a</q-item-label>
              <q-item-label caption lines="1">a</q-item-label>
            </q-item-section>

            <q-item-section side>
              <q-icon name="chat_bubble" color="green" />
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
import { useI18n } from 'vue-i18n'
import { ref } from 'vue';
const { t } = useI18n()
const search = ref("")
const input = ref("")
</script>
