<template>
  <div>
    <v-toolbar dense>
      <v-btn
        text
        @click="add"
      >
        <v-icon>playlist_add</v-icon>
      </v-btn>
      <v-btn
        text
        @click="moveUp"
      >
        <v-icon>move_up</v-icon>
      </v-btn>
      <v-btn
        text
        @click="moveDown"
      >
        <v-icon>move_down</v-icon>
      </v-btn>
      <v-btn
        text
        @click="remove"
      >
        <v-icon>delete_outline</v-icon>
      </v-btn>
      <v-btn
        text
        @click="save"
      >
        Сохранить
      </v-btn>
      <v-btn
        text
        @click="test"
      >
        Тест
      </v-btn>
    </v-toolbar>
    <v-container>
      <v-row>
        <v-col>
          <StackLine
            v-for="(line, index) in stack"
            :key="line.variable"
            :stack-line="line"
            :selected="index === selectedIndex"
            :current-vars="currentVars"
            :test-item="testItems[index] ? testItems[index] : null"
            @selected="selectLine(line, index)"
            @setFunc="setFunc"
            @change="change"
            @changeVarName="changeVarName"
          />
        </v-col>
      </v-row>
    </v-container>
    <DlgTestStack v-model="dlgTest" @start="runTest" />
  </div>
</template>

<script>

import {Timestamp} from 'google-protobuf/google/protobuf/timestamp_pb';
import {clientStack} from '../../grpc/instance';
import {stackToUpdateRequest, setValueByArgument} from '../../grpc/mappers';
import {GetRequest, TestRequest} from '../../grpc/gen/js/liderman/traderstack/stack/v1/stack_api_pb';
import {Item} from '../../grpc/gen/js/liderman/traderstack/stack/v1/stack_pb';

export default {
  name: 'StackPage',
  head: {
    title: 'Стек',
  },
  data() {
    return {
      stackId: "",
      stackName: "",
      selectedIndex: 0,
      currentVars: [],
      stack: [],
      testItems: [],
      dlgTest: false,
    }
  },
  mounted() {
    this.reload(this.$route.params.id);
  },
  methods:{
    test() {
      this.dlgTest = true;
    },
    runTest(account, dateTime) {
      const req = new TestRequest();
      req.setId(this.stackId);
      req.setAccountId(account);

      const timestamp = new Timestamp();
      timestamp.fromDate(dateTime);
      req.setTime(timestamp);

      clientStack.test(req, {})
        .then(response => {
          this.testItems = response.toObject().resultList;
        })
        .catch(err => {
          console.log(`Ошибка тестирования: ${err.message}`);
        })
        .finally(() => {

        })
    },
    reload(id) {
      const req = new GetRequest();
      req.setId(id);
      clientStack.get(req, {})
        .then(response => {
          const stack = response.toObject().stack;
          this.stackId = stack.id;
          this.stackName = stack.name;
          this.stack = stack.itemsList;
          console.log(stack);
        })
        .catch(err => {
          console.log(`Ошибка создания: ${err.message}`);
        })
        .finally(() => {

        })
    },
    save() {
      const req = stackToUpdateRequest(this.stack, this.stackId, this.stackName);

      clientStack.update(req, {})
        .then(response => {
          const stack = response.toObject().stack;
          this.stackId = stack.id;
          this.stackName = stack.name;
          console.log(stack);
        })
        .catch(err => {
          this.error = `Ошибка создания: ${err.message}`;
        })
        .finally(() => {

        })
    },
    changeVarName(name) {
      this.stack[this.selectedIndex].variable = name;
    },
    change(id, varType, val) {
      const args = this.stack[this.selectedIndex].stackFunc.argumentsList;
      let argIdx = -1;
      let arg = null;
      for (let i = 0; i < args.length; i++) {
        if (args[i].id === id) {
          argIdx = i;
          arg = args[i]
          break;
        }
      }

      if (arg === null) {
        console.log('function argument is not found');
        return;
      }

      const set = setValueByArgument(arg, varType, val);
      this.stack[this.selectedIndex].stackFunc.argumentsList[argIdx].variable = set.variable;
      this.stack[this.selectedIndex].stackFunc.argumentsList[argIdx].input = set.input;
      console.log('STACK', this.stack);
    },
    setFunc(stackFunc) {
      console.log(stackFunc);
      this.stack[this.selectedIndex].stackFunc = stackFunc;
    },
    selectLine(line, index) {
      const currentVars = [];
      for (let i = 0; i < index; i++) {
        currentVars.push(this.stack[i].variable);
      }
      this.currentVars = currentVars;

      this.selectedIndex = index;
    },
    add() {
      let maxVarNum = 0;
      this.stack.forEach(v => {
        if (!/^-?\d+$/.test(v.variable)) {
          return;
        }

        const varNum = parseInt(v.variable);
        if (varNum > maxVarNum) {
          maxVarNum = varNum;
        }
      });

      const item = new Item();
      item.setVariable(`${maxVarNum+1}`);

      this.stack.splice(this.selectedIndex+1, 0, item.toObject());

      this.selectedIndex++;
    },
    moveUp() {
      const toIndex = this.selectedIndex - 1;
      if (toIndex < 0) {
        return;
      }

      this.stack = this.move(this.stack, this.selectedIndex, toIndex);
      this.selectedIndex = toIndex;
    },
    moveDown() {
      const toIndex = this.selectedIndex + 1;
      if (toIndex > this.stack.length - 1) {
        return;
      }

      this.stack = this.move(this.stack, this.selectedIndex, toIndex);
      this.selectedIndex = toIndex;
    },
    remove() {
      let newIndex = this.selectedIndex - 1;
      if (newIndex < 0) {
        newIndex = 0;
      }

      this.stack.splice(this.selectedIndex, 1);
      this.selectedIndex = newIndex;
    },
    move(data, fromIndex, toIndex) {
      const el = data[fromIndex];
      data.splice(fromIndex, 1);
      data.splice(toIndex, 0, el);
      return data;
    }
  }
}
</script>
