<template>
  <v-select
    :value="value"
    :items="funcList"
    label="Функция"
    item-text="name"
    item-value="name"
    single-line
    return-object
    :loading="loading"
    @change="val => $emit('change', val)"
  ></v-select>
</template>

<script>
import {clientStack} from "~/grpc/instance";
import {FuncListRequest} from '~/grpc/gen/js/liderman/traderstack/stack/v1/stack_api_pb';

export default {
  name: 'SelectFunc',
  model: {
    prop: 'value',
    event: 'change',
  },
  props: {
    value: {
      type: [Object, null],
      required: false,
      default: null,
    },
  },
  data() {
    return {
      loading: false,
      funcList: [],
    }
  },
  mounted() {
    this.reload();
  },
  methods: {
    reload() {
      this.loading = true;
      clientStack.funcList(new FuncListRequest(), {})
        .then(response => {
          this.funcList = response.toObject().funcList;
        })
        .catch(err => {
          console.log(`Ошибка получения списка функций: ${err.message}`);
        })
        .finally(() => {
          this.loading = false;
        })
    },
  },
}
</script>
