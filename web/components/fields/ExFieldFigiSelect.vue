<template>
  <div
    class="field-string"
  >
    <v-autocomplete
      v-if="isEdit"
      v-model="value"
      :loading="loading"
      :items="items"
      :filter="filter"
      :search-input.sync="search"
      flat
      hide-no-data
      hide-details
      required
      :label="config.name"
      item-text="ticker"
      item-value="figi"
      solo-inverted
      append-outer-icon="code"
      @click:append-outer="$emit('switch')"
    >
      <template v-slot:item="{ item }">
        <v-list-item-content>
          <v-list-item-title v-text="item.ticker"></v-list-item-title>
          <v-list-item-subtitle v-text="item.name"></v-list-item-subtitle>
          <v-list-item-subtitle v-text="item.figi"></v-list-item-subtitle>
        </v-list-item-content>
      </template>
    </v-autocomplete>

    <span v-else>
      {{ value }}
    </span>
  </div>
</template>

<script>
import {clientInfo} from "~/grpc/instance";
import {SearchInstrumentRequest} from '~/grpc/gen/js/liderman/traderstack/info/v1/info_api_pb';

export default {
  name: 'ExFieldFigiSelect',
  props: {
    isEdit: Boolean,
    config: Object,
  },
  mounted() {
    if (this.config.input && this.config.input.string) {
      this.querySelections(this.config.input.string);
      this.value = this.config.input.string;
    }
  },
  data() {
    return {
      value: null,
      items: [],
      loading: false,
      search: null,
    }
  },
  methods: {
    filter(item, queryText, itemText) {
      return true;
    },
    querySelections(v) {
      const req = new SearchInstrumentRequest();
      req.setTicker(v);
      this.loading = true;
      clientInfo.searchInstrument(req, {})
        .then(response => {
          this.items = response.toObject().instrumentsList;
        })
        .catch(err => {
          console.log(`Ошибка получения списка функций: ${err.message}`);
        })
        .finally(() => {
          this.loading = false;
        })
    }
  },
  watch: {
    value(val) {
      if (this.config.input && this.config.input.string === val) {
        return;
      }

      this.$emit('change', val);
    },
    search (val) {
      console.log(val);
      val && val !== this.select && this.querySelections(val)
    },
  },
}
</script>

<style>
.field-string {
  display: inline-block;
}
</style>
