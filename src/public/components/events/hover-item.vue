<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    value?: string;
    href?: string;
    placement?: "top" | "bottom" | "left" | "right";
  }>(),
  {
    value: "",
    href: "",
    placement: "right",
  }
);
</script>

<template>
  <n-popover
    trigger="hover"
    style="padding: 2px 6px"
    to="body"
    :placement="props.placement"
  >
    <template #trigger>
      <span v-if="props.href?.length > 0" class="align-middle" v-bind="$attrs">
        <EventsLink :href="props.href">
          <slot name="icon" />
          <slot name="value">{{ props.value }}</slot>
        </EventsLink>
      </span>
      <span v-else class="align-middle max-w-[24ch]" v-bind="$attrs">
        <slot name="icon" />
        <slot name="value">{{ props.value }}</slot>
      </span>
    </template>
    <slot />
  </n-popover>
</template>
