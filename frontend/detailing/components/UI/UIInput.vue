<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: "UIInput",
  props: {
    modelValue: [String, Number],
    placeholder: String,
    type: {
      type: String,
      default: 'text'
    },
    disabled: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: ''
    },
    errorText: {
      type: String,
      default: ''
    }
  },
  emits: ['update:modelValue']
})
</script>

<template>
  <div class="custom-input">
    <div v-if="title" class="custom-input__title">
      {{ title }}
    </div>
    <input
        class="custom-input__element"
        :class="[
            {
              'custom-input__element-error' : errorText,
            }
        ]"
        :value="modelValue"
        :placeholder="placeholder"
        :type="type"
        :disabled="disabled"
        @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
        v-bind="$attrs"
    />
    <div v-if="errorText" class="custom-input__error">
      {{ errorText }}
    </div>
  </div>
</template>

<style scoped lang="scss">
.custom-input {
  display: flex;
  flex-direction: column;
  justify-content: end;
  gap: 8px;
  &__title {
    color: black;
  }

  &__element {
    padding: 12px;
    border: 1px solid rgba(0, 0, 0, .16);
    border-radius: 8px;
    font-size: 16px;
    background-color: white;
    width: 100%;
    box-sizing: border-box;

    &:focus {
      outline: none;
      border-color: #1073fe;
    }

    &:disabled {
      background-color: grey;
      cursor: not-allowed;
    }

    &-error {
      border: 1px solid #b51c1c;
      &::placeholder {
        color: #b51c1c;
      }
    }
  }

  &__error {
    color: #b51c1c;
  }
}
</style>
