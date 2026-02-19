<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'UIButton',
  props: {
    name: {
      type: String,
      default: ''
    },
    fontSize: {
      type: String as () => 'small' | 'medium' | 'large',
      default: 'medium'
    },
    isLoading: {
      type: Boolean,
      default: false
    },
    isDisabled: {
      type: Boolean,
      default: false
    },
    theme: {
      type: String as () => 'colored' | 'whited',
      default: 'colored'
    }
  }
})
</script>

<template>
  <button
      type="button"
      class="custom-button"
      :class="[
      `button--size_${fontSize}`,
      `${theme}-button`,
      { 'disabled-button': isLoading || isDisabled }
    ]"
      :disabled="isLoading || isDisabled"
  >
    <slot />
    {{ name }}

    <svg
        v-if="isLoading"
        class="loading-icon"
        :class="{
        'loading_icon-whited': theme === 'colored',
        'loading_icon-colored': theme === 'whited'
      }"
        xmlns="http://www.w3.org/2000/svg"
        width="20"
        height="20"
        viewBox="0 0 102 100"
        fill="none"
    >
      <g clip-path="url(#clip0_9487_51525)">
        <circle
            opacity="0.1"
            cx="51.3086"
            cy="50"
            r="42.5"
            stroke-width="15"
            stroke="white"
        />
        <mask id="path-2-inside-1_9487_51525" fill="white">
          <path d="M16.7382 25.2385C13.3813 22.8341 12.5734 18.1151 15.4511 15.1538C22.356 8.04854 31.2615 3.11172 41.069 1.05973C53.331 -1.50581 66.1093 0.625125 76.8752 7.03085C87.6412 13.4366 95.6102 23.6503 99.2054 35.6509C102.081 45.2492 101.991 55.4312 99.0404 64.8894C97.8107 68.8312 93.2779 70.3725 89.5632 68.5694C85.8485 66.7662 84.3871 62.2945 85.3628 58.2822C86.8231 52.2773 86.6801 45.9468 84.8812 39.9422C82.3612 31.5306 76.7754 24.3714 69.2292 19.8814C61.6829 15.3914 52.7261 13.8977 44.1313 15.696C37.9958 16.9797 32.3639 19.8741 27.7832 24.0223C24.7224 26.794 20.0952 27.643 16.7382 25.2385Z" />
        </mask>
        <path
            d="M16.7382 25.2385C13.3813 22.8341 12.5734 18.1151 15.4511 15.1538C22.356 8.04854 31.2615 3.11172 41.069 1.05973C53.331 -1.50581 66.1093 0.625125 76.8752 7.03085C87.6412 13.4366 95.6102 23.6503 99.2054 35.6509C102.081 45.2492 101.991 55.4312 99.0404 64.8894C97.8107 68.8312 93.2779 70.3725 89.5632 68.5694C85.8485 66.7662 84.3871 62.2945 85.3628 58.2822C86.8231 52.2773 86.6801 45.9468 84.8812 39.9422C82.3612 31.5306 76.7754 24.3714 69.2292 19.8814C61.6829 15.3914 52.7261 13.8977 44.1313 15.696C37.9958 16.9797 32.3639 19.8741 27.7832 24.0223C24.7224 26.794 20.0952 27.643 16.7382 25.2385Z"
            stroke-width="30"
            mask="url(#path-2-inside-1_9487_51525)"
        />
      </g>
      <defs>
        <clipPath id="clip0_9487_51525">
          <rect width="101" height="100" fill="white" transform="translate(0.5)" />
        </clipPath>
      </defs>
    </svg>
  </button>
</template>

<style scoped lang="scss">
.custom-button {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  padding: 15px 24px;
  border-radius: 8px;
  cursor: pointer;
  width: 100%;
  border: 0;
  max-height: 60px;
  font: 600 normal 20px/150% Regular;
  transition: opacity 0.3s ease;

  &:hover:not(.disabled-button) {
    opacity: 0.9;
  }
}

.button--size {
  &_small {
    font: 100 normal 16px/100% Regular;
    padding: 10px 20px;
    max-height: 40px;
  }

  &_medium {
    font: 600 normal 20px/150% Regular;
    padding: 15px 24px;
    max-height: 50px;
  }

  &_large {
    font: 700 normal 24px/150% Regular;
    padding: 20px 32px;
    max-height: 60px;
  }
}

.colored-button {
  color: white;
  background-color: #1073fe;
}

.whited-button {
  color: black;
  border: 1px solid rgba(0, 0, 0, .16);
  background-color: white;
}

.disabled-button {
  opacity: 0.8;
  cursor: not-allowed;
}

.loading-icon {
  width: 20px;
  height: 20px;
  margin-top: 1px;
  animation: rotate 1.5s linear infinite;
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.loading_icon-whited {
  g {
    path {
      stroke: white;
    }
    circle {
      stroke: white;
    }
  }
}

.loading_icon-colored {
  g {
    path {
      stroke: #1073fe;
    }
    circle {
      stroke: #1073fe;
    }
  }
}
</style>
