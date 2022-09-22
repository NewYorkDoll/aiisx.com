<script lang="ts" setup>
import { useBaseState } from "~~/stores/state";
import { branchMenuOptions } from "@/lib/core/navigation";
import { version as vueVersion } from "vue";

const state = useBaseState();
</script>
<template>
  <div class="flex items-stretch justify-center flex-auto w-full h-full lg:items-center">
    <div
      class="flex h-full w-full shrink grow-0 basis-auto flex-col items-stretch pt-[15px] md:items-center lg:max-h-[28rem] lg:max-w-3xl">
      <CoreNavigation />
      <CoreTerminal class="mb-4 flex flex-auto justify-center text-[38px] md:text-[45px]" path="~" :prefix="'#'"
        :value="'Hello!'" />
      <n-card content-style="padding: 0;display: flex;flex-direction: column;" class="flex flex-auto w-full h-full p-0">
        <slot />
        <n-layout-footer :bordered="true" class="bottom-bar">
          <span class="flex flex-auto">
            <n-tooltip trigger="hover">
              <template #trigger>
                <span class="bar-item misc">
                  <n-icon class="align-middle">
                    <div class="i-logos-visual-studio-code"></div>
                  </n-icon>
                </span>
              </template>
              build date: {{state.base!.version.date}}
            </n-tooltip>

            <n-popover placement="top-start" raw :show-arrow="false"
              class="rounded border border-solid border-zinc-700 shadow-none !m-0">
              <template #trigger>
                <span class="bar-item">
                  <n-icon class="align-middle">
                    <div class="i-mdi-source-branch"></div>
                  </n-icon>
                  {{
                  branchMenuOptions.find((o) => o.to.name === $route.name)!
                  .name
                  }}
                </span>
              </template>

              <ul class="flex flex-col flex-nowrap">
                <li v-for="link in branchMenuOptions" :key="link.name">
                  <router-link v-slot="{ isActive, href, navigate }" :to="link.to" active-class="bg:zinc-900/50" custom>
                    <a :href="href"
                      class="flex items-center px-1 rounded hover:bg-zinc-900/50 hover:text-zinc-300 text-zinc-400"
                      :class="{ 'bg-zinc-900/50 text-zinc-300': isActive }" @click="navigate">
                      <div class="i-mdi-source-branch"></div>
                      <span> {{ link.name }} </span>
                    </a>
                  </router-link>
                </li>
              </ul>
            </n-popover>

            <span class="bar-item misc">
              <n-icon class="mr-1 align-middle">
                <div class="i-logos-gopher"></div>
              </n-icon>
              {{ state.base!.version.goVersion.replace(/^go/, "") }}
            </span>
            <span class="bar-item misc">
              <n-icon class="mr-1 align-middle">
                <div class="i-logos-vue"></div>
              </n-icon>
              {{ vueVersion }}
            </span>

            <slot name="footer">
              <span class="ml-auto" />
            </slot>

            <n-tooltip trigger="hover">
              <template #trigger>
                <span class="bar-item misc"> spaces:4 </span>
              </template>
              ... or just gofmt
            </n-tooltip>
            <n-tooltip trigger="hover">
              <template #trigger>
                <a class="px-2 transition bg-blue-600 rounded-br-sm hover:bg-blue-800 hover:text-current"
                  style="color: white !important" :href="state.base?.githubUser.htmlurl!">
                  <n-icon class="align-middle mt-[-3px] mr-[-7px]">
                    <div class="i-mdi-github"></div>

                  </n-icon>
                  {{ state.base?.githubUser.login }}
                </a>
              </template>
              <p>
                <n-icon class="align-middle">
                  <div class="i-mdi-github"></div>
                </n-icon>
                {{ state.base?.githubUser.bio }} &middot;
                {{ state.base?.githubUser.name }}
              </p>
            </n-tooltip>
          </span>
        </n-layout-footer>
      </n-card>
    </div>
  </div>
</template>
<style>
.bottom-bar {
  line-height: 1.5;
  @apply text-[1.2em] md:text-[1em];
}

.bar-item {
  @apply pl-1.5 pr-2 rounded-br-sm inline-flex text-zinc-400 align-middle cursor-pointer transition hover:bg-zinc-800;
}

.bar-item.misc {
  @apply hidden sm:inline-flex;
}
</style>
  