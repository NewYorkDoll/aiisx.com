<script setup lang="ts">
import { h } from "vue";
import { NBadge } from "naive-ui";
import type { LabelWhereInput, Label } from "@/lib/api";
import { watchDebounced } from "@vueuse/core";
const props = withDefaults(
  defineProps<{
    modelValue: string | string[] | number | number[];
    field?: string;
    where?: LabelWhereInput | null;
    suggest?: string;
  }>(),
  {
    modelValue: (): string[] => [],
    field: "name",
    where: () => null,
    suggest: "",
  }
);
const emit = defineEmits(["update:modelValue"]);
const selected = computed({
  get: () => (Array.isArray(props.modelValue) ? props.modelValue : [props.modelValue]),
  set: (val) => emit("update:modelValue", val),
});
const labels = await useAsyncGql("getLabels", { where: props.where });
defineExpose({ refetch: labels.refresh });
type RenderedLabel<T> = {
  label: string;
  value: T;
  data: Label;
  popularity: number;
};
const options = computed(() =>
  labels.data?.value
    ?.labels!.edges!.map(
      (item: any) =>
        ({
          label: item!.node!.name,
          value: item!.node![props.field],
          data: item!.node!,
        } as RenderedLabel<typeof props.field>)
    )
    .sort((a, b) => b.popularity - a.popularity)
);
function renderLabel(option: RenderedLabel<typeof props.field>) {
  return [
    h(NBadge, {
      "show-zero": true,
      color: "grey",
      style: { "margin-right": "1ch" },
      value: option.popularity,
    }),
    option.data?.name,
  ];
}
const suggestions = ref<any[]>([]);
const suggest = computed(() => props.suggest);
watchDebounced(suggest, makeSuggestions, {
  debounce: 300,
  maxWait: 700,
  immediate: true,
});

function makeSuggestions(val: string | undefined) {
  if (!val || !options.value) return;
  const newSuggestions = [];
  for (const option of options.value) {
    if (selected.value.includes((option as any).data[props.field])) continue;
    const reg = new RegExp(`(^|\\W)${option.data.name}(\\W|$)`, "ig");
    if (suggest.value!.match(reg)) {
      newSuggestions.push(option);
    }
  }
  suggestions.value = newSuggestions;
}
function addSuggestion(option: any) {
  selected.value.push(option.data[props.field]);
  makeSuggestions(suggest.value);
}
</script>

<template>
  <n-select
    v-bind="$attrs"
    v-model:value="selected"
    :options="options"
    :render-label="renderLabel"
    :loading="labels.pending.value"
    clearable
    filterable
    multiple
    placeholder="Select labels"
  />
  <div v-if="suggestions.length > 0" class="flex flex-col mb-2">
    <span class="text-emerald-500">Suggestions</span>
    <div class="inline-flex flex-row flex-wrap gap-1">
      <n-tag
        v-for="label in suggestions"
        :key="label.data.id"
        class="cursor-pointer hover:bg-emerald-700"
        @click="addSuggestion(label)"
      >
        <n-badge show-zero color="grey" class="mr-0" :value="label.popularity" />
        {{ label.data.name }}
      </n-tag>
    </div>
  </div>
</template>
