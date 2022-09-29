<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core";

definePageMeta({
  layout: false,
});

const regenerate = () => {};

const { data, error } = await useAsyncGql("getPosts", {
  first: 100,
  before: null,
});
const posts = computed(() => data?.value?.posts!.edges!.map((v) => v!.node));
const deletePost = (val: any) => {
  console.log(val);
};
</script>

<template>
  <NuxtLayout name="admin">
    <div class="p-4 sm:container sm:mx-auto">
      <n-table bordered single-line striped size="small">
        <thead>
          <tr>
            <th>Title</th>
            <th class="hidden md:table-cell">Slug</th>
            <th class="hidden md:table-cell">Labels</th>
            <th class="hidden md:table-cell">Views</th>
            <th class="hidden md:table-cell">Published</th>
            <th>
              Actions
              <n-button class="ml-10" type="error" size="small" @click="regenerate">
                <n-icon>
                  <div class="i-mdi-reload"></div>
                </n-icon>
              </n-button>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="post in posts" :key="post!.id">
            <td>
              <router-link :to="{ name: 'p-slug', params: { slug: post!.slug } }">
                {{ post!.title }}
                <span v-if="!post!.public" class="text-yellow-500">[draft]</span>
              </router-link>
            </td>
            <td class="hidden md:table-cell">
              <router-link :to="{ name: 'p-slug', params: { slug: post!.slug } }">
                {{ post!.slug }}
              </router-link>
            </td>
            <td class="hidden md:flex">
              <n-popover
                style="max-height: 240px"
                trigger="hover"
                content-style="padding: 0;"
                scrollable
                placement="left"
              >
                <template #trigger>
                  <n-button size="small">{{ post!.labels.edges!.length }} tags</n-button>
                </template>

                <CoreObjectRender
                  :value="post!.labels"
                  linkable
                  class="grid mx-1 my-[2px]"
                />
              </n-popover>
            </td>
            <td class="hidden md:table-cell">
              <PostViewCount :value="post!.viewCount" />
            </td>
            <td class="hidden md:table-cell">
              {{ useTimeAgo(Date.parse(post!.publishedAt)).value }}
            </td>
            <td>
              <router-link :to="{ path: `admin-edit-post/${post!.id}` }">
                <n-button size="small" type="primary" tertiary>
                  <n-icon class="mr-1"><div class="i-mdi-pencil-outline"></div> </n-icon>
                  Edit
                </n-button>
              </router-link>
              <n-button
                size="small"
                type="error"
                tertiary
                class="ml-2"
                @click="deletePost(post!)"
              >
                <n-icon class="mr-1"><div class="i-mdi-trash-can-outline"></div></n-icon>
                Delete
              </n-button>
            </td>
          </tr>
        </tbody>
      </n-table>
    </div>
    <router-link
      :to="{ name: 'admin-new-post' }"
      class="absolute no-underline bottom-5 right-5"
    >
      <n-button tertiary circle size="large" type="primary" class="h-[13] w-[13]">
        <n-icon class="text-[2em]"><div class="i-mdi-pencil-outline"></div></n-icon>
      </n-button>
    </router-link>
  </NuxtLayout>
</template>
