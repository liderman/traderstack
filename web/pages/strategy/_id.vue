<template>
  <v-container>
    <v-row>
      <v-col cols="6">
        <v-card
          class="mx-auto"
          tile
        >
          <v-card-title class="text-h5">
            Стратегия
          </v-card-title>

          <v-card-text v-if="!loading">
            <SelectAccount v-model="accountId" :sandbox="false" />
            <SelectStack v-model="stackId"/>

            <v-text-field
              label="Периодичность запуска"
              v-model="runIntervalDuration"
              hint="Секунд"
              reqired
              type="number"
            ></v-text-field>

            <v-switch
              v-model="enabled"
              :label="`Стратегия: ${enabled ? 'Включена' : 'Выключена'}`"
            ></v-switch>
          </v-card-text>

          <v-card-actions>
            <v-btn
              color="primary"
              text
              @click="update"
            >
              Сохранить
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col cols="6">
        <v-card
          class="mx-auto"
          tile
        >
          <v-card-title class="text-h5">
            Логи
          </v-card-title>

          <v-data-table
            dense
            :headers="logsHeaders"
            :items="logs"
            class="elevation-1"
          >
            <template v-slot:[`item.time`]="{ item }">
              {{ item.time | datetime }}
            </template>
          </v-data-table>

          <v-card-actions>
            <v-btn
              color="primary"
              text
              @click="reloadLogs"
            >
              Обновить
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import {Duration} from 'google-protobuf/google/protobuf/duration_pb';
import {GetRequest, UpdateRequest, GetLogsRequest} from "@/grpc/gen/js/liderman/traderstack/strategy/v1/strategy_api_pb";
import {Strategy} from "@/grpc/gen/js/liderman/traderstack/strategy/v1/strategy_pb";
import {clientStrategy} from "@/grpc/instance";
import SelectAccount from "@/components/selects/SelectAccount";
import SelectStack from "@/components/selects/SelectStack";

export default {
  name: 'StrategyPage',
  components: {SelectStack, SelectAccount},
  head: {
    title: 'Стратегия',
  },
  data: () => {
    return {
      id: "",
      accountId: "",
      stackId: "",
      enabled: false,
      runIntervalDuration: 0,
      logs: [],
      loading: false,
      logsHeaders: [
        {
          text: 'Тип',
          align: 'start',
          sortable: false,
          value: 'logType',
        },
        { text: 'Время', value: 'time' },
        { text: 'Сообщение', value: 'message', sortable: false },
      ]
    }
  },
  mounted() {
    this.id = this.$route.params.id;
    this.reload();
  },
  methods: {
    reload() {
      const req = new GetRequest();
      req.setId(this.id);

      this.loading = true;
      clientStrategy.get(req, {})
        .then(response => {
          const strategy = response.toObject().strategy;
          console.log('strategy', strategy);
          this.accountId = strategy.accountId;
          this.stackId = strategy.stackId;
          this.enabled = strategy.enabled;
          this.runIntervalDuration = strategy.runIntervalDuration.seconds;
        })
        .catch(err => {
          console.log(`Ошибка получения стратегии: ${err.message}`);
        })
        .finally(() => {
          this.loading = false;
        })
    },
    update() {
      const strategy = new Strategy();
      strategy.setId(this.id);
      strategy.setAccountId(this.accountId);
      strategy.setStackId(this.stackId);
      strategy.setEnabled(this.enabled);

      const interval = new Duration();
      interval.setSeconds(this.runIntervalDuration);
      strategy.setRunIntervalDuration(interval);

      const req = new UpdateRequest();
      req.setStrategy(strategy);

      clientStrategy.update(req, {})
        .then(response => {
          this.reload();
        })
        .catch(err => {
          console.log(`Ошибка обновления: ${err.message}`);
        })
    },
    reloadLogs() {
      const req = new GetLogsRequest();
      req.setId(this.id);

      clientStrategy.getLogs(req, {})
        .then(response => {
          this.logs = response.toObject().logsList;
          console.log(response.toObject());
        })
        .catch(err => {
          console.log(`Ошибка получения логов: ${err.message}`);
        })
    },
  }
}
</script>
