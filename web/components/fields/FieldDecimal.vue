<template>
  <div
    class="field-decimal"
  >
    <v-text-field
      v-if="isEdit"
      :label="config.name"
      v-model="value"
      :rules="rules"
      :hint="config.desc"
      :reqired="config.required"
      append-outer-icon="code"
      @click:append-outer="$emit('switch')"
    ></v-text-field>

    <span v-else>
      {{ value }}
    </span>
  </div>
</template>

<script>

import Decimal from 'decimal.js'

export default {
  name: 'FiledDecimal',
  props: {
    isEdit: Boolean,
    config: Object,
  },
  mounted() {
    if (this.config.required) {
      this.rules.push(
        v => !!v || 'Обязательно'
      );
    }

    if (this.config.input && this.config.input.decimal) {
      const val = new Decimal(this.config.input.decimal);
      this.value = val.toString();
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
      if (this.config.input && this.config.input.decimal === val) {
        return;
      }

      let valDecimal = null;
      try {
        valDecimal = new Decimal(val);
      } catch (e) {
        if ( e instanceof Error && /DecimalError/.test(e.message) ) {
          return;
        }
      }

      this.$emit('change', valDecimal.toString());
    },
  },
}
</script>

<style>
.field-decimal {
  display: inline-block;
}
</style>
