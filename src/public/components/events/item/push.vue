  
<script setup lang="ts">
import { GithubEventNode } from '~~/lib/api/api';


const props = defineProps<{
    event: GithubEventNode
}>()

const repo = ref(props.event.repo)
const commits = ref<Record<string, any>[]>(props.event.payload.commits)
</script>
    
<template>
    <div>
        <span style="color: #d086ff">pushed</span>

        <span v-if="commits.length > 1">
            <EventsHoverItem :href="repo.name + '/compare/' + commits[0].sha + '...' + props.event.payload.head"
                :value="commits.length + ' commits'" style="max-width: 350px">
                <p v-for="commit in commits" :key="commit.sha" class="truncate" :title="commit.message">
                <div class="i-mdi-source-commit"></div>
                {{ commit.sha.slice(0, 7) }}:
                {{ commit.message.split("\n")[0] }}
                </p>
            </EventsHoverItem>
        </span>
        <span v-else>
            <EventsHoverItem :href="repo.name + '/commit/' + commits[0].sha" :value="commits[0].sha.slice(0, 7)">
                <div class="i-mdi-source-commit"></div>
                {{ commits[0].sha.slice(0, 7) }}:
                {{ commits[0].message.split("\n")[0] }}
            </EventsHoverItem>
        </span>
        to
        <EventsLink :href="repo.name" />

        <EventsBlame>
            {{ commits[0].message.split("\n")[0] }}
        </EventsBlame>
    </div>
</template>
