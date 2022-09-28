<script setup lang="ts">
  import type { ResultProps } from "naive-ui"
  
  const route = useRoute()
  console.log(route);
  
  const source = computed(() => route.params.catchall)
  const errorCode = ref<ResultProps["status"]>()
  const errorTitle = ref<string>("")
  const supported = ["info", "success", "warning", "error", "404", "403", "500", "418"]
  
    for (const item of source.value) {
      if (supported.includes(item)) {
        errorCode.value = item as ResultProps["status"]
      } else if (item.match(/^[45][0-9]+$/)) {
        errorCode.value = "error"
      } else if (item == "CombinedError") {
        errorCode.value = "error"
        errorTitle.value = "query error"
      } else {
        errorCode.value = "404"
        errorTitle.value = "not found"
        break
      }
    }
  
    if (errorTitle.value == "") {
      errorTitle.value = errorCode.value!
    }
  </script>
  

<template>
  <div class="flex flex-col justify-center flex-auto">
    <n-result
      :status="errorCode"
      :title="'Error code: ' + errorTitle"
      :description="'You know life is always ridiculous.'"
    >
      <template #footer>
        <n-button-group>
          <n-button @click="$router.back()">
            <n-icon class="mr-1"><div class="i-mdi-undo-variant"></div></n-icon>
            Go back
          </n-button>
          <n-button @click="$router.push('/')">
            <n-icon class="mr-1"><div class="i-mdi-home"></div></n-icon>
            Home
          </n-button>
        </n-button-group>
      </template>
    </n-result>
  </div>
</template>

