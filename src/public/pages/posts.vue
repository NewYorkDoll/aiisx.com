<script setup lang="ts">
definePageMeta({
  layout: false,
});
const { data } = await useAsyncGql("base");
console.log(data);
</script>
<template>
  <NuxtLayout name="posts">
    <div class="text-red-600">
      Posts Pages

      <n-button type="primary"> Error </n-button>
    </div>
  </NuxtLayout>
</template>
