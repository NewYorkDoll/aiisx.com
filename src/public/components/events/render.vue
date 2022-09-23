<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core";
import { vInfiniteScroll } from "@vueuse/components";
import EventItemPush from "./item/push.vue"
import EventItemCreate from "./item/create.vue"
import EventItemFork from "./item/fork.vue"

const eventMap: { [name: string]: any } = {
  "CreateEvent": EventItemCreate,
  "PushEvent": EventItemPush,
  "ForkEvent": EventItemFork
}



const emit = defineEmits<{
  (e: "eventCount", value: number): void;
}>();

const hasNextPage = ref(true);
const cursor = ref<string>("");
const nextCursor = ref<string>("");

const scrollContainer = ref<HTMLElement | null>(null);

function fetchEvents() {
  console.log("fetchEvents");

  if (!hasNextPage.value) return;
  cursor.value = nextCursor.value;
}

const events = await useAsyncGql("getEvents")
const { githubevents } = events.data.value!
const edges = githubevents.edges!.map((i) => i!.node!)
type node = typeof edges

const fetched = ref<node>([]);

watch(() => events, (val) => {
  if (!val.data.value) return
  const { githubevents } = val.data.value!
  hasNextPage.value = githubevents.pageInfo.hasNextPage
  nextCursor.value = githubevents.pageInfo.endCursor
  fetched.value = [...fetched.value, ...githubevents.edges!.map((item) => item?.node!)]
  emit("eventCount", fetched.value.length)
  setTimeout(() => {
    if (
      hasNextPage.value &&
      scrollContainer.value!.scrollHeight <=
      scrollContainer.value!.scrollTop + scrollContainer.value!.clientHeight
    ) {
      fetchEvents()
    }
  }, 300)
}, {
  immediate: true
})

</script>

<template>
  <div id="main" ref="scrollContainer" v-infinite-scroll="[fetchEvents, { distance: 40 }]">
    <TransitionGroup name="stepped" appear>
      <div v-for="(e, i) in fetched" :key="e.id" :style="{ '--i': fetched.length - i, '--total': fetched.length }"
        class="flex flex-auto flex-row items-center gap-x-1 px-1 hover:bg-zinc-500/10 text-zinc-400 transition duration-75 ease-out border-b-[1px] border-b-gray-100">
        <a :href="e.actor.login" target="_blank">
          <n-avatar round :size="15" :src="e.actor.avatarURL" class="mr-1 align-middle" />
        </a>
        <client-only>
          <component :is="eventMap[e.eventType]" :event="e" class="flex items-center gap-2 truncate grow" />
        </client-only>
        <div class="flex-none">
          <n-popover trigger="hover" style="padding: 2px 6px" :to="false" placement="left">
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
