<script setup lang="ts">
import { useBaseState } from "@/stores/state";
import { darkTheme } from "naive-ui";
const state = useBaseState();

useAsyncData(() => state.getBaseData());
</script>

<template>
  <n-config-provider :theme="darkTheme" abstract>
    <NuxtLayout>
      <n-message-provider>
        <NuxtPage />
      </n-message-provider>
    </NuxtLayout>
  </n-config-provider>
</template>
