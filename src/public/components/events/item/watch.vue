<template>
    <div class="text-amber-300">
      {{ props.event.payload.action == "started" ? "starred" : "unstarred" }}
      <EventsLink :href="props.event.repo.name" />
    </div>
  </template>
  
  <script setup lang="ts">
  import type { GithubEvent } from "@/lib/api"
  const props = defineProps<{
    event: GithubEvent
  }>()
  </script>