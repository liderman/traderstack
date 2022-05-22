<template>
  <div>
    <div
      class="stack-line"
      :class="{'stack-line-active': selected}"
      @click="select()"
    >
      <div class="stack-line-name">
        {{ !selected ? '$' + name + ' :' : ''}}
        <v-text-field
          v-if="selected"
          label="Имя"
          v-model="name"
          prefix="$"
          suffix=" :"
          @change="$emit('changeVarName', name)"
        ></v-text-field>
      </div>
      <div class="stack-line-data">
        <div class="stack-line-data-func">
          <StackFunc
            v-if="stackLine.stackFunc"
            :is-edit="selected"
            :config="stackLine.stackFunc"
            :var-list="currentVars"
            @change="(param, type, val) => $emit('change', param, type, val)"
          />
          <select-func
            v-model="funcSelect"
            v-else
          />
        </div>

        <div
          v-if="testItem"
          class="stack-line-data-result"
          :class="{'stack-line-data-result-error': testItem.error !== ''}"
        >
          {{ testItem.error === '' ? getValue(testItem) : testItem.error }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'StackLine',
  props: {
    stackLine: Object,
    testItem: Object,
    currentVars: Array,
    selected: Boolean,
  },
  data() {
    return {
      line: null,
      funcSelect: null,
      name: '',
      funcConfig: null,
    }
  },
  watch: {
    funcSelect() {
      if (this.funcSelect) {
        this.$emit('setFunc', this.funcSelect);
      }
    }
  },
  mounted() {
    this.name = this.stackLine.variable;
    this.line = this.stackLine;
  },
  methods: {
    select() {
      this.$emit('selected');
    },
    getValue(testItem) {
      switch (testItem.baseType) {
        case "string":
          return `"${testItem.result.string}"`;
        case "integer":
          return testItem.result.integer;
        case "decimal":
          return testItem.result.decimal;
        case "boolean":
          return testItem.result.pb_boolean ? 'True' : 'False';
      }
      return '';
    }
  }
}
</script>

<style>
.stack-line {
  min-height: 32px;
  padding: 8px;
  margin-bottom: 10px;
  border: transparent 1px solid;
  border-left-width: 8px;
  display: flex;
  flex-direction: row;
  align-items: stretch;
}

.stack-line:hover {
  border-color: #404040;
}

.stack-line-active {
  border-color: aquamarine;
}

.stack-line-active:hover {
  border-color: #69d3ae;
}

.stack-line-name {
  width: 10%;
  text-align: right;
  margin-right: 20px;
}

.stack-line-data {
  width: 90%;
}

.stack-line-data-func {
  border: #4d4d4d 1px solid;
  background-color: #232323;
  padding: 4px;
}

.stack-line-data-result {
  padding-top: 10px;
}

.stack-line-data-result-error {
  color: #f83f64;
}
</style>
