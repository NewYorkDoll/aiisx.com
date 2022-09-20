<script setup lang="ts">
import { darkTheme } from 'naive-ui'

</script>

<template>
  <div>
    <n-config-provider :theme="darkTheme">
      <div>123</div>
    </n-config-provider>
  </div>
</template>
