<template>
  <v-container>
    <v-row class="mb-6">
      <v-col
        v-for="item in menu"
        :key="item.title"
        :md="4"
        :sm="6"
      >
        <v-hover>
          <template #default="{ hover }">
            <v-card
              :to="{ path: item.to }"
              class="pa-2"
              :elevation="hover ? 12 : 2"
              tile
              height="160px"
            >
              <v-card-title>{{ item.title }}</v-card-title>
              <v-card-text>{{ item.text }}</v-card-text>
            </v-card>
          </template>
        </v-hover>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'IndexPage',
  head: {
    title: 'Старт',
  },
  data: () => {
    return {
      menu: [
        {
          icon: 'view_stream',
          title: 'Стеки',
          text: 'Позволяет задать алгоритм торговой стратегии в виде стека функций',
          to: '/stacks',
        },
        {
          icon: 'insights',
          title: 'Стратегии',
          text: 'Запускайте свои созданные стеки с нужной периодичностью и отслеживайте активность в логе',
          to: '/strategies',
        },
        {
          icon: 'manage_accounts',
          title: 'Песочница',
          text: 'Управление счетами и балансом песочницы',
          to: '/sandbox',
        },
      ],
    }
  }
}
</script>
