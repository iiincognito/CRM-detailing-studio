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
      authState: 'registration' as TAuthState,
      userData: {
        name: '',
        email: '',
        phone: '',
        password: '',
        date_of_birth: {
          first: '',
          second: '',
          third: '',
          result: ''
        }
      },
      identificationCode: {
        first: '',
        second: '',
        third: '',
        four: ''
      },
      userErrors: {
        name: '',
        email: '',
        phone: '',
        password: '',
        date_of_birth: {
          text: ''
        }
      },
      identificationCodeError: {
        first: '',
        second: '',
        third: '',
        four: '',
        text: ''
      },
    }
  },
  methods: {
    resetErrors() {
      this.userErrors =  {
        name: '',
            email: '',
            phone: '',
            password: '',
            date_of_birth: {
          text: ''
        }
      }
      this.identificationCodeError = {
        first: '',
            second: '',
            third: '',
            four: '',
            text: ''
      }
    },
    validation(authState: TAuthState) {
      if (authState == 'registration') {
        if (this.userData.name.length === 0) {
          this.userErrors.name = 'Введите имя'
          return false;
        } else if (this.userData.name.length < 2) {
          this.userErrors.name = 'Короткое имя'
          return false;
        }

        if (this.userData.email.length === 0) {
          this.userErrors.email = 'Введите электронную почту'
          return false;
        } else if (!/^[\w\.-]+@[\w\.-]+\.[a-zA-Z]{2,}$/.test(this.userData.email)) {
          this.userErrors.email = 'Неверный формат email';
          return false;
        }

        if (this.userData.phone.length === 0) {
          this.userErrors.phone = 'Введите номер телефона'
          return false;
        } else if (this.userData.phone.length < 11) {
          this.userErrors.phone = 'Введите корректный сотовый телефон'
          return false;
        }

        if (this.userData.password.length === 0) {
          this.userErrors.password = 'Введите пароль'
          return false;
        } else if (this.userData.password.length < 6) {
          this.userErrors.password = 'Данный пароль короткий'
          return false;
        }

        if (this.userData.date_of_birth.first.length === 0 &&
            this.userData.date_of_birth.second.length === 0 &&
            this.userData.date_of_birth.third.length === 0) {
          this.userErrors.date_of_birth.text = 'Введите дату рождения'
          return false;
        } else if (this.userData.date_of_birth.first.length === 0 ||
            this.userData.date_of_birth.second.length === 0 ||
            this.userData.date_of_birth.third.length === 0) {
          this.userErrors.date_of_birth.text = 'Заполните дату рождения до конца'
          return false;
        }
      } else if (authState === 'authorization') {
        if (this.userData.email.length === 0) {
          this.userErrors.email = 'Введите электронную почту'
          return false;
        } else if (!/^[\w\.-]+@[\w\.-]+\.[a-zA-Z]{2,}$/.test(this.userData.email)) {
          this.userErrors.email = 'Неверный формат email';
          return false;
        }

        if (this.userData.password.length === 0) {
          this.userErrors.password = 'Введите пароль'
          return false;
        }
      } else if (authState === 'identification') {
        const code = this.identificationCode;
        const error = this.identificationCodeError;

        const hasEmpty = ['first', 'second', 'third', 'four'].some(field => {
          if (code[field as keyof typeof code].length === 0) {
            error[field as keyof typeof error] = '+';
            return true;
          }
          return false;
        });

        if (hasEmpty) return false;
      }
      return true;
    },
    async submit() {
      if (this.authState === 'registration') {
        const forRequest = {
          FIO: 'pidoras',
          Phone: '89503215476',
          Password: 'almaz_pidr',
          Birthday: '2002-02-28',
          Email: 'aidar-dev@mail.ru'
        }
        const result = await $fetch('https://lot-initial-silver-gourmet.trycloudflare.com/api/v1/auth/register', {
          method: 'POST',
          body: forRequest
        })
        console.log('Успешно отправлено:', result)
        this.authState = 'identification'
      } else if (this.authState === 'authorization') {
        console.log('Здесь будет асинхронный запрос на авторизацию')
        this.authState = 'identification'
      } else if (this.authState === 'identification') {
        console.log('Здесь будет асинхронный запрос на идентификацию')
      }
    },
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
          <div @click="authState = 'registration'" class="auth-mode">
            Регистрация
          </div>
        </div>
        <div class="auth-body">
          <UIInput
              title="Эл. почта"
              :error-text="userErrors.email"
              v-model="userData.email"
              placeholder="Введите эл.почту"
          />
          <UIInput
              title="Пароль"
              :error-text="userErrors.password"
              v-model="userData.password"
              placeholder="Введите пароль"
          />
        </div>
        <div class="auth-footer">
          <UIButton
              name="Войти"
              theme="colored"
              @click="submit"
          />
          <UIButton
              name="У меня еще нет аккаунта"
              font-size="small"
              theme="whited"
              @click="authState = 'registration'"
          />
        </div>
      </div>
      <div v-if="authState == 'registration'" class="auth-wrapper">
        <div class="auth-header">
          <div class="auth-title">
            Регистрация
          </div>
          <div @click="authState = 'authorization'" class="auth-mode">
            Вход
          </div>
        </div>
        <div class="auth-body">
          <UIInput
              title="Имя"
              :error-text="userErrors.name"
              v-model="userData.name"
              placeholder="Имя"
              @onInput="resetErrors"
          />
          <UIInput
              title="Эл. почта"
              :error-text="userErrors.email"
              v-model="userData.email"
              placeholder="Эл. почта"
              @onInput="resetErrors"
          />
          <UIInput
              title="Номер телефона"
              :error-text="userErrors.phone"
              v-model="userData.phone"
              placeholder="Номер телефона"
              @onInput="resetErrors"
          />
          <UIInput
              title="Пароль"
              v-model="userData.password"
              :error-text="userErrors.password"
              placeholder="Пароль"
              @onInput="resetErrors"
          />
          <div class="auth-body">
            <div class="auth-row">
              <UIInput
                  title="Дата рождения"
                  :is-show-error="userErrors.date_of_birth.text.length > 0"
                  v-model="userData.date_of_birth.first"
                  placeholder="ДД"
                  @onInput="resetErrors"
              />
              <UIInput
                  title=" "
                  :is-show-error="userErrors.date_of_birth.text.length > 0"
                  v-model="userData.date_of_birth.second"
                  placeholder="ММ"
                  @onInput="resetErrors"
              />
              <UIInput
                  title=" "
                  :is-show-error="userErrors.date_of_birth.text.length > 0"
                  v-model="userData.date_of_birth.third"
                  placeholder="ГГГГ"
                  @onInput="resetErrors"
              />
            </div>
            <div v-if="userErrors.date_of_birth.text" class="auth-error__text">
              {{ userErrors.date_of_birth.text }}
            </div>
          </div>
        </div>
        <div class="auth-footer">
          <UIButton
              name="Зарегистрироваться"
              theme="colored"
              @click="submit"
          />
          <UIButton
              name="У меня уже есть аккаунт пользователя"
              font-size="small"
              theme="whited"
              @click="authState = 'authorization'"
          />
        </div>
      </div>
      <div v-if="authState == 'identification'" class="auth-wrapper">
        <div class="auth-body">
          <div class="auth-row">
            <UIInput
                v-model="identificationCode.first"
                placeholder="0"
                @onInput="resetErrors"
                :is-show-error="identificationCodeError.first.length > 0"
            />
            <UIInput
                v-model="identificationCode.second"
                placeholder="0"
                @onInput="resetErrors"
                :is-show-error="identificationCodeError.second.length > 0"
            />
            <UIInput
                v-model="identificationCode.third"
                placeholder="0"
                @onInput="resetErrors"
                :is-show-error="identificationCodeError.third.length > 0"
            />
            <UIInput
                v-model="identificationCode.four"
                placeholder="0"
                @onInput="resetErrors"
                :is-show-error="identificationCodeError.four.length > 0"
            />
          </div>
          <div v-if="identificationCodeError.text" class="auth-error__text">
            {{ identificationCodeError.text }}
          </div>
        </div>
        <div class="auth-footer">
          <UIButton
              name="Подтвердить"
              theme="colored"
              @click="submit"
          />
          <UIButton
              name="Выслать код повторно"
              font-size="small"
              theme="whited"
          />
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