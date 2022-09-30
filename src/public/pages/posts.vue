<script setup lang="ts">
import { useRouteQuery } from "@vueuse/router";

import { resetCursor, usePagination } from "@/lib/util/pagination";
import { useSorter } from "@/lib/util/sorter";
import { LabelWhereInput } from "~~/.nuxt/gql-sdk";
definePageMeta({
  layout: false,
});
const cursor = useRouteQuery<string | null>("cur", null);
const labels = useRouteQuery<Array<string>>("label", []);
const search = useRouteQuery<string>("q", "");
const filterSearch = refDebounced<string>(search, 300);
const direction = useRouteQuery<string>("dir", "desc");
const field = useRouteQuery<string>("sort", "date");

resetCursor(cursor, [labels, search, direction, field]);

const useVariables = () => ({
  ...usePagination(cursor, 10),
  where: {
    or: [
      { titleContainsFold: filterSearch.value },
      { summaryContainsFold: filterSearch.value },
    ],
    hasLabelsWith: labels.value as LabelWhereInput[],
  },
});

const variables = ref(useVariables());
const { data, error, pending, refresh } = await useAsyncGql({
  operation: "getPosts",
  variables: variables.value,
  options: {
    initialCache: true,
  },
});
watchThrottled(
  filterSearch,
  () => {
    variables.value = useVariables();
    useAsyncGql({
      operation: "getPosts",
      variables: variables.value,
      options: {
        initialCache: true,
      },
    }).then((res) => {
      data.value = res.data.value;
    });
  },
  { throttle: 1000 }
);
</script>
<template>
  <NuxtLayout name="sidebar">
    <div class="flex flex-auto gap-2 mt-1 mb-8">
      <n-input
        v-model:value="search"
        :loading="pending"
        type="text"
        clearable
        placeholder="Search for a post"
      >
        <template #prefix>
          <n-icon>
            <div class="i-mdi-search"></div>
          </n-icon>
        </template>
      </n-input>
    </div>
    <CoreObjectRender
      v-if="data?.posts"
      :value="data.posts"
      type="post"
      linkable
      show-empty
      divider
    />
  </NuxtLayout>
</template>
