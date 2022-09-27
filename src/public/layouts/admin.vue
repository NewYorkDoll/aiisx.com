<script setup lang="ts">
import { adminSidebarOptions } from "@/lib/core/navigation";
import { useBaseState } from "~~/stores/state";
const props = reactive({
  loading: false,
  error: "",
});

const state = useBaseState();
const router = useRouter();
if (!state.base || !state.base.self) {
  router.replace("/404");
}
console.log(state.base?.self);
</script>

<template>
  <n-layout
    position="absolute"
    content-style="display: flex;flex: 1 1 auto;flex-direction: column"
  >
    <n-layout has-sider>
      <n-layout-sider
        :collapsed-width="54"
        :width="220"
        bordered
        show-trigger
        collapse-mode="width"
      >
        <n-menu
          :options="adminSidebarOptions"
          :root-indent="16"
          value="post-admin"
          :indent="12"
          :collapsed-width="54"
          :collapsed-icon-size="22"
        />
      </n-layout-sider>
      <n-layout
        embedded
        content-style="max-height: 100%;display: flex;flex: 1 1 auto;flex-direction: column"
      >
        <!-- <CoreBreadcrumbs class="hidden pt-2 pl-3 md:block" /> -->

        <n-alert
          v-if="props.error"
          title="An error occurred"
          type="error"
          class="m-2 md:m-6"
        >
          {{ props.error }}
        </n-alert>
        <template v-else-if="props.loading">
          <n-spin class="flex-auto">
            <template #description> Loading... </template>
          </n-spin>
        </template>
        <template v-else>
          <slot></slot>
        </template>
      </n-layout>
    </n-layout>
  </n-layout>
</template>
