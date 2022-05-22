<template>
  <div class="stack-func" v-if="config">
      <span :class="{'stack-func-name': isEdit}">
        {{ config.name }} (
      </span>

      <span
        v-for="(param, index) in config.argumentsList"
        :key="param.id"
      >
        <span :class="{'stack-func-param': !isEdit}">
          <AnyField
            :config="param"
            :is-edit="isEdit"
            :var-list="varList"
            @change="(argumentId, varType, val) => changeValue(argumentId, varType, val)"
          />
        </span>
        <span
          v-if="index !== config.argumentsList.length - 1"
          :class="{'stack-func-name': isEdit}"
        >
          ,
        </span>
      </span>

      <span :class="{'stack-func-name': isEdit}">
        )
      </span>
  </div>
</template>

<script>
import AnyField from "~/components/fields/AnyField";
export default {
  name: 'StackFunc',
  components: {AnyField},
  props: {
    isEdit: Boolean,
    config: Object,
    varList: Array,
  },
  data() {
    return {
      varSelect: null,
    }
  },
  mounted() {
    console.log("SFC", this.config);
  },
  methods: {
    select(line) {
      this.selected = line.name;
    },
    changeValue(param, type, val) {
      this.$emit('change', param, type, val)
    }
  }
}
</script>

<style>
.stack-func {
  display: inline-block;
  background-color: darkslategray;
  border: #1e3333 solid 1px;
  border-radius: 3px;
  padding: 5px;
}

.stack-func-name {
  font-size: 18px;
}

.stack-func-param {
  padding: 2px;
  background-color: #404040;
}
</style>
