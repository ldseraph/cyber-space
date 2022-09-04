<template>
  <q-list class="text-gray-600" padding>
    <template v-for="(item, index) in items">
      <q-item-label
        class="ml-2"
        :key="`header-${index}`"
        v-if="'header' in item"
        header
      >
        {{ t(item.header) }}
      </q-item-label>
      <template v-else>
        <q-item
          v-if="!item.submenu"
          :key="`item-${index}`"
          :to="item.url"
          :disable="item.isDisabled"
          exact
          clickable
          v-ripple
        >
          <q-item-section class="pl-2" avatar>
            <q-icon :name="item.icon" size='1.5rem'/>
          </q-item-section>
          <q-item-section class="truncate">
            {{ t(item.name) }}
          </q-item-section>
          <q-item-section v-if="item.tag" side>
            <q-chip
              class="chip-shadow text-white"
              :color="item.tag?.color || 'primary'"
            >
              {{ t(item.tag.text) }}
            </q-chip>
          </q-item-section>
        </q-item>
        <template v-else>
          <q-expansion-item :key="`group-${index}`">
            <template v-slot:header>
              <q-item-section class="pl-2" avatar>
                <q-icon :name="item.icon" />
              </q-item-section>
              <q-item-section class="truncate">
                {{ t(item.name) }}
              </q-item-section>
            </template>
            <rex-sidebar-group
              class="pl-2"
              :items="item.submenu"
            />
          </q-expansion-item>
        </template>
      </template>
    </template>
  </q-list>
</template>

<script lang="ts" setup>
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
  }
  submenu?: sidebarItemsMenu[];
}

import { useI18n } from 'vue-i18n'
const { t } = useI18n()

defineProps<{
  items?: Array<sidebarHeaderMenu | sidebarItemsMenu>;
}>();

</script>
