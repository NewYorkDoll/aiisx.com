<script setup lang="ts">
import { useBaseState } from "@/stores/state";
import { menuOptions } from "@/lib/core/navigation";

const state = useBaseState();

const sudoUrl = computed(() => {
  return state.base && state.base.self ? "/admin" : "/-/auth";
});
</script>
<template>
  <ul class="flex flex-wrap justify-center md:justify-start md:gap-x-2">
    <li v-for="link in menuOptions" :key="link.name">
      <router-link :to="link.to" active-class="active">
        {{ link.alias }}
      </router-link>
    </li>
    <li>
      <a :href="state.base?.githubUser.htmlurl!" target="_blank">github</a>
    </li>

    <li><a :href="sudoUrl">sudo</a></li>
  </ul>
</template>

<style scoped>
ul :deep(a) {
  position: relative;
  display: inline-flex;
  padding-left: 5ch;
  padding-right: 4ch;
  @apply py-3;
  @apply !text-violet-400 transition duration-100;
}

ul :deep(a):hover,
ul :deep(a).active {
  padding-right: 4ch;
  @apply !text-violet-500;
}

ul :deep(a)::before {
  content: "func";
  @apply absolute left-0 text-red-400;
}

ul :deep(a)::after {
  content: "()";
  @apply absolute right-[2ch] text-yellow-300 transition duration-100;
}

ul :deep(a):hover::after {
  content: "(ctx)";
  @apply absolute right-0 text-red-400;
}

ul :deep(a):hover::after {
  content: "(go)";
  @apply absolute right-0 text-zinc-400;
}
</style>
