<script setup lang="ts">
import { useBaseState } from "~~/stores/state";

const props = defineProps<{
  loading?: boolean;
  error?: Error | string;
}>();
const state = useBaseState();
</script>

<template>
  <!-- TODO: less flex? -->
  <div class="z-[1] relative flex flex-auto flex-col">
    <div
      class="sm:container flex flex-auto flex-col pt-[15px] lg:pt-[70px] pb-[60px] max-w-[100vw] px-4 md:px-0 xl:px-[200px] sm:mx-auto"
    >
      <CoreNavigation />

      <template v-if="props.loading">
        <n-spin class="flex-auto">
          <template #description> Loading... </template>
        </n-spin>
      </template>
      <n-alert
        v-else-if="props.error"
        title="An error occurred"
        type="error"
        class="m-2 md:m-6"
      >
        {{ props.error }}
      </n-alert>
      <template v-else>
        <slot></slot>
      </template>
    </div>

    <span class="p-2 text-center">
      Made with
      <n-icon class="align-middle text-emerald-600">
        <div class="i-mdi-heart"></div>
      </n-icon>
      by
      <a
        :href="state.base!.githubUser.htmlurl!"
        target="_blank"
        >{{ state.base!.githubUser.login }}</a
      >
      <br />
      <a href="https://beian.miit.gov.cn/" target="_blank">宁ICP备16000711号</a>
    </span>
  </div>
</template>
