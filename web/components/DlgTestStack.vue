<template>
  <v-dialog
    :value="show"
    persistent
    max-width="600px"
  >
    <v-card>
      <v-card-title>
        <span class="text-h5">Протестировать стек</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-alert
              text
              type="info"
            >
              Выполняет одну итерацию запуска в песочнице эмулируя заявку в указанное время.
              В режиме тестирования работа биржи определяется на основании наличия свечи в указанное время (если свеча
              есть - значит биржа работала)
            </v-alert>
            <v-text-field
              label="Дата и время"
              required
              v-model="dateTime"
            ></v-text-field>
            <SelectAccount v-model="account" :sandbox="true" />
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
          color="blue darken-1"
          text
          @click="$emit('change', false)"
        >
          Отмена
        </v-btn>
        <v-btn
          color="blue darken-1"
          text
          :disabled="account === ''"
          @click="start"
        >
          Старт
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>

import format from 'date-fns/format';
import parse from 'date-fns/parse';
import SelectAccount from "@/components/selects/SelectAccount";

export default {
  name: 'DlgTestStack',
  components: {SelectAccount},
  model: {
    prop: 'show',
    event: 'change',
  },
  props: {
    show: Boolean,
  },
  data() {
    return {
      dateTime: "",
      account: "",
      error: "",
      loading: false,
    }
  },
  mounted() {
    this.dateTime = format(new Date(), "yyyy.MM.dd HH:mm:ss");
  },
  methods: {
    start() {
      const dateTime = parse(this.dateTime, "yyyy.MM.dd HH:mm:ss", new Date());
      this.$emit('start', this.account, dateTime);
      this.$emit('change', false);
    }
  }
}
</script>
