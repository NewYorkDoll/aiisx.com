import type { Ref } from "vue";

const desc = "desc";
const asc = "asc";
const directions: string[] = [desc, asc];

type SortFields<T> = {
  [key: string]: T;
};

type Filter = {
  order: string;
  orderBy: string;
};

export type Sorter = {
  toggle: (field: string) => void;
  fields: SortFields<string>;
  refs: { [key: string]: Ref<string> };
  filter: Filter;
};

/**
 * useSorter - generates the necessary sorting and ordering data for a graphql
 * query.
 *
 * @export
 * @param {SortFields<string>} fields
 * @param {Ref<string>} direction
 * @param {Ref<string>} field
 * @param {string} [defaultField=null] - the default field will be the first key in 'fields'.
 * @param {string} [defaultDirection=desc]
 * @return {*}  {Sorter}
 */
export function useSorter(
  fields: SortFields<string>,
  direction: Ref<string>,
  field: Ref<string>,
  defaultField = null as unknown as string,
  defaultDirection = desc
): Sorter {
  if (!defaultField) {
    defaultField = Object.keys(fields)[0];
  }

  return {
    toggle: (newField: string) => {
      if (newField === field.value) {
        direction.value = direction.value === desc ? asc : desc;
      } else {
        field.value = newField;
        // https://github.com/vueuse/vueuse/issues/901
        setTimeout(() => {
          nextTick(() => {
            direction.value = defaultDirection;
          });
        }, 1);
      }
    },
    fields: fields,
    refs: {
      direction,
      field,
    },
    filter: {
      order: computed(() => {
        return (
          directions.includes(direction.value)
            ? direction.value
            : defaultDirection
        ).toUpperCase();
      }).value,
      orderBy: computed(() => {
        return (
          Object.keys(fields).includes(field.value!)
            ? field.value!
            : defaultField!
        ).toUpperCase();
      }).value,
    },
  };
}
