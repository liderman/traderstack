<template>
  <div
    class="field-string"
  >
    <v-select
      v-if="isEdit"
      v-model="value"
      :items="varList"
      :label="config.name"
      required
      single-line
      prefix="$"
      append-outer-icon="code"
      @click:append-outer="$emit('switch')"
    ></v-select>
    <span v-else>
      ${{ value }}
    </span>
  </div>
</template>

<script>
export default {
  name: 'FiledVariable',
  props: {
    isEdit: Boolean,
    config: Object,
    varList: Array,
  },
  mounted() {
    if (this.config.variable && this.config.variable.name) {
      this.value = this.config.variable.name;
    }
  },
  data() {
    return {
      value: '',
      rules: [],
    }
  },
  watch: {
    value(val) {
      if (this.config.variable && this.config.variable.name === val) {
        return;
      }

      this.$emit('change', val);
    },
  },
}
</script>

<style>
.field-string {
  display: inline-block;
}
</style>
