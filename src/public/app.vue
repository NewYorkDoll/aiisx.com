<script setup lang="ts">
  import { useBaseState } from "@/stores/state";
import { darkTheme } from "naive-ui";
const state = useBaseState();

useAsyncData(()=>state.getBaseData())

</script>

<template>
  <n-config-provider :theme="darkTheme" abstract>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </n-config-provider>
</template>
