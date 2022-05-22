<template>
  <v-container>
    <v-row>
      <v-col cols="4">
        <v-card
          class="mx-auto"
          max-width="400"
          tile
        >
          <v-card-title class="text-h5">
            Счета
          </v-card-title>

          <v-list-item
            v-for="(item) in accounts"
            :key="item.id"
            two-line
            @click="loadPositions(item.id)"
          >
            <v-list-item-content>
              <v-list-item-title>{{ item.name ? item.name : 'Без имени (' + item.id.slice(-12) + ')' }}</v-list-item-title>
              <v-list-item-subtitle>{{ item.id }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>

          <v-card-actions>
            <v-btn
              color="primary"
              text
              @click="createAccount"
            >
              Создать счёт
            </v-btn>
          </v-card-actions>
        </v-card>

      </v-col>
      <v-col cols="4">
        <v-card
          v-if="accountId"
          class="mx-auto"
          max-width="400"
          tile
        >
          <v-card-title class="text-h5">
            Баланс
          </v-card-title>

          <v-list-item
            v-for="(item, index) in positions"
            :key="index"
            two-line
          >
            <v-list-item-content>
              <v-list-item-title>{{ item.value }}</v-list-item-title>
              <v-list-item-subtitle>{{ item.currency }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>

          <v-card-actions>
            <v-btn
              color="primary"
              text
              @click="payIn('1000')"
            >
              Пополнить на 1000
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col cols="4">

      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import {SandboxAccountsRequest, OpenSandboxAccountRequest, GetSandboxPositionsRequest, SandboxPayInRequest} from "@/grpc/gen/js/liderman/traderstack/info/v1/info_api_pb";
import {clientInfo} from "@/grpc/instance";

export default {
  name: 'SandboxPage',
  head: {
    title: 'Песочница',
  },
  data: () => {
    return {
      accountId: "",
      accounts: [],
      positions: [],
    }
  },
  mounted() {
    this.reloadAccounts();
  },
  methods: {
    createAccount() {
      const req = new OpenSandboxAccountRequest();

      clientInfo.openSandboxAccount(req, {})
        .then(() => {
          this.reloadAccounts();
        })
        .catch(err => {
          console.log(`Ошибка получения создания аккаунта: ${err.message}`);
        })
    },
    reloadAccounts() {
      const req = new SandboxAccountsRequest();

      clientInfo.sandboxAccounts(req, {})
        .then(response => {
          this.accounts = response.toObject().accountsList;
        })
        .catch(err => {
          console.log(`Ошибка получения списка аккаунтов: ${err.message}`);
        })
    },
    loadPositions(id) {
      this.accountId = id;
      const req = new GetSandboxPositionsRequest();
      req.setAccountId(id);
      clientInfo.getSandboxPositions(req, {})
        .then(response => {
          this.positions = response.toObject().moneyList;
        })
        .catch(err => {
          console.log(`Ошибка получения списка позиций: ${err.message}`);
        })
    },
    payIn(amount) {
      const req = new SandboxPayInRequest();
      req.setAccountId(this.accountId);
      req.setAmount(amount)
      clientInfo.sandboxPayIn(req, {})
        .then(() => {
          this.loadPositions(this.accountId);
        })
        .catch(err => {
          console.log(`Ошибка пополнения: ${err.message}`);
        })
    },
  }
}
</script>
