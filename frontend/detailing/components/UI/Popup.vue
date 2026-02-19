<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: "Popup"
})
</script>

<template>
  <div class="popup">
    <div class="popup-wrapper">
      <slot />
    </div>
  </div>
</template>

<style scoped lang="scss">
  .popup {
    width: 100%;
    height: 100%;
    position: fixed;
    top: 0;
    left: 0;
    backdrop-filter: blur(16px);
    background: #ffffff52;
    display: flex;
    align-items: center;
    justify-content: center;
    &-wrapper {
      width: 400px;
      height: auto;
      padding: 20px 20px 12px;
      backdrop-filter: blur(8px);
      background: #ffffffa3;
      border-radius: 12px;
      box-sizing: border-box;
      position: relative;
      box-shadow: 0 12px 24px 0 #00345729;
    }
  }
</style>