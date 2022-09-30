<script setup lang="ts">
import AdminPostCreate from "@/components/admin/post/creat.vue";
import { useMessage } from "naive-ui";
import { UpdatePostInput } from "~~/.nuxt/gql-sdk";

import { Post } from "~~/lib/api";

definePageMeta({
  layout: false,
});
const route = useRoute();
const message = useMessage();

const router = useRouter();
const id = route.params.id as string;
const { data, error, pending } = await useAsyncGql("getPost", {
  id,
});

const post = computed(() => data.value?.node as NonNullable<Post>);

const updatePost = (val: UpdatePostInput, labelIDs: string[]) => {
  const addedLabels = labelIDs.filter(
    (id) => !post.value.labels?.edges?.map((item) => item!.node!.id).includes(id)
  );
  const removedLabels = post.value.labels?.edges
    ?.map((item) => item!.node!.id)
    .filter((id) => !labelIDs.includes(id));

  useAsyncGql("updatePost", {
    id,
    input: {
      title: val.title,
      content: val.content,
      slug: val.slug,
      addLabelIDs: addedLabels,
      removeLabelIDs: removedLabels,
      publishedAt: val.publishedAt,
      public: val.public,
    },
  }).then((result) => {
    if (!result.error.value) {
      message.success("Post created successfully");
      router.replace({ name: "admin-posts" });
    } else {
      message.error(result.error.value.gqlErrors[0].message);
    }
  });
};
</script>
<template>
  <NuxtLayout name="admin">
    <div class="p-4 sm:container sm:mx-auto lg:p-0">
      <n-page-header :subtitle="post.title" class="hidden mt-4 mb-8 lg:block">
        <template #avatar>
          <n-icon :size="40"> <div class="i-mdi-pencil-outline"></div> </n-icon>
        </template>
        <template #title>
          <a href="#" class="no-underline capitalize" style="color: inherit">
            Editing post #{{ post.id }}
          </a>
        </template>
      </n-page-header>

      <n-spin :show="pending">
        <n-alert v-if="error" title="Error pending post" type="error">
          {{ error }}
        </n-alert>

        <AdminPostCreate v-if="post" :post="post" @update:post="updatePost" />
      </n-spin>
    </div>
  </NuxtLayout>
</template>
