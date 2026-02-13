<script lang="ts">
import { defineComponent } from 'vue'
import UIButton from '~/components/UI/UIButton.vue';
import type { IUser } from '~/types/IUser.ts'
import type { IOrder } from "~/types/IOrder.ts"

export default defineComponent({
  name: 'MainPage',
  components: { UIButton },
  data() {
    const orders: IOrder[] = [{
      id: 1,
      service_name: 'Химчистка легкового автомобиля',
      price: 20000,
      date_of_order: '10.02.2025',
      buyer_id: 1
    }]

    return {
      orders,
      user: {
        id: 1,
        name: 'Aidar',
        fullname: 'Zaynutdinov',
        surname: 'Dinarovich',
        email: 'aidar-dev@mail.ru',
        phone: '89872846800',
        birthday: '07.11.2001',
        orders
      } as IUser
    }
  }
})
</script>

<template>
  <div>
    Страница главная<br><br>
    Имя {{ user.name }}<br><br>
    Фамилия {{ user.fullname }}<br><br>
    Отчество {{ user.surname }}<br><br>
  </div>
</template>

<style></style>