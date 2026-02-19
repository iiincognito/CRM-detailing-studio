import type { IOrder } from "~/types/IOrder"

export interface IUser {
  id: number,
  name: string,
  fullname: string,
  surname: string,
  email: string,
  phone: string,
  birthday: string,
  orders: IOrder[]
}
