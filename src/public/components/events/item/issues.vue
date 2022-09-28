<script setup lang="ts">
import { GithubEventNode } from "~~/lib/api/api";
const props = defineProps<{
  event: GithubEventNode;
}>();
const repo = ref(props.event.repo);
const action = ref<string>(props.event.payload.action);
const issue = ref<Record<string, any>>(props.event.payload.issue);
</script>
<template>
  <div>
    <span class="text-cyan-400">
      <span v-if="['opened', 'edited', 'closed', 'reopened'].includes(action)">
        {{ action }}
      </span>
      <span v-else-if="['assigned', 'unassigned'].includes(action)">
        changed assignees on
      </span>
      <span v-else-if="['labeled', 'unlabeled'].includes(action)">
        edited labels on
      </span>
    </span>

    issue
    <EventsHoverItem :href="issue.html_url" :value="'#' + issue.number">
      {{ issue.title }}
    </EventsHoverItem>

    on <EventsLink :href="repo.name" />

    <EventsBlame>{{ issue.title }}</EventsBlame>
  </div>
</template>
