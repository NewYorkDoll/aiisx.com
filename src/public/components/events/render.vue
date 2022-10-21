<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core";
import { vInfiniteScroll } from "@vueuse/components";
import EventItemPush from "./item/push.vue";
import EventItemCreate from "./item/create.vue";
import EventItemFork from "./item/fork.vue";
import EventItemIssues from "./item/issues.vue";
import EventIssueComment from "./item/issue-comment.vue";
import EventItemWatch from "./item/watch.vue";
import { _AsyncData } from "nuxt/dist/app/composables/asyncData";
import { GithubEventNode } from "@/lib/api/api";
import { GetEventsQuery } from "~~/.nuxt/gql-sdk";

const eventMap: { [name: string]: any } = {
  CreateEvent: EventItemCreate,
  PushEvent: EventItemPush,
  ForkEvent: EventItemFork,
  IssuesEvent: EventItemIssues,
  IssueCommentEvent: EventIssueComment,
  WatchEvent: EventItemWatch,
};

const fetched = ref<GithubEventNode[]>([]);

const pushData = (val: GetEventsQuery) => {
  const { githubevents } = val;
  hasNextPage.value = githubevents.pageInfo.hasNextPage;
  nextCursor.value = githubevents.pageInfo.endCursor;
  fetched.value = [...fetched.value, ...githubevents.edges!.map((item) => item?.node!)];
  emit("eventCount", fetched.value.length);
};

const emit = defineEmits<{
  (e: "eventCount", value: number): void;
}>();

const hasNextPage = ref(true);
const cursor = ref<string>();
const nextCursor = ref<string>();

const scrollContainer = ref<HTMLElement | null>(null);

async function fetchEvents() {
  if (!hasNextPage.value) return;
  cursor.value = nextCursor.value;
  const events = await useAsyncGql("getEvents", {
    count: 15,
    cursor: cursor.value,
  });
  if (events.data.value) {
    pushData(events.data.value);
  }
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
          <n-avatar round :size="15" :src="e.actor.avatarURL" class="mr-1 align-middle" />
        </a>
        <div class="flex items-center gap-2 truncate grow">
          <!-- <client-only fallbackTag="div"> -->
          <component :is="eventMap[e.eventType]" :event="e" />
          <!-- </client-only> -->
        </div>

        <div class="flex-none">
          <n-popover
            trigger="hover"
            style="padding: 2px 6px"
            :to="false"
            placement="left"
          >
            <template #trigger>
              <div class="i-mdi-clock-time-two-outline timestamp"></div>
            </template>
            {{ useTimeAgo(e.createdAt).value }}
          </n-popover>
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<style scoped lang="postcss">
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
