<script setup lang="ts">

const props = defineProps<{
    event: any
}>()

const repo = ref(props.event.repo)
const fork = ref<Record<string, any>>(props.event.payload.forkee)
</script>
    
<template>
    <div>
        <span class="text-purple-400">forked</span>
        <EventsLink :href="repo.name" />
        to
        <EventsLink :href="fork.full_name" />

        <EventsBlame>{{ fork.description }}</EventsBlame>
    </div>
</template>
  
