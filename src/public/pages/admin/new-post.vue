<script setup lang="ts">
import AdminPostCreate from "@/components/admin/post/creat.vue";
definePageMeta({
  layout: false,
});
const router = useRouter();
function createPost(val: any) {
  console.log(val);
}
</script>

<template>
  <NuxtLayout name="admin">
    <n-page-header class="hidden px-5 mt-4 mb-8 md:block">
      <template #avatar>
        <n-icon :size="40">
          <div class="i-mdi-pencil-outline"></div>
        </n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit">
          Create new post
        </a>
      </template>
    </n-page-header>

    <div class="p-4 sm:container sm:mx-auto lg:p-0">
      <ClientOnly>
        <AdminPostCreate create @update:post="createPost"> </AdminPostCreate>
      </ClientOnly>
    </div>
  </NuxtLayout>
</template>
