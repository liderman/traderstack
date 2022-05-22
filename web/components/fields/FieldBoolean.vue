<template>
  <div
    class="field-boolean"
  >
    <v-select
      v-if="isEdit"
      v-model="value"
      :items="items"
      :label="config.name"
      required
      single-line
      append-outer-icon="code"
      @click:append-outer="$emit('switch')"
    ></v-select>
    <span v-else>
      {{ value }}
    </span>
  </div>
</template>

<script>
export default {
  name: 'FieldBoolean',
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

    if (this.config.input) {
      this.value = this.config.input.pb_boolean ? 'True' : 'False';
    }
  },
  data() {
    return {
      value: '',
      rules: [],
      items: ['True', 'False'],
    }
  },
  watch: {
    value(val) {
      const boolVal = val === 'True';
      if (this.config.input && this.config.input.pb_boolean === boolVal) {
        return;
      }

      this.$emit('change', val === 'True');
    },
  },
}
</script>

<style>
.field-boolean {
  display: inline-block;
}
</style>
