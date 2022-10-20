<template>
  <div>
    <n-button-group>
      <n-button
        ghost
        :disabled="!page?.hasPreviousPage"
        @click="cursor = 'b.' + page?.startCursor"
      >
        <template #icon>
          <n-icon><div class="i-mdi-chevron-left"></div></n-icon>
        </template>
      </n-button>
      <n-button
        ghost
        :disabled="!page?.hasNextPage"
        @click="cursor = 'a.' + page?.endCursor"
      >
        <template #icon>
          <n-icon><div class="i-mdi-chevron-right"></div></n-icon>
        </template>
      </n-button>
    </n-button-group>
  </div>
</template>
<script setup lang="ts">
import { GetPostsQuery } from "~~/.nuxt/gql-sdk";

const props = defineProps<{
  modelValue: string | null;
  pageInfo?: GetPostsQuery["posts"]["pageInfo"];
}>();
const emit = defineEmits(["update:modelValue"]);
const cursor = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
});
const page = computed(() => props.pageInfo || null);
</script>
