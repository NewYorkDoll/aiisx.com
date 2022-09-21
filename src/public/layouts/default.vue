<script lang="ts" setup>
console.log("base default");
</script>
<template>
  <div
    class="flex items-stretch justify-center flex-auto w-full h-full lg:items-center"
  >
    <div
      class="flex h-full w-full shrink grow-0 basis-auto flex-col items-stretch pt-[15px] md:items-center lg:max-h-[28rem] lg:max-w-3xl"
    >
      <CoreNavigation />
      <slot />
    </div>
  </div>
</template>
