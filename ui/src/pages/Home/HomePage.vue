<template>
  <q-page class="flex flex-col" padding>
    <q-card v-for="n in 110" :key="n" class="flex-1">
      Row {{ n }} / 150
    </q-card>
  </q-page>
</template>

<script lang="ts" setup>

</script>
