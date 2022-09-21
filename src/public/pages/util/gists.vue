<script setup lang="ts">
definePageMeta({
  layout: "terminal",
});

const eventCount = ref<number>();
</script>
<template>
  <EventsRender
    class="relative w-full h-full overflow-x-hidden grow basis-0"
    @event-count="(e) => (eventCount = e)"
  />
</template>
