<script lang="ts">
import Popup from '~/components/UI/Popup.vue'
import UIInput from "~/components/UI/UIInput.vue";

type TAuthState = 'authorization' | 'registration' | 'identification';

export default defineComponent({
  name: "Authentification",
  components: {
    Popup,
    UIInput
  },
  data() {
    return {
      authState: 'authorization' as TAuthState,
      userData: {
        name: '',
        email: '',
        phone: '',
        password: '',
        date_of_birth: ''
      },
      userErrors: {
        name: '',
        email: '',
        phone: '',
        password: '',
        date_of_birth: ''
      }
    }
  }
})
</script>

<template>
  <div class="auth">
    <Popup>
      <div v-if="authState == 'authorization'" class="auth-wrapper">
        <div class="auth-header">
          <div class="auth-title">
            Вход
          </div>
          <div class="auth-mode">
            Регистрация
          </div>
        </div>
        <div class="auth-body">
          <UIInput error-text="pizda" title="Эл. почта" v-model="userData.email" placeholder="Введите эл.почту" />
          <UIInput title="Пароль" v-model="userData.password" placeholder="Введите пароль" />
        </div>
        <div class="auth-footer">
          <UIButton name="Войти" theme="colored" />
          <UIButton name="У меня еще нет аккаунта" font-size="small" theme="whited" />
        </div>
      </div>
      <div v-if="authState == 'registration'" class="auth-wrapper">
        <div class="auth-header">
          <div class="auth-title">
            Регистрация
          </div>
          <div class="auth-mode">
            Вход
          </div>
        </div>
        <div class="auth-body">
          <UIInput title="Имя" v-model="userData.name" placeholder="Имя" />
          <UIInput title="Эл. почта" v-model="userData.email" placeholder="Эл. почта" />
          <UIInput title="Номер телефона" v-model="userData.email" placeholder="Номер телефона" />
          <div class="auth-row">
            <UIInput title="Дата рождения" v-model="userData.name" placeholder="ДД" />
            <UIInput title=" " v-model="userData.email" placeholder="ММ" />
            <UIInput title=" " v-model="userData.email" placeholder="ГГГГ" />
          </div>
        </div>
        <div class="auth-footer">
          <UIButton name="Зарегистрироваться" theme="colored" />
          <UIButton name="У меня уже есть аккаунт пользователя" font-size="small" theme="whited" />
        </div>
      </div>
      <div v-if="authState == 'identification'" class="auth-wrapper">
        <div class="auth-row">
          <UIInput v-model="userData.email" placeholder="0" />
          <UIInput v-model="userData.password" placeholder="0" />
          <UIInput v-model="userData.email" placeholder="0" />
          <UIInput v-model="userData.password" placeholder="0" />
        </div>
        <div class="auth-footer">
          <UIButton name="Подтвердить" theme="colored" />
          <UIButton name="Выслать код повторно" font-size="small" theme="whited" />
        </div>
      </div>
    </Popup>
  </div>
</template>

<style scoped lang="scss">
.auth {
  &-wrapper {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  &-header {
    display: flex;
    justify-content: space-between;
  }

  &-title {
    color: black;
  }

  &-mode {
    color: #1073fe;
    cursor: pointer;
  }

  &-body {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  &-row {
    display: flex;
    gap: 8px;
  }

  &-footer {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
}
</style>