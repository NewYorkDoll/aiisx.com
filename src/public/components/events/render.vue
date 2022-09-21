<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core";
import { vInfiniteScroll } from "@vueuse/components";

const emit = defineEmits<{
  (e: "eventCount", value: number): void;
}>();

const fetched = ref<any[]>([]);
const hasNextPage = ref(true);
const cursor = ref<string>("");
const nextCursor = ref<string>("");

const scrollContainer = ref<HTMLElement | null>(null);

function fetchEvents() {
  if (!hasNextPage.value) return;
  cursor.value = nextCursor.value;
}

onMounted(() => {
  fetchEvents();
});
</script>

<template>
  <div
    id="main"
    ref="scrollContainer"
    v-infinite-scroll="[fetchEvents, { distance: 40 }]"
  >
    <TransitionGroup name="stepped" appear>
      <div
        v-for="(e, i) in fetched"
        :key="e.id"
        :style="{ '--i': fetched.length - i, '--total': fetched.length }"
        class="flex flex-auto flex-row items-center gap-x-1 px-1 hover:bg-zinc-500/10 text-zinc-400 transition duration-75 ease-out border-b-[1px] border-b-gray-100"
      >
        <a :href="e.actor.login" target="_blank">
          <n-avatar
            round
            :size="15"
            :src="e.actor.avatarURL"
            class="mr-1 align-middle"
          />
        </a>

        <!-- <component :event="e" class="flex items-center gap-2 truncate grow" /> -->
        <div class="flex-none">
          <!-- <EventHoverItem placement="left">
            <template #value>
              <i-mdi-clock-time-two-outline class="timestamp" />
            </template>

            {{ useTimeAgo(e.createdAt).value }}
          </EventHoverItem> -->
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<style scoped>
#main {
  font-size: 1.05em;
}

@screen md {
  #main {
    font-size: 1.15em;
  }
}

.timestamp {
  opacity: max(0.2, calc(var(--i) / var(--total)));
  @apply transition duration-100;
}

.timestamp:hover {
  opacity: 1;
}

div :deep(.icon) {
  font-size: 0.8em;
}
</style>
