<template>
  <div class="field-integer">
    <v-text-field
      v-if="isEdit"
      :label="config.name"
      v-model="value"
      :rules="rules"
      :hint="config.desc"
      :reqired="config.required"
      type="number"
      append-outer-icon="code"
      @click:append-outer="$emit('switch')"
    ></v-text-field>

    <span v-else>
      {{ value }}
    </span>
  </div>
</template>

<script>
export default {
  name: 'FieldInteger',
  props: {
    isEdit: Boolean,
    config: Object,
  },
  data() {
    return {
      value: 0,
      rules: [],
    }
  },
  mounted() {
    if (this.config.required) {
      this.rules.push(
        v => !!v || 'Обязательно'
      );
    }

    if (this.config.input && this.config.input.integer) {
      this.value = this.config.input.integer;
    }
  },
  watch: {
    value(val) {
      if (this.config.input && this.config.input.integer === val) {
        return;
      }

      this.$emit('change', val);
    },
  },
}
</script>

<style>
.field-integer {
  display: inline-block;
}
</style>
