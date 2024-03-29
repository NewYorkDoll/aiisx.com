<script setup lang="ts">
import { useBaseState } from "~~/stores/state";

definePageMeta({
  layout: false,
});
const state = useBaseState();

interface LanguageBucket {
  language: string;
  totalSeconds: number;
  percentage?: number;
  titleLength?: number;
}

const codingStats = computed(() => {
  const out: LanguageBucket[] = [];
  let maxTitleLength = 5;

  // Bucket by language with a cap.
  for (const stat of state.base!.codingStats.languages!) {
    if (out.length === 6) {
      out[5].language = "Other";
      out[5].totalSeconds += stat.totalSeconds;
      continue;
    }

    maxTitleLength = Math.max(maxTitleLength, stat.language.length);
    out.push(stat);
  }

  // Calculate percentages.
  for (const stat of out) {
    stat.titleLength = maxTitleLength;
    stat.percentage = Math.round(
      (stat.totalSeconds / state.base!.codingStats.totalSeconds) * 100
    );
  }

  return out;
});

const eventCount = ref<number>(0);
const eventCountChange = (e: number) => {
  eventCount.value = e;
};
</script>
<template>
  <NuxtLayout name="terminal">
    <EventsRender
      class="relative w-full h-full overflow-x-hidden grow basis-0"
      @event-count="eventCountChange"
    />

    <template #footer>
      <n-popover
        placement="top-start"
        :width="250"
        raw
        :show-arrow="false"
        class="px-2 py-1 rounded border border-solid border-zinc-700 shadow-none !m-0"
      >
        <template #trigger>
          <span class="bar-item misc">
            <n-icon class="mr-1 align-middle text-violet-400">
              <div class="i-mdi-history"></div>
            </n-icon>
            {{ state.base!.codingStats.totalDuration }}
          </span>
        </template>

        <p class="text-violet-400">
          last {{ state.base!.codingStats.calculatedDays }} day coding stats
        </p>

        <div
          v-for="stat in codingStats"
          :key="stat.language"
          class="flex flex-row items-center flex-auto"
        >
          <n-tooltip trigger="hover">
            <template #trigger>
              <div
                class="text-right shrink-0 mr-[1ch] max-w-[10ch] truncate"
                :style="{ width: stat.titleLength + 'ch' }"
              >
                {{ stat.language }}
              </div>
            </template>
            <div>{{ stat.language }}</div>
          </n-tooltip>

          <div class="w-full rounded bg-zinc-900">
            <div
              class="h-2 rounded bg-gradient-to-r from-fuchsia-600 to-pink-600"
              :style="{ width: stat.percentage + '%' }"
            />
          </div>
          <div class="shrink-0 ml-[1ch] w-[3ch]">{{ stat.percentage }}%</div>
        </div>
      </n-popover>

      <span class="ml-auto" />
      <ClientOnly>
        <span v-if="eventCount > 0" class="bar-item misc"> ln:{{ eventCount }} </span>
      </ClientOnly>
    </template>
  </NuxtLayout>
</template>
