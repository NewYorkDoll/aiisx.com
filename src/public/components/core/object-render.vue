<script setup lang="ts">
import { h } from "vue";
import PostObject from "@/components/post/object.vue";
import LabelObject from "@/components/label/object.vue";
import { GetPostsQuery } from "~~/.nuxt/gql-sdk";
type Label = NonNullable<
  NonNullable<NonNullable<GetPostsQuery["posts"]["edges"]>[number]>["node"]
>["labels"];

type Post = GetPostsQuery["posts"];
const props = defineProps<{
  value: Label | Post;
  showEmpty?: boolean;
  divider?: boolean;
  type?: string;
}>();

const objects = computed(() => {
  if (!props.value) {
    return [];
  }
  const results = typeMapper(props.value.edges);
  if (!Array.isArray(results)) {
    return [results];
  }
  return results;
});
const typeMapper = (o: any): any => {
  if (props.type === "post") {
    return o.map((item: any) => ({
      component: h(PostObject as any, { value: item.node }),
      object: item.node,
    }));
  } else {
    return o.map((item: any) => ({
      component: h(LabelObject as Label, { value: item.node }),
      object: item.node,
    }));
  }
};
</script>

<template>
  <n-empty
    v-if="showEmpty && objects.length < 1"
    description="No items found matching filters"
    class="!flex"
  />
  <TransitionGroup v-else-if="objects.length > 0" appear name="fade">
    <div
      v-for="(object, i) in objects"
      :key="object.object.id"
      :style="{ '--i': i, '--total': objects.length }"
    >
      <component :is="object.component" v-bind="$attrs" />
      <n-divider v-if="divider && i != objects.length - 1" />
    </div>
  </TransitionGroup>
  <n-empty v-else description="not found" class="!flex" />
</template>

<style scoped>
.fade-move,
.fade-enter-active,
.fade-leave-active {
  transition: all 0.1s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
