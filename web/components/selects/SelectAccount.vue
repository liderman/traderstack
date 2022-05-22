<template>
  <v-select
    v-model="value"
    :items="items"
    :loading="loading"
    item-text="name"
    item-value="id"
    label="Аккаунт"
    required
    single-line
  >
    <template v-slot:item="{ item }">
      <v-list-item-content>
        <v-list-item-title>{{item.name ? item.name : 'Без имени (' + item.id.slice(-12) + ')'}}</v-list-item-title>
        <v-list-item-subtitle v-text="item.id"></v-list-item-subtitle>
      </v-list-item-content>
    </template>
    <template v-slot:selection="{ item }">
      {{item.name ? item.name : 'Без имени (' + item.id.slice(-12) + ')'}}
    </template>
  </v-select>
</template>

<script>
import {clientInfo} from "~/grpc/instance";
import {SandboxAccountsRequest, AccountsRequest} from '~/grpc/gen/js/liderman/traderstack/info/v1/info_api_pb';

export default {
  name: 'SelectAccount',
  model: {
    prop: 'input',
    event: 'change',
  },
  props: {
    input: String,
    sandbox: Boolean,
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
      let fRes = null;
      if (this.sandbox) {
        fRes = clientInfo.sandboxAccounts(new SandboxAccountsRequest(), {});
      } else {
        fRes = clientInfo.accounts(new AccountsRequest(), {});
      }

      this.loading = true;
      fRes
        .then(response => {
          this.items = response.toObject().accountsList;
          console.log('ACC', response.toObject());
        })
        .catch(err => {
          console.log(`Ошибка получения списка аккаунтов: ${err.message}`);
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
