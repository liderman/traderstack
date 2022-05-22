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
            Список стеков
          </v-card-title>

          <v-list-item
            v-for="(stack) in stacks"
            :key="stack.id"
            two-line
            @click="showStack(stack.id)"
          >
            <v-list-item-content>
              <v-list-item-title>{{ stack.name }}</v-list-item-title>
              <v-list-item-subtitle>Кол-во позиций: {{ stack.itemsList.length }}</v-list-item-subtitle>
            </v-list-item-content>

            <v-list-item-action>
              <v-btn icon @click.stop="removeStack(stack.id)">
                <v-icon>delete_outline</v-icon>
              </v-btn>
            </v-list-item-action>
          </v-list-item>

          <v-card-actions>
            <v-btn
              color="primary"
              text
              @click="showCreateStack = true"
            >
              Создать стек
            </v-btn>
          </v-card-actions>
          <dlg-create-stack v-model="showCreateStack"/>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import {GetAllRequest, DeleteRequest} from "@/grpc/gen/js/liderman/traderstack/stack/v1/stack_api_pb";
import {clientStack} from "@/grpc/instance";

export default {
  name: 'StacksPage',
  head: {
    title: 'Стеки',
  },
  data: () => {
    return {
      showCreateStack: false,
      stacks: [],
    }
  },
  mounted() {
    this.reload();
  },
  methods: {
    reload() {
      const req = new GetAllRequest();

      clientStack.getAll(req, {})
        .then(response => {
          this.stacks = response.toObject().stacksList;
        })
        .catch(err => {
          console.log(`Ошибка получения стеков: ${err.message}`);
        })
    },
    showStack(stackId) {
      this.$router.push({name: 'stack-id', params:{id: stackId}});
    },
    removeStack(stackId) {
      if (!confirm('Вы уверены, что хотите удалить стек?')) {
        return;
      }
      const req = new DeleteRequest();
      req.setId(stackId);

      clientStack.delete(req, {})
        .then(() => {
          this.reload();
        })
        .catch(err => {
          console.log(`Ошибка удаления стека: ${err.message}`);
        })
    }
  }
}
</script>
