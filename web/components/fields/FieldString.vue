<template>
  <div
    class="field-string"
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
      &quot;{{ value }}&quot;
    </span>
  </div>
</template>

<script>
export default {
  name: 'FiledString',
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

    if (this.config.input && this.config.input.string) {
      this.value = this.config.input.string;
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
      if (this.config.input && this.config.input.string === val) {
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
