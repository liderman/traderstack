<template>
  <div class="field">
    <div
      :class="{'field-edit': isEdit, 'field-show': !isEdit}"
    >
      <ExFieldFigiSelect
        v-if="config.input !== undefined && config.extendedType === 'figi-select'"
        :is-edit="isEdit"
        :config="config"
        @change="(val) => $emit('change', config.id, 'value', val)"
        @switch="() => $emit('change', config.id, 'variable', undefined)"
      />
      <FieldsFieldString
        v-else-if="config.input !== undefined && config.baseType === 'string'"
        :is-edit="isEdit"
        :config="config"
        @change="(val) => $emit('change', config.id, 'value', val)"
        @switch="() => $emit('change', config.id, 'variable', undefined)"
      />
      <FieldsFieldInteger
        v-else-if="config.input !== undefined && config.baseType === 'integer'"
        :is-edit="isEdit"
        :config="config"
        @change="(val) => $emit('change', config.id, 'value', val)"
        @switch="() => $emit('change', config.id, 'variable', undefined)"
      />
      <FieldsFieldDecimal
        v-else-if="config.input !== undefined && config.baseType === 'decimal'"
        :is-edit="isEdit"
        :config="config"
        @change="(val) => $emit('change', config.id, 'value', val)"
        @switch="() => $emit('change', config.id, 'variable', undefined)"
      />
      <FieldsFieldBoolean
        v-else-if="config.input !== undefined && config.baseType === 'boolean'"
        :is-edit="isEdit"
        :config="config"
        @change="(val) => $emit('change', config.id, 'value', val)"
        @switch="() => $emit('change', config.id, 'variable', undefined)"
      />
      <FieldsFieldVariable
        v-else-if="config.variable !== undefined"
        :is-edit="isEdit"
        :config="config"
        :var-list="varList"
        @change="(val) => $emit('change', config.id, 'variable', val)"
        @switch="() => $emit('change', config.id, 'value', undefined)"
      />
      <span
        v-else
      >
        <FieldsFieldVariable
          :is-edit="isEdit"
          :config="config"
          :var-list="varList"
          @change="(val) => $emit('change', config.id, 'variable', val)"
        />
      </span>
    </div>
  </div>
</template>

<script>
import ExFieldFigiSelect from "@/components/fields/ExFieldFigiSelect";
export default {
  name: 'AnyField',
  components: {ExFieldFigiSelect},
  props: {
    isEdit: Boolean,
    config: Object,
    varList: Array,
  },
}
</script>

<style>
.field, .field-edit, .field-show {
  display: inline-block;
}

.field-edit {
  width: 150px;
}

.field-show {
  padding: 2px;
  background-color: #404040;
}
</style>
