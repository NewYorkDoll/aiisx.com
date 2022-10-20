<script setup lang="ts">
import { useRouteQuery } from "@vueuse/router";
import { nextTick } from "vue";
import { resetCursor, usePagination } from "@/lib/util/pagination";
import { useSorter } from "@/lib/util/sorter";
import { GetPostsQueryVariables } from "~~/.nuxt/gql-sdk";

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

let sorter = useSorter(
  {
    date: "date",
    title: "title",
    view_count: "popularity",
  },
  direction,
  field
);

const useVariables = () =>
  ({
    ...usePagination(cursor, 10),
    ...sorter.filter,
    where: {
      or: [
        { titleContainsFold: filterSearch.value },
        { summaryContainsFold: filterSearch.value },
      ],
      hasLabelsWith: labels.value.length ? { nameIn: labels.value } : null,
    },
  } as GetPostsQueryVariables);

const variables = ref(useVariables());

const { data, error, pending, refresh } = await useAsyncGql({
  operation: "getPosts",
  variables: variables.value,
  options: {
    initialCache: false,
  },
});

const getData = () => {
  console.log(direction);

  sorter = useSorter(
    {
      date: "date",
      title: "title",
      view_count: "popularity",
    },
    direction,
    field
  );
  variables.value = useVariables();

  useAsyncGql({
    operation: "getPosts",
    variables: variables.value,
    options: {
      initialCache: false,
    },
  }).then((res) => {
    data.value = res.data.value;
  });
};
watchThrottled(filterSearch, () => getData(), { throttle: 1000 });
const route = useRoute();

watch(
  () => route.query,
  () => getData()
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
      <CorePagination v-model="cursor" :page-info="data?.posts?.pageInfo" />
    </div>
    <CoreObjectRender
      v-if="data?.posts"
      :value="data.posts"
      type="post"
      linkable
      show-empty
      divider
    />

    <template #sidebar>
      <div class="text-center md:text-left">
        <div class="text-emerald-500">Sort posts</div>
        <CoreSorter :sorter="sorter" class="pb-4" />
        <div class="text-emerald-500">Filter by label</div>
        <LabelSelect v-model="labels" :where="{ hasPosts: true }" />
      </div>
    </template>
  </NuxtLayout>
</template>
