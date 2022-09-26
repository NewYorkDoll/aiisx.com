<script setup lang="ts">
import { GithubEventNode } from "~~/lib/api/api";

const props = defineProps<{
  event: GithubEventNode;
}>();
</script>

<template>
  <div>
    <span class="text-lime-500">created</span>
    <template v-if="props.event.payload.ref">
      <EventsHoverItem :value="props.event.payload.ref" class="truncate text-zinc-200">
        <template #icon>
          <div
            class="i-mdi-source-branch inline-block align-middle"
            v-if="props.event.payload.ref_type == 'branch'"
          ></div>
          <div v-else class="i-mdi-tag inline-block align-middle"></div>
        </template>

        {{ props.event.payload.ref }}
      </EventsHoverItem>
      on
    </template>
    <EventsLink :href="props.event.repo.name" />
  </div>
</template>
