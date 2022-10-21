<template>
  <div>
    <span class="text-lime-300" v-if="action == 'created'"> added </span>
    <span class="text-lime-300" v-else-if="action == 'edited'"> updated </span>
    <span class="text-lime-300" v-if="action == 'deleted'"> removed </span>

    a
    <EventsLink
      :href="comment.html_url"
      class="text-cyan-400 hover:text-cyan-500"
      value="comment"
    />
    to {{ issue.pull_request ? "pr" : "issue" }}

    <EventsHoverItem :href="issue.html_url" :value="'#' + issue.number">
      {{ issue.title }}
    </EventsHoverItem>

    on <EventsLink :href="repo.name" />

    <EventsBlame>{{ issue.title }}</EventsBlame>
  </div>
</template>

<script setup lang="ts">
import type { GithubEvent } from "@/lib/api";
const props = defineProps<{
  event: GithubEvent;
}>();
const repo = ref(props.event.repo);
const action = ref<string>(props.event.payload.action);
const comment = ref<Record<string, any>>(props.event.payload.comment);
const issue = ref<Record<string, any>>(props.event.payload.issue);
</script>
