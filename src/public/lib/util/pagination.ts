import { shallowEqual } from "@/lib/util/equal";

import type { ComputedRef, Ref, WatchStopHandle } from "vue";

type Filter = {
  first: number | null;
  last: number | null;
  before: string | null;
  after: string | null;
};

/**
 * usePagination - generates the necessary pagination data for a graphql query,
 * based off the provided cursor in the form of "<a|b>.<cursor>", where a = after,
 * and b = before.
 *
 * @export
 * @param {Ref<string>} cursor
 * @param {number} [size=10]
 * @return {Filter}
 */
export function usePagination(cursor: Ref<string | null>, size = 10): Filter {
  return {
    first: computed(() =>
      !cursor.value ? size : cursor.value?.startsWith("a.") ? size : null
    ).value,
    last: computed(() => (cursor.value?.startsWith("b.") ? size : null)).value,
    before: computed(() =>
      cursor.value?.startsWith("b.") ? cursor.value.split(".")[1] : null
    ).value,
    after: computed(() =>
      cursor.value?.startsWith("a.") ? cursor.value.split(".")[1] : null
    ).value,
  };
}

/**
 * resetCursor - reset the cursor to the first page when any of the provided refs
 * change.
 *
 * @export
 * @param {Ref<any>} cursor
 * @param {Ref<any>[]} refs
 * @return {WatchStopHandle}
 */
export function resetCursor(
  cursor: Ref<string | null>,
  refs: Ref<any>[]
): WatchStopHandle {
  return watch(refs, (newv, oldv) => {
    for (let i = 0; i < newv.length; i++) {
      if (!shallowEqual(newv[i], oldv[i])) {
        cursor.value = null;
        return;
      }
    }
  });
}
