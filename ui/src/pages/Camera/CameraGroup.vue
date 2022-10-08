<template>
  <q-splitter v-model="splitterModel" before-class="p-4 min-w-fit">
    <template v-slot:before>
      <div class="flex flex-nowrap">
        <div class="pr-10">
          <q-tree ref="treeRef" :nodes="groupList" label-key="name" node-key="id" selected-color="primary"
            v-model:selected="selected" default-expand-all @lazy-load="groupLazyLoad" />
        </div>

        <q-separator vertical />

        <q-list class="pl-4" :bordered="false">
          <template v-for="(item, index) in cameraList" :key="`${index}`">
            <q-item-label>
              {{ item.name }}
            </q-item-label>
          </template>
        </q-list>
      </div>
    </template>

    <template v-slot:after>

    </template>
  </q-splitter>
</template>
<script lang="ts" setup>
import { ref, watch } from 'vue';
import { Client, Record } from '@/utils/api';
import cyrb53 from '@/utils/cyrb53';

// import { useI18n } from 'vue-i18n';
// const { t } = useI18n();

const splitterModel = ref(15);
const treeRef = ref<{ $el: HTMLElement } | null>(null);
const selected = ref('')
const cameraList = ref<Record[]>()

watch(selected, async (newValue: string) => {
  if (newValue == null) {
    return
  }
  console.log(newValue)
  let ids = newValue.split('@')
  let id = ids[ids.length - 1]

  try {
    const returnedData = await Client.records.getList('cameras', 1, 50, {
      filter: '(groups~"' + id + '")',
      expand: 'channels, groups'
    });
    cameraList.value = returnedData.items
  } catch (error) {
    cameraList.value = []
  }
})

const groupList = ref<Record[]>(await Client.records.getFullList('groups', 50, {
  filter: '(isRoot=true)',
  expand: 'subGroups'
}).then(function (d) {
  return d.map(function (item: Record) {
    item.children = childrenGroup(item.id, item)
    item.lazy = true
    if (selected.value == '') {
      selected.value = item.id
    }
    return item
  })
}))

function childrenGroup(id: string, item: Record) {
  if (item.subGroups.length) {
    return item['@expand'].subGroups.map(function (item1: Record) {
      item1.id = cyrb53(id) + '@' + item1.id
      item1.lazy = true
      return item1
    })
  }
  return []
}

interface Details {
  node: Record
  key: string
  done: (r: Record[]) => void
  fail: () => void
}

async function groupLazyLoad(details: Details) {
  let id = details.node.id.split('@')
  details.done(childrenGroup(details.node.id, await Client.records.getOne('groups', id[id.length - 1], {
    expand: 'subGroups'
  })))
}
</script>
