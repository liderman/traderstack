<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card
          class="mx-auto"
          max-width="800"
          tile
        >
          <v-card-title class="text-h5">
            Список стратегий
          </v-card-title>

          <v-list-item
            v-for="(strategy) in strategies"
            :key="strategy.id"
            two-line
            @click="showStrategy(strategy.id)"
          >
            <v-list-item-content>
              <v-list-item-title>{{ strategy.id }}</v-list-item-title>
              <v-list-item-subtitle>
                Состояние: {{ strategy.enabled ? 'Включена' : 'Выключена' }}},
                Период запуска: {{ strategy.runIntervalDuration.seconds }} сек.
              </v-list-item-subtitle>
            </v-list-item-content>

            <v-list-item-action>
              <v-btn icon @click.stop="removeStrategy(strategy.id)">
                <v-icon>delete_outline</v-icon>
              </v-btn>
            </v-list-item-action>
          </v-list-item>

          <v-card-actions>
            <v-btn
              color="primary"
              text
              @click="createStrategy()"
            >
              Создать стратегию
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import {GetAllRequest, DeleteRequest, CreateRequest} from "@/grpc/gen/js/liderman/traderstack/strategy/v1/strategy_api_pb";
import {clientStrategy} from "@/grpc/instance";

export default {
  name: 'StrategiesPage',
  head: {
    title: 'Стратегии',
  },
  data: () => {
    return {
      showCreateStack: false,
      strategies: [],
    }
  },
  mounted() {
    this.reload();
  },
  methods: {
    reload() {
      const req = new GetAllRequest();

      clientStrategy.getAll(req, {})
        .then(response => {
          this.strategies = response.toObject().strategiesList;
        })
        .catch(err => {
          console.log(`Ошибка получения стратегий: ${err.message}`);
        })
    },
    createStrategy() {
      const req = new CreateRequest();

      clientStrategy.create(req, {})
        .then(response => {
          const strategy = response.toObject().strategy;
          this.showStrategy(strategy.id);
        })
        .catch(err => {
          console.log(`Ошибка создания стратеги: ${err.message}`);
        })
    },
    showStrategy(sid) {
      this.$router.push({name: 'strategy-id', params:{id: sid}});
    },
    removeStrategy(id) {
      if (!confirm('Вы уверены, что хотите удалить стратегию?')) {
        return;
      }
      const req = new DeleteRequest();
      req.setId(id);

      clientStrategy.delete(req, {})
        .then(() => {
          this.reload();
        })
        .catch(err => {
          console.log(`Ошибка удаления стратегии: ${err.message}`);
        })
    }
  }
}
</script>
