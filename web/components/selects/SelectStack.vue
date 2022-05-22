<template>
  <v-select
    v-model="value"
    :items="items"
    :loading="loading"
    item-text="name"
    item-value="id"
    label="Стек"
    required
    single-line
  >
    <template v-slot:item="{ item }">
      <v-list-item-content>
        <v-list-item-title>{{ item.name }}</v-list-item-title>
        <v-list-item-subtitle>
          Кол-во позиций: {{ item.itemsList.length }}
        </v-list-item-subtitle>
      </v-list-item-content>
    </template>
  </v-select>
</template>

<script>
import {clientStack} from "~/grpc/instance";
import {GetAllRequest} from '~/grpc/gen/js/liderman/traderstack/stack/v1/stack_api_pb';

export default {
  name: 'SelectStack',
  model: {
    prop: 'input',
    event: 'change',
  },
  props: {
    input: String,
  },
  mounted() {
    if (this.input) {
      this.value = this.input;
    }

    this.reload();
  },
  data() {
    return {
      value: null,
      items: [],
      loading: false,
    }
  },
  methods: {
    reload() {
      const req = new GetAllRequest();

      this.loading = true;
      clientStack.getAll(req, {})
        .then(response => {
          this.items = response.toObject().stacksList;
          console.log('Stacks', response.toObject());
        })
        .catch(err => {
          console.log(`Ошибка получения списка стеков: ${err.message}`);
        })
        .finally(() => {
          this.loading = false;
        })
    }
  },
  watch: {
    value(val) {
      if (this.input === val) {
        return;
      }

      this.$emit('change', val);
    },
  },
}
</script>
