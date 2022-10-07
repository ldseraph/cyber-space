<template>
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

    <template v-slot:body-cell-info="props">
      <q-td :props="props">
        <div class="flex flex-nowrap items-center space-x-4">
          <q-btn outline color="primary" label="RTSP">
            <q-popup-proxy>
              <q-card>
                <q-card-section>
                  <div class="flex space-x-2">
                    <div>
                      {{ t('camera.profile.rtsp.port') }}:
                    </div>
                    <div>
                      {{ props.row.rtspPort }}
                    </div>
                  </div>
                  <div class="flex space-x-2">
                    <div>
                      {{ t('camera.profile.rtsp.username') }}:
                    </div>
                    <div>
                      {{ props.row.rtspUsername }}
                    </div>
                  </div>
                  <div class="flex space-x-2">
                    <div>
                      {{ t('camera.profile.rtsp.pwd') }}:
                    </div>

                    <div class="flex flex-nowrap items-center">
                      <q-icon :name="props.row.rtspPwdIsShow ? 'visibility_off' : 'visibility'"
                        class="flex-1 cursor-pointer" @click="props.row.rtspPwdIsShow = !props.row.rtspPwdIsShow" />
                      <div class="pl-2" v-show="props.row.rtspPwdIsShow">{{ props.row.rtspPwd }}</div>
                    </div>
                  </div>
                </q-card-section>
              </q-card>
            </q-popup-proxy>
          </q-btn>
          <q-btn outline color="primary" label="GB28181">
            <q-popup-proxy>
            </q-popup-proxy>
          </q-btn>
          <q-btn outline color="primary" label="码流">
            <q-popup-proxy>
            </q-popup-proxy>
          </q-btn>
        </div>
      </q-td>
    </template>
  </q-table>
</template>

<script lang="ts" setup>
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

interface Camera {
  id: string;
  name: string;
  serialNumber: string;
  host: string;
  rtspPort: number;
  status: string;
  description: string;
  created: string;
  rtspUsername: string;
  rtspPwd: string;
  rtspPwdIsShow?: boolean;
}

interface Columns {
  /**
   * Unique id, identifies column, (used by pagination.sortBy, 'body-cell-[name]' slot, ...)
   */
  name: string;
  /**
   * Label for header
   */
  label: string;
  /**
   * Row Object property to determine value for this column or function which maps to the required property
   * @param row The current row being processed
   * @returns Value for this column
   */
  field: string | ((row: Camera) => string);
  /**
   * If we use visible-columns, this col will always be visible
   */
  required?: boolean;
  /**
   * Horizontal alignment of cells in this column
   * Default value: right
   */
  align?: 'left' | 'right' | 'center';
  /**
   * Tell QTable you want this column sortable
   */
  sortable?: boolean;
  /**
   * Compare function if you have some custom data or want a specific way to compare two rows
   * @param a Value of the first comparison term
   * @param b Value of the second comparison term
   * @param rowA Full Row object in which is contained the first term
   * @param rowB Full Row object in which is contained the second term
   * @returns Comparison result of term 'a' with term 'b'. Less than 0 when 'a' should come first; greater than 0 if 'b' should come first; equal to 0 if their position must not be changed with respect to each other
   */
  sort?: (a: Camera, b: Camera, rowA: Camera, rowB: Camera) => number;
  /**
   * Set column sort order: 'ad' (ascending-descending) or 'da' (descending-ascending); Overrides the 'column-sort-order' prop
   * Default value: ad
   */
  sortOrder?: 'ad' | 'da';
  /**
   * Function you can apply to format your data
   * @param val Value of the cell
   * @param row Full Row object in which the cell is contained
   * @returns The resulting formatted value
   */
  format?: (val: Camera, row: Camera) => string;
  /**
   * Style to apply on normal cells of the column
   * @param row The current row being processed
   */
  style?: string | ((row: Camera) => string);
  /**
   * Classes to add on normal cells of the column
   * @param row The current row being processed
   */
  classes?: string | ((row: Camera) => string);
  /**
   * Style to apply on header cells of the column
   */
  headerStyle?: string;
  /**
   * Classes to add on header cells of the column
   */
  headerClasses?: string;
}

const columns: Columns[] = [
  {
    name: 'serialNumber',
    label: t('camera.profile.serialNumber'),
    field: (row: Camera) => row.serialNumber,
    required: true,
    align: 'right',
    headerClasses: 'q-table--col-auto-width',
    sortable: true,
  },
  {
    name: 'name',
    label: t('camera.profile.name'),
    field: (row: Camera) => row.name,
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
  },
  {
    name: 'host',
    label: t('camera.profile.host'),
    field: (row: Camera) => row.host,
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
  },
  {
    name: 'status',
    label: t('camera.profile.status'),
    field: (row: Camera) =>
      t('camera.status.' + row.status, t('other.unknown')),
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
  },
  {
    name: 'created',
    label: t('camera.profile.created'),
    field: function (row: Camera) {
      return d(row.created + 'Z', 'long')
    },
    align: 'center',
    headerClasses: 'q-table--col-auto-width',
    sortable: true,
  },
  {
    name: 'description',
    label: t('camera.profile.description'),
    field: (row: Camera) => row.description || '',
    align: 'left',
  },
  {
    name: 'info',
    label: t('camera.profile.info'),
    field: () => '',
    align: 'left'
  }
  // {
  //   name: 'rtspUsername',
  //   label: t('camera.profile.rtspUsername'),
  //   field: (row: Camera) => row.rtspUsername,
  //   align: 'center',
  //   headerClasses: 'q-table--col-auto-width',
  // },
  // {
  //   name: 'rtspPwd',
  //   label: t('camera.profile.rtspPwd'),
  //   field: (row: Camera) => {
  //     return row.rtspPwd
  //   },
  //   align: 'center',
  //   headerClasses: 'q-table--col-auto-width',
  // },
];

async function fetchFromServer(page: number, rowsPerPage: number, filter: string, sortBy: string, descending: boolean) {
  return await Client.records.getList(
    'cameras',
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

