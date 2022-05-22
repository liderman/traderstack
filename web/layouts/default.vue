<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="false"
      :clipped="false"
      fixed
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar fixed app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-btn icon @click="$router.push({name: 'sandbox'})">
        <v-icon>manage_accounts</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <Nuxt />
    </v-main>
    <v-footer app>
      <span>TraderStack &copy; {{ new Date().getFullYear() }}. Используя данное приложение вы принимаете все риски на себя.</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: 'DefaultLayout',
  data() {
    return {
      title: '',
      drawer: false,
      items: [
        {
          icon: 'home',
          title: 'Главная',
          to: '/',
        },
        {
          icon: 'view_stream',
          title: 'Стеки',
          to: '/stacks',
        },
        {
          icon: 'insights',
          title: 'Стратегии',
          to: '/strategies',
        },
        {
          icon: 'manage_accounts',
          title: 'Песочница',
          to: '/sandbox',
        },
      ],
    }
  },
  head() {
    return {
      changed: (info) =>  {
        this.title = info.title;
      },
    };
  },
}
</script>
