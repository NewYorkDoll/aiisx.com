<script setup lang="ts">
import { useMessage } from "naive-ui";
import LabelSelect from "./select.vue";
const message = useMessage();

const props = defineProps<{
  modelValue: string[];
  suggest?: string;
}>();
const emit = defineEmits(["update:modelValue"]);

type label = typeof LabelSelect;
const selectRef = ref<label | null>(null);

const selected = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
});

const newLabelInput = ref("");

function createNewLabel(val: string) {
  useAsyncGql("createLabel", {
    input: {
      name: val,
    },
  }).then((result) => {
    if (!result.error) {
      selectRef.value!.refetch().then(() => {
        selected.value = [...(selected.value ?? []), result.data.value!.createLabel.id];
      });
      newLabelInput.value = "";
      message.success("Created label");
    } else {
      message.error(result.error.toString());
    }
  });
}
</script>
<template>
  <div>
    <span class="text-emerald-500">Update tags</span>
    <LabelSelect
      ref="selectRef"
      v-model="selected"
      field="id"
      :suggest="props.suggest"
      class="mb-3"
    />

    <n-input
      v-model:value="newLabelInput"
      placeholder="Create label"
      @keyup.enter="createNewLabel(newLabelInput)"
    >
      <template #prefix>
        <n-icon><div class="i-mdi-tag"></div></n-icon>
      </template>
    </n-input>
  </div>
</template>
