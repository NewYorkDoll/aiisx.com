<script setup lang="ts">
import { darkTheme } from "naive-ui";
</script>

<template>
  <n-config-provider :theme="darkTheme" abstract>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </n-config-provider>
</template>
