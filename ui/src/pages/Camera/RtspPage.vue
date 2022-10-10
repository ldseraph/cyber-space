<template>
  <q-page class="flex flex-col" padding>
    <RexBreadcrumb />
    <div class="flex-1">
      <q-table ref="tableRef" :rows="rows" :columns="columns" row-key="id" v-model:pagination="pagination"
        :loading="loading" color="primary" :filter="filter" binary-state-sort @request="onRequest">
        <template v-slot:top-right>
          <q-input borderless dense debounce="300" v-model="filter" :placeholder="t('other.search')">
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </template>

        <template v-slot:body-cell-status="props">
          <q-td :props="props">
            <div class="flex flex-nowrap items-center justify-center">
              <div v-if="props.row.status == 'online'" class="w-3 h-3 rounded-full bg-green-500"></div>
              <div v-else-if="props.row.status == 'offline'" class="w-3 h-3 rounded-full bg-red-500"></div>
              <div v-else class="w-3 h-3 rounded-full bg-stone-500"></div>
              <div class="pl-2">
                {{ props.value }}
              </div>
            </div>
          </q-td>
        </template>

        <template v-slot:body-cell-password="props">
          <q-td :props="props">
            <div class="flex flex-nowrap items-center">
              <q-icon :name="props.row.passwordIsShow ? 'visibility_off' : 'visibility'" class="flex-1 cursor-pointer"
                @click="props.row.passwordIsShow = !props.row.passwordIsShow" />
              <div class="pl-2" v-show="props.row.passwordIsShow">{{ props.row.password }}</div>
            </div>
          </q-td>
        </template>
      </q-table>
    </div>
  </q-page>
</template>

<script lang="ts" setup>
import RexBreadcrumb from '@/components/RexBreadcrumb.vue';
import { ref, onMounted } from 'vue';
import { Client, Record } from '@/utils/api';
import { useI18n } from 'vue-i18n';
const { t, d } = useI18n();

const tableRef = ref()
const rows = ref<Record[]>([])
const filter = ref('')
const loading = ref(false)
const pagination = ref({
  sortBy: 'id',
  descending: false,
  page: 1,
  rowsPerPage: 10,
  rowsNumber: 0
})

interface RtspCamera {
  id: string;
  created: string;
  name: string;
  host: string;
  port: number;
  user: string;
  password: string;
  description: string;
  status: 'offline' | 'online';
  passwordIsShow?: boolean;
}

interface Columns {
  name: string;
  label: string;
  field: string | ((row: RtspCamera) => string);
  required?: boolean;
  align?: 'left' | 'right' | 'center';
  sortable?: boolean;
  sort?: (a: RtspCamera, b: RtspCamera, rowA: RtspCamera, rowB: RtspCamera) => number;
  sortOrder?: 'ad' | 'da';
  format?: (val: RtspCamera, row: RtspCamera) => string;
  style?: string | ((row: RtspCamera) => string);
  classes?: string | ((row: RtspCamera) => string);
  headerStyle?: string;
  headerClasses?: string;
}

const columns: Columns[] = [
  {
    name: 'name',
    label: t('rtsp.profile.name'),
    field: (row: RtspCamera) => row.name,
    required: true,
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
    sortable: true,
  },
  {
    name: 'created',
    label: t('rtsp.profile.created'),
    field: function (row: RtspCamera) {
      return d(row.created + 'Z', 'long')
    },
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
    sortable: true,
  },
  {
    name: 'host',
    label: t('rtsp.profile.host'),
    field: (row: RtspCamera) => row.host,
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
  },
  {
    name: 'port',
    label: t('rtsp.profile.port'),
    field: (row: RtspCamera) => row.port.toString(),
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
  },
  {
    name: 'user',
    label: t('rtsp.profile.user'),
    field: (row: RtspCamera) => row.user,
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
  },
  {
    name: 'password',
    label: t('rtsp.profile.password'),
    field: (row: RtspCamera) => row.password,
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
  },
  {
    name: 'status',
    label: t('rtsp.profile.status'),
    field: (row: RtspCamera) =>
      t('rtsp.status.' + row.status, t('other.unknown')),
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
  },
  {
    name: 'description',
    label: t('rtsp.profile.description'),
    field: (row: RtspCamera) => row.description,
    align: 'center',
  }
];

async function fetchFromServer(page: number, rowsPerPage: number, filter: string, sortBy: string, descending: boolean) {
  return await Client.records.getList(
    'rtsp_cameras',
    page,
    rowsPerPage,
    {
      sort: (descending ? '-' : '') + sortBy,
      filter: filter == '' ? null : '(name~"' + filter + '")',
      expand: 'channels, groups'
    }
  );
}

interface Request {
  /**
   * Pagination object
   */
  pagination: {
    /**
     * Column name (from column definition)
     */
    sortBy: string;
    /**
     * Is sorting in descending order?
     */
    descending: boolean;
    /**
     * Page number (1-based)
     */
    page: number;
    /**
     * How many rows per page? 0 means Infinite
     */
    rowsPerPage: number;
  };
  /**
   * String/Object to filter table with (the 'filter' prop)
   */
  filter?: string;
}

async function onRequest(props: Request) {
  const { page, sortBy, rowsPerPage, descending } = props.pagination
  const filter = props.filter || ''

  loading.value = true
  const returnedData = await fetchFromServer(page, rowsPerPage, filter, sortBy, descending)
  pagination.value.rowsNumber = returnedData.totalItems
  pagination.value.page = page
  pagination.value.rowsPerPage = rowsPerPage
  pagination.value.descending = descending
  pagination.value.sortBy = sortBy
  let items = returnedData.items.map(function (item) {
    item.pwd_is_show = false
    return item
  })
  rows.value.splice(0, rows.value.length, ...items)
  loading.value = false
}

onMounted(() => {
  tableRef.value.requestServerInteraction()
})

</script>
