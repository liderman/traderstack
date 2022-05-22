<template>
  <v-dialog
    :value="show"
    persistent
    max-width="600px"
  >
    <v-card>
      <v-card-title>
        <span class="text-h5">Создать стек</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-alert
              text
              type="info"
            >
              Стек — это список список определённых вами функций, которые будут выполнены сверху вниз,
              что позволяет создать собственную стратегию.
            </v-alert>
            <v-text-field
              label="Название стека"
              required
              v-model="name"
            ></v-text-field>
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
          :loading="loading"
          :disabled="name === '' || loading"
          @click="create"
        >
          Создать...
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>

import {clientStack} from '../grpc/instance';
import {CreateRequest} from '../grpc/gen/js/liderman/traderstack/stack/v1/stack_api_pb';

export default {
  name: 'DlgCreateStack',
  model: {
    prop: 'show',
    event: 'change',
  },
  props: {
    show: Boolean,
  },
  data() {
    return {
      name: "",
      error: "",
      loading: false,
    }
  },
  methods: {
    create() {
      const req = new CreateRequest();
      req.setName(this.name);

      this.loading = true;
      this.error = "";
      clientStack.create(req, {})
        .then(response => {
          const stack = response.toObject().stack;
          this.$emit('change', false);
          this.$router.push({name: 'stack-id', params:{id: stack.id}});
        })
        .catch(err => {
          this.error = `Ошибка создания: ${err.message}`;
        })
        .finally(() => {
          this.loading = false;
        })
    }
  }
}
</script>
