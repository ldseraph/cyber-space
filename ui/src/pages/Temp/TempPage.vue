<template>
  <q-page class="flex" padding>
    <div class="flex-auto bg-white w-64" v-for="i in 1000" :key="i">
      aaaa {{ i }}
      <!-- <div class="font-bold w-full pb-4">
        <q-icon name="star" class="text-amber-400" />
        {{ t('收藏') }}
      </div>
      <q-input class="w-full pb-4" filled v-model="input" label="Find a cyperware..." />
      <q-list class="w-full">
        <q-separator />
        <q-item v-for="cpberware in cpberwares" :key="cpberware.name" clickable v-ripple>
          <q-item-section avatar>
            <q-img :src="cpberware.img" />
          </q-item-section>

          <q-item-section>
            <q-item-label overline class="text-warning font-semibold">
              {{ cpberware.type }}
            </q-item-label>
            <q-item-label class="text-negative font-semibold" caption lines="1">
              {{ cpberware.name }}
            </q-item-label>
          </q-item-section>

          <q-item-section side>
            <q-item-label class="text-warning font-semibold">
              {{ cpberware.quality }}
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-list> -->
    </div>
  </q-page>
</template>

<script lang="ts" setup>
</script>
